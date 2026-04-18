# TDTL — tKeel Data Transformation Language

TDTL is a query and data-transformation DSL for the [tKeel](https://github.com/tkeel-io) IoT platform. It lets you select, filter, and reshape JSON data from one or more entities using an SQL-like syntax, then write the result to a target entity.

## Features

- SQL-like `INSERT INTO … SELECT … AS` syntax for field projection and aliasing
- Arithmetic, comparison, and logical operators with automatic type coercion
- `CASE / WHEN / THEN / ELSE` conditional expressions
- JSON path access with dot notation and array indexing (`a.b[0]`, `a[#]`)
- Built-in functions: `abs()`, `base64()`
- Extensible function registry via `ContextFunc`
- Composable multi-source context (`MutilContext`)

## Quick Start

```go
import "github.com/tkeel-io/tdtl"

// Parse a TDTL statement once
t, err := tdtl.NewTDTL(
    `insert into target select entity1.temperature + 1 as temp`,
    nil,
)

// Execute against incoming data
result, err := t.Exec(map[string]tdtl.Node{
    "entity1.temperature": tdtl.IntNode(50),
})
// result["temp"] == tdtl.IntNode(51)
```

## Syntax

### Statement structure

```
INSERT INTO <target> SELECT <field_list>
```

Every TDTL statement must start with `INSERT INTO` followed by the target entity name, then `SELECT` followed by one or more field expressions separated by commas.

### Field expressions

```sql
insert into target select
    entity1.temperature + 1         as temp,
    entity1.color                   as metadata.color,
    entity1.temperature + '°F'      as label
```

A field expression without an `AS` alias is evaluated but not written to the output.

### Operators

| Category    | Operators                          |
|-------------|------------------------------------|
| Arithmetic  | `+` `-` `*` `/` `%`               |
| Comparison  | `=` `!=` `<>` `<` `<=` `>` `>=`  |
| Logical     | `AND` `OR` `NOT`                   |
| Grouping    | `( … )`                            |

String `+` concatenates; numeric strings are coerced automatically.

### JSON path syntax

| Expression  | Meaning                        |
|-------------|-------------------------------|
| `a.b`       | nested object field           |
| `a[0]`      | array element at index 0      |
| `a[#]`      | array length                  |
| `a[#].b`    | pluck field `b` from every element |

### CASE expression

Keywords (`CASE`, `WHEN`, `THEN`, `ELSE`) require surrounding whitespace.

```sql
 case color
 when 'red'   then 'stop'
 when 'green' then 'go'
 else 'wait'
```

### Filter (WHERE-style evaluation)

Use `ParseFilter` + `EvalFilter` to evaluate boolean filter expressions:

```go
expr, _ := tdtl.ParseFilter(`temperature > 30 and color = 'red'`)
ok := tdtl.EvalFilter(ctx, expr)
```

### Built-in functions

| Function     | Description                         |
|--------------|-------------------------------------|
| `abs(x)`     | Absolute value of a number          |
| `base64(x)`  | Base-64 encode a string or JSON value |

Register custom functions through `NewMapContext` or the `extFunc` parameter of `NewTDTL`.

## API

### Parsing

```go
// Full INSERT INTO … SELECT statement
expr, err := tdtl.Parse(sql)

// Bare expression (for embedding in larger systems)
expr, err := tdtl.ParseExpr(exprStr)

// Filter / WHERE expression
expr, err := tdtl.ParseFilter(filterStr)
```

### Evaluation

```go
// Evaluate a full statement, returning a JSON object of aliased fields
result := tdtl.EvalRuleQL(ctx, expr)

// Evaluate a filter expression, returning a bool
pass := tdtl.EvalFilter(ctx, expr)
```

### Context

```go
// JSON string as data source
ctx := tdtl.NewJSONContext(`{"temperature":50,"color":"red"}`)

// In-memory map as data source
ctx := tdtl.NewMapContext(
    map[string]tdtl.Node{"temperature": tdtl.IntNode(50)},
    nil,
)

// Combine multiple sources (checked in order)
ctx := tdtl.MutilContext{ctx1, ctx2}
```

### JSON collection utilities

`tdtl.New` wraps raw JSON and exposes a fluent manipulation API:

```go
c := tdtl.New(`{"a":1}`)
c.Set("b", tdtl.IntNode(2))           // {"a":1,"b":2}
c.Append("tags", tdtl.NewString("x")) // {"a":1,"b":2,"tags":["x"]}
c.Del("a")                            // {"b":2,"tags":["x"]}
sub := c.Get("tags[0]")               // "x"

arr := tdtl.New(`[{"k":"a"},{"k":"b"}]`)
grouped := arr.GroupBy("k")           // {"a":[…],"b":[…]}
keyed   := arr.KeyBy("k")            // {"a":{…},"b":{…}}
merged  := arr.MergeBy("k")          // {…merged by k…}
```

## Architecture

```
SQL string
  │
  ▼
ANTLR4 Lexer / Parser  (parser/)
  │
  ▼
TDTLListener  →  AST (Expr tree)   (parse.go, types.go)
  │
  ▼
Evaluator  →  Node                 (evaluator.go)
  │
  ▼
Context  →  data lookup            (context.go, context_map.go, context_json.go)
```

Node types: `BoolNode`, `IntNode`, `FloatNode`, `StringNode`, `JSONNode`.

## Development

Regenerate the parser after editing `TDTL.g4`:

```sh
antlr4 -Dlanguage=Go -o parser TDTL.g4
```

Run tests:

```sh
go test ./...
```

## License

Apache 2.0 — see [LICENSE](LICENSE).
