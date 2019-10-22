package pointer

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/pointer -p github.com/martinohmann/collections-gen/examples/pointer

// +collections-gen=true
//
// Generate FooCollection with *Foo element type.
//
// +collections-gen:options=pointer
//
// Foo is a custom type.
type Foo struct {
	Bar string
}
