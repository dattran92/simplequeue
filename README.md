# Simple Queue in Go

### About
I'm working mostly with NodeJS in backend. It's single threaded, so that I don't have to care about thread-safe. So that, I picked Golang to test thread-safe in an simple implementation of Queue.

### How it works
I implement a struct with a slice for queueing and a lock to make sure the same thread cannot input data at the same time

### Run example

```
go run main.go
```
