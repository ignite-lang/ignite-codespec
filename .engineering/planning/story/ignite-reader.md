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
- depends_on: story:ignite-writer
scope:
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: examples/ledger/harness.toml
- confidence: inferred
  path: src/read.ignite
revision: 6
---
# Story: ignite-reader

## Outcome

C3: `codespec-read-ignite`, an Ignite program that runs `ignite emit --typed-json FILE` through `Launcher` and prints one `cir/1` module, Ignite-only constructs as `foreign` plus a refusal. `examples/ledger/harness.toml` gains the entries `go-from-ignite` and `rust-from-ignite`: codespec's Go and Rust writers on the reader's document of `reference.ignite`, under the Go and Rust drivers of `story:ledger-example`.

## Context

- External dependency, not expressible as an edge: ignite-lang/ignite `story:typed-program-json` (milestone C2) provides `ignite emit --typed-json`. This story starts when that story is merged in a released or pinned ignite.
- Example, drivers and reference: `story:ledger-example`. Depends on `story:ignite-writer` because both edit `examples/ledger/harness.toml` and the `check` task of `Taskfile.yml`.

## Acceptance

- `codespec read --lang ignite examples/ledger/reference.ignite` passes `codespec validate --canonical`, and reading it twice gives the same bytes.
- `task harness` over every entry of the manifest, `go-from-ignite` and `rust-from-ignite` included, reports byte-equal output on all 200 seeds.
- `task check` runs both and is green.

## Out of Scope

- The writer and the round trip.
