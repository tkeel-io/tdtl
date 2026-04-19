---
stepsCompleted: ['step-01-init','step-02-discovery','step-02b-vision','step-02c-executive-summary','step-03-success','step-04-journeys','step-05-domain','step-06-innovation','step-07-project-type','step-08-scoping','step-09-functional','step-10-nonfunctional','step-11-polish','step-12-complete']
inputDocuments: ['README.md']
workflowType: 'prd'
classification:
  projectType: developer_tool
  domain: iot_digital_twin
  complexity: medium
  projectContext: brownfield
---

# Product Requirements Document — TDTL v2

**Author:** Root
**Date:** 2026-04-19
**Version:** 2.0
**Status:** Draft

---

## Executive Summary

TDTL v2 is a Go library providing **schema-less JSON-native expression evaluation** for IoT and digital twin platforms. It enables developers to compute derived properties from heterogeneous, multi-source JSON entity data without upfront schema registration.

**Target users:** Go developers building IoT data pipelines, digital twin runtimes (e.g., tKeel core), and edge computing systems that transform raw device telemetry into computed properties.

**Core problem solved:** IoT devices produce heterogeneous JSON payloads with dynamic, unknown-at-compile-time schemas. Existing Go expression engines (cel-go, expr-lang) require typed schema or struct binding before evaluation. TDTL evaluates directly against raw JSON bytes, making it the only Go expression library designed for schema-less, multi-entity IoT data.

### What Makes TDTL Special

Three capabilities unavailable in any single Go expression library:

1. **Zero-schema evaluation** — expressions run against raw `[]byte` JSON; no struct, no protobuf, no `map[string]any` conversion required
2. **Multi-entity JSON path syntax** — `entity1.temperature + entity2.offset` is first-class grammar, not a workaround
3. **IoT array operations** — `avg(sensors[#].value)`, `sum(meters[#].kwh)` address the most common IoT data shape (arrays of device readings)

### Architectural Direction (v1 → v2)

TDTL v1 mixed three orthogonal concerns into one DSL:
- Expression evaluation (`SELECT source.field`)
- Routing (`INSERT INTO target`)
- Property binding (`AS alias`)

TDTL v2 separates these cleanly: the library owns **only expression evaluation**. Routing and property binding are the caller's responsibility (e.g., tKeel core runtime).

## Project Classification

- **Project Type:** Developer tool (Go open-source library)
- **Domain:** IoT / Digital Twin / Edge Computing
- **Complexity:** Medium — no compliance requirements, but real-time and concurrent-use constraints apply
- **Context:** Brownfield — refactoring existing TDTL v1 codebase
- **Innovation signal:** DSL creation with IoT-native path syntax

---

## Success Criteria

### User Success

- A Go developer integrates TDTL v2 and evaluates their first expression in < 30 minutes using only the README Quick Start
- Zero struct definitions required: expressions evaluate against any JSON payload without type registration
- Parse errors return descriptive `error` values with position info; no call to any exported function triggers `panic`

### Technical Success

- `go test ./... -race` passes with 0 failures and 0 data races
- Test coverage ≥ 90% on `evaluator.go`, `context*.go`, and `collectjs.go` as measured by `go test -cover`
- `CompiledExpr` evaluation completes in < 10μs at 99th percentile for primitive-result expressions, measured by `go test -bench`
- Library compiles and tests pass on `linux/amd64`, `linux/arm64`, `darwin/arm64`

### Adoption Success

- At least one external IoT/digital twin project (beyond tKeel) adopts TDTL v2 within 6 months of release
- v2 core API (`Parse`, `Eval`, `EvalFilter`) is backwards-compatible with existing v1 call sites

---

## Product Scope

### MVP — Phase 1 (v2.0)

Core architectural cleanup and highest-impact missing capabilities:

