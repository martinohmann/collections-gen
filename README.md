collections-gen
===============

[![Build Status](https://travis-ci.org/martinohmann/collections-gen.svg?branch=master)](https://travis-ci.org/martinohmann/collections-gen)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinohmann/collections-gen?style=flat)](https://goreportcard.com/report/github.com/martinohmann/collections-gen)
[![GoDoc](https://godoc.org/github.com/martinohmann/collections-gen?status.svg)](https://godoc.org/github.com/martinohmann/collections-gen)

Code generator for mutable and immutable golang collections. `collections-gen`
is built on top of [gengo](https://github.com/kubernetes/gengo) and can
generate collections code based on tags in the doc blocks preceeding a type. It
identifies the types needing code generation by the comment tag
`+collections-gen`.

Simple example
--------------

Given you have a type `Foo`, add the `+collections-gen=true` tag to the 1st or
2nd doc block preceeding the type definition:

```go
// +collections-gen=true
//
// Foo is a custom type.
type Foo struct {
    ...
}
```

To generate a collection for the `Foo` type, run `collections-gen` with the input and output packages as arguments:

```sh
collections-gen -i pkg/path/to/foo -o pkg/path/to/generated/collection
```

For example:

```sh
collections-gen -i github.com/martinohmann/mylib/foo -o github.com/martinohmann/mylib/foo/collections
```

The code generator will generate the file `foo_collection.go` containing the
`FooCollection` type in `github.com/martinohmann/mylib/foo/collections`.

Check out [examples/simple](examples/simple/) for the example code.

**Hint:** You can also make use of `go generate` by adding a machine readable comment to your source files:

```go
//go:generate collections-gen -i pkg/path/to/foo -o pkg/path/to/generated/collection
```

Afterwards you can generate your custom collections by running:

```sh
go generate ./...
```

Naming
------

By default the generated collections are named `ElementTypeCollection` and
`ImmutableElementTypeCollection` for immutable collections, where `ElementType`
is the element type of the collection. The default output files follow the same
convention just in snake case: `element_type_collection.go` and
`immutable_element_type_collection.go`. For overriding the naming behaviour
have a look at the `name` and `out-name` options in the [customizing generated
collections](#customizing-generated-collections) section.

Customizing generated collections
---------------------------------

To customize the generator behaviour for a given type, the comment tag
`+collections-gen:options` can be added. Adding it multiple times causes the
generation of multiple collection types. The `+collections-gen:options` tag
accepts configuration flags of the following form:

```go
// +collections-gen:options=flag1=value1,boolflag1,flag2=value2,boolflag2
```

### Supported flags

| Flag            | Type     | Description                                                                                                                                                                     |
| ----            | ----     | -----------                                                                                                                                                                     |
| `mutable`       | `bool`   | If set, the generated collection will be mutable (default).                                                                                                                     |
| `immutable`     | `bool`   | If set, the generated collection will be immutable.                                                                                                                             |
| `pointer`       | `bool`   | If set, the collection elements are pointers.                                                                                                                                   |
| `underlying`    | `bool`   | If the type is an alias, use the underlying type for the collection elements.                                                                                                   |
| `equality-func` | `string` | Specify a custom func for equality checks. If left out, `reflect.DeepEqual` will be used for slice, map and func types. For everything else plain `==` comparison will be used. |
| `name`          | `string` | Specify a custom name for the generated collection type.                                                                                                                        |
| `out-name`      | `string` | Specify a custom name for the generated go file (without the `.go` extension).                                                                                                  |
| `prefix`        | `string` | The name prefix for the generated collection. Defaults to `immutable` if the collection is immutable. Will be ignored if `name` is set.                                         |
| `suffix`        | `string` | The name suffix for the generated collection. Defaults to `collection`. Will be ignored if `name` is set.                                                                       |
| `noprefix`      | `bool`   | Disables adding the collection name prefix.                                                                                                                                     |
| `nosuffix`      | `bool`   | Disables adding the collection name suffix.                                                                                                                                     |

### Custom options example

```go
// +collections-gen=true
//
// Generate FooCollection with *Foo element type.
//
// +collections-gen:options=pointer
//
// Generate an immutable collection with Foo element type and name
// BarCollection and write it into collection.go in the output package.
//
// +collections-gen:options=immutable,name=BarCollection,out-name=collection
//
// Foo is a custom type.
type Foo struct {
	Bar string
}
```

Check out the examples in the [examples](examples/) directory to see all
options in action.

License
-------

The source code of collections-gen is released under the MIT License. See the
bundled LICENSE file for details.
