# TDTL v2 — Epics & Stories

**Date:** 2026-04-19
**Source:** PRD v2.0 + Architecture v2.0

---

## Epic Overview

| Epic | Name | Priority | FRs Covered |
|------|------|----------|-------------|
| E1 | Architecture Cleanup | P0 — Blocker | ADR-001, ADR-002 |
| E2 | Array Aggregation Functions | P0 — Core Differentiator | FR11–FR16 |
| E3 | CASE Expression Fix | P0 — Usability Bug | FR3 |
| E4 | Error-First API (No Panics) | P0 — Reliability | FR5, NFR4 |
| E5 | Compile API & Goroutine Safety | P1 — Performance | FR23–FR25, NFR5 |
| E6 | Test Coverage & Benchmarks | P1 — Quality Gate | NFR1, NFR3, NFR6 |
| E7 | Documentation & Migration | P1 — Adoption | Developer Tool requirements |

---

## E1 — Architecture Cleanup

**Goal:** Remove `INSERT INTO ... SELECT` routing from expression core. TDTL becomes a pure expression evaluator. Routing moves to caller.

**Acceptance Criteria:**
- `INSERT INTO ... SELECT ... AS` DSL is removed from `TDTL.g4` and generated parser
- `SelectStatementExpr`, `FieldExpr`, `FieldsExpr` removed from `expression.go`
- `ExitTargetAsElem`, `ExitFieldElemSource`, `ExitRoot` removed from `parse.go`
- `EvalRuleQL` now returns `Node` (result of single expression) not JSON object of aliased fields, OR is deprecated with documented replacement
- `go test ./...` passes with 0 failures after removal

---

### Story E1-S1: Remove INSERT INTO grammar rule

**As a** library maintainer
**I want to** remove the `insert_stmt` rule from `TDTL.g4`
**So that** TDTL is scoped to expression evaluation only

**Acceptance Criteria:**
- [ ] `TDTL.g4`: `insert_stmt`, `fields`, `field_elem`, `target`, `target_name` rules removed
- [ ] Grammar regenerated: `parser/` files updated via `antlr4 -Dlanguage=Go -o parser TDTL.g4`
- [ ] `go build ./...` succeeds after regeneration

**Dependencies:** None
**Estimate:** S (2–4h)

---

### Story E1-S2: Remove routing-related listener methods

**As a** library maintainer
**I want to** remove `ExitTargetAsElem`, `ExitFieldElemSource`, `ExitRoot`, and related fields from `TDTLListener`
**So that** `parse.go` only builds expression ASTs, not routing structures

**Acceptance Criteria:**
- [ ] `ExitTargetAsElem` removed from `parse.go`
- [ ] `ExitFieldElemSource` removed from `parse.go`
- [ ] `ExitRoot` removed or simplified to no-op
- [ ] `TDTLListener.target`, `.fields` fields removed
- [ ] `SelectStatementExpr`, `FieldExpr`, `FieldsExpr` removed from `expression.go`
- [ ] `go test ./...` passes

**Dependencies:** E1-S1
**Estimate:** S (2–4h)

---

### Story E1-S3: Deprecate EvalRuleQL, introduce Eval

**As a** TDTL user
**I want to** call `tdtl.Eval(ctx, expr)` returning a single `Node`
**So that** I have a clean, routing-free evaluation API

**Acceptance Criteria:**
- [ ] `tdtl.Eval(ctx Context, expr Expr) Node` exported function added to `tdtl.go`
- [ ] `EvalRuleQL` marked `// Deprecated: use Eval` with a note in godoc
- [ ] `EvalRuleQL` kept for backwards compatibility (wraps `Eval`)
- [ ] `go test ./...` passes

**Dependencies:** E1-S2
**Estimate:** XS (1–2h)

---

## E2 — Array Aggregation Functions

