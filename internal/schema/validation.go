package schema

import (
	"github.com/pkg/errors"

	"k8s.io/gengo/types"
)

var (
	errKindMismatch      = errors.New("kinds do not match")
	errNameMismatch      = errors.New("names do not match")
	errParameterMismatch = errors.New("func parameters do not match")
	errResultMismatch    = errors.New("func results do not match")
	errVariadic          = errors.New("one of the funcs is variadic, the other is not")
	errNilType           = errors.New("one of the types is nil")
)

// ValidateType validates a given type against an expected schema. Returns an
// error if the schema of given differs from expected.
func ValidateType(expected, given *types.Type) error {
	return validateType(expected, given)
}

func validateType(a, b *types.Type) error {
	if a == b {
		return nil
	}

	if a == nil || b == nil {
		return errNilType
	}

	if a.Kind != b.Kind {
		return errors.Wrapf(errKindMismatch, "%s != %s", a.Kind, b.Kind)
	}

	if a.Kind != types.Func && a.Name != b.Name {
		return errors.Wrapf(errNameMismatch, "%s != %s", a.Name, b.Name)
	}

	switch a.Kind {
	case types.Slice, types.Map, types.Pointer, types.Chan:
		return validateType(a.Elem, b.Elem)
	case types.Alias, types.DeclarationOf:
		return validateType(a.Underlying, b.Underlying)
	case types.Func:
		return validateSignature(a.Signature, b.Signature)
	}

	return nil
}

func validateSignature(a, b *types.Signature) error {
	if a.Variadic != b.Variadic {
		return errors.Wrapf(errVariadic, "expected variadic=%v, got %v", a.Variadic, b.Variadic)
	}

	if len(a.Parameters) != len(b.Parameters) {
		return errors.Wrapf(errParameterMismatch, "expected %d, got %d", len(a.Parameters), len(b.Parameters))
	}

	if len(a.Results) != len(b.Results) {
		return errors.Wrapf(errResultMismatch, "expected %d, got %d", len(a.Results), len(b.Results))
	}

	for i, param := range a.Parameters {
		err := validateType(param, b.Parameters[i])
		if err != nil {
			return err
		}
	}

	for i, result := range a.Results {
		err := validateType(result, b.Results[i])
		if err != nil {
			return err
		}
	}

	return nil
}
