package customtype

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/custom-type -p github.com/martinohmann/collections-gen/examples/custom-type

// +collections-gen=true
// +collections-gen:options=pointer,equality-func=Equal
// +collections-gen:options=pointer,immutable,equality-func=Equal

// Type is just a dummy type for demonstrating a custom type example.
type Type struct {
	Name string
}

// Equal is just a dummy func for demonstrating a custom equals func.
func Equal(a, b *Type) bool {
	return a == b
}
