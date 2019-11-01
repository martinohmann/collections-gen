package schema

import (
	"k8s.io/gengo/types"
)

// Collection returns the type schema for a collection with given name and
// element type.
func Collection(name types.Name, elemType *types.Type) *types.Type {
	sliceType := sliceType(elemType)
	collectionType := collectionType(name)
	collectionType.Methods = collectionMethods(collectionType, sliceType, elemType)

	return collectionType
}

func collectionType(name types.Name) *types.Type {
	return &types.Type{
		Kind: types.Pointer,
		Elem: &types.Type{
			Kind: types.Struct,
			Name: name,
		},
		Name: types.Name{
			Name: "*" + name.String(),
		},
	}
}

func sliceType(elemType *types.Type) *types.Type {
	elemName := elemType.Name

	return &types.Type{
		Kind: types.Slice,
		Name: types.Name{Name: "[]" + elemName.String()},
		Elem: elemType,
	}
}

func collectionMethods(collectionType, sliceType, elemType *types.Type) map[string]*types.Type {
	return map[string]*types.Type{
		"All": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					types.Bool,
				},
			},
		},
		"Any": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					types.Bool,
				},
			},
		},
		"Append": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Variadic: true,
				Parameters: []*types.Type{
					sliceType,
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Cap": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					types.Int,
				},
			},
		},
		"Collect": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Contains": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					elemType,
				},
				Results: []*types.Type{
					types.Bool,
				},
			},
		},
		"Copy": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Cut": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
					types.Int,
				},
				Results: []*types.Type{
					sliceType,
				},
			},
		},
		"Each": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
						},
					},
				},
			},
		},
		"EachIndex": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
								types.Int,
							},
						},
					},
				},
			},
		},
		"Filter": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Find": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					elemType,
				},
			},
		},
		"FindOk": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					elemType,
					types.Bool,
				},
			},
		},
		"First": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					elemType,
				},
			},
		},
		"FirstN": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
				},
				Results: []*types.Type{
					sliceType,
				},
			},
		},
		"Get": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
				},
				Results: []*types.Type{
					elemType,
				},
			},
		},
		"IndexOf": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					elemType,
				},
				Results: []*types.Type{
					types.Int,
				},
			},
		},
		"InsertItem": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					elemType,
					types.Int,
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Interface": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					{
						Kind: types.Interface,
						Name: types.Name{Name: "interface{}"},
					},
				},
			},
		},
		"IsSorted": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					types.Bool,
				},
			},
		},
		"Items": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					sliceType,
				},
			},
		},
		"Last": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					elemType,
				},
			},
		},
		"LastN": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
				},
				Results: []*types.Type{
					sliceType,
				},
			},
		},
		"Len": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					types.Int,
				},
			},
		},
		"Map": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								elemType,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"MapIndex": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
								types.Int,
							},
							Results: []*types.Type{
								elemType,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Nth": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
				},
				Results: []*types.Type{
					elemType,
				},
			},
		},
		"Partition": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
					collectionType,
				},
			},
		},
		"Prepend": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Variadic: true,
				Parameters: []*types.Type{
					sliceType,
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Reduce": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
								elemType,
							},
							Results: []*types.Type{
								elemType,
							},
						},
					},
				},
				Results: []*types.Type{
					elemType,
				},
			},
		},
		"Reject": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Remove": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"RemoveItem": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					elemType,
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Reverse": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Results: []*types.Type{
					collectionType,
				},
			},
		},
		"Slice": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					types.Int,
					types.Int,
				},
				Results: []*types.Type{
					sliceType,
				},
			},
		},
		"Sort": &types.Type{
			Kind: types.Func,
			Signature: &types.Signature{
				Parameters: []*types.Type{
					{
						Kind: types.Func,
						Signature: &types.Signature{
							Parameters: []*types.Type{
								elemType,
								elemType,
							},
							Results: []*types.Type{
								types.Bool,
							},
						},
					},
				},
				Results: []*types.Type{
					collectionType,
				},
			},
		},
	}
}
