package ir

import "fmt"

// OpType represents operation types
type OpType int

const (
	// Arithmetic operations
	OpAdd OpType = iota
	OpSub
	OpMul
	OpDiv
	
	// Comparison operations
	OpEQ  // ==
	OpNE  // !=
	OpLT  // <
	OpLE  // <=
	OpGT  // >
	OpGE  // >=
	
	// Logical operations
	OpOr  // ||
	OpAnd // &&
	
	// Custom binary operations (for function-like operators)
	OpCustom
)

func (op OpType) String() string {
	switch op {
	case OpAdd:
		return "+"
	case OpSub:
		return "-"
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	case OpEQ:
		return "=="
	case OpNE:
		return "!="
	case OpLT:
		return "<"
	case OpLE:
		return "<="
	case OpGT:
		return ">"
	case OpGE:
		return ">="
	case OpOr:
		return "||"
	case OpAnd:
		return "&&"
	case OpCustom:
		return "custom"
	default:
		return fmt.Sprintf("op(%d)", int(op))
	}
}

// IsArithmetic returns true if the operation is arithmetic
func (op OpType) IsArithmetic() bool {
	return op >= OpAdd && op <= OpDiv
}

// IsComparison returns true if the operation is comparison
func (op OpType) IsComparison() bool {
	return op >= OpEQ && op <= OpGE
}

// ParseOp converts string operator to OpType
func ParseOp(op string) OpType {
	switch op {
	case "+":
		return OpAdd
	case "-":
		return OpSub
	case "*":
		return OpMul
	case "/":
		return OpDiv
	case "==":
		return OpEQ
	case "!=":
		return OpNE
	case "<":
		return OpLT
	case "<=":
		return OpLE
	case ">":
		return OpGT
	case ">=":
		return OpGE
	case "||":
		return OpOr
	case "&&":
		return OpAnd
	default:
		return OpCustom
	}
}