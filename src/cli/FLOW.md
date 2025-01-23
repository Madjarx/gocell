# Flow

Lifecycle goes like this:

1. Transaction.go calls the `builder` to make the command
2. Transaction.go calls the `executor` to execute the command