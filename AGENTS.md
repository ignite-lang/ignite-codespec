# AGENTS.md — ignite-codespec

Ignite adapters for codespec's `cir/1`: `codespec-write-ignite` (`cir/1` → Ignite) and
`codespec-read-ignite` (Ignite → `cir/1`). Design: [docs/design.md](docs/design.md). Humans:
[README.md](README.md).

## Sources of truth

- `cir/1` is defined by `github.com/beyond10x/codespec` (`crates/model/src/lib.rs`, schema
  `schemas/cir-1.schema.json`). This repository never changes the format; a gap is a refusal here
  and a request there.
- Ignite is defined by `github.com/ignite-lang/ignite` (`docs/slice.md`). Pin the compiler by commit
  or release tag: `ignite.pin` holds the commit every task runs with as `IGNITE`, the first with
  `ignite emit --typed-json` (format `ignite-typed/0.1`, `docs/slice.md` section 24), which
  `codespec-read-ignite` runs. The reader's inline tests hold that command's documents for small
  source files; regenerate them with the source when the pin moves.

## Rules

- Written in Ignite. No Rust or Python in this repository; Bash only for one-line invocations in
  `Taskfile.yml`. Go only as test input: a package under `examples/*/source/` (read by `codespec
  read --lang go`) and its harness driver under `examples/*/drivers/go/`, nowhere else.
- `examples/<name>/harness.toml` is codespec's harness manifest; `task harness` runs it with the
  `ignite` of the `IGNITE` var first on `PATH`, so an entry builds with a bare `ignite`. An Ignite
  driver sees the implementation under test as its sibling module through a symbolic link beside
  it (`examples/ledger/drivers/ignite/ledger.ignite`).
- `src/cir.ignite` is generated from the pinned schema by `tools/gen_model.ignite`; never edit it by
  hand, regenerate it with the generator change.
- Strict reading, canonical writing, deterministic output: the same input gives the same bytes.
- What an adapter cannot express is refused with a position and a reason, never approximated.
- Exception: a program that traps in its source language (Ignite Int overflow) may be written as one that does not; nothing else is approximated.
- A translation is accepted by codespec's comparison harness only, never by a claim.
- No comments in `Taskfile.yml` or any other config file.
- Commit titles follow Conventional Commits; every commit has a bullet body.

## Planning

AEP store in `.engineering/` (`aep plan artifact list`).
