# grab
grab — The Go REST API Builder

## Introduction

The idea behind grab is to quickly create a REST API. In addition to
scaffolding out a starting REST API, the idea is that grab will also
provide generators so that you can add endpoints quickly.

## Documentation

```bash
$ grab new myapp
$ cd myapp
$ go build
$ ./myapp
[negroni] listening on :9090
```

## Dependencies

- [negroni][]
- [gorilla/mux][mux] 

[mux]: http://www.gorillatoolkit.org/pkg/mux
[negroni]: https://github.com/codegangsta/negroni
