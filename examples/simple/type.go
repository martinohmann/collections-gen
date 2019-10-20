package simple

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/simple -p github.com/martinohmann/collections-gen/examples/simple

// +collections-gen=true
//
// Foo is a custom type.
type Foo struct {
	Bar string
}
