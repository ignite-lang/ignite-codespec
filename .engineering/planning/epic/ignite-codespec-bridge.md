---
format: aep.planning-md/3
id: epic:ignite-codespec-bridge
kind: epic
status: implemented
title: Ignite and codespec, both directions
revision: 7
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T05:09:24Z", actor: "human:timo", revision: 5}
- {from: "proposed", to: "active", at: "2026-11-06T05:09:24Z", actor: "human:timo", revision: 6}
- {from: "active", to: "implemented", at: "2026-11-06T05:09:24Z", actor: "human:timo", revision: 7}
---
# Epic: ignite-codespec-bridge

## Outcome

Ignite becomes a language of codespec: Ignite is read into `cir/1` and `cir/1` is written as Ignite, so codespec moves code between Ignite and every language it has adapters for, with adapters written in Ignite and every translation accepted by codespec's comparison harness only.

## Context

- Requested by Timo, 2026-11-06: "we need 1.) codespec -> ignite and 2.) ignite -> codespec", in a new public repository, written in Ignite.
- Design and mapping: `docs/design.md`.
- Milestone C2 lives in ignite-lang/ignite as `story:typed-program-json`; this store cannot hold an edge to it.
- External prerequisite: codespec's harness and writers are on its branch `m1-writers-harness` (harness b73ffce, Rust writer 3b2f6eb), not yet on its main.

## Acceptance

- Milestones C0 to C4 of `docs/design.md` section 6 are done.
