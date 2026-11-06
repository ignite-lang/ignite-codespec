---
format: aep.planning-md/3
id: story:ignite-pin-v0-9
kind: story
status: implemented
title: Adapters on Ignite v0.9.0 without copied code
revision: 5
transitions:
- {from: "draft", to: "proposed", at: "2026-11-06T15:02:26Z", actor: "human:timo", revision: 2}
- {from: "proposed", to: "active", at: "2026-11-06T15:02:26Z", actor: "human:timo", revision: 3}
- {from: "active", to: "implemented", at: "2026-11-06T15:24:50Z", actor: "human:timo", revision: 5, decided_on: {"recorded":{"test_result":1}}}
---
# Story: ignite-pin-v0-9

## Outcome

The adapters build and test with the released Ignite v0.9.0 (tag `v0.9.0`, commit `2658ba053bf3fb493052321362825d5750ba2190`) instead of the unreleased commit `f64d0f12b`, and the code that `src/write.ignite` copied from other modules because sibling imports did not resolve is imported from its one home instead.

## Context

- `ignite.pin` holds `f64d0f12b681b08c363c6b6ba833d4c4b01902a3`, the first commit with `ignite emit --typed-json`.
- Sibling module imports across `src/` work since ignite PR 61 (`story:codespec-friction`), released in v0.8.0.
- The reader's inline tests hold `ignite emit --typed-json` documents; AGENTS.md says to regenerate them when the pin moves.

## Acceptance

## Acceptance

- `ignite.pin` is `2658ba053bf3fb493052321362825d5750ba2190`, and AGENTS.md names the tag `v0.9.0` beside it.
- No function under `src/` is a copy of one in another `src/` file: `read`, `has`, `cut`, `canonical` and `pretty` have one definition each and are imported elsewhere. Functions of the same name with different bodies in the two adapters (`stmt`, `ty`, `val`, `block`, …) are separate code and stay (revised 2026-11-06: the first wording counted names, which these share by design).
- The reader's typed-json test documents equal what v0.9.0 prints for their sources.
- `task check` with `IGNITE` set to a v0.9.0 build is green, `task harness` included (200 seeds byte-equal on every entry).
- Grown during the unit: v0.9.0's built-in `Int.wrapping_add`/`wrapping_sub` (E041 against the writer's own helper) means the writer emits `a.wrapping_add(b)` and the reader reads it as `+`.

## Out of Scope

- New `cir/1` nodes, new examples, behaviour changes of either adapter.
