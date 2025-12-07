package semantic

import "fmt"

// Type represents the type of a value in mlang
type Type interface {
	String() string
	Equals(Type) bool
	IsCompatibleWith(Type) bool
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

// IsCompatibleWith checks if this type is compatible with another type
func (b *BasicType) IsCompatibleWith(other Type) bool {
	// Check if other is a NumericType and this is a numeric type
	if n, ok := other.(*NumericType); ok && n.Name == "numeric" {
		return b.Name == "int" || b.Name == "float"
	}
	// Any type is compatible with AnyType
	if other.Equals(AnyType) {
		return true
	}
	return b.Equals(other)
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

func (a *ArrayType) IsCompatibleWith(other Type) bool {
	// Any array type is compatible with AnyType
	if other.Equals(AnyType) {
		return true
	}
	return a.Equals(other)
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

func (d *DictType) IsCompatibleWith(other Type) bool {
	// Any dict type is compatible with AnyType
	if other.Equals(AnyType) {
		return true
	}
	return d.Equals(other)
}

// FunctionType represents function types
type FunctionType struct {
	ParamTypes   []Type
	ReturnType   Type
	IsVariadic   bool // 是否为可变参数函数
	VariadicType Type // 可变参数的类型（当IsVariadic为true时有效）
}

func (f *FunctionType) String() string {
	params := "("
	for i, p := range f.ParamTypes {
		if i > 0 {
			params += ", "
		}
		params += p.String()
	}
	if f.IsVariadic {
		if len(f.ParamTypes) > 0 {
			params += ", "
		}
		params += "..." + f.VariadicType.String()
	}
	params += ")"
	return fmt.Sprintf("%s -> %s", params, f.ReturnType.String())
}

func (f *FunctionType) Equals(other Type) bool {
	if o, ok := other.(*FunctionType); ok {
		if len(f.ParamTypes) != len(o.ParamTypes) {
			return false
		}
		if f.IsVariadic != o.IsVariadic {
			return false
		}
		for i, p := range f.ParamTypes {
			if !p.Equals(o.ParamTypes[i]) {
				return false
			}
		}
		if f.IsVariadic && !f.VariadicType.Equals(o.VariadicType) {
			return false
		}
		return f.ReturnType.Equals(o.ReturnType)
	}
	return false
}

func (f *FunctionType) IsCompatibleWith(other Type) bool {
	return f.Equals(other)
}

// BinaryOpType represents binary operator types
type BinaryOpType struct {
	LeftType   Type
	RightType  Type
	ReturnType Type
}

func (b *BinaryOpType) String() string {
	return fmt.Sprintf("(%s, %s) -> %s", b.LeftType.String(), b.RightType.String(), b.ReturnType.String())
}

func (b *BinaryOpType) Equals(other Type) bool {
	if o, ok := other.(*BinaryOpType); ok {
		return b.LeftType.Equals(o.LeftType) && 
			   b.RightType.Equals(o.RightType) && 
			   b.ReturnType.Equals(o.ReturnType)
	}
	return false
}

func (b *BinaryOpType) IsCompatibleWith(other Type) bool {
	return b.Equals(other)
}

// NumericType represents the supertype for int and float
type NumericType struct {
	Name string
}

func (n *NumericType) String() string {
	return n.Name
}

func (n *NumericType) Equals(other Type) bool {
	if o, ok := other.(*NumericType); ok {
		return n.Name == o.Name
	}
	return false
}

// IsCompatibleWith checks if a type is compatible with this numeric type
func (n *NumericType) IsCompatibleWith(other Type) bool {
	if n.Name == "numeric" {
		if b, ok := other.(*BasicType); ok {
			return b.Name == "int" || b.Name == "float"
		}
	}
	return n.Equals(other)
}

// Predefined types
var (
	IntType     = &BasicType{"int"}
	FloatType   = &BasicType{"float"}
	NumericBaseType = &NumericType{"numeric"}
	StringType  = &BasicType{"string"}
	BooleanType = &BasicType{"boolean"}
	VoidType    = &BasicType{"void"}
	AnyType     = &BasicType{"any"}
)