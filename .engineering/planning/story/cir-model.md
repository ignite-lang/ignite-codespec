---
format: aep.planning-md/3
id: story:cir-model
kind: story
status: implemented
title: The cir/1 model in Ignite, generated from codespec's schema
relations:
- decomposes: epic:ignite-codespec-bridge
scope:
- confidence: cited
  path: Taskfile.yml
- confidence: cited
  path: codespec.pin
- confidence: inferred
  path: src/cir.ignite
- confidence: cited
  path: src/cir_json.ignite
- confidence: inferred
  path: tools/gen_model.ignite
revision: 11
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T01:42:46Z", actor: "human:timo", revision: 9, decided_on: {"recorded":{"review_outcome":1}}}
- {from: "proposed", to: "active", at: "2026-11-06T01:42:46Z", actor: "human:timo", revision: 10, decided_on: {"recorded":{"review_outcome":1}}}
- {from: "active", to: "implemented", at: "2026-11-06T02:30:21Z", actor: "human:timo", revision: 11, decided_on: {"recorded":{"test_result":1,"review_outcome":1}}}
---
# Story: cir-model

## Outcome

C0: `tools/gen_model.ignite` generates `src/cir.ignite` from codespec's `schemas/cir-1.schema.json` at the codespec commit pinned in `codespec.pin`; `src/cir_json.ignite` reads `cir/1` documents strictly and writes them canonically.

## Context

- `docs/design.md` section 2.
- The checked-in codespec documents are `examples/orders/orders.read.cir.json` on main and `orders.cir.json` on `m1-writers-harness`; the pin names the commit both are read from.

## Acceptance

- `task model-check` regenerates `src/cir.ignite` from the pinned schema and fails if the file differs from the checked-in one by one byte.
- `task roundtrip-check` reads every `*.cir.json` of the pinned codespec commit's `examples/`, re-encodes it, and fails unless the bytes are equal; on the LLVM target (the default).
- Inline tests in `src/cir_json.ignite` refuse a document with an unknown field, an unknown kind and an expression without a type, each with its JSON path in the error.
- `task check` runs both tasks and the inline tests, and is green.

## Out of Scope

- The writer and the reader.