**Must-Have Capabilities:**
- Remove `INSERT INTO ... SELECT` from expression core; caller handles routing
- Explicit `Compile(exprStr) (CompiledExpr, error)` API separating parse from eval
- Array aggregation functions: `sum`, `avg`, `min`, `max`, `count` operating on `[#]` pluck results
- CASE expression fix: keywords (`CASE`, `WHEN`, `THEN`, `ELSE`) work without mandatory leading whitespace
- All panics converted to `error` returns: `ExitInteger`, `ExitFloat`, `ExitTargetAsElem`, `Expr()` empty-stack guard
- `CompiledExpr` is goroutine-safe (no mutable shared state during evaluation)

**Core User Journeys Supported:**
- IoT developer computing device properties (Journey 1)
- Architect evaluating multi-entity aggregations (Journey 2)
- Library maintainer extending functions (Journey 3)

### Growth — Phase 2 (v2.1+)

- Stateful context: sliding window, rate-of-change, last-N-values
- String functions: `upper`, `lower`, `trim`, `substr`, `replace`, `len`
- Type conversion: `to_int`, `to_float`, `to_string`, `coalesce`
- Time functions: `now()`, `timestamp()`, `date_format()`
- JSONPath filter expressions: `sensors[?value > 30]`

### Vision — Phase 3 (v3.0+)

- Property binding layer: declare computed properties with typed dependency graph
- DAG-based change propagation engine (what tKeel core currently implements manually)
- DTDL schema compatibility layer for computed property definitions

### Risk Mitigation

- **Technical risk:** ANTLR4 grammar changes can break generated parser; mitigate by adding grammar regression tests before any grammar edit
- **API compatibility risk:** removing `INSERT INTO` is a breaking change for existing users; mitigate by providing `tdtl/v2` module path and migration guide
- **Concurrency risk:** ANTLR4 parser instances are not goroutine-safe; mitigate by making `Compile` return an immutable AST that is safe to share

---

## User Journeys

### Journey 1: IoT Developer Computing Device Properties (Primary — Happy Path)

**Persona:** Wei, backend engineer at a factory monitoring startup using tKeel
**Situation:** Receives JSON telemetry from 50+ heterogeneous sensor models. Each model has slightly different field names and units. Needs to normalize and compute derived values before writing to digital twin state.
**Goal:** Define rules like `celsius = (fahrenheit - 32) * 5 / 9` and `status = case when temp > 80 then 'critical' else 'normal'` that work against raw device JSON without writing struct types for each sensor model.
**Obstacle:** Today he'd need to unmarshal each payload to a Go struct or map, write type assertions, handle missing fields manually.

**Journey:**
1. Wei adds `go get github.com/tkeel-io/tdtl/v2` — zero extra dependencies beyond Go toolchain
2. He calls `tdtl.Compile("(fahrenheit - 32) * 5 / 9")` once at application startup for each rule
3. For each incoming MQTT message, he calls `tdtl.Eval(tdtl.NewJSONContext(msg.Payload), compiled)`
4. TDTL returns `FloatNode(26.67)` from raw JSON `{"fahrenheit": 80}` — no struct, no unmarshal
5. If a payload lacks the `fahrenheit` field, TDTL returns `UNDEFINED_RESULT`; Wei handles that case in one place

**Resolution:** Wei defines 30 sensor normalization rules in 40 lines of Go; previously needed 300+ lines with per-model struct types.

### Journey 2: Digital Twin Architect — Multi-Entity Aggregation (Primary — Complex Path)

**Persona:** Lin, platform architect at tKeel designing building-level digital twins
**Situation:** A building twin must expose `avg_floor_temp` computed from 10 floor entity temperatures. Each floor entity is a separate JSON document in the state store.
**Goal:** Define `avg_floor_temp` as an expression that spans multiple entities, recomputed whenever any floor entity changes.

**Journey:**
1. Lin creates `MutilContext{floor1ctx, floor2ctx, ..., floor10ctx}` combining 10 JSON contexts
2. Compiles `avg(floors[#].temperature)` — the `[#].temperature` plucks the `temperature` field from each context
3. On any floor temperature change, the runtime evaluates the compiled expression against the updated multi-context
4. TDTL returns the average as `FloatNode` — Lin writes it back to the building twin without custom aggregation code

**Resolution:** What previously required 50 lines of manual aggregation is one compiled expression evaluated by TDTL.

