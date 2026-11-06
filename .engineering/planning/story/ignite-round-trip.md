---
format: aep.planning-md/3
id: story:ignite-round-trip
kind: story
status: draft
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
revision: 4
---
# Story: ignite-round-trip

## Outcome

C4: `examples/ledger/reference.ignite` read by `codespec-read-ignite` and written back by `codespec-write-ignite` behaves the same as the original.

## Context

- `docs/design.md` section 6, C4.

## Acceptance

- `task roundtrip` writes `examples/ledger/out/roundtrip.ignite`, and `codespec harness examples/ledger --seeds 200 --commands 300` with the implementations `reference` (the hand-written Ignite) and `roundtrip` reports byte-equal output on all 200 seeds.
- `task check` runs it and is green.

## Out of Scope

- Examples other than `examples/ledger`.
