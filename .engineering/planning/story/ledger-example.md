---
format: aep.planning-md/3
id: story:ledger-example
kind: story
status: draft
title: 'examples/ledger: a channel-free Go example with Go and Ignite harness drivers'
relations:
- decomposes: epic:ignite-codespec-bridge
scope:
- confidence: cited
  path: AGENTS.md
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: examples/ledger
revision: 4
---
# Story: ledger-example

## Outcome

`examples/ledger/` holds the first example every direction is judged on: a Go package with no channel, no goroutine and no shared pointer (structs, slices, a map of values, loops, `switch`, integer arithmetic, `(T, error)` results), its Go driver and its Ignite driver, and a `harness.toml` naming the implementations `original` (Go) and `ignite` (the writer's output), so `codespec harness examples/ledger` can run.

## Context

- The harness is codespec's (`codespec harness DIR --seeds N --commands M --build-dir DIR`), on its branch `m1-writers-harness` at 3b2f6eb: an example directory with `source/`, `drivers/<impl>/` and a `harness.toml` of `[[implementation]]` entries with `build` and `run` commands; a driver reads one trace on stdin and prints one line per command result (`examples/orders/drivers/go/main.go`, trace format in codespec `crates/harness`).
- codespec's only example, `orders`, uses a map of pointers, a channel of `any` and a goroutine (`orders.go:53`, `:58`, `:132`); codespec's own algorithmic example (its epic `algorithmic-example`, M3) is an empty draft, so this repository owns its first example.
- The Go source and its Go driver are test inputs; `AGENTS.md` allows Go under `examples/*/source/` and `examples/*/drivers/go/` only.

## Acceptance

- `codespec harness examples/ledger --seeds 200 --commands 300` with only the `original` implementation and the Ignite driver run against a hand-written `examples/ledger/reference.ignite` exits 0 and reports byte-equal output on all 200 seeds.
- `codespec read --lang go examples/ledger/source` passes `codespec validate --canonical` and reports no refusals.
- `task harness` runs the first bullet; `task check` runs it.

## Out of Scope

- The writer, reader and round trip; this story supplies their input and their harness entry.
