# hello-service

A small Go program for trying out coding agents end to end: they clone it,
change it in a sandboxed box, and the change comes back as a pull request.
Its issues are the tickets.

The code is the `hello` module of [golang/example](https://github.com/golang/example)
at 7f05d21, under its BSD licence (see LICENSE and PATENTS).

```bash
cd hello && go vet ./... && go test ./...
```
