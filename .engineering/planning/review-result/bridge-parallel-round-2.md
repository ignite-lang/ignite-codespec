---
format: aep.planning-md/3
id: review-result:bridge-parallel-round-2
kind: review-result
status: active
title: parallel critic, round 2
relations:
- reviews: epic:ignite-codespec-bridge
- reviews: story:cir-model
- reviews: story:ledger-example
- reviews: story:ignite-writer
- reviews: story:ignite-reader
- reviews: story:ignite-round-trip
revision: 1
---
needs-revision

- story:ledger-example — it and story:cir-model both land on `Taskfile.yml` (cited, both scopes: each acceptance has `task check` run its own tasks), no `depends_on` edge orders them, and neither body names the other, so the store's own `waves` prints `collision: story:cir-model story:ledger-example Taskfile.yml`. Fix, without choosing: add an ordering edge whose stated reason is the shared `check` task, or split the surface so only one story edits `Taskfile.yml`. — .engineering/planning/story/ledger-example.md:7-10 and :32, .engineering/planning/story/cir-model.md:6-10 and :27
- story:ignite-reader — its acceptance, "`task check` runs both and is green", needs an edit to the single `check` task in `Taskfile.yml`, and so does story:ignite-writer's identical line. Neither scope lists that file, the two have no edge between them (both sit in wave 3), and neither body names the other. The surface is inferred from the acceptance lines and the one-task `Taskfile.yml`. Fix: add `Taskfile.yml` to both scopes and record an ordering edge or a split. — .engineering/planning/story/ignite-reader.md:32, .engineering/planning/story/ignite-writer.md:30, Taskfile.yml:7-10

**What I read:** 12 artifacts, 6 in the set. I read the round-1 record `review-result:bridge-parallel-round-1` and `show` on the five codespec stories plus typed-program-json and static-linking. I ran `aep plan artifact graph`, `relations` and `waves` in the codespec store, `validate` in both, and `list` in the ignite-lang store. I also read `Taskfile.yml` at head 4d1641b. Heads confirmed: codespec 4d1641b, ignite-lang 0fa97d555.

**Round-1 findings, one by one**
- **typed-program-json against static-linking: fixed.**
  - An edge `depends_on story:static-linking` now exists.
  - The body names the shared `main.rs`, `lib.rs` and `docs/slice.md` as the reason, and puts the serialiser in a new module, `crates/ignite/src/typed_json.rs` (cited).
  - The `crates/ignite/tests` overlap is covered by the same edge.
- **ignite-reader against typed-program-json: fixed.**
  - Context now says "External dependency, not expressible as an edge" and names `story:typed-program-json` (C2).
  - That is the stated-external-dependency remedy round 1 allowed.
- **ignite-round-trip scope: fixed.** It now declares `Taskfile.yml` and `examples/ledger/harness.toml`, both cited. `waves` reports 0 unassessed.

**Numbers for part 3**
- **Surfaces:** 4 cited (cir-model, ledger-example, round-trip, typed-program-json), 2 inferred (writer, reader), 0 unplaceable. Writer also has one cited file, `examples/orders-refusals.expected.json`.
- **Acceptance reads traced to a producer in the set: 9.**
  - 7 are direct edges: writer to cir-model, writer to ledger-example, reader to cir-model, reader to ledger-example, round-trip to writer, round-trip to reader, typed-program-json to static-linking.
  - 1 is covered transitively: round-trip's `harness.toml` read, produced by ledger-example, via the writer.
  - 1 is an external prose dependency, reader to typed-program-json, which the store cannot hold as an edge.

**What I could not establish**
- Whether the store accepts a cross-store id. I did not try and I wrote nothing.
- The codespec harness (`codespec harness`, branch `m1-writers-harness` at 3b2f6eb) is outside the set, so I did not check that it exists as the acceptances assume.
- Out of lane (design/acceptance): ledger-example's Outcome does not list `examples/ledger/reference.ignite`, though its acceptance and the reader's Context treat it as a deliverable.
- Out of lane (design/acceptance): ledger-example reads codespec at a commit it cites itself, while cir-model creates `codespec.pin`. Nothing ties the two, so they could skew.
- Out of lane (scope): typed-program-json's `crates/ignite-types/src` is also claimed by `story:stable-node-ids` (draft, outside the set). Its body says it does not need that story.
- `story:llvm-lowering-copies` has no scope in the ignite-lang store. It is outside the set, so I judged nothing about it.

```findings
- file: .engineering/planning/story/ledger-example.md
  line: 7
  category: parallel-safety
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "this story and story:cir-model both land on Taskfile.yml (cited in both scopes; each acceptance has task check run its own tasks), no depends_on edge orders them and neither body names the other, so add an ordering edge recording the shared check task or split the surface, without choosing"
- file: .engineering/planning/story/ignite-reader.md
  line: 32
  category: parallel-safety
  severity: warning
  verdict: needs-revision
  origin: undecided
  message: "the acceptance 'task check runs both and is green' needs an edit to the single check task in Taskfile.yml, as does story:ignite-writer's identical line, but neither scope lists the file and the two share no edge and are in one wave; the surface is inferred from the acceptance lines, so list Taskfile.yml in both and record an ordering edge or a split, without choosing"
```
