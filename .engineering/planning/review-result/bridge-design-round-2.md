---
format: aep.planning-md/3
id: review-result:bridge-design-round-2
kind: review-result
status: active
title: design critic, round 2
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

- story:ledger-example — its outcome names no Rust driver, no `go` or `rust` implementation entry and no `reference.ignite`, yet `story:ignite-reader`'s second acceptance line runs codespec's Go and Rust writers "with the Go and Rust drivers of `examples/ledger`", and the reader, round-trip and ledger-example acceptances all read `reference.ignite`. No story in the set produces the Rust driver or the `go` and `rust` manifest entries, and the reader's scope is `src/read.ignite` only. The fix is for ledger-example's outcome to own `drivers/rust`, the Go overlay, the `go` and `rust` entries and `reference.ignite`, or to put them in the reader's scope with an edge. This is the round-1 finding (a driver and entry that nothing produces), only partly closed. — .engineering/planning/story/ignite-reader.md:30, .engineering/planning/story/ledger-example.md:22, `git -C /home/timo/beyond10x/codespec show m1-writers-harness:examples/orders/harness.toml` (the orders example has its `go` and `rust` entries and a `drivers/rust`)

- story:ledger-example — `examples/ledger/harness.toml` is one manifest that `codespec harness` builds and compares in full, with no way to select entries. Four stories each name a different subset of it:
  - ledger-example's outcome defines `ignite` as "the writer's output", but its own acceptance runs the Ignite driver "against a hand-written `reference.ignite`".
  - The writer uses `original` and `ignite`.
  - The reader uses `original` plus the Go and Rust writer outputs.
  - The round trip adds `reference` and `roundtrip`, and only it has the manifest in scope.

  Before the writer lands, `ignite` has nothing to build. The entry that runs `reference.ignite` is named by no story until the round trip creates it. The body of ledger-example has to say which entry runs the hand-written file, where `ignite` points, and which story adds each later entry, so the entries are not split across four stories. — .engineering/planning/story/ledger-example.md:22 and :32, .engineering/planning/story/ignite-round-trip.md:15 and :30, /home/timo/beyond10x/codespec `git show m1-writers-harness:crates/harness/src/lib.rs` (`build` loops over every `[[implementation]]`; `Harness` takes only example, seeds, commands and build-dir)

**What you read:**
- Six artifacts: the epic and its five stories (cir-model, ledger-example, ignite-writer, ignite-reader, ignite-round-trip) in the codespec store.
- `story:typed-program-json` in the ignite-lang store.
- `review-result:bridge-design-round-1`, `docs/design.md`, `AGENTS.md` and `Taskfile.yml`.
- codespec's harness manifest and `crates/harness/src/lib.rs` on `m1-writers-harness`.
- Commands: `aep plan artifact show` on each id, `relations`, `graph` and `validate` (valid, both stores) in both stores, and `waves` (3 Taskfile.yml collisions) in the codespec store.

**Graph and acceptance reads:**
- I walked 9 declared edges in the codespec store: 5 `decomposes` and 4 `depends_on`, and 1 `depends_on` in the ignite-lang store from `typed-program-json` to `static-linking`. I walked outside the set to the ignite-lang graph. There is no cycle and no serialising chain.
- The shape is `cir-model` (wave 1), then `ledger-example`, then `ignite-reader` and `ignite-writer` in parallel, then `ignite-round-trip`. `ledger-example` sits in wave 2 only because of the `Taskfile.yml` collision, not an edge.
- The `typed-program-json` edge to `static-linking` is recorded with its reason, a shared file, in its body (`story:typed-program-json` Context), so it is not a chain finding.
- I traced 14 acceptance reads against a producer in the set. 12 are recorded by an edge, directly or through the writer. 2 have no producer in the set (the Rust driver, and the `go` and `rust` entries), and they are the first finding. The `ignite` and `reference` entries are produced in two places, and that is the second finding.
- The round-1 split-example finding is mostly closed: the writer, reader and round trip now depend on `ledger-example`. The two findings above are what remains.

**What you could not establish:**
- Out of my lane (scope and acceptance): the harness's trace generator and output format are fixed to the orders commands (`new`, `place`, `pay`, `ship`, `cancel`, `get`, `pending`, with events from a channel) in `crates/harness/src/lib.rs` on `m1-writers-harness`. `ledger-example` does not say how a channel-free ledger answers those commands. This may mean the ledger needs an orders-shaped API or a codespec change (its M3 epic). I did not set the verdict on it.
- Out of my lane (parallel safety): the writer and the reader both need `examples/ledger/harness.toml` and `Taskfile.yml` in wave 3, and neither declares them in scope.
- Outside the set, not a defect: the reader's wait on `story:typed-program-json` (another store, recorded in prose at `ignite-reader.md:24`), and codespec's harness being on an unmerged branch.

```findings
- file: /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ignite-reader.md
  line: 30
  category: design
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "story:ledger-example — its outcome names no Rust driver, no go or rust implementation entry and no reference.ignite, yet the reader's acceptance reads the Go and Rust drivers of examples/ledger and the reader, round trip and ledger-example all read reference.ignite; no story in the set produces the Rust driver or the go and rust entries, so ledger-example's outcome must own them or the reader's scope must add them with an edge"
- file: /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ledger-example.md
  line: 22
  category: design
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "story:ledger-example — examples/ledger/harness.toml is one manifest the harness builds and compares in full with no entry selection, but ledger-example defines ignite as the writer's output while its own acceptance runs reference.ignite, and the writer, reader and round trip each name a different subset of entries with only the round trip holding the manifest in scope; the body must say which entry runs the hand-written file, where ignite points and which story adds each later entry"
```
