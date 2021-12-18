# Go Private Monorepo

This repository is a test/guide for importing libraries from private monorepos.

## Project Layout
```
.
├── app1        # The main app that imports the libraries
├── lib1        # The first library, adjacent to app1
└── nested      # Just a directory for organisation
   └── lib2     # The second library, nested from app1
```

## GOPRIVATE

Go provides a useful environment variable `GOPRIVATE` for telling go that module(s)
are not publicly accessible. More info on this can be found with
`go help private`.

- For this project we set `GOPRIVATE=github.com/jonnobrow/go-private-monorepo/*`.
- You can also set for an organisation/user `GOPRIVATE=github.com/jonnobrow/*`

Setting this will mean modules with a full path matching `GOPRIVATE` will be
treated as private and Go will require authentication to access them. 

## Git Setup

By default Go uses Git with HTTPS to get the modules from the private repository.

To setup SSH you can add the following snippet to you git configuration 
(`~/.gitconfig` or `~/.config/git/config`).

```ini
[url "ssh://git@github.com:"]
    insteadOf = https://github.com/
```

## The `replace` block in `go.mod`

The replace block allows you to tell the compiler where modules live locally.
For this project for example I can tell the compiler to traverse up a directory
and into `lib1` for the `lib1` package and something similar for `lib2`. 

The block below shows how to do this in your `go.mod` file.

```go
replace (
    github.com/jonnobrow/go-private-monorepo/lib1 => ../lib1
    github.com/jonnobrow/go-private-monorepo/nested/lib2 => ../nested/lib2
)
```
