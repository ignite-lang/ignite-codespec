---
format: aep.planning-md/3
id: story:ignite-reader
kind: story
status: active
title: 'codespec-read-ignite: Ignite to cir/1'
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:cir-model
- depends_on: story:ledger-example
- depends_on: story:ignite-writer
scope:
- confidence: cited
  path: Taskfile.yml
- confidence: inferred
  path: src/read.ignite
revision: 10
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T01:42:46Z", actor: "human:timo", revision: 9, decided_on: {"recorded":{"review_outcome":6}}}
- {from: "proposed", to: "active", at: "2026-11-06T04:00:44Z", actor: "human:timo", revision: 10, decided_on: {"recorded":{"review_outcome":6}}}
---
# Story: ignite-reader

## Outcome

C3: `codespec-read-ignite`, an Ignite program that runs `ignite emit --typed-json FILE` through `Launcher` and prints one `cir/1` module, Ignite-only constructs as `foreign` plus a refusal.

## Context

- External dependency, not expressible as an edge: ignite-lang/ignite `story:typed-program-json` (milestone C2) provides `ignite emit --typed-json`. This story starts when that story is merged in a released or pinned ignite.
- Example: `examples/ledger/reference.ignite` from `story:ledger-example`. Its behaviour through `cir/1` is judged by `story:ignite-round-trip`; this story judges the document. Depends on `story:ignite-writer` for the shared `check` task of `Taskfile.yml`.

## Acceptance

- `codespec read --lang ignite examples/ledger/reference.ignite` passes `codespec validate --canonical`, and reading it twice gives the same bytes.
- Inline tests in `src/read.ignite` map one Ignite construct of each row of `docs/design.md` section 3 to its `cir/1` node, and an Ignite-only construct to `foreign` plus a refusal with its position.
- `task check` runs both and is green.

## Out of Scope

- The writer and the round trip.
- Any other language: codespec's own writers turn the document into Go or Rust.
