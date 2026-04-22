# TDTL — tKeel Data Transformation Language

TDTL is a schema-less, JSON-native expression engine for IoT and digital-twin scenarios. It evaluates computed properties directly against raw JSON payloads — no struct registration, no schema definition required.

**Key differentiators vs other Go expression engines:**
- Works on raw `[]byte` JSON without any type registration
- Native multi-entity paths: `device1.temp + device2.offset` resolved across separate JSON contexts
- IoT-oriented array aggregations: `sum`, `avg`, `min`, `max`, `count` over `[#].field` pluck results
- Parse once, eval concurrently — `CompiledExpr` is goroutine-safe

## Quick Start

```go
import "github.com/tkeel-io/tdtl"

// Compile once at startup — goroutine-safe, reuse across calls
expr, err := tdtl.Compile(`(fahrenheit - 32) * 5 / 9`)

// Eval against each incoming payload
ctx := tdtl.NewJSONContext(`{"fahrenheit":98.6}`)
result := tdtl.Eval(ctx, expr)
// result → FloatNode(37.0)
```

## Core API (v2)

### Compile + Eval

```go
// Parse once
expr, err := tdtl.Compile(exprStr)   // returns CompiledExpr (= Expr), error

// Eval many — safe to call concurrently
node := tdtl.Eval(ctx, expr)
```

`Eval` automatically includes all built-in functions (`abs`, `base64`, `sum`, `avg`, `min`, `max`, `count`).

### Context

```go
// From raw JSON bytes/string
ctx := tdtl.NewJSONContext(`{"temperature":50,"color":"red"}`)

// From an in-memory map (useful in tests or when data is already parsed)
ctx := tdtl.NewMapContext(
    map[string]tdtl.Node{"temperature": tdtl.IntNode(50)},
    map[string]tdtl.ContextFunc{"double": func(args ...tdtl.Node) tdtl.Node {
        return args[0].(tdtl.IntNode) * 2
    }},
)

// Multi-entity: chain contexts, first non-undefined match wins
ctx := tdtl.MutilContext{ctx1, ctx2, ctx3}
```

### Filter evaluation

```go
expr, err := tdtl.ParseFilter(`temperature > 30 and color = 'red'`)
pass := tdtl.EvalFilter(ctx, expr)   // bool
```

### Node types

| Type | Go type | Example |
|------|---------|---------|
| `Int` | `IntNode` (int64) | `IntNode(42)` |
| `Float` | `FloatNode` (float64) | `FloatNode(3.14)` |
| `Bool` | `BoolNode` (bool) | `BoolNode(true)` |
| `String` | `StringNode` (string) | `StringNode("hello")` |
| `JSON` / `Array` / `Object` | `JSONNode` | from context lookups |
| `Undefined` | `UNDEFINED_RESULT` | missing field |
| `Null` | `NULL_RESULT` | JSON null |

## Expression Syntax

### Operators

| Category   | Operators |
|------------|-----------|
| Arithmetic | `+` `-` `*` `/` `%` |
| Comparison | `=` `!=` `<>` `<` `<=` `>` `>=` |
| Logical    | `AND` `OR` `NOT` |
| Grouping   | `( … )` |

String `+` concatenates. Numeric strings coerce automatically (`'111' - 11` → `IntNode(100)`).

### JSON path syntax

| Path | Meaning |
|------|---------|
| `a.b` | nested field |
| `a[0]` | array element at index 0 |
| `a[#]` | array length |
| `a.#.b` | pluck field `b` from every element → JSON array |

### CASE / WHEN

```
case color
when 'red'   then 'stop'
when 'green' then 'go'
else 'wait'
```

No `END` keyword. No leading space required (fixed in v2).

### Built-in functions

| Function | Description |
|----------|-------------|
| `abs(x)` | Absolute value |
| `base64(x)` | Base-64 encode a string or JSON value |
| `sum(arr)` | Sum of numeric values in a JSON array |
| `avg(arr)` | Average of numeric values in a JSON array |
| `min(arr)` | Minimum of numeric values in a JSON array |
| `max(arr)` | Maximum of numeric values in a JSON array |
| `count(arr)` | Count of elements in a JSON array |

Aggregation example:

```go
// Payload: {"sensors":[{"temp":20},{"temp":25},{"temp":30}]}
expr, _ := tdtl.Compile(`avg(sensors.#.temp)`)
ctx := tdtl.NewJSONContext(payload)
result := tdtl.Eval(ctx, expr)   // FloatNode(25.0)
```

## Legacy SQL API (v1, deprecated)

The `INSERT INTO … SELECT … AS` SQL syntax is still supported for backwards compatibility but deprecated in v2. Prefer `Compile`/`Eval` for new code.

```go
t, err := tdtl.NewTDTL(
    `insert into target select entity1.temperature + 1 as temp`,
    nil,
)
result, _ := t.Exec(map[string]tdtl.Node{
    "entity1.temperature": tdtl.IntNode(50),
})
// result["temp"] == tdtl.IntNode(51)
```

## JSON Collection Utilities

`tdtl.New` wraps raw JSON with a fluent mutation API:

```go
c := tdtl.New(`{"a":1}`)
c.Set("b", tdtl.IntNode(2))            // {"a":1,"b":2}
c.Append("tags", tdtl.NewString("x"))  // {"a":1,"b":2,"tags":["x"]}
c.Del("a")                             // {"b":2,"tags":["x"]}
sub := c.Get("tags[0]")                // "x"

arr := tdtl.New(`[{"k":"a"},{"k":"b"}]`)
grouped := arr.GroupBy("k")   // {"a":[…],"b":[…]}
keyed   := arr.KeyBy("k")    // {"a":{…},"b":{…}}
merged  := arr.MergeBy("k")  // merged by key k
```

## Architecture

```
┌────────────────────────────────────────────────────┐
│          Layer 1: Expression Engine                 │
│  Compile / Eval / ParseFilter / EvalFilter          │
│  parse.go  evaluator.go  types.go                   │
└───────────────┬──────────────────┬─────────────────┘
                │                  │
┌───────────────▼──────┐  ┌────────▼───────────────┐
│  Layer 2: Context    │  │  Layer 3: JSON Utils    │
│  JSONContext          │  │  (collectjs)            │
│  MapContext           │  │  Set/Get/Del/Append     │
│  MutilContext         │  │  GroupBy/KeyBy/MergeBy  │
└──────────────────────┘  └────────────────────────┘
```

`CompiledExpr` (the AST) is immutable after `Compile`. `Eval` takes a `Context` argument — create one per goroutine/call. No shared mutable state.

## Development

```sh
# Run tests
go test ./...

# Run tests with race detector
go test -race ./...

# Coverage report
go test -coverprofile=cover.out ./... && go tool cover -html=cover.out

# Run benchmarks
go test -bench=. -benchmem github.com/tkeel-io/tdtl

# Regenerate parser (requires ANTLR4)
antlr4 -Dlanguage=Go -o parser TDTL.g4
```

## Migrating from v1 to v2

| v1 | v2 |
|----|----|
| `ParseExpr(s)` + `EvalRuleQL(ctx, expr)` | `Compile(s)` + `Eval(ctx, expr)` |
| `EvalRuleQL` wraps DefaultValue | `Eval` wraps DefaultValue |
| CASE needs leading space: ` case x when...` | CASE works without leading space |
| No aggregations | `sum/avg/min/max/count` built in |
| `NewTDTL` for routing | Deprecated; implement routing in caller |

`MutilContext` spelling preserved (typo from v1 — renaming would be a breaking change).

## License

Apache 2.0 — see [LICENSE](LICENSE).
