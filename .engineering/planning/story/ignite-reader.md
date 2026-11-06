---
format: aep.planning-md/3
id: story:ignite-reader
kind: story
status: draft
title: 'codespec-read-ignite: Ignite to cir/1'
relations:
- decomposes: epic:ignite-codespec-bridge
- depends_on: story:cir-model
scope:
- confidence: inferred
  path: src/read.ignite
revision: 2
---
# Story: ignite-reader

## Outcome

C3: an Ignite program that runs `ignite emit --typed-json` on a file through `Launcher` and prints one `cir/1` module, Ignite-only constructs as `foreign` plus a refusal.

## Context

- `docs/design.md` (2026-11-06).

## Acceptance

- An Ignite example read to `cir/1` passes `codespec validate --canonical`, and codespec's Go and Rust writers on it give programs the harness finds byte-equal to the Ignite original.
- Reading the same file twice gives the same bytes.
- `task check` green.

## Out of Scope

- The other milestones of `docs/design.md` section 6.
