package codegen

import (
	"fmt"
	"strings"
	
	"github.com/hatlonely/mlang/go/ir"
	"github.com/hatlonely/mlang/go/semantic"
)

// GoGenerator generates Go expressions from IR
type GoGenerator struct{}

// NewGoGenerator creates a new Go expression generator
func NewGoGenerator() *GoGenerator {
	return &GoGenerator{}
}

// GenerateExpression generates a Go expression from IR expression
func (g *GoGenerator) GenerateExpression(expr ir.IRExpr) (string, error) {
	return g.generateExpression(expr)
}

// GenerateProgram generates Go expression from the last statement in a program
func (g *GoGenerator) GenerateProgram(program *ir.Program) (string, error) {
	if len(program.Statements) == 0 {
		return "", fmt.Errorf("empty program")
	}
	
	// Get the last statement
	lastStmt := program.Statements[len(program.Statements)-1]
	if exprStmt, ok := lastStmt.(*ir.ExprStmt); ok {
		return g.generateExpression(exprStmt.Expr)
	}
	
	return "", fmt.Errorf("last statement is not an expression")
}

func (g *GoGenerator) generateExpression(expr ir.IRExpr) (string, error) {
	switch e := expr.(type) {
	case *ir.IntLiteral:
		return fmt.Sprintf("%d", e.Value), nil
		
	case *ir.FloatLiteral:
		return fmt.Sprintf("%g", e.Value), nil
		
	case *ir.StringLiteral:
		return fmt.Sprintf("%q", e.Value), nil
		
	case *ir.BooleanLiteral:
		return fmt.Sprintf("%t", e.Value), nil
		
	case *ir.Variable:
		return e.Name, nil
		
	case *ir.BinaryOp:
		return g.generateBinaryOp(e)
		
	case *ir.FunctionCall:
		return g.generateFunctionCall(e)
		
	case *ir.ArrayLiteral:
		return g.generateArrayLiteral(e)
		
	case *ir.DictLiteral:
		return g.generateDictLiteral(e)
	
	case *ir.FieldAccess:
		return g.generateFieldAccess(e)
		
	default:
		return "", fmt.Errorf("unsupported expression type: %T", expr)
	}
}

func (g *GoGenerator) generateBinaryOp(op *ir.BinaryOp) (string, error) {
	left, err := g.generateExpression(op.Left)
	if err != nil {
		return "", err
	}
	
	right, err := g.generateExpression(op.Right)
	if err != nil {
		return "", err
	}
	
	// Apply casts if needed
	if op.LeftCast != nil {
		left = g.generateCast(left, op.LeftCast)
	}
	
	if op.RightCast != nil {
		right = g.generateCast(right, op.RightCast)
	}
	
	// Generate the operation
	switch op.Op {
	case ir.OpAdd:
		return fmt.Sprintf("(%s + %s)", left, right), nil
	case ir.OpSub:
		return fmt.Sprintf("(%s - %s)", left, right), nil
	case ir.OpMul:
		return fmt.Sprintf("(%s * %s)", left, right), nil
	case ir.OpDiv:
		return fmt.Sprintf("(%s / %s)", left, right), nil
	case ir.OpEQ:
		return fmt.Sprintf("(%s == %s)", left, right), nil
	case ir.OpNE:
		return fmt.Sprintf("(%s != %s)", left, right), nil
	case ir.OpLT:
		return fmt.Sprintf("(%s < %s)", left, right), nil
	case ir.OpLE:
		return fmt.Sprintf("(%s <= %s)", left, right), nil
	case ir.OpGT:
		return fmt.Sprintf("(%s > %s)", left, right), nil
	case ir.OpGE:
		return fmt.Sprintf("(%s >= %s)", left, right), nil
	default:
		return "", fmt.Errorf("unsupported binary operator: %s", op.Op)
	}
}

