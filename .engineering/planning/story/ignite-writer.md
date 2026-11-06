---
format: aep.planning-md/3
id: story:ignite-writer
kind: story
status: draft
title: 'codespec-write-ignite: cir/1 to Ignite'
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:cir-model
scope:
- confidence: inferred
  path: src/write.ignite
revision: 2
---
# Story: ignite-writer

## Outcome

C1: an Ignite program that reads one `cir/1` module on stdin and prints Ignite source, mapping each node per `docs/design.md` section 3 or refusing it with position and reason.

## Context

- `docs/design.md` (2026-11-06).

## Acceptance

- An example without channels and shared pointers: `ignite check` is silent on the output, and codespec's harness compares the Go original and the Ignite output byte-equal on its seeded traces.
- `examples/orders`: the output lists every refusal with position and reason; nothing is dropped.
- `task check` green.

## Out of Scope

- The other milestones of `docs/design.md` section 6.
