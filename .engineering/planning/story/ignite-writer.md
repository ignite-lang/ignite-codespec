---
format: aep.planning-md/3
id: story:ignite-writer
kind: story
status: proposed
title: 'codespec-write-ignite: cir/1 to Ignite'
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:cir-model
- depends_on: story:ledger-example
scope:
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: examples/ledger/harness.toml
- confidence: cited
  path: examples/orders-refusals.expected.json
- confidence: inferred
  path: src/write.ignite
revision: 8
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T01:42:46Z", actor: "human:timo", revision: 8, decided_on: {"recorded":{"review_outcome":2}}}
---
# Story: ignite-writer

## Outcome

C1: `codespec-write-ignite`, an Ignite program that reads one `cir/1` module (with the facts of `codespec pass`) on stdin and prints Ignite source, mapping each node per `docs/design.md` section 3 or refusing it with position and reason; refusals go to stderr as one JSON line each and into a `refusals.json` beside the output. `examples/ledger/harness.toml` gains the entry `ignite` (the writer's output under the Ignite driver).

## Context

- Pipeline: `codespec read --lang go` → `codespec pass` (facts `shared`, `fallible`, `comma_ok`, `messages`, `unordered_iteration`) → `codespec write --lang ignite`. codespec's Rust writer refuses a module without facts (3b2f6eb).
- Example, drivers and the first entries: `story:ledger-example`. `story:ignite-reader` edits the same manifest and `Taskfile.yml` after this story.

## Acceptance

- `examples/ledger`: the writer's output passes `ignite check` with no diagnostic, and `task harness` over the entries `original`, `go`, `rust`, `reference`, `ignite` reports byte-equal output on all 200 seeds.
- `examples/orders` of the pinned codespec commit: `refusals.json` holds exactly the refusals for the map of pointers (`orders.go:53`), the channel (`orders.go:58`) and the goroutine (`orders.go:132`), checked against a checked-in `examples/orders-refusals.expected.json`; any other construct of `orders.go` is either written or added to that expected file with its reason in the same commit.
- `task check` runs both and is green.

## Out of Scope

- The reader and the round trip.