func (g *GoGenerator) generateFunctionCall(call *ir.FunctionCall) (string, error) {
	// 不再区分内置函数，所有函数都当作用户注册的函数
	args := make([]string, len(call.Args))
	for i, arg := range call.Args {
		argCode, err := g.generateExpression(arg)
		if err != nil {
			return "", err
		}
		
		// Apply cast if needed
		if i < len(call.ArgCasts) && call.ArgCasts[i] != nil {
			argCode = g.generateCast(argCode, call.ArgCasts[i])
		}
		
		args[i] = argCode
	}
	
	return fmt.Sprintf("%s(%s)", call.Name, strings.Join(args, ", ")), nil
}

func (g *GoGenerator) generateArrayLiteral(arr *ir.ArrayLiteral) (string, error) {
	if len(arr.Elements) == 0 {
		return "[]interface{}{}", nil
	}
	
	elements := make([]string, len(arr.Elements))
	for i, elem := range arr.Elements {
		elemCode, err := g.generateExpression(elem)
		if err != nil {
			return "", err
		}
		
		// Apply cast if needed
		if i < len(arr.ElementCasts) && arr.ElementCasts[i] != nil {
			elemCode = g.generateCast(elemCode, arr.ElementCasts[i])
		}
		
		elements[i] = elemCode
	}
	
	goType := g.generateGoType(arr.ElementType)
	return fmt.Sprintf("[]%s{%s}", goType, strings.Join(elements, ", ")), nil
}

func (g *GoGenerator) generateDictLiteral(dict *ir.DictLiteral) (string, error) {
	if len(dict.Keys) == 0 {
		return "map[interface{}]interface{}{}", nil
	}
	
	pairs := make([]string, len(dict.Keys))
	for i := range dict.Keys {
		keyCode, err := g.generateExpression(dict.Keys[i])
		if err != nil {
			return "", err
		}
		
		valueCode, err := g.generateExpression(dict.Values[i])
		if err != nil {
			return "", err
		}
		
		// Apply casts if needed
		if i < len(dict.KeyCasts) && dict.KeyCasts[i] != nil {
			keyCode = g.generateCast(keyCode, dict.KeyCasts[i])
		}
		if i < len(dict.ValueCasts) && dict.ValueCasts[i] != nil {
			valueCode = g.generateCast(valueCode, dict.ValueCasts[i])
		}
		
		pairs[i] = fmt.Sprintf("%s: %s", keyCode, valueCode)
	}
	
	keyType := g.generateGoType(dict.KeyType)
	valueType := g.generateGoType(dict.ValueType)
	return fmt.Sprintf("map[%s]%s{%s}", keyType, valueType, strings.Join(pairs, ", ")), nil
}

func (g *GoGenerator) generateCast(value string, toType semantic.Type) string {
	goType := g.generateGoType(toType)
	return fmt.Sprintf("%s(%s)", goType, value)
}

func (g *GoGenerator) generateGoType(t semantic.Type) string {
	switch typ := t.(type) {
	case *semantic.BasicType:
		switch typ.Name {
		case "int":
			return "int"
		case "float":
			return "float64"
		case "string":
			return "string"
		case "boolean":
			return "bool"
		default:
			return "interface{}"
		}
	case *semantic.ArrayType:
		elemType := g.generateGoType(typ.ElementType)
		return fmt.Sprintf("[]%s", elemType)
	case *semantic.DictType:
		keyType := g.generateGoType(typ.KeyType)
		valueType := g.generateGoType(typ.ValueType)
		return fmt.Sprintf("map[%s]%s", keyType, valueType)
	default:
		return "interface{}"
	}
}

func (g *GoGenerator) generateFieldAccess(field *ir.FieldAccess) (string, error) {
	object, err := g.generateExpression(field.Object)
	if err != nil {
		return "", err
	}
	
	// 生成字段访问代码: object.FieldName
	return fmt.Sprintf("%s.%s", object, field.FieldName), nil
}