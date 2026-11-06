# AGENTS.md — ignite-codespec

Ignite adapters for codespec's `cir/1`: `codespec-write-ignite` (`cir/1` → Ignite) and
`codespec-read-ignite` (Ignite → `cir/1`). Design: [docs/design.md](docs/design.md). Humans:
[README.md](README.md).

## Sources of truth

- `cir/1` is defined by `github.com/beyond10x/codespec` (`crates/model/src/lib.rs`, schema
  `schemas/cir-1.schema.json`). This repository never changes the format; a gap is a refusal here
  and a request there.
- Ignite is defined by `github.com/ignite-lang/ignite` (`docs/slice.md`). Pin the compiler by commit
  or release tag.

## Rules

- Written in Ignite. No Rust, Go or Python in this repository; Bash only for one-line invocations
  in `Taskfile.yml`.
- `src/cir.ignite` is generated from the pinned schema by `tools/gen_model.ignite`; never edit it by
  hand, regenerate it with the generator change.
- Strict reading, canonical writing, deterministic output: the same input gives the same bytes.
- What an adapter cannot express is refused with a position and a reason, never approximated.
- A translation is accepted by codespec's comparison harness only, never by a claim.
- No comments in `Taskfile.yml` or any other config file.
- Commit titles follow Conventional Commits; every commit has a bullet body.

## Planning

AEP store in `.engineering/` (`aep plan artifact list`).
