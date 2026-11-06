---
format: aep.planning-md/3
id: story:cir-model
kind: story
status: draft
title: The cir/1 model in Ignite, generated from codespec's schema
relations:
- decomposes: epic:ignite-codespec-bridge
scope:
- confidence: inferred
  path: src/cir.ignite
- confidence: inferred
  path: tools/gen_model.ignite
revision: 3
---
# Story: cir-model

## Outcome

C0: `tools/gen_model.ignite` generates `src/cir.ignite` from codespec's `schemas/cir-1.schema.json` at a pinned commit; strict reading and canonical writing of `cir/1` documents.

## Context

- `docs/design.md` (2026-11-06).

## Acceptance

- Every checked-in codespec example document is read and re-encoded to the same bytes, on both Ignite targets.
- A document with an unknown field, an unknown kind or an expression without a type is refused with its JSON path.
- `task check` green.

## Out of Scope

- The other milestones of `docs/design.md` section 6.
