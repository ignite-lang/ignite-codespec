---
format: aep.planning-md/3
id: review-result:bridge-scope-round-2
kind: review-result
status: active
title: scope critic, round 2
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

`story:ledger-example` — "Go and Rust drivers of `examples/ledger`" is required by the reader's C3 acceptance, but no item claims the Rust driver. This story's outcome (line 22) supplies only a Go driver, an Ignite driver and `harness.toml` entries `original` and `ignite`. Its context (line 28) allows only Go under `drivers/go/`. The reader's scope is `src/read.ignite` alone, and the round-trip story adds only `reference` and `roundtrip` entries. So no item owns a Rust driver or the `[[implementation]]` entries that run the Go and Rust writers' output. codespec's own example has all three: `examples/orders/drivers/rust`, plus `go` and `rust` entries in its `harness.toml`. The story should say who supplies them, or the C3 promise should be narrowed explicitly. Citation: `.engineering/planning/story/ignite-reader.md:30`, and `docs/design.md:88` ("codespec's Go and Rust writers, harness byte-equal"). The fix belongs in the story that "supplies their input and their harness entry" (`ledger-example.md:38`). `.engineering/planning/epic/ignite-codespec-bridge.md:24` (C0 to C4 done) is the parent promise that carries it.

Round 1's three scope gaps are closed:
- The epic context line 19 now names `story:typed-program-json` as the claimant of C2.
- `story:ledger-example` now claims the Ignite harness driver, which the round 1 scope report said no story carried.
- `story:cir-model` now says "on the LLVM target (the default)", so "both Ignite targets" is gone.

Promise count: I extracted 8 promises from the epic (`ignite-codespec-bridge.md:13`, `:24`) and traced all 8 to an item. The Rust-driver finding above is a sub-promise of C3, not a separate count. The 8:
- both directions through `cir/1`: writer and reader
- adapters in Ignite: writer and reader
- harness-only acceptance: writer, reader, round-trip and ledger
- C0: `story:cir-model`
- C1: `story:ignite-writer`
- C2: `story:typed-program-json`
- C3: `story:ignite-reader`
- C4: `story:ignite-round-trip`

The new public repository already exists. "Every language codespec reads or writes" is Go in and Go plus Rust out, which is what codespec has on `m1-writers-harness`. I found no reach beyond the parent. `story:ledger-example` is the example C1, C3 and C4 need. The extra acceptance lines in `story:typed-program-json` serve its own outcome. Two items claim no one outcome twice.

What I read: 7 artifacts, plus `docs/design.md` section 6 and the codespec branch tree. Commands: `aep plan artifact show` on the epic, the five stories and `review-result:bridge-scope-round-1`, `show` on `story:typed-program-json` in the ignite-lang store, `aep plan artifact graph` and `list`, and `git ls-tree` on codespec `m1-writers-harness`.

What I could not establish:
- Out of my lane (design or parallel-safety): `story:typed-program-json` is still `proposed` and depends on `story:static-linking`, and `story:ignite-reader` has no edge to it. The epic and the reader's context name the dependency, which is the stated limit of the store.
- Out of my lane (design): the writer and reader acceptances still depend on codespec's unmerged branch.

```findings
- file: .engineering/planning/story/ignite-reader.md
  line: 30
  category: scope
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "story:ledger-example — the reader's C3 acceptance requires 'the Go and Rust drivers of examples/ledger' and harness runs of the Go and Rust writers' output, but no item claims a Rust driver or the go/rust [[implementation]] entries (ledger-example's outcome lists only Go and Ignite drivers and entries original and ignite; its context allows only drivers/go; the reader's scope is src/read.ignite), so the promise of design.md:88 is claimed by nothing"
```
