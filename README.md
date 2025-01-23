# Cellframe CLI done via Golang

## This document is not up to date
I will update this doc after running the tests myself
For now forget that it exists



Lets try to avoid passing in strings as one missed zero place can lead to
a lot of headache.


### Structure
```
src/
  ├── cli/
  │   └── manager.go      # CLI command implementations
  ├── config/
  │   └── paths.go        # Manages imports and other important fs paths
  │   
  └── main.go             # Example of how to use the CLI commands
```

### Explanation & what we need

- transactions vs orders vs datum signing etc. (partial purchase)
- how do wallets work (wallet activation)
- how do keys work (certificates, etc.)


### How to run

```go run src/examples/**``` - its empty as of now

or if its not implemented then:

```go run src/main.go```