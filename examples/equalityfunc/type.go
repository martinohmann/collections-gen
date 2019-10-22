package equalityfunc

//go:generate collections-gen -i github.com/martinohmann/collections-gen/examples/equalityfunc -p github.com/martinohmann/collections-gen/examples/equalityfunc/collections

// +collections-gen=true
//
// Generate FooCollection with equality func from another package.
//
// +collections-gen:options=pointer,equality-func=github.com/martinohmann/collections-gen/examples/equalityfunc/somepkg.Equal
//
// Generate BarCollection with equality func same package.
//
// +collections-gen:options=pointer,name=BarCollection,equality-func=Equal

// Type is just a dummy type for demonstrating a custom type example.
type Foo struct {
	Name string
}

// Equal is just a dummy func for demonstrating a custom equals func.
func Equal(a, b *Foo) bool {
	return a == b
}