### Journey 3: Library Maintainer — Registering Custom Functions (Operations Path)

**Persona:** tKeel core team member adding `rate_of_change` for edge analytics
**Situation:** Operators want to detect rapid pressure changes using `rate_of_change(pressure.current, pressure.previous, pressure.interval)`.
**Goal:** Register a custom function callable from TDTL expressions without modifying the TDTL core library.

**Journey:**
1. Maintainer implements `func rateOfChange(args []tdtl.Node) tdtl.Node` following `ContextFunc` signature
2. Registers via `tdtl.NewMapContext(data, map[string]tdtl.ContextFunc{"rate_of_change": rateOfChange})`
3. Expression `rate_of_change(pressure.current, pressure.previous, pressure.interval)` evaluates correctly
4. Unit tests use `NewMapContext` with known values; no mocking of TDTL internals required

**Resolution:** Custom domain functions are first-class without forking TDTL.

### Journey 4: Error Recovery Path

**Persona:** Wei from Journey 1, receiving a malformed payload
**Situation:** A sensor sends `{"fahrenheit": "N/A"}` (non-numeric string) during calibration.
**Goal:** Expression evaluation degrades gracefully; system continues processing other payloads.

**Journey:**
1. `tdtl.Eval(ctx, compiled)` returns `UNDEFINED_RESULT` (not panic, not error)
2. Wei checks `result.Type() == tdtl.Undefined` and logs the event
3. No goroutine crash, no recovery needed, processing continues for next payload

### Journey Requirements Summary

| Capability Required | Journey |
|---|---|
| Zero-schema JSON evaluation | 1, 2, 4 |
| Multi-entity context | 2 |
| Array pluck + aggregation | 2 |
| Custom function registration | 3 |
| Graceful undefined propagation | 1, 4 |
| Compile-once, eval-many pattern | 1, 2 |
| Concurrent evaluation safety | 1, 2 |

---

## Domain-Specific Requirements

### IoT / Edge Deployment Constraints

- **Air-gap compatibility:** Zero network I/O during library operation. No calls to external services, no DNS lookups, no telemetry. Library must function in factory networks with no internet access.
- **Resource constraints:** Suitable for `linux/arm64` edge nodes (Raspberry Pi class hardware). Binary size contribution < 5MB. No CGO dependencies.
- **Resilience:** Any malformed, truncated, or structurally unexpected JSON input returns UNDEFINED_RESULT or an error. No undefined behavior, no memory corruption vectors.

### Digital Twin Semantics

- **Missing field → UNDEFINED:** In digital twins, an entity that hasn't yet reported a property is semantically different from an entity that reported `null`. TDTL preserves this via `UNDEFINED_RESULT` vs `NULL_RESULT` distinction.
- **Multi-source consistency:** When multiple contexts provide the same key, the first context in `MutilContext` wins. This is documented and predictable.

---

## Innovation & Novel Patterns

### Detected Innovation: Schema-less IoT DSL in Go

No existing Go expression library combines:
1. Raw JSON byte evaluation (no struct/schema requirement)
2. IoT-specific array path syntax (`[#]`, `[#].field`) as grammar
3. Multi-entity context as first-class concept

TDTL's innovation is in the **intersection** — individually, gjson handles JSON paths, cel-go handles expressions, but neither addresses the multi-entity, schema-less, IoT-native combination.

### Market Context

- **cel-go** (Google): requires protobuf/struct type registration; adopted for API gateway policies and Kubernetes admission controllers
- **expr-lang**: requires Go struct or `map[string]any`; adopted for business rule engines with known schemas
- **govaluate**: no IoT array syntax; adopted for simple dynamic calculations
- **TDTL v2 niche**: IoT platforms where device schemas change frequently, data arrives as raw JSON, and aggregating across entity instances is the primary computation pattern

### Validation Approach

The innovation is validated by the existing tKeel core adoption. v2 validation: if the tKeel core subscription.go can replace its manual JSON manipulation with TDTL v2 expression evaluation using fewer lines of code, the design is confirmed.

---

## Developer Tool Specific Requirements

### Language & Platform Matrix

