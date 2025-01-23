# Cellframe CLI ran via Golang

**Note - sometimes the CLI might report stuff that isnt true, double check it

**Note - make sure to read the comments, i've left some useful information there (i.e. format of the `value` and `fee` fields)

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


- Transactions - they are used to transfer Tokens (https://wiki.cellframe.net/Resources/Terms+and+definitions/Token). 
- Exchanges - Cellframe has an exchange "built in" - it is an order-book based system. People publish orders, other people buy them
- Wallets are "node-based" meaning in order to use your `FreshWallet#2` 
  - you need to have access to it via node - check via `wallet list` and **activate it** - `wallet activate -w FreshWallet#2 -password 123
  - you have to create it with `wallet new`
  - **those wallets need to be active and accessible in order for you to use them via the cli**


### How to test the integrations

Check the file and comment out regions that you do not want to run.
Pass in the proper values you want to run the tests with

```go run src/main.go```