**Goal:** Add `sum`, `avg`, `min`, `max`, `count` as built-in functions operating on `[#].field` pluck results. This is the primary IoT-differentiating feature of v2.

**Acceptance Criteria:**
- All five functions registered in `DefaultValue` context
- Functions accept a `JSONNode` array argument (output of `arr[#].field`)
- Functions return correct typed `Node` or `UNDEFINED_RESULT` for empty/undefined input
- Unit tests cover: integer array, float array, mixed array, empty array, undefined elements

---

### Story E2-S1: Implement array aggregation helpers

**As a** TDTL developer
**I want to** implement `sumAgg`, `avgAgg`, `minAgg`, `maxAgg`, `countAgg` helper functions
**So that** aggregation logic is testable in isolation before registration

**Acceptance Criteria:**
- [ ] Functions implemented in `context.go` (alongside `absFuncHandle`, `base64FuncHandle`)
- [ ] Each function signature: `func(args []Node) Node`
- [ ] `sum`: sums all `IntNode`/`FloatNode` values in `args[0]` (a `JSONNode` array); returns `IntNode` if all int, `FloatNode` if any float
- [ ] `avg`: returns `FloatNode` mean; `UNDEFINED_RESULT` for empty
- [ ] `min` / `max`: return minimum/maximum numeric value; `UNDEFINED_RESULT` for empty
- [ ] `count`: returns `IntNode` count of non-undefined elements
- [ ] Unit tests for each function with: all-int input, all-float input, mixed, empty array, array with undefined elements

**Dependencies:** None
**Estimate:** M (4–8h)

---

### Story E2-S2: Register aggregation functions in DefaultValue

**As a** TDTL user
**I want to** call `sum(sensors[#].value)` in an expression string
**So that** I can aggregate IoT array data without custom function registration

**Acceptance Criteria:**
- [ ] `sum`, `avg`, `min`, `max`, `count` keys added to `DefaultValue` context's function map
- [ ] `EvalRuleQL` / `Eval` resolves these functions via `MutilContext{DefaultValue, userCtx}`
- [ ] Integration test: `avg(sensors[#].value)` on `{"sensors":[{"value":10},{"value":20},{"value":30}]}` returns `FloatNode(20.0)`
- [ ] Integration test: `sum([#])` on empty array `[]` returns `UNDEFINED_RESULT`

**Dependencies:** E2-S1
**Estimate:** S (2–4h)

---

### Story E2-S3: Regression tests for array aggregation edge cases

**As a** TDTL maintainer
**I want to** have comprehensive tests for aggregation edge cases
**So that** I can refactor safely and prevent regressions

**Acceptance Criteria:**
- [ ] Test: mixed int/float array → `avg` returns `FloatNode`
- [ ] Test: array with one undefined element → `sum` counts only defined elements
- [ ] Test: non-array argument to aggregation → returns `UNDEFINED_RESULT`
- [ ] Test: deeply nested pluck `data.sensors[#].readings[0].value` + `avg`
- [ ] Tests added to `evaluator_regression_test.go`

**Dependencies:** E2-S2
**Estimate:** S (2–4h)

---

## E3 — CASE Expression Fix

**Goal:** CASE/WHEN/THEN/ELSE keywords work without mandatory leading whitespace. `case x when 1 then 'a'` at the start of an input string is valid.

**Acceptance Criteria:**
- `case x when 1 then 'a' else 'b'` parses correctly with no leading whitespace
- Existing CASE tests in `evaluator_regression_test.go` pass (currently require ` case` with space)
- No regression in other keyword handling

---

### Story E3-S1: Fix CASE keyword grammar rule

**As a** TDTL user
**I want to** write `case x when 1 then 'a'` without a leading space
**So that** CASE expressions are ergonomic at the start of an expression string

