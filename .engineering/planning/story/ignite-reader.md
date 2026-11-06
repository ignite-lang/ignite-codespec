---
format: aep.planning-md/3
id: story:ignite-reader
kind: story
status: draft
title: 'codespec-read-ignite: Ignite to cir/1'
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:cir-model
- depends_on: story:ledger-example
scope:
- confidence: inferred
  path: src/read.ignite
revision: 3
---
# Story: ignite-reader

## Outcome

C3: `codespec-read-ignite`, an Ignite program that runs `ignite emit --typed-json FILE` through `Launcher` and prints one `cir/1` module, Ignite-only constructs as `foreign` plus a refusal.

## Context

- External dependency, not expressible as an edge: ignite-lang/ignite `story:typed-program-json` (milestone C2) provides `ignite emit --typed-json`. This story starts when that story is merged in a released or pinned ignite.
- Example: `examples/ledger/reference.ignite` from `story:ledger-example`.

## Acceptance

- `codespec read --lang ignite examples/ledger/reference.ignite` passes `codespec validate --canonical`, and reading it twice gives the same bytes.
- After `codespec pass`, codespec's Go and Rust writers on that document give programs that `codespec harness examples/ledger --seeds 200 --commands 300` finds byte-equal to the `original` implementation, with the Go and Rust drivers of `examples/ledger`.
- `task check` runs both and is green.

## Out of Scope

- The writer and the round trip.
