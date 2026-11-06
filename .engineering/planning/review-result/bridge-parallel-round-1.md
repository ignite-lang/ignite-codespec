---
format: aep.planning-md/3
id: review-result:bridge-parallel-round-1
kind: review-result
status: active
title: parallel critic, round 1
relations:
- reviews: epic:ignite-codespec-bridge
- reviews: story:cir-model
- reviews: story:ignite-writer
- reviews: story:ignite-reader
- reviews: story:ignite-round-trip
revision: 1
---
needs-revision

- story:typed-program-json — it and story:static-linking, which is active, both land on `crates/ignite/src/main.rs`, `crates/ignite/src/lib.rs` and `docs/slice.md` (surfaces cited in both bodies), and neither body names the other; the typed-json story's new `crates/ignite/tests` test file is inferred, and static-linking claims that directory. Fix, without choosing: add an ordering edge or a prose note that records the shared files as its reason, or split the surface. The usual split is a new module for the serialiser, so that `main.rs` and `lib.rs` get only a flag and dispatch line. — .engineering/planning/story/typed-program-json.md:8-16, .engineering/planning/story/static-linking.md:8-22
- story:ignite-reader — its outcome and its first acceptance line need `ignite emit --typed-json`, which story:typed-program-json (the codespec design's C2) creates in another store. The body names neither the story nor C2, and no `depends_on` edge exists. So `aep plan artifact waves` in the codespec store puts the reader in wave 2 beside the writer, as if nothing external gated it. The fix is to name story:typed-program-json in Context and record the prerequisite, as an edge if the store accepts a cross-store id and otherwise as a stated external dependency. — .engineering/planning/story/ignite-reader.md:19,27, .engineering/planning/story/typed-program-json.md:25
- story:ignite-round-trip — it has no scope, and its body names no file, test or Taskfile target where the round-trip check lives. The store lists it as "unassessed", and unassessed does not mean safe. Fix: name the surface (a fixture, a task, a test module) so it can be placed against the writer, the reader and the shared `Taskfile.yml`. — .engineering/planning/story/ignite-round-trip.md:1-12

**What I read:** 6 artifacts. I ran `aep plan artifact list`, `show` on the epic and all five stories, `graph`, `waves` and `validate` in both stores. I also read `docs/design.md`, `Taskfile.yml`, and the cited lines in `crates/ignite/src/main.rs`, `crates/ignite-types/src/lib.rs` and `crates/ignite/tests/main.rs`.

**Numbers for part 3**
- Surfaces established: 2 cited (story:cir-model, story:typed-program-json), 2 inferred (story:ignite-writer and story:ignite-reader, whose bodies point at `docs/design.md`, and only its section 1 table names `src/write.ignite` and `src/read.ignite`), 1 unplaced (story:ignite-round-trip).
- Acceptance reads traced to a producer in the set: 5.
  - writer → cir-model: edge recorded.
  - reader → cir-model: edge recorded.
  - round-trip → writer: edge recorded.
  - round-trip → reader: edge recorded.
  - reader → typed-program-json: no edge (finding 2).
- Edges recorded: 4 of 5.

**What I could not establish**
- The harness that the writer, reader and round-trip acceptances read lives on codespec's branch `m1-writers-harness`, not `main`. It is outside the set, so I could not check whether it exists in the form assumed. If it is wanted as a prerequisite, that is for the design and scope critics.
- The writer and reader land on separate files, `src/write.ignite` and `src/read.ignite`. I found no shared helper file and I do not claim one.
- Out of lane: story:typed-program-json's scope names `crates/ignite-types/src`, which story:stable-node-ids also claims. That story is outside my set, and typed-program-json's body says it does not need it.
- Out of lane: the epic's acceptance lists C2 but no story in the codespec store carries it. That is a scope-critic matter.
- Out of lane: `task check` runs `ignite check` over `src` and `tools` only. Whether the acceptance lines saying "`task check` green" actually verify the adapters' behaviour is an acceptance-critic matter.

```findings
- file: .engineering/planning/story/typed-program-json.md
  line: 8
  category: parallel-safety
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "this story and the active story:static-linking both land on crates/ignite/src/main.rs, crates/ignite/src/lib.rs and docs/slice.md (cited in both scopes), and the typed-json test under crates/ignite/tests is inferred to meet static-linking's claim on that directory; neither body names the other, so add an ordering edge recording the shared files or split the surface, without choosing"
- file: .engineering/planning/story/ignite-reader.md
  line: 19
  category: parallel-safety
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "the outcome and the first acceptance line need ignite emit --typed-json, which story:typed-program-json creates in another store; the body does not name it and no depends_on edge exists, so the codespec store schedules the reader in wave 2 beside the writer"
- file: .engineering/planning/story/ignite-round-trip.md
  line: 1
  category: parallel-safety
  severity: warning
  verdict: needs-revision
  origin: introduced
  message: "the story declares no scope and its body names no file, test or task, so the store lists it as unassessed and it cannot be placed against the writer, the reader or Taskfile.yml"
```
