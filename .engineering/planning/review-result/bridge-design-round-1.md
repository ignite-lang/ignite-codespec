---
format: aep.planning-md/3
id: review-result:bridge-design-round-1
kind: review-result
status: active
title: design critic, round 1
relations:
- reviews: epic:ignite-codespec-bridge
- reviews: story:cir-model
- reviews: story:ignite-writer
- reviews: story:ignite-reader
- reviews: story:ignite-round-trip
revision: 1
---
needs-revision

- story:ignite-writer — its acceptance (the harness compares the Go original and the Ignite output byte-equal) reads an Ignite-side driver, a `harness.toml` implementation entry and an example without channels or shared pointers. No story in the set produces any of the three, and the three acceptances of ignite-writer, ignite-reader and ignite-round-trip all need them with no edge. The fix is one story that owns the Ignite example, its driver and the codespec harness entry, with `depends_on` edges from the writer, the reader and the round trip, or the same scope added to one existing story that the other two then depend on. — /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ignite-writer.md:18, /home/timo/beyond10x/codespec (git show m1-writers-harness:examples/orders/harness.toml, which lists one `[[implementation]]` per language, each with its own driver under `examples/orders/drivers/`), /home/timo/projects/ignite-lang/ignite-codespec/docs/design.md:93

**What you read:**
- Five artifacts in the codespec store: the epic and four stories, each shown whole.
- `story:typed-program-json` in the ignite-lang store, shown whole.
- `aep plan artifact relations` and `graph` in both stores.
- `aep plan artifact validate` in both stores: codespec reports `valid`. Ignite-lang reports `valid` plus one warning about `story:llvm-lowering-copies` having no scope, which is not this set's.
- Both design documents (`docs/design.md` in ignite-codespec and in codespec).
- codespec's harness story, harness epic and `harness.toml` on branch `m1-writers-harness`.

**Graph and acceptance reads:**
- I walked 8 edges in the codespec store: 4 `decomposes` and 4 `depends_on`. There is no cycle and no serialising chain.
- The shape is cir-model, then writer and reader in parallel, then round-trip.
- I also walked outside the set to `story:typed-program-json`. It has no edges at all in the ignite-lang graph.
- I traced 9 acceptance reads that land on a producer in the set. Of those, 6 are recorded by an edge: writer to cir-model, reader to cir-model, and round-trip to writer and reader, twice each. The 3 unrecorded reads are the Ignite driver, entry and example the finding names, and no story in the set produces them, so no edge could record them.
- `story:typed-program-json` is a clean vertical slice. Its acceptance reads only the document it emits, and it has no ordering dependency on `stable-node-ids`.

**What I could not establish:**
- Out of my lane (scope): the epic's acceptance requires milestone C2, but no story under the epic covers it. C2 lives only in the ignite-lang store, with no recorded link to the epic.
- Out of my lane (scope): `ignite-writer` needs a channel-free example. codespec's only example is `orders`, and its second example, M3, is a separate draft epic in codespec.
- Not a design defect: `ignite-reader` needs `ignite emit --typed-json`, which is `story:typed-program-json` in another store, and its body does not name that story. This is a dependency outside the set. The reader cannot start until that story lands, and no wave computation in this store will show it.
- Out of my lane (acceptance): `ignite-reader` and `ignite-round-trip` say nothing about the `codespec pass` step. codespec's Rust writer refuses a module that has no facts (commit 3b2f6eb).
- Out of my lane (acceptance): `cir-model` reads "every checked-in example document". On `main` of codespec only `orders.read.cir.json` exists, and `orders.cir.json` is on an unmerged branch.
- Out of my lane (parallel safety): `ignite-round-trip` declares no scope.

```findings
- file: /home/timo/projects/ignite-lang/ignite-codespec/.engineering/planning/story/ignite-writer.md
  line: 18
  category: design
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "its acceptance (the harness compares the Go original and the Ignite output byte-equal) reads an Ignite-side driver, a harness.toml implementation entry and an example without channels or shared pointers that no story in the set produces, while ignite-reader and ignite-round-trip read the same three; one story must own them with depends_on edges from the writer, the reader and the round trip, or the scope must be added to one existing story that the other two depend on"
```