**Acceptance Criteria:**
- [ ] `TDTL.g4`: `CASE`, `WHEN`, `THEN`, `ELSE` token rules updated to not require leading `STUFF` (whitespace)
- [ ] Grammar regenerated: `parser/` updated
- [ ] `ParseExpr("case 1 when 1 then 'one' else 'other'")` returns non-nil `Expr` with nil error
- [ ] Existing CASE tests updated to remove the leading space workaround
- [ ] `go test ./...` passes

**Dependencies:** None
**Estimate:** S (2–4h)

---

### Story E3-S2: Expand CASE test coverage

**As a** TDTL maintainer
**I want to** have tests covering all CASE/WHEN/THEN/ELSE keyword capitalizations and whitespace variants
**So that** grammar changes don't introduce silent regressions

**Acceptance Criteria:**
- [ ] Tests: `CASE`, `Case`, `case` all parse correctly (case-insensitive)
- [ ] Tests: CASE at start of string (no leading space)
- [ ] Tests: CASE inside a larger expression: `abs(case x when 1 then 0-1 else 1)`
- [ ] Tests: CASE with no ELSE (returns `UNDEFINED_RESULT` for no match)
- [ ] Tests in `evaluator_regression_test.go`

**Dependencies:** E3-S1
**Estimate:** XS (1–2h)

---

## E4 — Error-First API (No Panics)

**Goal:** Replace all `panic` calls in exported and parse-callback code with error returns or graceful UNDEFINED_RESULT. Zero panics under any input.

**Acceptance Criteria:**
- `go test -run TestFuzz` (or equivalent adversarial inputs) produces no panics
- `ExitInteger`, `ExitFloat` return errors on parse failure (requires refactoring ANTLR4 listener pattern)
- `ExitTargetAsElem` panic removed (superseded by E1)
- `Expr()` empty-stack guard in place (already fixed in v1 regression, confirm in v2)

---

### Story E4-S1: Convert ExitInteger/ExitFloat parse panics to errors

**As a** TDTL user
**I want to** receive a descriptive error when an integer or float literal cannot be parsed
**So that** my application never crashes due to malformed expressions

**Acceptance Criteria:**
- [ ] `ExitInteger` in `parse.go`: `strconv.ParseInt` failure sets `l.errors` instead of panicking
- [ ] `ExitFloat` in `parse.go`: `strconv.ParseFloat` failure sets `l.errors` instead of panicking
- [ ] `l.error()` propagates to `Parse`/`ParseExpr`/`ParseFilter` return value
- [ ] Test: `ParseExpr("99999999999999999999999999")` (overflow int64) returns error, not panic

**Dependencies:** None
**Estimate:** XS (1–2h)

---

### Story E4-S2: Adversarial input fuzz test

**As a** TDTL maintainer
**I want to** have a fuzz test that runs random/adversarial expression strings through the parser
**So that** no input causes a panic or undefined behavior

**Acceptance Criteria:**
- [ ] `FuzzParseExpr(f *testing.F)` added to `parser_test.go` or new `fuzz_test.go`
- [ ] Seed corpus includes: empty string, very long string, all operators, nested expressions, unicode strings, binary data
- [ ] `go test -fuzz=FuzzParseExpr -fuzztime=60s` completes without panic
- [ ] `go test -run=FuzzParseExpr` (corpus-only run) passes in CI

**Dependencies:** E4-S1
**Estimate:** S (2–4h)

---

## E5 — Compile API & Goroutine Safety

**Goal:** Add `Compile(str) (CompiledExpr, error)` as the recommended production API. Document and test concurrent evaluation safety.

**Acceptance Criteria:**
- `tdtl.Compile` exported and documented
- `CompiledExpr` is the same `Expr` interface, wrapped in a type alias with godoc noting goroutine safety
- Concurrent evaluation test passes under `-race`

---

### Story E5-S1: Export Compile function

**As a** TDTL user
**I want to** call `tdtl.Compile(exprStr)` once and reuse the result across goroutines
**So that** I avoid repeated ANTLR4 parse overhead in hot paths

