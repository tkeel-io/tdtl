# TDTL v2 — Architecture Document

**Date:** 2026-04-19
**Version:** 2.0
**Status:** Draft

---

## Architecture Overview

TDTL v2 is structured as three independent, composable layers. Each layer has a single responsibility and can be used standalone.

```
┌─────────────────────────────────────────────────────┐
│               Caller (e.g., tKeel core)             │
│  Routing, property binding, change propagation      │
└───────────────────────┬─────────────────────────────┘
                        │  uses
┌───────────────────────▼─────────────────────────────┐
│          Layer 1: Expression Engine                  │
│  Parse / Compile / Eval / EvalFilter                 │
│  Input: string → Expr  |  Context + Expr → Node     │
│  Files: parse.go, evaluator.go, types.go             │
└──────────────┬──────────────────┬───────────────────┘
               │ reads            │ reads
┌──────────────▼──────┐  ┌───────▼──────────────────┐
│  Layer 2: Context   │  │  Layer 3: JSON Utilities  │
│  Adapters           │  │  (collectjs)              │
│  JSONContext        │  │  New / Set / Get / Del    │
│  MapContext         │  │  Append / GroupBy         │
│  MutilContext       │  │  KeyBy / MergeBy          │
│  Files:             │  │  Files:                   │
│  context.go         │  │  collectjs.go             │
│  context_json.go    │  │  collectjs_types.go       │
│  context_map.go     │  │  collectjs_patch.go       │
└─────────────────────┘  └──────────────────────────┘
```

---

## Layer 1: Expression Engine

### Responsibility

Parse expression strings into an immutable AST (`Expr`), then evaluate that AST against a `Context` to produce a `Node`.

### Key Design Decisions

**1. Parse once, eval many**

`Compile(str) (CompiledExpr, error)` is the primary entry point for production use. It returns an immutable, goroutine-safe expression that can be evaluated concurrently. `ParseExpr`/`ParseFilter` are convenience wrappers that return the same underlying type.

```
exprStr  ─────►  Lexer (ANTLR4)  ─────►  Parser (ANTLR4)
                                              │
                                      TDTLListener
                                    (ExitBinary, ExitString, …)
                                              │
                                         Expr (AST)
                                       (immutable tree)
                                              │
                              Eval(ctx, expr) │
                                              ▼
                                 Context.Get(path) ──► Node
                                              │
                                           Node
                               (Int / Float / Bool / String / JSON / Undefined / Null)
```

**2. No INSERT INTO in the engine**

`INSERT INTO target SELECT ... AS alias` routing is removed from the grammar and listener. Callers that need routing implement it themselves using the `Eval` return value and their own state store.

**3. CASE keyword whitespace fix**

The ANTLR4 grammar `STUFF` fragment currently requires whitespace on both sides of keywords. Fix: restructure `switch_stmt` rule to treat `case`/`when`/`then`/`else` as reserved identifiers that don't require leading whitespace. Affects `TDTL.g4` and regenerated parser.

### AST Node Types

```
Expr (interface)
├── BinaryExpr      { Op int, LHS Expr, RHS Expr }
├── JSONPathExpr    { path string }
├── CallExpr        { key string, args []Expr }
├── SwitchExpr      { exp Expr, list []*CaseExpr, last Expr }
├── CaseExpr        { when Expr, then Expr }
├── SelectStatementExpr  { fields FieldsExpr, filter *FilterExpr }
└── Literal nodes: IntNode, FloatNode, BoolNode, StringNode
```

### Evaluator

`eval(ctx Context, expr Expr) Node` is the recursive dispatcher:

```
eval(ctx, BinaryExpr)
  → eval LHS, eval RHS
  → if isBooleanOP(op): evalBooleanOP(lhs, rhs)
  → if isLogicOP(op):   evalLogicOP(lhs, rhs)
  → else:               evalBinaryInt / evalBinaryFloat / evalBinaryString

eval(ctx, JSONPathExpr)
  → ctx.Get(path) → Node

eval(ctx, CallExpr)
  → ctx.Get(key) → ContextFunc
  → evaluate args
  → call func(args)

eval(ctx, SwitchExpr)
  → eval subject
  → iterate cases: eval when, compare, if match: eval then
  → if no match: eval last (else) or UNDEFINED_RESULT
```

### New: Array Aggregation Functions

Registered in `DefaultValue` context alongside `abs` and `base64`:

```go
"sum":   aggFuncHandle(sumAgg),
"avg":   aggFuncHandle(avgAgg),
"min":   aggFuncHandle(minAgg),
"max":   aggFuncHandle(maxAgg),
"count": aggFuncHandle(countAgg),
```

`aggFuncHandle` unpacks a `JSONNode` array argument and applies the aggregation. Works with the output of `[#].field` pluck expressions.

