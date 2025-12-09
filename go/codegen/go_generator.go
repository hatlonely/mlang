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

// GenerateProgram generates Go code from program IR
func (g *GoGenerator) GenerateProgram(program *ir.Program) (string, error) {
	if len(program.Statements) == 0 {
		return "", fmt.Errorf("empty program")
	}
	
	// Get the last statement
	lastStmt := program.Statements[len(program.Statements)-1]
	
	switch stmt := lastStmt.(type) {
	case *ir.ExprStmt:
		// For expression statements, return the expression
		return g.generateExpression(stmt.Expr)
	case *ir.AssignmentStmt:
		// For assignment statements, generate the assignment
		return g.generateAssignment(stmt)
	default:
		return "", fmt.Errorf("unsupported statement type: %T", stmt)
	}
}

// generateAssignment generates Go code for assignment statement
func (g *GoGenerator) generateAssignment(stmt *ir.AssignmentStmt) (string, error) {
	// Check if the left-hand side is a property variable
	if varLvalue, ok := stmt.Lvalue.(*ir.VariableLvalue); ok {
		if propType, ok := varLvalue.ValueType.(*semantic.PropertyType); ok {
			// Generate right-hand side
			value, err := g.generateExpression(stmt.Value)
			if err != nil {
				return "", err
			}
			
			// Apply cast if needed
			if stmt.ValueCast != nil {
				value = g.generateCast(value, stmt.ValueCast)
			}
			
			// For property variables, generate setter function call
			return fmt.Sprintf("%s(ds, %s)", propType.Setter, value), nil
		}
	}
	
	// Generate left-hand side for normal variables
	lvalue, err := g.generateLvalue(stmt.Lvalue)
	if err != nil {
		return "", err
	}
	
	// Generate right-hand side
	value, err := g.generateExpression(stmt.Value)
	if err != nil {
		return "", err
	}
	
	// Apply cast if needed
	if stmt.ValueCast != nil {
		value = g.generateCast(value, stmt.ValueCast)
	}
	
	// Generate assignment
	return fmt.Sprintf("%s = %s", lvalue, value), nil
}

// generateLvalue generates Go code for left-value expression
func (g *GoGenerator) generateLvalue(lvalue ir.IRLvalue) (string, error) {
	switch lval := lvalue.(type) {
	case *ir.VariableLvalue:
		return lval.Name, nil
	case *ir.IndexLvalue:
		return g.generateIndexLvalue(lval)
	case *ir.FieldLvalue:
		return g.generateFieldLvalue(lval)
	default:
		return "", fmt.Errorf("unsupported lvalue type: %T", lvalue)
	}
}

// generateIndexLvalue generates Go code for index lvalue (array[index] or dict[key])
func (g *GoGenerator) generateIndexLvalue(lval *ir.IndexLvalue) (string, error) {
	// Generate object
	object, err := g.generateLvalue(lval.Object)
	if err != nil {
		return "", err
	}
	
	// Generate index
	index, err := g.generateExpression(lval.Index)
	if err != nil {
		return "", err
	}
	
	// Apply cast if needed
	if lval.IndexCast != nil {
		index = g.generateCast(index, lval.IndexCast)
	}
	
	return fmt.Sprintf("%s[%s]", object, index), nil
}

// generateFieldLvalue generates Go code for field lvalue (struct.field)
func (g *GoGenerator) generateFieldLvalue(lval *ir.FieldLvalue) (string, error) {
	// Generate object
	object, err := g.generateLvalue(lval.Object)
	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("%s.%s", object, lval.FieldName), nil
}

