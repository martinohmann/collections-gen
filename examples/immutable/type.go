package immutable

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/immutable -p github.com/martinohmann/collections-gen/examples/immutable

// +collections-gen=true
//
// Generate an immutable collection of Foos.
//
// +collections-gen:options=immutable
//
// Foo is a custom type.
type Foo struct {
	Bar string
}