---

## Layer 2: Context Adapters

### Responsibility

Resolve JSON path lookups and function lookups for the evaluator. Context is read-only during evaluation.

### Interface

```go
type Context interface {
    Get(path string) Node
}
```

### Implementations

**JSONContext** (`context_json.go`)
- Wraps raw JSON `[]byte`
- Path resolution via embedded gjson (`pkg/json/gjson`)
- `Get("a.b[0].c")` → gjson path lookup → Node conversion
- JSON type mapping: `number` (int/float) → `IntNode`/`FloatNode`, `string` → `StringNode`, `true/false` → `BoolNode`, `null` → `NULL_RESULT`, missing → `UNDEFINED_RESULT`

**MapContext** (`context_map.go`)
- Wraps `map[string]Node` for in-process testing and custom data
- Also holds `map[string]ContextFunc` for user-registered functions
- Function lookup: `Get(name)` checks `funcs` map, returns wrapped callable node

**DefaultValue** (`context.go`)
- Static context providing built-in functions: `abs`, `base64`, `sum`, `avg`, `min`, `max`, `count`
- Always included last in `EvalRuleQL`'s `MutilContext`

**MutilContext** (`context.go`)
- `[]Context` slice; `Get` iterates in order, returns first non-undefined result
- Enables multi-entity expressions: `entity1.temp + entity2.offset` resolved by placing each entity's context at the appropriate index

### Context Lookup Flow

```
MutilContext.Get("entity1.temperature")
  → try ctx[0].Get("entity1.temperature") → UNDEFINED_RESULT (not found)
  → try ctx[1].Get("entity1.temperature") → FloatNode(23.5) ✓
  → return FloatNode(23.5)
```

---

## Layer 3: JSON Collection Utilities

### Responsibility

Provide a fluent API for reading and mutating JSON documents as typed `Node` trees. This layer is independent of the expression engine and can be used standalone.

### Design

`Collect` struct wraps `[]byte` with an `error` field (railway-oriented error propagation: operations on an errored `Collect` are no-ops).

```go
type Collect struct {
    data    []byte
    err     error
    datatype CollectType  // Object, Array, String, Number, Bool, Null, Undefined
}
```

Path operations delegate to:
- **Read:** `pkg/json/gjson` for zero-copy path access
- **Write:** `pkg/json/sjson` for immutable JSON patch operations

### Key Methods

| Method | Description |
|--------|-------------|
| `Get(path)` | Returns Node at dot/bracket path |
| `Set(path, node)` | Returns new Collect with value set at path |
| `Del(path)` | Returns new Collect with path removed |
| `Append(path, node)` | Appends to array at path; creates array if absent |
| `GroupBy(key)` | `[{k:a,…},{k:b,…}]` → `{a:[…],b:[…]}` |
| `KeyBy(key)` | `[{k:a,…},{k:b,…}]` → `{a:{…},b:{…}}` |
| `MergeBy(key)` | Merges array elements sharing the same key |

### Error Propagation

```go
// Errors propagate silently — callers check at the end
result := New(`[{"id":"a"}]`).GroupBy("id")
if result.Error() != nil { ... }
```

---

## Data Flow: End-to-End Example

```
Incoming device JSON: {"fahrenheit": 80, "humidity": 65}

1. CALLER: loads CompiledExpr for "(fahrenheit - 32) * 5 / 9"
   (compiled at startup, reused for every message)

2. CALLER: ctx = tdtl.NewJSONContext(devicePayload)

3. TDTL: Eval(ctx, compiledExpr)
   → eval BinaryExpr{*, LHS: BinaryExpr{-, …}, RHS: FloatNode(9)}
   → eval BinaryExpr{-, LHS: JSONPathExpr{"fahrenheit"}, RHS: IntNode(32)}
   → ctx.Get("fahrenheit") → IntNode(80)   [gjson lookup]
   → evalBinaryInt(-, 80, 32) → IntNode(48)
   → evalBinaryFloat(*, 48.0, 5.0/9.0) → FloatNode(26.67)

4. CALLER: writes FloatNode(26.67) to digital twin state store
```

---

## Key Architectural Constraints

### Goroutine Safety

`CompiledExpr` (the AST) is immutable after creation. `Eval` takes a `Context` parameter — callers create a new `Context` per goroutine/call. No shared mutable state during evaluation.

⚠️ `NewJSONContext` and `NewMapContext` are not goroutine-safe to create concurrently (standard Go map behavior). Create per-call or use appropriate synchronization in the caller.

### No Global State Mutations

Built-in functions in `DefaultValue` are registered once at package init. User-registered functions are per `MapContext` instance. No global function registry to avoid races.

### ANTLR4 Parser Lifecycle

