---
format: aep.planning-md/3
id: epic:ignite-codespec-bridge
kind: epic
status: draft
title: Ignite and codespec, both directions
revision: 2
---
# Epic: ignite-codespec-bridge

## Outcome

Code moves between Ignite and every language codespec reads or writes, in both directions, through `cir/1`, with adapters written in Ignite and every translation accepted by codespec's comparison harness only.

## Context

- Requested by Timo, 2026-11-06: "we need 1.) codespec -> ignite and 2.) ignite -> codespec", in a new public repository, written in Ignite.
- Design and mapping: `docs/design.md`.
- Milestone C2 lives in ignite-lang/ignite as `story:typed-program-json`; this store cannot hold an edge to it.
- External prerequisite: codespec's harness and writers are on its branch `m1-writers-harness` (3b2f6eb), not yet on its main.

## Acceptance

- Milestones C0 to C4 of `docs/design.md` section 6 are done.
