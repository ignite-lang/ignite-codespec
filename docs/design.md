# ignite-codespec — design (2026-11-06)

## Why

codespec (`github.com/beyond10x/codespec`) defines `cir/1`, a typed code intermediate
representation, with readers that lift a language into it, passes that rewrite it and writers that
emit a language from it. Every translation is checked by running the original and the output on the
same seeded inputs and comparing the outputs byte for byte.

This repository makes Ignite one of codespec's languages, in both directions, and is written in Ignite. It knows `cir/1` and Ignite and no other language; moving code between Ignite and Go, Rust or any other language is codespec's, through that language's adapters:

```
cir/1 ──codespec-write-ignite──▶ Ignite      (1, codespec → Ignite)
Ignite ──codespec-read-ignite──▶ cir/1       (2, Ignite → codespec)
```

## 1. Adapters

codespec runs an adapter as a program named `codespec-read-<lang>` or `codespec-write-<lang>` from
its `--adapters` directory (`codespec/crates/cli/src/main.rs`). A reader prints one `cir/1` module
on stdout; a writer reads one on stdin. `codespec read` and `codespec write` validate every document
that crosses that boundary, so this repository cannot weaken the format.

| program | built from | direction |
|---|---|---|
| `codespec-write-ignite` | `src/write.ignite` | `cir/1` on stdin → Ignite source on stdout |
| `codespec-read-ignite` | `src/read.ignite` | Ignite source file → `cir/1` on stdout |

Both are Ignite executables built by `ignite build`, statically linked once `ignite build --static`
lands (ignite `story:static-linking`), so the adapters directory needs no runtime.

## 2. The model in Ignite

`src/cir.ignite` holds the `cir/1` types as Ignite types, generated from codespec's
`schemas/cir-1.schema.json` at a pinned codespec commit by `tools/gen_model.ignite` (an Ignite
program). It is not written by hand: the schema is generated from codespec's Rust model, so the
Ignite model follows the one definition. Reading is strict as codespec's is: an unknown field or
kind, a missing expression type, or a format other than `cir/1` is refused. Writing is canonical:
the same module re-encoded gives the same bytes as codespec's canonical JSON.

## 3. codespec → Ignite (the writer)

`cir/1` is shaped by Go, its first reader. Each node maps to Ignite or is refused with a position
and a reason; nothing is approximated. "Verified" rows cite the Ignite contract (`docs/slice.md` in
`github.com/ignite-lang/ignite`); "inferred" rows are to be confirmed by the first writer test.

| `cir/1` | Ignite | basis |
|---|---|---|
| `int {width: 64, signed: true}`, `int {width: platform}` | `Int` (i64) | verified, slice.md "`Int` → `i64`" |
| other int widths | `Int` with the range checked at every conversion, or refused | inferred |
| `bool`, `string`, `unit` | `Bool`, `Text`, `()` | inferred |
| `slice`, `map`, `tuple` | `List<T>`, `Map<K, V>`, tuple | inferred; Ignite's `Map` iterates in key order, so a loop the `unordered_iteration` fact marks is written over sorted keys and the harness confirms the order does not reach the output |
| `named`, `struct`, `enum`, `union` | `type` records, enums, enums with data | inferred |
| `ptr` to a value the `shared` fact clears | the value (Ignite has value semantics) | inferred |
| `ptr` the `shared` fact marks, `chan`, `send`, `any` | refused | Ignite has no shared mutable reference and no channel |
| `let`, `var` | `let`, `var` | verified, slice.md "`let x = e` / `var x = e`" |
| `assign`, `incr`, `decr` | `x = e`, `x = x + 1`, `x = x - 1` | inferred |
| `if`, `for_range`, `switch`, `return`, `expr` | `if`, `for x in`, `match`, `return`, expression statement | verified for the keywords (slice.md section 3), inferred for the mapping |
| `lookup_else` | `match m.get(k) { None => …, Some(v) => … }` | inferred |
| a function the `fallible` fact marks | returns `Result<T, E>` | inferred |
| a `(T, bool)` the `comma_ok` fact marks | `Option<T>` | inferred |
| `foreign` | refused, with its source text and reason | codespec rule |

## 4. Ignite → codespec (the reader)

An Ignite program cannot parse and type Ignite source itself. The compiler provides it: ignite
`story:typed-program-json` adds `ignite emit --typed-json FILE`, which prints the parsed module with
a type on every expression. `codespec-read-ignite` runs it through `Launcher` (std.process) and maps
the typed JSON to `cir/1`. Ignite constructs with no `cir/1` node are `foreign` statements plus a
`refusals` entry: declarations of systems, components, protocols and contracts; `spawn` and `await`;
capability resources; closures until `cir/1` has a node for them.

## 5. Evidence

A translation is accepted only by codespec's comparison harness (codespec M1): the original and the
output run on the same seeded traces and their outputs are compared byte for byte. The first example
is codespec's `examples/orders` (Go, 137 lines): it uses a map of pointers, a channel of `any` and a
goroutine (`orders.go:53`, `:58`, `:132`), so the writer's first result on it is a set of refusals,
and the refusals are the measurement.

## 6. Milestones

| | content | done when |
|---|---|---|
| C0 | repository, generated `cir.ignite`, strict reading and canonical writing | every checked-in codespec example document read and re-encoded gives the same bytes |
| C1 | `codespec-write-ignite` | an example without channels and shared pointers: `ignite check` silent on the output, harness byte-equal against the Go original; `orders`: the refusal list, each with position and reason |
| C2 | ignite `typed-program-json` (ignite repository) | `ignite emit --typed-json` on every std and example file, a type on every expression |
| C3 | `codespec-read-ignite` | an Ignite example → `cir/1` that passes `codespec validate --canonical`, deterministic; behaviour judged by C4 |
| C4 | round trip | Ignite → `cir/1` → Ignite, harness byte-equal |

## 7. Open questions

- The harness interface (codespec M1, branch `m1-writers-harness` at 3b2f6eb): how an
  implementation receives a trace and prints its output; the Ignite output needs the same driver.
- Whether `cir/1` gains nodes for `match` on data, `Option`/`Result` and closures, or Ignite keeps
  them as refusals. A node added is `cir/2` by codespec's rules; that decision is codespec's.