| Target | Support Level |
|---|---|
| Go 1.21+ | Full support |
| `linux/amd64` | Primary CI target |
| `linux/arm64` | Required (edge devices) |
| `darwin/arm64` | Required (developer machines) |
| CGO | Not used |

### Installation

```bash
go get github.com/tkeel-io/tdtl/v2
```

Single command, no system dependencies, no build tags required.

### API Surface (v2)

```go
// Parse and compile for repeated evaluation
compiled, err := tdtl.Compile(exprStr)

// One-shot parse (convenience wrapper)
expr, err := tdtl.Parse(exprStr)
expr, err := tdtl.ParseExpr(exprStr)
expr, err := tdtl.ParseFilter(exprStr)

// Evaluate
result := tdtl.Eval(ctx, expr)           // returns Node
pass   := tdtl.EvalFilter(ctx, expr)     // returns bool

// Context construction
ctx := tdtl.NewJSONContext(jsonStr)
ctx := tdtl.NewMapContext(map[string]tdtl.Node{}, extFuncs)
ctx := tdtl.MutilContext{ctx1, ctx2}

// Node types
tdtl.IntNode(42)
tdtl.FloatNode(3.14)
tdtl.BoolNode(true)
tdtl.StringNode("hello")
tdtl.New(`{"a":1}`)   // JSONNode / collection
```

### Migration Guide (v1 → v2)

**Non-breaking changes (v1 call sites continue to work):**
- `Parse`, `ParseExpr`, `ParseFilter` signatures unchanged
- `Eval`, `EvalFilter`, `EvalRuleQL` unchanged
- `NewJSONContext`, `NewMapContext`, `MutilContext` unchanged

**Breaking changes:**
- `INSERT INTO ... SELECT` DSL removed from expression engine; routing moved to caller
- Panics in parse callbacks → descriptive `error` returns (callers not catching panics are unaffected; callers with `recover()` may need updates)
- Module path: `github.com/tkeel-io/tdtl/v2` (major version bump)

**New in v2:**
- `tdtl.Compile(exprStr) (CompiledExpr, error)` — parse + goroutine-safe result
- Array aggregation functions: `sum`, `avg`, `min`, `max`, `count`

---

## Functional Requirements

### Expression Parsing

- **FR1:** Developers can parse arithmetic expressions (`+`, `-`, `*`, `/`, `%`) returning an `Expr` without providing schema or type information
- **FR2:** Developers can parse comparison expressions (`=`, `!=`, `<>`, `<`, `<=`, `>`, `>=`) returning an `Expr` that evaluates to a `BoolNode`
- **FR3:** Developers can parse CASE/WHEN/THEN/ELSE expressions; keywords are recognized case-insensitively and require no leading whitespace
- **FR4:** Developers can parse function call expressions with 0 or more arguments
- **FR5:** `Parse`, `ParseExpr`, `ParseFilter`, and `Compile` return a non-nil `error` for syntactically invalid input; no function panics on any input

### JSON Path Access

- **FR6:** Expressions access nested JSON object fields via dot notation (`entity.field.subfield`) of arbitrary depth
- **FR7:** Expressions access JSON array elements by zero-based index (`arr[0]`, `arr[2]`)
- **FR8:** Expressions retrieve array length with `arr[#]`, returning `IntNode`
- **FR9:** Expressions pluck a field from every array element with `arr[#].field`, returning a `JSONNode` array
- **FR10:** Path access on a missing or null field returns `UNDEFINED_RESULT`; evaluation continues without error

### Array Aggregation Functions

- **FR11:** `sum(arr[#].field)` sums all numeric values in the plucked array, returning `IntNode` or `FloatNode`
- **FR12:** `avg(arr[#].field)` computes the arithmetic mean, returning `FloatNode`
- **FR13:** `min(arr[#].field)` returns the minimum numeric value as `IntNode` or `FloatNode`
- **FR14:** `max(arr[#].field)` returns the maximum numeric value as `IntNode` or `FloatNode`
- **FR15:** `count(arr[#].field)` returns the count of non-undefined elements as `IntNode`
- **FR16:** All aggregation functions return `UNDEFINED_RESULT` for empty arrays or arrays where all elements are undefined

