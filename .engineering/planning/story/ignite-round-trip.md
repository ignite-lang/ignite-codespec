---
format: aep.planning-md/3
id: story:ignite-round-trip
kind: story
status: proposed
title: Ignite to cir/1 to Ignite
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:ignite-writer
- depends_on: story:ignite-reader
scope:
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: examples/ledger/harness.toml
revision: 6
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T01:42:46Z", actor: "human:timo", revision: 6, decided_on: {"recorded":{"review_outcome":2}}}
---
# Story: ignite-round-trip

## Outcome

C4: `examples/ledger/reference.ignite` read by `codespec-read-ignite` and written back by `codespec-write-ignite` behaves the same as the original; `examples/ledger/harness.toml` gains the entry `roundtrip`.

## Context

- `docs/design.md` section 6, C4. The entry `reference` comes from `story:ledger-example`.

## Acceptance

- `task roundtrip` writes `examples/ledger/out/roundtrip.ignite`, and `task harness` over every entry of the manifest, `roundtrip` included, reports byte-equal output on all 200 seeds.
- `task check` runs it and is green.

## Out of Scope

- Examples other than `examples/ledger`.
