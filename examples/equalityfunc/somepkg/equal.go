package somepkg

import "github.com/martinohmann/collections-gen/examples/equalityfunc"

// Equal is just a dummy func for demonstrating a custom equals func.
func Equal(a, b *equalityfunc.Foo) bool {
	return a == b
}
