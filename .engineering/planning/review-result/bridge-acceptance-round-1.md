---
format: aep.planning-md/3
id: review-result:bridge-acceptance-round-1
kind: review-result
status: active
title: acceptance critic, round 1
relations:
- reviews: epic:ignite-codespec-bridge
- reviews: story:cir-model
- reviews: story:ignite-writer
- reviews: story:ignite-reader
- reviews: story:ignite-round-trip
revision: 1
---
needs-revision

story:cir-model — no acceptance bullet observes the generator, though the Outcome names `tools/gen_model.ignite` generating `src/cir.ignite` from the schema "at a pinned commit"; nothing says which commit or that regenerating gives a byte-identical file — .engineering/planning/story/cir-model.md:28
story:cir-model — the read-and-re-encode and refusal bullets name no command or test, and `task check` only runs `ignite check` over src and tools, so it is already green on a tree with no `src/cir.ignite` — /home/timo/projects/ignite-lang/ignite-codespec/Taskfile.yml:8
story:ignite-writer — "An example without channels and shared pointers" names no example, and none exists: the only checked-in codespec example is `orders`, which has channels, and its harness drivers are Go and Rust only — .engineering/planning/story/ignite-writer.md:27
story:ignite-writer — "nothing is dropped" has no reference list, so a complete refusal list can't be told from a short one, and where the output lists them is unstated; pin the expected set (`orders.go:53`, `:58`, `:132`, per design.md section 5) — .engineering/planning/story/ignite-writer.md:28
story:ignite-reader — "An Ignite example" is unnamed and the bullet joins two independent outcomes (`codespec validate --canonical` passes; Go and Rust writers' programs are byte-equal to the Ignite original), and the second needs an Ignite harness driver that does not exist (docs/design.md:93) — .engineering/planning/story/ignite-reader.md:27
story:ignite-round-trip — "an Ignite example" is unnamed and the bullet names no harness command, seeds or commands count, or Ignite driver, so "byte-equal on the seeded traces" can't be run — .engineering/planning/story/ignite-round-trip.md:25
story:typed-program-json — the Outcome promises every declaration, statement and expression with its source position, and format `ignite-typed/0.1`, but the acceptance checks only that expressions carry a `type`, so positions, declarations, statements and the format string go unchecked — .engineering/planning/story/typed-program-json.md:41

What I read: 6 artifacts given or implied (epic:ignite-codespec-bridge, story:cir-model, story:ignite-writer, story:ignite-reader, story:ignite-round-trip, story:typed-program-json), all 6 opened in full with `aep plan artifact show`. I also ran `aep plan artifact kinds` and `lifecycle story`, read docs/design.md, Taskfile.yml and AGENTS.md in ignite-codespec, and checked codespec's tree, its `m1-writers-harness` branch (`story:harness`, Taskfile) and the CLI's `--canonical` flag. In the ignite repo I grepped for `ignite-model/0.1` and `formatVersion`. Both stores' `aep plan artifact validate` return valid, or only an unrelated story's missing scope.

What I could not establish:
- Whether codespec's harness can take an Ignite implementation at all. `story:harness` at 3b2f6eb lists only the original, Go output and Rust output for `examples/orders`; design.md:93 lists this as open.
- The epic's acceptance ("Milestones C0 to C4 of `docs/design.md` section 6 are done") is checkable through design.md's "done when" column. I left it unflagged.
- Outside my lane: the reader and round-trip stories depend on the C2 story in another store, and no `depends_on` can express that. The writer's design relies on facts (`shared`, `fallible`), and the Rust writer refuses a module without them, while the orders document is read-only output. The epic's Outcome says "every language codespec reads or writes" but C0 to C4 cover Go and Rust only. These belong to the design, scope and parallel-safety critics.

```findings
- file: .engineering/planning/story/cir-model.md
  line: 28
  category: acceptance
  severity: warning
  verdict: needs-revision
  origin: introduced
  message: "no acceptance bullet observes the generator, though the Outcome names tools/gen_model.ignite generating src/cir.ignite from the schema at a pinned commit; nothing says which commit or that regenerating gives a byte-identical file"
- file: /home/timo/projects/ignite-lang/ignite-codespec/Taskfile.yml
  line: 8
  category: acceptance
  severity: warning
  verdict: needs-revision
  origin: introduced
  message: "story:cir-model's read-and-re-encode and refusal bullets name no command or test, and task check only runs ignite check over src and tools, so it is already green before the work"
- file: .engineering/planning/story/ignite-writer.md
  line: 27
  category: acceptance
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "'An example without channels and shared pointers' names no example and none exists: the only checked-in codespec example is orders, which has channels, and its harness drivers are Go and Rust only"
- file: .engineering/planning/story/ignite-writer.md
  line: 28
  category: acceptance
  severity: warning
  verdict: needs-revision
  origin: introduced
  message: "'nothing is dropped' has no reference list, so a complete refusal list cannot be told from a short one, and where the output lists them is unstated; pin the expected set (orders.go:53, :58, :132 per design.md section 5)"
- file: .engineering/planning/story/ignite-reader.md
  line: 27
  category: acceptance
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "'An Ignite example' is unnamed and the bullet joins two independent outcomes (codespec validate --canonical passes; Go and Rust writers' programs are byte-equal to the Ignite original), the second needing an Ignite harness driver that does not exist (docs/design.md:93)"
- file: .engineering/planning/story/ignite-round-trip.md
  line: 25
  category: acceptance
  severity: blocker
  verdict: needs-revision
  origin: introduced
  message: "'an Ignite example' is unnamed and the bullet names no harness command, seeds or commands count, or Ignite driver, so 'byte-equal on the seeded traces' cannot be run"
- file: .engineering/planning/story/typed-program-json.md
  line: 41
  category: acceptance
  severity: warning
  verdict: needs-revision
  origin: introduced
  message: "the Outcome promises every declaration, statement and expression with its source position, and format ignite-typed/0.1, but the acceptance checks only that expressions carry a type, so positions, declarations, statements and the format string go unchecked"
```
