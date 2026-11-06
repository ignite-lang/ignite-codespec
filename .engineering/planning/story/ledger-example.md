---
format: aep.planning-md/3
id: story:ledger-example
kind: story
status: draft
title: 'examples/ledger: a channel-free Go example with Go and Ignite harness drivers'
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:cir-model
scope:
- confidence: cited
  path: AGENTS.md
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: examples/ledger
revision: 5
---
# Story: ledger-example

## Outcome

`examples/ledger/` holds the example every direction is judged on: a Go package that implements codespec's orders interface (`Place`, `Pay`, `Ship`, `Cancel`, `Get`, `PendingOrders`) with no channel, no goroutine and no shared pointer (a map of values, slices, loops, `switch`, integer arithmetic, `(T, error)` results; events appended to a slice the driver drains), its drivers `drivers/go`, `drivers/rust` and `drivers/ignite`, the hand-written `examples/ledger/reference.ignite`, and a `harness.toml` with the entries `original` (the Go source), `go` and `rust` (codespec's Go and Rust writers on the Go source, as codespec's orders example has them) and `reference` (the Ignite driver over `reference.ignite`).

## Context

- The harness is codespec's `codespec harness DIR --seeds N --commands M --build-dir DIR` (commit b73ffce on `m1-writers-harness`; the Rust writer is 3b2f6eb under it). It has one trace generator and one output format, both orders-shaped (`crates/harness/src/lib.rs:20-56`), and no implementation filter: it builds and compares every `[[implementation]]` (`:221-245`). So the ledger speaks that trace and that output, and needs no codespec change; its drivers print the events from the slice in the order the channel would have delivered them.
- codespec is read at the commit in `codespec.pin` (from `story:cir-model`); this story depends on it for that file and for the shared `check` task in `Taskfile.yml`.
- `harness.toml` grows by story, and every acceptance runs the whole manifest as it stands: this story writes `original`, `go`, `rust`, `reference`; `story:ignite-writer` adds `ignite`; `story:ignite-reader` adds `go-from-ignite` and `rust-from-ignite`; `story:ignite-round-trip` adds `roundtrip`.
- codespec's only example, `orders`, uses a map of pointers, a channel of `any` and a goroutine (`orders.go:53`, `:58`, `:132`); its algorithmic example (epic `algorithmic-example`, M3) is an empty draft.
- The Go source, the Go driver and the Rust driver are test inputs; `AGENTS.md` allows Go under `examples/*/source/` and `examples/*/drivers/go/` and Rust under `examples/*/drivers/rust/` only.

## Acceptance

- `task harness` runs `codespec harness examples/ledger --seeds 200 --commands 300` over the four entries `original`, `go`, `rust`, `reference`; it exits 0 with byte-equal output on all 200 seeds.
- `codespec read --lang go examples/ledger/source` passes `codespec validate --canonical` and reports no refusals.
- `task check` runs `task harness` and is green.

## Out of Scope

- The writer, reader and round trip and their harness entries.
