package ir

import (
	"fmt"
	
	"github.com/hatlonely/mlang/go/semantic"
)

// IRNode represents a node in the intermediate representation
type IRNode interface {
	Type() semantic.Type
	String() string
}

// IRExpr represents an expression node
type IRExpr interface {
	IRNode
	isExpr()
}

// IRStmt represents a statement node
type IRStmt interface {
	IRNode
	isStmt()
}

// Literal nodes
type IntLiteral struct {
	Value int64
}

func (i *IntLiteral) Type() semantic.Type { return semantic.IntType }
func (i *IntLiteral) String() string      { return fmt.Sprintf("int(%d)", i.Value) }
func (i *IntLiteral) isExpr()             {}

type FloatLiteral struct {
	Value float64
}

func (f *FloatLiteral) Type() semantic.Type { return semantic.FloatType }
func (f *FloatLiteral) String() string      { return fmt.Sprintf("float(%g)", f.Value) }
func (f *FloatLiteral) isExpr()             {}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) Type() semantic.Type { return semantic.StringType }
func (s *StringLiteral) String() string      { return fmt.Sprintf("string(%q)", s.Value) }
func (s *StringLiteral) isExpr()             {}

type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) Type() semantic.Type { return semantic.BooleanType }
func (b *BooleanLiteral) String() string      { return fmt.Sprintf("bool(%t)", b.Value) }
func (b *BooleanLiteral) isExpr()             {}

// Variable reference
type Variable struct {
	Name      string
	ValueType semantic.Type
}

func (v *Variable) Type() semantic.Type { return v.ValueType }
func (v *Variable) String() string      { return fmt.Sprintf("var(%s:%s)", v.Name, v.ValueType) }
func (v *Variable) isExpr()             {}

// Field access
type FieldAccess struct {
	Object    IRExpr
	FieldName string
	FieldType semantic.Type
}

func (f *FieldAccess) Type() semantic.Type { return f.FieldType }
func (f *FieldAccess) String() string      { return fmt.Sprintf("field(%s.%s:%s)", f.Object, f.FieldName, f.FieldType) }
func (f *FieldAccess) isExpr()             {}

// Binary operation
type BinaryOp struct {
	Left       IRExpr
	Right      IRExpr
	Op         OpType
	ResultType semantic.Type
	LeftCast   semantic.Type // nil if no cast needed
	RightCast  semantic.Type // nil if no cast needed
}

func (b *BinaryOp) Type() semantic.Type { return b.ResultType }
func (b *BinaryOp) String() string {
	return fmt.Sprintf("binop(%s %s %s:%s)", b.Left, b.Op, b.Right, b.ResultType)
}
func (b *BinaryOp) isExpr() {}

// Function call
type FunctionCall struct {
	Name       string
	Args       []IRExpr
	ResultType semantic.Type
	ArgCasts   []semantic.Type // nil if no cast needed for that arg
	IsBuiltin  bool
	IsVariadic bool
}

func (f *FunctionCall) Type() semantic.Type { return f.ResultType }
func (f *FunctionCall) String() string {
	return fmt.Sprintf("call(%s(...): %s)", f.Name, f.ResultType)
}
func (f *FunctionCall) isExpr() {}

// Array literal
type ArrayLiteral struct {
	Elements    []IRExpr
	ElementType semantic.Type
	ElementCasts []semantic.Type // nil if no cast needed for that element
}

func (a *ArrayLiteral) Type() semantic.Type {
	return &semantic.ArrayType{ElementType: a.ElementType}
}
func (a *ArrayLiteral) String() string {
	return fmt.Sprintf("array([%d]%s)", len(a.Elements), a.ElementType)
}
func (a *ArrayLiteral) isExpr() {}

// Dictionary literal
type DictLiteral struct {
	Keys       []IRExpr
	Values     []IRExpr
	KeyType    semantic.Type
	ValueType  semantic.Type
	KeyCasts   []semantic.Type // nil if no cast needed
	ValueCasts []semantic.Type // nil if no cast needed
}

func (d *DictLiteral) Type() semantic.Type {
	return &semantic.DictType{KeyType: d.KeyType, ValueType: d.ValueType}
}
func (d *DictLiteral) String() string {
	return fmt.Sprintf("dict({%s: %s})", d.KeyType, d.ValueType)
}
func (d *DictLiteral) isExpr() {}

// Expression statement
type ExprStmt struct {
	Expr IRExpr
}

func (e *ExprStmt) Type() semantic.Type { return e.Expr.Type() }
func (e *ExprStmt) String() string      { return fmt.Sprintf("expr_stmt(%s)", e.Expr) }
func (e *ExprStmt) isStmt()             {}

// Program represents the entire program
type Program struct {
	Statements []IRStmt
	ResultType semantic.Type // Type of the last expression
}

func (p *Program) Type() semantic.Type { return p.ResultType }
func (p *Program) String() string      { return fmt.Sprintf("program(%d stmts)", len(p.Statements)) }