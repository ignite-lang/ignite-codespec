---
format: aep.planning-md/3
id: story:ignite-round-trip
kind: story
status: implemented
title: Ignite to cir/1 to Ignite
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:ignite-writer
- depends_on: story:ignite-reader
scope:
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: examples/ledger
- confidence: cited
  path: examples/ledger/harness.toml
- confidence: cited
  path: src/write.ignite
revision: 11
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T01:42:46Z", actor: "human:timo", revision: 6, decided_on: {"recorded":{"review_outcome":2}}}
- {from: "proposed", to: "active", at: "2026-11-06T04:39:51Z", actor: "human:timo", revision: 10, decided_on: {"recorded":{"review_outcome":2}}}
- {from: "active", to: "implemented", at: "2026-11-06T05:09:21Z", actor: "human:timo", revision: 11, decided_on: {"recorded":{"test_result":1,"review_outcome":2}}}
---
# Story: ignite-round-trip

## Outcome

C4: Ignite read by `codespec-read-ignite` and written back by `codespec-write-ignite` behaves the same as the original, on the Ignite that the writer produces: `examples/ledger/out/ignite/ledger.ignite` → cir/1 → Ignite is `examples/ledger/out/roundtrip/ledger.ignite`, and `examples/ledger/harness.toml` gains the entry `roundtrip`.

## Context

- `docs/design.md` section 6, C4.
- Input changed on 2026-11-06 after the reader landed: `examples/ledger/reference.ignite` uses a generic `Step<T>`, a union taken apart by `match` and a `Result` with a non-enum error. cir/1 has no node for union values, `match` on data, ordered map iteration or checked overflow, so the reader refuses 9 of its declarations (`docs/design.md` "As read (C3)"). The writer's output is cir/1-expressible Ignite by construction, so it is the round trip's input. The 9 reference refusals stay recorded in `examples/ledger/reference-refusals.expected.json`; the missing cir/1 nodes are a request to codespec, not a change here.
- Known writer defect to fix first: a Go `switch` without `default` becomes an if-chain with no final `else`, so a function such as `label` does not end in a value (`ledger.ignite:59:10: error[E040]` on the reader's output written back).
- Conformance of reader and writer against other languages is tested on the codespec side (Timo, 2026-11-06).

## Acceptance

- `task roundtrip` writes `examples/ledger/out/roundtrip/ledger.ignite`; `ignite check` is silent on it, and `task harness` over every entry of the manifest, `roundtrip` included, reports byte-equal output on all 200 seeds.
- Fixed point: reading `out/roundtrip/ledger.ignite` and writing it again gives the same bytes.
- `codespec read --lang ignite examples/ledger/reference.ignite` refusals equal `examples/ledger/reference-refusals.expected.json`.
- The writer writes a `switch` without `default` so the function ends in a value, with an inline test.
- `task check` runs all of it and is green.

## Out of Scope

- Examples other than `examples/ledger`.
- New cir/1 nodes.