ANTLR4 `Lexer`+`Parser` objects are created per `Parse`/`Compile` call (they are not goroutine-safe internally). The resulting `Expr` AST is plain Go interfaces — goroutine-safe, allocation-free to share.

For high-throughput scenarios, callers should pool parse calls using `sync.Pool` if needed, or (preferred) simply compile once and reuse.

---

## File Structure (v2)

```
tdtl/
├── types.go              # Node interface, type constants, UNDEFINED/NULL sentinels
├── parse.go              # TDTLListener, AST builder, Parse/ParseExpr/ParseFilter
├── evaluator.go          # eval(), isBooleanOP(), evalBinary*, EvalRuleQL, EvalFilter
├── context.go            # DefaultValue (built-ins), MutilContext, ContextFunc type
├── context_json.go       # JSONContext
├── context_map.go        # MapContext
├── collectjs.go          # Collect struct, Set/Get/Del/Append/GroupBy/KeyBy/MergeBy
├── collectjs_types.go    # CollectType, UNDEFINED_RESULT, NULL_RESULT
├── collectjs_patch.go    # JSON patch helpers
├── expression.go         # Expr interface, SelectStatementExpr, FieldExpr, etc.
├── tdtl.go               # NewTDTL, public API surface
├── codec.go              # Node ↔ JSON serialization helpers
├── utils.go              # typeOf, shared utilities
├── parser/               # ANTLR4-generated lexer/parser (do not edit manually)
│   ├── tdtl_lexer.go
│   ├── tdtl_parser.go
│   └── ...
├── pkg/json/
│   ├── gjson/            # Vendored gjson for read path
│   └── sjson/            # Vendored sjson for write path
└── TDTL.g4               # Grammar source (edit here, regenerate parser/)
```

---

## v2 Migration: What Changes in Each File

| File | Change |
|------|--------|
| `TDTL.g4` | Remove `insert_stmt` rule; fix CASE keyword whitespace; add array aggregation function names to built-in list |
| `parser/` | Regenerate after grammar change |
| `parse.go` | Remove `ExitTargetAsElem`, `ExitFieldElemSource`, `ExitRoot` (routing-related); fix `ExitSwitch_stmt` for no-whitespace CASE |
| `evaluator.go` | Add `sum/avg/min/max/count` dispatch in `CallExpr` evaluation |
| `context.go` | Register `sum/avg/min/max/count` in `DefaultValue`; add `aggFuncHandle` |
| `types.go` | Add `CompiledExpr` type alias / wrapper for concurrency documentation |
| `tdtl.go` | Add `Compile(str) (CompiledExpr, error)` public API; deprecate `NewTDTL` |
| `expression.go` | Remove `SelectStatementExpr`, `FieldExpr`, `FieldsExpr`, `FilterExpr` |

---

## Architecture Decision Records

### ADR-001: Remove INSERT INTO from Expression Core

**Decision:** `INSERT INTO target SELECT ... AS alias` is removed from the TDTL grammar and evaluator.

**Rationale:** Mixes three orthogonal concerns (expression evaluation, routing, property binding) in one statement. Each concern has a different lifecycle and owner. The routing layer belongs to the caller (e.g., tKeel runtime), not the expression engine.

**Consequences:** Breaking change for callers using the SQL DSL. Mitigated by major version bump (`v2`) and migration guide.

### ADR-002: Compile Returns Immutable AST (Not Bytecode)

**Decision:** `Compile` returns the same `Expr` AST as `Parse`, wrapped in a type alias marking it as safe for concurrent use.

**Rationale:** ANTLR4 produces an AST of plain Go interfaces. These are immutable after construction. A bytecode VM would improve throughput but adds significant complexity. Profile first; add bytecode compilation only if benchmark evidence demands it.

**Consequences:** `CompiledExpr` is goroutine-safe. Performance improvement vs `Parse` is only the elimination of repeated ANTLR4 invocations.

### ADR-003: Array Aggregations as Built-in Functions (not Grammar)

**Decision:** `sum`, `avg`, `min`, `max`, `count` are registered as `ContextFunc` in `DefaultValue`, not as grammar keywords.

**Rationale:** No grammar change = no parser regeneration = smaller blast radius. The `CallExpr` path already handles `abs()` and `base64()` with the same pattern.

**Consequences:** These names are not reserved keywords; a user could override them via `MapContext`. Document this as a feature (user override) not a bug.

### ADR-004: MutilContext Spelling Preserved (No Rename)

**Decision:** `MutilContext` is not renamed to `MultiContext` in v2.

**Rationale:** Renaming would be a breaking API change with no functional benefit. The typo is unfortunate but not worth the churn.

**Consequences:** Document the name in migration guide with a note acknowledging the typo.
