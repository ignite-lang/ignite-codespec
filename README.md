# ignite-codespec

Translate code into and out of [Ignite](https://github.com/ignite-lang/ignite) through
[codespec](https://github.com/beyond10x/codespec)'s typed code representation `cir/1`, written in
Ignite.

```
Go / Rust ──▶ cir/1 ──codespec-write-ignite──▶ Ignite
Ignite ──codespec-read-ignite──▶ cir/1 ──▶ Go / Rust
```

Every translation is checked by codespec's harness: the original and the output run on the same
seeded inputs and their outputs are compared byte for byte.

Status: design only ([docs/design.md](docs/design.md)); milestones C0 to C4 there.

## License

Apache-2.0.
