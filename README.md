# lamp

analyzer collections.

## install run

```
$ go install github.com/kijimaD/lamp@main
$ lamp sample/main.go
──── 3 scopes ────
(*types.Scope)(0xc00001e600)(package "main" scope 0xc00001e600 {
.  func main.main()
}
)
(*types.Scope)(0xc00001e660)( scope 0xc00001e660 {
.  package fmt
}
)
(*types.Scope)(0xc0002f3c80)(function scope 0xc0002f3c80 {
.  const message untyped string
}
)
```

## docker run

```
$ docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/lamp:latest sample/main.go
```
