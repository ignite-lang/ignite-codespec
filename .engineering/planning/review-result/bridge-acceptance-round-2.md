---
format: aep.planning-md/3
id: review-result:bridge-acceptance-round-2
kind: review-result
status: active
title: acceptance critic, round 2
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

ignite-codespec-bridge/story:ledger-example — the acceptance runs `codespec harness examples/ledger`, but the harness only generates orders commands, so no ledger driver can read its traces — /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ledger-example.md:32
story:ledger-example — the first bullet runs "only the `original` implementation and the Ignite driver … against `reference.ignite`", but `codespec harness` has no implementation filter and builds every entry of `harness.toml`; the Outcome also reserves the entry `ignite` for the writer's output, which does not exist when this story closes — /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ledger-example.md:32
story:ignite-reader — the second bullet needs "the Go and Rust drivers of `examples/ledger`" and `harness.toml` entries for the Go-writer and Rust-writer outputs, yet `ledger-example` delivers only a Go driver and an Ignite driver, and no story's scope covers those entries (the reader's scope is `src/read.ignite` only) — /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ignite-reader.md:30

Citations for the first two lines:

- `crates/harness/src/lib.rs:22-28` on `m1-writers-harness` at b73ffce: the trace table holds only `new`, `place &lt;customer&gt; &lt;total&gt; &lt;gift&gt;`, `pay`, `ship`, `cancel`, `get` and `pending`.
- The same file, lines 49-57, gives the output format as orders-shaped.
- `crates/cli/src/main.rs:63-75` on that branch: `Harness` takes only `example`, `--seeds`, `--commands` and `--build-dir`.
- `crates/harness/src/lib.rs:221-245` on that branch: `build` loops over every `[[implementation]]`.

Round 1 status:
- Fixed, no repeat:
  - cir-model: the generator and read/re-encode bullets now name `task model-check`, `task roundtrip-check` and a byte diff. Those tasks do not exist in the Taskfile today (it has only `check`), so the transition is observable.
  - ignite-writer: the named example is now `ledger`, and the three orders refusals are pinned to `orders.go:53`, `:58` and `:132` through a checked-in expected file.
  - ignite-round-trip: the bullet now names `task roundtrip`, `out/roundtrip.ignite`, 200 seeds and the entries `reference` and `roundtrip`.
  - typed-program-json: the acceptance now checks `formatVersion`, a `span` on every declaration, statement and expression, a `type` on every expression, AST counts, byte-stable output and the `docs/slice.md` subsection.
- Still open: the reader's unnamed Ignite example and its Ignite driver are fixed. The reader's second bullet now fails on the missing Rust driver and the unscoped `harness.toml` entries, as the third finding says.

What I read: 7 artifacts given or implied, 7 opened in full. They are story:cir-model, story:ledger-example, story:ignite-writer, story:ignite-reader, story:ignite-round-trip and epic:ignite-codespec-bridge (`aep plan artifact show` in ignite-codespec, head 4d1641b), plus story:typed-program-json (in ignite-lang, head 0fa97d555). I also read review-result:bridge-acceptance-round-1, `aep plan artifact kinds` and `lifecycle story`, `docs/design.md` and the Taskfile. In codespec I read the harness crate and CLI on `m1-writers-harness` (b73ffce), the `story:harness` body and the orders `harness.toml`.

What I could not establish:
- Whether the ledger stories intend to extend codespec's trace generator or to map the orders commands onto ledger operations. Either way the story must say so. The line-22 rule of the harness doc ("no field is empty or holds a space") and the `Audit` note suggest the generator is per-example only in the sense of orders.
- Whether `refusals.json` for orders is exactly the three named refusals. The writer does not exist, so that bullet cannot be run yet. Its "or added to that expected file" clause lets the expected file grow. I judged that visible in the commit diff and did not flag it.
- The stories cite the harness at 3b2f6eb, but the harness commit is b73ffce (3b2f6eb is the Rust writer under it). That is a citation slip, not an acceptance defect.
- Out of my lane, and they set no part of the verdict:
  - The reader and round-trip stories depend on C2 in another store, which no edge can express.
  - Whether the harness change the first finding needs is a scope question for codespec.
  - Whether the epic's "every language codespec reads or writes" is covered by C0 to C4, which cover Go and Rust only.

```findings
- file: .engineering/planning/story/ledger-example.md
  line: 32
  category: acceptance
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "the acceptance runs `codespec harness examples/ledger`, but the harness only generates orders commands (place, pay, ship, cancel, get, pending), so no ledger driver can read its traces and the story names neither a harness change nor how a ledger driver maps them"
- file: .engineering/planning/story/ledger-example.md
  line: 32
  category: acceptance
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "the first bullet runs 'only the `original` implementation and the Ignite driver ... against reference.ignite', but `codespec harness` has no implementation filter and builds every harness.toml entry, and the Outcome reserves the entry `ignite` for the writer's output, which does not exist when this story closes"
- file: .engineering/planning/story/ignite-reader.md
  line: 30
  category: acceptance
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "the second bullet needs the Go and Rust drivers of examples/ledger and harness.toml entries for the Go-writer and Rust-writer outputs, but ledger-example delivers only a Go driver and an Ignite driver and no story's scope covers those entries"
```