### Multi-Entity Context

- **FR17:** `MutilContext` evaluates path lookups against each contained context in declaration order; first non-undefined result wins
- **FR18:** Developers can pass `MutilContext` anywhere a `Context` interface is accepted
- **FR19:** Developers register custom functions as `map[string]ContextFunc` in `NewMapContext`; functions are callable from any expression evaluated against that context

### Built-in Functions

- **FR20:** `abs(x)` returns the absolute value for `IntNode` and `FloatNode`; returns `UNDEFINED_RESULT` for other node types
- **FR21:** `base64(x)` returns a base64-encoded `StringNode` for any non-undefined node
- **FR22:** Custom functions registered via `ContextFunc` receive `[]Node` arguments and return `Node`; they are resolved before built-ins (caller override is possible)

### Expression Compilation

- **FR23:** `Compile(exprStr) (CompiledExpr, error)` parses and returns an immutable compiled expression
- **FR24:** A `CompiledExpr` can be passed to `Eval` and `EvalFilter` concurrently from multiple goroutines without data races
- **FR25:** `Compile` returns a descriptive error for invalid expressions including the token position

### Type System & Coercion

- **FR26:** Numeric string operands (`"42"`, `"3.14"`) are automatically coerced to `IntNode` or `FloatNode` in arithmetic operations
- **FR27:** `+` with at least one non-numeric string operand performs string concatenation
- **FR28:** `UNDEFINED_RESULT` propagates through all arithmetic operations: any arithmetic with undefined returns undefined
- **FR29:** Division and modulo by zero return `UNDEFINED_RESULT`
- **FR30:** `NULL_RESULT.Type()` returns `Null` (not `Undefined`); the two sentinel values are distinguishable

### JSON Collection Utilities

- **FR31:** `collection.Set(path, node)` sets the value at the given dot-notation path, creating intermediate objects as needed
- **FR32:** `collection.Get(path)` retrieves the `Node` at the given path; returns `UNDEFINED_RESULT` for missing paths
- **FR33:** `collection.Del(path)` removes the key at the given path; no-op for missing paths
- **FR34:** `collection.Append(path, node)` appends a value to the array at the given path, creating the array if absent
- **FR35:** `collection.GroupBy(key)` partitions a JSON array into a JSON object keyed by the given field value
- **FR36:** `collection.KeyBy(key)` converts a JSON array to a JSON object keyed by the given field, keeping one element per key
- **FR37:** `collection.MergeBy(key)` merges objects in a JSON array sharing the same key value into single objects

---

## Non-Functional Requirements

### Performance

- **NFR1:** Evaluation of a `CompiledExpr` against a `JSONContext` with a primitive result (`Int`, `Float`, `Bool`, `String`) completes in < 10μs at 99th percentile under single-goroutine sustained load, as measured by `go test -bench=./...`
- **NFR2:** `Compile(exprStr)` for expressions ≤ 200 characters completes in < 1ms, as measured by benchmarks
- **NFR3:** Each evaluation of a compiled primitive-result expression allocates ≤ 5 heap objects, as measured by `go test -bench -benchmem`

### Reliability

- **NFR4:** No call to any exported function panics under any valid or invalid input; all failure modes produce `error` returns or `UNDEFINED_RESULT` nodes
- **NFR5:** `CompiledExpr` is safe for concurrent evaluation across goroutines: `go test -race ./...` reports zero data races

### Correctness

- **NFR6:** Line coverage ≥ 90% for `evaluator.go`, `context*.go`, `collectjs.go`, and `parse.go` as measured by `go test -cover`
- **NFR7:** Integer arithmetic that would overflow 64-bit signed range returns `UNDEFINED_RESULT`

### Portability

- **NFR8:** `go test ./...` passes without modification on `linux/amd64`, `linux/arm64`, and `darwin/arm64`
- **NFR9:** The library performs no network I/O; suitable for air-gapped edge deployments
- **NFR10:** No CGO dependencies; library compiles with `CGO_ENABLED=0`
