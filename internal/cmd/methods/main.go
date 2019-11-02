package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/martinohmann/collections-gen/internal/schema"
	"k8s.io/gengo/types"
)

var (
	name     = flag.String("name", "CollectionType", "Name of the collection type")
	elemName = flag.String("elem-name", "ElemType", "Name of the element type")
)

func main() {
	flag.Parse()

	name := types.Name{Name: *name}
	elemType := &types.Type{Name: types.Name{Name: *elemName}, Kind: types.Struct}

	schema := schema.Collection(name, elemType)

	methodNames := make([]string, 0, len(schema.Methods))

	for methodName := range schema.Methods {
		methodNames = append(methodNames, methodName)
	}

	sort.Strings(methodNames)

	for _, methodName := range methodNames {
		m := schema.Methods[methodName]

		fmt.Printf("func (*%s) %s%s\n", name, methodName, funcSignature(m))
	}
}

func funcSignature(t *types.Type) string {
	return fmt.Sprintf("(%s)%s", funcParameters(t.Signature), funcResults(t.Signature))
}

func funcParameters(s *types.Signature) string {
	names := make([]string, len(s.Parameters))
	for i, p := range s.Parameters {
		if p.Kind == types.Func {
			names[i] = "func" + funcSignature(p)
		} else if i == len(s.Parameters)-1 && s.Variadic {
			names[i] = "..." + p.String()[2:]
		} else {
			names[i] = p.String()
		}
	}

	return strings.Join(names, ", ")
}

func funcResults(s *types.Signature) string {
	names := make([]string, len(s.Results))
	for i, r := range s.Results {
		if r.Kind == types.Func {
			names[i] = "func" + funcSignature(r)
		} else {
			names[i] = r.String()
		}
	}

	out := strings.Join(names, ", ")

	if strings.Contains(out, " ") {
		out = "(" + out + ")"
	}

	if len(out) > 0 {
		out = " " + out
	}

	return out
}