**Acceptance Criteria:**
- [ ] `func Compile(exprStr string) (CompiledExpr, error)` added to `tdtl.go`
- [ ] `CompiledExpr` is a type alias `type CompiledExpr = Expr` with godoc: "CompiledExpr is safe for concurrent use across goroutines."
- [ ] `Compile` internally calls `ParseExpr` and returns the result
- [ ] Godoc example showing compile-once, eval-many pattern

**Dependencies:** None
**Estimate:** XS (1–2h)

---

### Story E5-S2: Goroutine-safety test for CompiledExpr

**As a** TDTL maintainer
**I want to** have a test proving compiled expressions are safe to evaluate concurrently
**So that** users can rely on the concurrency guarantee in production

**Acceptance Criteria:**
- [ ] Test spawns 100 goroutines, each evaluating the same `CompiledExpr` against independent `JSONContext` instances
- [ ] `go test -race` passes with zero data races
- [ ] Test covers: arithmetic expr, CASE expr, function call expr
- [ ] Added to `evaluator_regression_test.go` or new `concurrency_test.go`

**Dependencies:** E5-S1
**Estimate:** XS (1–2h)

---

## E6 — Test Coverage & Benchmarks

**Goal:** Achieve ≥ 90% line coverage on core files. Add benchmark suite proving NFR1 (< 10μs eval), NFR2 (< 1ms compile), NFR3 (≤ 5 allocs).

---

### Story E6-S1: Coverage audit and gap-fill

**As a** TDTL maintainer
**I want to** achieve ≥ 90% line coverage on `evaluator.go`, `context*.go`, `collectjs.go`, `parse.go`
**So that** NFR6 is met and regressions are caught quickly

**Acceptance Criteria:**
- [ ] `go test -cover ./...` reports ≥ 90% for each target file
- [ ] New tests added for any uncovered paths discovered during audit
- [ ] Coverage report generated and saved to `_bmad-output/implementation-artifacts/coverage.txt`

**Dependencies:** E1, E2, E3, E4
**Estimate:** M (4–8h)

---

### Story E6-S2: Benchmark suite

**As a** TDTL maintainer
**I want to** have benchmarks for compile time and evaluation time
**So that** performance regressions are caught in CI and NFR1–NFR3 are provably met

**Acceptance Criteria:**
- [ ] `BenchmarkCompile` in `evaluator_test.go` or `benchmark_test.go`: measures single `Compile` call
- [ ] `BenchmarkEvalInt`: measures `Eval` on `CompiledExpr` returning `IntNode` (simple arithmetic)
- [ ] `BenchmarkEvalJSONPath`: measures `Eval` on deep JSON path expression
- [ ] `BenchmarkEvalAvgAgg`: measures `Eval` on `avg(sensors[#].value)` with 10-element array
- [ ] `go test -bench=. -benchmem` output shows < 10μs/op for Eval benchmarks, ≤ 5 allocs/op for primitive results
- [ ] Benchmark results committed to `_bmad-output/implementation-artifacts/benchmarks.txt`

**Dependencies:** E5
**Estimate:** S (2–4h)

---

### Story E6-S3: Multi-platform CI validation

**As a** TDTL maintainer
**I want to** validate that tests pass on `linux/amd64`, `linux/arm64`, and `darwin/arm64`
**So that** NFR8 (portability) and NFR10 (no CGO) are continuously verified

**Acceptance Criteria:**
- [ ] GitHub Actions workflow updated (or created) with matrix: `[linux/amd64, linux/arm64, darwin/arm64]`
- [ ] Each matrix job runs `CGO_ENABLED=0 go test ./...`
- [ ] `linux/arm64` tested via QEMU emulation if native runner unavailable
- [ ] Workflow passes on all three targets

**Dependencies:** E6-S1
**Estimate:** S (2–4h)

---

## E7 — Documentation & Migration Guide

