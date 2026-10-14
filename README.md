# Go Studies

This is one of my study repositories. My focus here is practicing [Go](https://go.dev/) by building small projects.

The project ideas come from [florinpop17/app-ideas](https://github.com/florinpop17/app-ideas), a collection of application ideas organized by difficulty level.

## Structure

A single Go module, one folder per project, grouped by the app-ideas tiers:

```
beginner/<project>/
intermediate/<project>/
advanced/<project>/
```

Each project is its own `package main`, with the logic kept apart from `main.go` so it can be tested.

## Usage

```sh
go run ./beginner/hello   # run one project
go test ./...             # test everything
```