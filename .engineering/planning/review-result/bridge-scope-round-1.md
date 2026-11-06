---
format: aep.planning-md/3
id: review-result:bridge-scope-round-1
kind: review-result
status: active
title: scope critic, round 1
relations:
- reviews: story:cir-model
- reviews: epic:ignite-codespec-bridge
- reviews: story:ignite-writer
- reviews: story:ignite-reader
- reviews: story:ignite-round-trip
revision: 1
---
approve

**What you read:** 6 artifacts (the epic, four stories, `story:typed-program-json`), plus `docs/design.md`, `AGENTS.md`, the drafter's scratch in `.engineering/drafts/`, both stores' graphs and the codespec repo's `main` and `m1-writers-harness`. Commands: `aep plan artifact show epic:ignite-codespec-bridge`, `show` on each story, `aep plan artifact graph` in both stores, `git ls-tree`, `gh repo view ignite-lang/ignite-codespec`.

**Promise count.** I extracted 10 promises from the epic, at `.engineering/planning/epic/ignite-codespec-bridge.md:13` and `:22`, and traced all 10.

| Promise | Claimed by |
|---|---|
| Code moves both directions through `cir/1` | `story:ignite-writer`, `story:ignite-reader` |
| Adapters written in Ignite | the same two stories, and `AGENTS.md` |
| Translation accepted by the harness only | harness acceptance in the writer, reader and round-trip stories; `AGENTS.md` |
| C0 | `story:cir-model` |
| C1 | `story:ignite-writer` |
| C2 | `story:typed-program-json` in the ignite-lang store. Its context names "milestone C2" of `docs/design.md`. Its acceptance is "every file under `std/` and `examples/`, every expression carries a `type`", which matches design.md:87. |
| C3 | `story:ignite-reader` |
| C4 | `story:ignite-round-trip` |
| New public repository | Already exists: `gh repo view` shows `ignite-lang/ignite-codespec`, PUBLIC. |
| "every language codespec reads or writes" | The adapters take `cir/1`, so they are language-neutral. Codespec today reads Go only and writes Go and Rust, and C3 names Go and Rust writers. |

No two items claim the same outcome, and no item reaches past the epic. The extra acceptance lines in `story:typed-program-json` are the format subsection in `docs/slice.md`, the version key, determinism and the `std.json` read test. All serve its own outcome, "tools written in Ignite can read Ignite programs", so they are not reach.

**What you could not establish:**
- The epic's context, and the context of `story:ignite-reader`, never name `story:typed-program-json` as the claimant of C2. The codespec store alone shows no item for C2, and nothing links the two stores. It is covered, so I did not count it, but a one-line pointer would let the epic be closed from its own store.
- Nothing claims the harness driver for Ignite programs: how an Ignite original or output receives a trace and prints its output. C1, C3 and C4 all need it. `docs/design.md:93` names it as an open question, so it is not a gap, but no story's context or exclusions carry it.
- `story:cir-model` adds "on both Ignite targets", which is in neither the epic nor design.md:85. This is minor and only a note.
- Out of my lane:
  - `story:ignite-reader` has no `depends_on` or context for the C2 story. (design, parallel-safety)
  - The writer and reader acceptances depend on the codespec writers and harness, which exist only on the unmerged branch `m1-writers-harness` (3b2f6eb). Codespec `main` has only a Go reader. (design)

```findings
[]
```
