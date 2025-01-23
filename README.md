# Cellframe CLI ran via Golang

Note - sometimes the CLI might report stuff that isnt true, double check it

### Structure
```
src
├── cli
│   ├── command_builder.go  # builds actual command based on params:  abc=123 --> "wallet info abc"
│   ├── executioner.go      # Runner / daemon / executes command
│   ├── FLOW.md             # WIP doc
│   ├── transaction.go      # main file, bundles everything together
│   └── types.go            # .
├── config
│   └── paths.go            # Path to CLI binary
└── main.go

```

### Explanation & what we need

- transactions vs orders vs datum signing etc. (partial purchase)
- how do wallets work (wallet activation)
- how do keys work (certificates, etc.)


### How to run

```go run src/examples/**``` - its empty as of now

or if its not implemented then:

```go run src/main.go```