package semantic

import "fmt"

// Type represents the type of a value in mlang
type Type interface {
	String() string
	Equals(Type) bool
}

// BasicType represents primitive types
type BasicType struct {
	Name string
}

func (b *BasicType) String() string {
	return b.Name
}

func (b *BasicType) Equals(other Type) bool {
	if o, ok := other.(*BasicType); ok {
		return b.Name == o.Name
	}
	return false
}

// ArrayType represents array types
type ArrayType struct {
	ElementType Type
}

func (a *ArrayType) String() string {
	return fmt.Sprintf("[%s]", a.ElementType.String())
}

func (a *ArrayType) Equals(other Type) bool {
	if o, ok := other.(*ArrayType); ok {
		return a.ElementType.Equals(o.ElementType)
	}
	return false
}

// DictType represents dictionary types
type DictType struct {
	KeyType   Type
	ValueType Type
}

func (d *DictType) String() string {
	return fmt.Sprintf("{%s: %s}", d.KeyType.String(), d.ValueType.String())
}

func (d *DictType) Equals(other Type) bool {
	if o, ok := other.(*DictType); ok {
		return d.KeyType.Equals(o.KeyType) && d.ValueType.Equals(o.ValueType)
	}
	return false
}

// FunctionType represents function types
type FunctionType struct {
	ParamTypes []Type
	ReturnType Type
}

func (f *FunctionType) String() string {
	params := "("
	for i, p := range f.ParamTypes {
		if i > 0 {
			params += ", "
		}
		params += p.String()
	}
	params += ")"
	return fmt.Sprintf("%s -> %s", params, f.ReturnType.String())
}

func (f *FunctionType) Equals(other Type) bool {
	if o, ok := other.(*FunctionType); ok {
		if len(f.ParamTypes) != len(o.ParamTypes) {
			return false
		}
		for i, p := range f.ParamTypes {
			if !p.Equals(o.ParamTypes[i]) {
				return false
			}
		}
		return f.ReturnType.Equals(o.ReturnType)
	}
	return false
}

// Predefined types
var (
	NumberType  = &BasicType{"number"}
	StringType  = &BasicType{"string"}
	BooleanType = &BasicType{"boolean"}
	VoidType    = &BasicType{"void"}
	AnyType     = &BasicType{"any"}
)