**Goal:** README accurately reflects v2 API. `v1 → v2` migration guide covers all breaking changes. GoDoc examples run as tests.

---

### Story E7-S1: Update README for v2 API

**As a** developer evaluating TDTL
**I want to** read a README that shows the v2 `Compile`/`Eval` pattern with realistic IoT examples
**So that** I can integrate TDTL in < 30 minutes

**Acceptance Criteria:**
- [ ] Quick Start updated: shows `Compile` + `Eval` pattern, not `NewTDTL`
- [ ] New section: "Array Aggregation" with `avg(sensors[#].value)` example
- [ ] New section: "Multi-Entity Context" with `MutilContext` example
- [ ] CASE expression example updated (no leading space required)
- [ ] Module import path updated to `github.com/tkeel-io/tdtl/v2`
- [ ] Architecture section updated to reflect three-layer design

**Dependencies:** E1, E2, E3, E5
**Estimate:** S (2–4h)

---

### Story E7-S2: v1 → v2 migration guide

**As an** existing TDTL v1 user
**I want to** read a migration guide covering all breaking changes
**So that** I can upgrade without trial-and-error

**Acceptance Criteria:**
- [ ] `MIGRATION.md` created with sections:
  - Module path change (`tdtl` → `tdtl/v2`)
  - `INSERT INTO ... SELECT` removal: what to do instead (routing in caller)
  - `EvalRuleQL` deprecation: replacement with `Eval`
  - Panic → error change: impact on existing `recover()` blocks
  - New `Compile` API: recommended usage
  - `NewTDTL` deprecation path
- [ ] Each breaking change has a before/after code example

**Dependencies:** E1, E5
**Estimate:** S (2–4h)

---

### Story E7-S3: GoDoc examples as runnable tests

**As a** TDTL developer
**I want to** have `Example*` functions in `_test.go` files for key API surfaces
**So that** documentation examples are tested and cannot drift from real behavior

**Acceptance Criteria:**
- [ ] `ExampleCompile()` in `tdtl_test.go` with `// Output:` assertion
- [ ] `ExampleEval_arithmetic()` with simple math expression
- [ ] `ExampleEval_jsonPath()` with JSON path access
- [ ] `ExampleEvalFilter()` with filter expression
- [ ] `ExampleMutilContext()` with multi-entity lookup
- [ ] All examples follow Go 1.24 naming rules (no `Example_UpperSuffix` pattern)
- [ ] `go test ./...` runs and passes all examples

**Dependencies:** E1, E5, E7-S1
**Estimate:** S (2–4h)

---

## Sprint Suggestion

### Sprint 1 — Foundation (Epics E1, E3, E4)
Remove routing, fix CASE, eliminate panics. Establishes clean v2 core.

| Story | Epic | Estimate |
|-------|------|----------|
| E1-S1 | E1 | S |
| E1-S2 | E1 | S |
| E1-S3 | E1 | XS |
| E3-S1 | E3 | S |
| E3-S2 | E3 | XS |
| E4-S1 | E4 | XS |
| E4-S2 | E4 | S |

### Sprint 2 — Core Differentiator (Epic E2, E5)
Array aggregations + Compile API. Delivers the IoT-differentiating features.

| Story | Epic | Estimate |
|-------|------|----------|
| E2-S1 | E2 | M |
| E2-S2 | E2 | S |
| E2-S3 | E2 | S |
| E5-S1 | E5 | XS |
| E5-S2 | E5 | XS |

### Sprint 3 — Quality & Launch (Epics E6, E7)
Coverage, benchmarks, docs. Ships v2.0.

| Story | Epic | Estimate |
|-------|------|----------|
| E6-S1 | E6 | M |
| E6-S2 | E6 | S |
| E6-S3 | E6 | S |
| E7-S1 | E7 | S |
| E7-S2 | E7 | S |
| E7-S3 | E7 | S |
