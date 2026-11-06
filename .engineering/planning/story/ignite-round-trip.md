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
revision: 1
---
# Story: ignite-round-trip

## Outcome

C4: an Ignite example read by `codespec-read-ignite` and written back by `codespec-write-ignite` behaves the same.

## Context

- `docs/design.md` (2026-11-06).

## Acceptance

- The harness finds the original and the round-tripped program byte-equal on the seeded traces.
- `task check` green.

## Out of Scope

- The other milestones of `docs/design.md` section 6.