// generateVariable generates code for variable access, handling PropertyType
func (g *GoGenerator) generateVariable(v *ir.Variable) (string, error) {
	// 检查变量类型是否是 PropertyType
	if propType, ok := v.ValueType.(*semantic.PropertyType); ok {
		// 对于属性变量，生成 getter 函数调用
		// 假设数据源对象是全局变量 "ds"
		return fmt.Sprintf("%s(ds)", propType.Getter), nil
	}
	
	// 普通变量直接返回名称
	return v.Name, nil
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
		return g.generateVariable(e)
		
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
		
	case *ir.IndexAccess:
		return g.generateIndexAccess(e)
		
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
	var result string
	switch op.Op {
	case ir.OpAdd:
		result = fmt.Sprintf("(%s + %s)", left, right)
	case ir.OpSub:
		result = fmt.Sprintf("(%s - %s)", left, right)
	case ir.OpMul:
		result = fmt.Sprintf("(%s * %s)", left, right)
	case ir.OpDiv:
		result = fmt.Sprintf("(%s / %s)", left, right)
	case ir.OpEQ:
		result = fmt.Sprintf("(%s == %s)", left, right)
	case ir.OpNE:
		result = fmt.Sprintf("(%s != %s)", left, right)
	case ir.OpLT:
		result = fmt.Sprintf("(%s < %s)", left, right)
	case ir.OpLE:
		result = fmt.Sprintf("(%s <= %s)", left, right)
	case ir.OpGT:
		result = fmt.Sprintf("(%s > %s)", left, right)
	case ir.OpGE:
		result = fmt.Sprintf("(%s >= %s)", left, right)
	case ir.OpOr:
		// || 运算符需要短路求值: 如果左侧不为零值，返回左侧，否则返回右侧
		result = g.generateOrExpression(left, right, op.ResultType)
	case ir.OpAnd:
		// && 运算符需要短路求值: 如果左侧为零值，返回右值的零值，否则返回右值
		result = g.generateAndExpression(left, right, op.ResultType, op.Left.Type())
	case ir.OpCustom:
		// Custom binary operators are generated as function calls
		if op.OpName == "" {
			return "", fmt.Errorf("custom binary operator missing OpName")
		}
		result = fmt.Sprintf("%s(%s, %s)", op.OpName, left, right)
	default:
		return "", fmt.Errorf("unsupported binary operator: %s", op.Op)
	}
	
	// Apply NOT if negated
	if op.Negated {
		result = fmt.Sprintf("!(%s)", result)
	}
	
	return result, nil
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

func (g *GoGenerator) generateIndexAccess(index *ir.IndexAccess) (string, error) {
	object, err := g.generateExpression(index.Object)
	if err != nil {
		return "", err
	}
	
	indexExpr, err := g.generateExpression(index.Index)
	if err != nil {
		return "", err
	}
	
	// 如果需要索引类型转换，应用转换
	if index.IndexCast != nil {
		indexExpr = g.generateCast(indexExpr, index.IndexCast)
	}
	
	// 生成索引访问代码: object[index]
	return fmt.Sprintf("%s[%s]", object, indexExpr), nil
}

// generateOrExpression 生成 || 运算符的短路求值表达式
func (g *GoGenerator) generateOrExpression(left, right string, resultType semantic.Type) string {
	// 根据结果类型生成不同的零值检查和短路逻辑
	switch typ := resultType.(type) {
	case *semantic.BasicType:
		switch typ.Name {
		case "int":
			return fmt.Sprintf("(func() int { if _left := %s; _left != 0 { return _left } else { return %s } })()", left, right)
		case "float":
			return fmt.Sprintf("(func() float64 { if _left := %s; _left != 0.0 { return _left } else { return %s } })()", left, right)
		case "string":
			return fmt.Sprintf("(func() string { if _left := %s; _left != \"\" { return _left } else { return %s } })()", left, right)
		case "boolean":
			// 对于布尔类型，直接使用Go的||运算符，它本身就支持短路求值
			return fmt.Sprintf("(%s || %s)", left, right)
		default:
			// 对于其他类型，使用interface{}和反射判断
			return fmt.Sprintf("(func() interface{} { if _left := %s; !isZeroValue(_left) { return _left } else { return %s } })()", left, right)
		}
	case *semantic.ArrayType:
		// 数组的零值是nil或长度为0
		elemType := g.generateGoType(typ.ElementType)
		return fmt.Sprintf("(func() []%s { if _left := %s; _left != nil && len(_left) > 0 { return _left } else { return %s } })()", elemType, left, right)
	case *semantic.DictType:
		// 字典的零值是nil或长度为0
		keyType := g.generateGoType(typ.KeyType)
		valueType := g.generateGoType(typ.ValueType)
		return fmt.Sprintf("(func() map[%s]%s { if _left := %s; _left != nil && len(_left) > 0 { return _left } else { return %s } })()", keyType, valueType, left, right)
	default:
		// 默认情况，使用interface{}
		return fmt.Sprintf("(func() interface{} { if _left := %s; !isZeroValue(_left) { return _left } else { return %s } })()", left, right)
	}
}

// generateAndExpression 生成 && 运算符的短路求值表达式
func (g *GoGenerator) generateAndExpression(left, right string, resultType semantic.Type, leftType semantic.Type) string {
	// 特殊优化：当左右操作数都是布尔类型时，直接使用Go的原生&&运算符
	if leftTyp, ok := leftType.(*semantic.BasicType); ok && leftTyp.Name == "boolean" {
		if resultTyp, ok := resultType.(*semantic.BasicType); ok && resultTyp.Name == "boolean" {
			return fmt.Sprintf("(%s && %s)", left, right)
		}
	}
	
	zeroValue := g.generateZeroValue(resultType)
	goType := g.generateGoType(resultType)
	
	// 关键：零值检查应该基于左操作数的类型，而不是结果类型
	var zeroCheck string
	switch leftTyp := leftType.(type) {
	case *semantic.BasicType:
		switch leftTyp.Name {
		case "int":
			zeroCheck = "_left == 0"
		case "float":
			zeroCheck = "_left == 0.0"
		case "string":
			zeroCheck = `_left == ""`
		case "boolean":
			zeroCheck = "!_left"
		default:
			// 对于未知的基础类型，使用接口检查
			zeroCheck = fmt.Sprintf("_left == nil || _left == %s", g.generateZeroValue(leftType))
		}
	case *semantic.ArrayType:
		zeroCheck = "_left == nil || len(_left) == 0"
	case *semantic.DictType:
		zeroCheck = "_left == nil || len(_left) == 0"
	default:
		// 对于复杂类型，使用接口检查
		zeroCheck = "_left == nil"
	}
	
	return fmt.Sprintf("(func() %s { if _left := %s; %s { return %s } else { return %s } })()", goType, left, zeroCheck, zeroValue, right)
}

// generateZeroValue 生成指定类型的零值
func (g *GoGenerator) generateZeroValue(t semantic.Type) string {
	switch typ := t.(type) {
	case *semantic.BasicType:
		switch typ.Name {
		case "int":
			return "0"
		case "float":
			return "0.0"
		case "string":
			return `""`
		case "boolean":
			return "false"
		default:
			return "nil"
		}
	case *semantic.ArrayType:
		elemType := g.generateGoType(typ.ElementType)
		return fmt.Sprintf("[]%s{}", elemType)
	case *semantic.DictType:
		keyType := g.generateGoType(typ.KeyType)
		valueType := g.generateGoType(typ.ValueType)
		return fmt.Sprintf("map[%s]%s{}", keyType, valueType)
	case *semantic.StructType:
		// 结构体零值需要类型信息，这里返回nil
		return "nil"
	default:
		return "nil"
	}
}