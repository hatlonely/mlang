package ir

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/antlr4-go/antlr/v4"
	"github.com/hatlonely/mlang/go/semantic"
	parser "github.com/hatlonely/mlang/gen/go"
)

// Builder converts analyzed AST to IR
type Builder struct {
	symbolTable *semantic.SymbolTable
	errors      []string
}

// NewBuilder creates a new IR builder
func NewBuilder(symbolTable *semantic.SymbolTable) *Builder {
	return &Builder{
		symbolTable: symbolTable,
		errors:      make([]string, 0),
	}
}

// BuildProgram converts program AST to IR
func (b *Builder) BuildProgram(ctx *parser.ProgContext) *Program {
	var statements []IRStmt
	var lastType semantic.Type = semantic.VoidType

	for _, statCtx := range ctx.AllStat() {
		if exprCtx := statCtx.Expr(); exprCtx != nil {
			expr := b.BuildExpression(exprCtx)
			statements = append(statements, &ExprStmt{Expr: expr})
			lastType = expr.Type()
		}
	}

	return &Program{
		Statements: statements,
		ResultType: lastType,
	}
}

// BuildExpression converts expression AST to IR
func (b *Builder) BuildExpression(ctx parser.IExprContext) IRExpr {
	switch expr := ctx.(type) {
	case *parser.NumberContext:
		return b.buildNumber(expr)
	case *parser.StringContext:
		return b.buildString(expr)
	case *parser.BooleanContext:
		return b.buildBoolean(expr)
	case *parser.IdContext:
		return b.buildVariable(expr)
	case *parser.ArrayContext:
		return b.buildArray(expr)
	case *parser.DictionaryContext:
		return b.buildDictionary(expr)
	case *parser.ParensContext:
		return b.BuildExpression(expr.Expr())
	case *parser.MulDivContext:
		return b.buildBinaryOp(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.AddSubContext:
		return b.buildBinaryOp(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.CompareSymbolContext:
		return b.buildBinaryOp(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.CompareFuncInfixContext:
		return b.buildCustomBinaryOp(expr)
	case *parser.NotCompareFuncInfixContext:
		return b.buildNotCustomBinaryOp(expr)
	case *parser.FunctionCallContext:
		return b.buildFunctionCall(expr)
	case *parser.FieldAccessContext:
		return b.buildFieldAccess(expr)
	case *parser.IndexAccessContext:
		return b.buildIndexAccess(expr)
	default:
		b.addError(ctx, fmt.Sprintf("unsupported expression type: %T", expr))
		return &IntLiteral{Value: 0} // fallback
	}
}

func (b *Builder) buildNumber(ctx *parser.NumberContext) IRExpr {
	text := ctx.GetText()
	if strings.Contains(text, ".") {
		if value, err := strconv.ParseFloat(text, 64); err == nil {
			return &FloatLiteral{Value: value}
		}
	} else {
		if value, err := strconv.ParseInt(text, 10, 64); err == nil {
			return &IntLiteral{Value: value}
		}
	}
	b.addError(ctx, "invalid number format: "+text)
	return &IntLiteral{Value: 0}
}

func (b *Builder) buildString(ctx *parser.StringContext) IRExpr {
	text := ctx.GetText()
	// Remove quotes and handle escape sequences
	if text[0] == '"' {
		// Double-quoted string with escape sequences
		value, err := strconv.Unquote(text)
		if err != nil {
			b.addError(ctx, "invalid string format: "+text)
			return &StringLiteral{Value: ""}
		}
		return &StringLiteral{Value: value}
	} else if text[0] == '`' {
		// Backtick string without escape sequences
		return &StringLiteral{Value: text[1 : len(text)-1]}
	}
	b.addError(ctx, "invalid string format: "+text)
	return &StringLiteral{Value: ""}
}

func (b *Builder) buildBoolean(ctx *parser.BooleanContext) IRExpr {
	text := strings.ToLower(ctx.GetText())
	return &BooleanLiteral{Value: text == "true"}
}

func (b *Builder) buildVariable(ctx *parser.IdContext) IRExpr {
	name := ctx.GetText()
	symbol, exists := b.symbolTable.Lookup(name)
	if !exists {
		b.addError(ctx, "undefined identifier: "+name)
		return &Variable{Name: name, ValueType: semantic.AnyType}
	}
	return &Variable{Name: name, ValueType: symbol.Type}
}

func (b *Builder) buildArray(ctx *parser.ArrayContext) IRExpr {
	if ctx.ArrayElements() == nil {
		return &ArrayLiteral{
			Elements:    []IRExpr{},
			ElementType: semantic.AnyType,
		}
	}

	elements := ctx.ArrayElements().AllExpr()
	if len(elements) == 0 {
		return &ArrayLiteral{
			Elements:    []IRExpr{},
			ElementType: semantic.AnyType,
		}
	}

	// Build all elements
	irElements := make([]IRExpr, len(elements))
	for i, elemCtx := range elements {
		irElements[i] = b.BuildExpression(elemCtx)
	}

	// Determine unified element type (first element's type)
	elementType := irElements[0].Type()
	
	// Check if type promotion is needed
	elementCasts := make([]semantic.Type, len(irElements))
	for i, elem := range irElements {
		elemType := elem.Type()
		if !elemType.Equals(elementType) {
			// Determine if we can promote types
			if promotedType := b.promoteTypes(elementType, elemType); promotedType != nil {
				elementType = promotedType
				// Mark elements that need casting
				for j := 0; j <= i; j++ {
					if !irElements[j].Type().Equals(elementType) {
						elementCasts[j] = elementType
					}
				}
			}
		}
	}

	return &ArrayLiteral{
		Elements:     irElements,
		ElementType:  elementType,
		ElementCasts: elementCasts,
	}
}

func (b *Builder) buildDictionary(ctx *parser.DictionaryContext) IRExpr {
	if ctx.DictElements() == nil {
		return &DictLiteral{
			Keys:      []IRExpr{},
			Values:    []IRExpr{},
			KeyType:   semantic.AnyType,
			ValueType: semantic.AnyType,
		}
	}

	pairs := ctx.DictElements().AllDictPair()
	if len(pairs) == 0 {
		return &DictLiteral{
			Keys:      []IRExpr{},
			Values:    []IRExpr{},
			KeyType:   semantic.AnyType,
			ValueType: semantic.AnyType,
		}
	}

	keys := make([]IRExpr, len(pairs))
	values := make([]IRExpr, len(pairs))
	
	for i, pair := range pairs {
		keys[i] = b.BuildExpression(pair.Expr(0))
		values[i] = b.BuildExpression(pair.Expr(1))
	}

	// Determine unified types
	keyType := keys[0].Type()
	valueType := values[0].Type()
	
	keyCasts := make([]semantic.Type, len(keys))
	valueCasts := make([]semantic.Type, len(values))
	
	// Type promotion for keys and values
	for i := 1; i < len(keys); i++ {
		if promotedType := b.promoteTypes(keyType, keys[i].Type()); promotedType != nil {
			keyType = promotedType
		}
		if promotedType := b.promoteTypes(valueType, values[i].Type()); promotedType != nil {
			valueType = promotedType
		}
	}
	
	// Mark elements that need casting
	for i := 0; i < len(keys); i++ {
		if !keys[i].Type().Equals(keyType) {
			keyCasts[i] = keyType
		}
		if !values[i].Type().Equals(valueType) {
			valueCasts[i] = valueType
		}
	}

	return &DictLiteral{
		Keys:       keys,
		Values:     values,
		KeyType:    keyType,
		ValueType:  valueType,
		KeyCasts:   keyCasts,
		ValueCasts: valueCasts,
	}
}

func (b *Builder) buildBinaryOp(leftCtx, rightCtx parser.IExprContext, opStr string) IRExpr {
	left := b.BuildExpression(leftCtx)
	right := b.BuildExpression(rightCtx)
	op := ParseOp(opStr)
	
	leftType := left.Type()
	rightType := right.Type()
	
	// Determine result type and necessary casts
	var resultType semantic.Type
	var leftCast, rightCast semantic.Type
	
	if op.IsArithmetic() {
		// Arithmetic operations: promote numeric types
		resultType = b.promoteNumericTypes(leftType, rightType)
		if resultType != nil {
			if !leftType.Equals(resultType) {
				leftCast = resultType
			}
			if !rightType.Equals(resultType) {
				rightCast = resultType
			}
		} else {
			resultType = semantic.AnyType
		}
	} else if op.IsComparison() {
		// Comparison operations always return boolean
		resultType = semantic.BooleanType
		// For numeric comparisons, promote to common type
		if b.isNumericType(leftType) && b.isNumericType(rightType) {
			promotedType := b.promoteNumericTypes(leftType, rightType)
			if promotedType != nil {
				if !leftType.Equals(promotedType) {
					leftCast = promotedType
				}
				if !rightType.Equals(promotedType) {
					rightCast = promotedType
				}
			}
		}
	}
	
	return &BinaryOp{
		Left:       left,
		Right:      right,
		Op:         op,
		ResultType: resultType,
		LeftCast:   leftCast,
		RightCast:  rightCast,
	}
}

func (b *Builder) buildCustomBinaryOp(ctx *parser.CompareFuncInfixContext) IRExpr {
	funcName := ctx.BinaryOp().GetText()
	left := b.BuildExpression(ctx.Expr(0))
	right := b.BuildExpression(ctx.Expr(1))
	
	// Try to resolve binary operator overload
	symbol, err := b.symbolTable.ResolveBinaryOpOverload(funcName, left.Type(), right.Type())
	if err != nil {
		b.addError(ctx, fmt.Sprintf("cannot resolve binary operator %s for types (%s, %s): %v", funcName, left.Type(), right.Type(), err))
		return &BinaryOp{
			Left:       left,
			Right:      right,
			Op:         OpCustom,
			ResultType: semantic.BooleanType,
		}
	}
	
	binaryOpType, ok := symbol.Type.(*semantic.BinaryOpType)
	if !ok {
		b.addError(ctx, funcName+" is not a binary operator")
		return &BinaryOp{
			Left:       left,
			Right:      right,
			Op:         OpCustom,
			ResultType: semantic.BooleanType,
		}
	}
	
	// No type casts needed - the overload resolution should have found an exact or compatible match
	return &FunctionCall{
		Name:       funcName,
		Args:       []IRExpr{left, right},
		ResultType: binaryOpType.ReturnType,
		ArgCasts:   []semantic.Type{nil, nil}, // No casting
		IsBuiltin:  false,
		IsVariadic: false,
	}
}

func (b *Builder) buildNotCustomBinaryOp(ctx *parser.NotCompareFuncInfixContext) IRExpr {
	funcName := ctx.BinaryOp().GetText()
	left := b.BuildExpression(ctx.Expr(0))
	right := b.BuildExpression(ctx.Expr(1))
	
	// Try to resolve binary operator overload
	symbol, err := b.symbolTable.ResolveBinaryOpOverload(funcName, left.Type(), right.Type())
	if err != nil {
		b.addError(ctx, fmt.Sprintf("cannot resolve binary operator %s for types (%s, %s): %v", funcName, left.Type(), right.Type(), err))
		return &BinaryOp{
			Left:       left,
			Right:      right,
			Op:         OpCustom,
			OpName:     funcName,
			ResultType: semantic.BooleanType,
			Negated:    true,
		}
	}
	
	binaryOpType, ok := symbol.Type.(*semantic.BinaryOpType)
	if !ok {
		b.addError(ctx, funcName+" is not a binary operator")
		return &BinaryOp{
			Left:       left,
			Right:      right,
			Op:         OpCustom,
			OpName:     funcName,
			ResultType: semantic.BooleanType,
			Negated:    true,
		}
	}
	
	// Check that the binary operator returns boolean type (required for NOT)
	if binaryOpType.ReturnType != semantic.BooleanType {
		b.addError(ctx, fmt.Sprintf("NOT operator can only be applied to binary operators that return boolean, but %s returns %s", funcName, binaryOpType.ReturnType))
		return &BinaryOp{
			Left:       left,
			Right:      right,
			Op:         OpCustom,
			OpName:     funcName,
			ResultType: semantic.BooleanType,
			Negated:    true,
		}
	}
	
	// Create a BinaryOp node with Negated=true to represent "NOT binaryOp"
	return &BinaryOp{
		Left:       left,
		Right:      right,
		Op:         OpCustom,
		OpName:     funcName,
		ResultType: semantic.BooleanType, // NOT always returns boolean
		Negated:    true,
	}
}

func (b *Builder) buildFunctionCall(ctx *parser.FunctionCallContext) IRExpr {
	funcName := ctx.Func_().GetText()
	
	// Build arguments
	var args []IRExpr
	if ctx.ExprList() != nil {
		for _, exprCtx := range ctx.ExprList().AllExpr() {
			args = append(args, b.BuildExpression(exprCtx))
		}
	}
	
	// Get argument types for overload resolution
	argTypes := make([]semantic.Type, len(args))
	for i, arg := range args {
		argTypes[i] = arg.Type()
	}
	
	// Try to resolve function overload
	symbol, err := b.symbolTable.ResolveFunctionOverload(funcName, argTypes)
	if err != nil {
		// Fallback to old lookup method for backward compatibility
		var exists bool
		symbol, exists = b.symbolTable.Lookup(funcName)
		if !exists {
			b.addError(ctx, "undefined function: "+funcName)
			return &FunctionCall{
				Name:       funcName,
				Args:       args,
				ResultType: semantic.AnyType,
			}
		}
	}
	
	funcType, ok := symbol.Type.(*semantic.FunctionType)
	if !ok {
		b.addError(ctx, funcName+" is not a function")
		return &FunctionCall{
			Name:       funcName,
			Args:       args,
			ResultType: semantic.AnyType,
		}
	}
	
	// Determine argument casts with smart type inference
	argCasts := make([]semantic.Type, len(args))
	
	// Fixed parameters
	for i := 0; i < len(funcType.ParamTypes) && i < len(args); i++ {
		expectedType := funcType.ParamTypes[i]
		argType := args[i].Type()
		if !argType.Equals(expectedType) && !expectedType.Equals(semantic.AnyType) {
			// Check if conversion is possible
			if b.canCastTo(argType, expectedType) {
				argCasts[i] = expectedType
			}
		}
	}
	
	// Variadic parameters with smart type inference
	if funcType.IsVariadic && len(args) > len(funcType.ParamTypes) {
		expectedType := funcType.VariadicType
		
		// If the expected type is still NumericBaseType, infer the concrete type
		if expectedType.Equals(semantic.NumericBaseType) {
			inferredType := b.inferConcreteNumericType(args[len(funcType.ParamTypes):])
			if inferredType != nil {
				expectedType = inferredType
			}
		}
		
		for i := len(funcType.ParamTypes); i < len(args); i++ {
			argType := args[i].Type()
			if !argType.Equals(expectedType) && !expectedType.Equals(semantic.AnyType) {
				if b.canCastTo(argType, expectedType) {
					argCasts[i] = expectedType
				}
			}
		}
	}
	
	// Update result type if it was NumericBaseType
	resultType := funcType.ReturnType
	if resultType.Equals(semantic.NumericBaseType) && funcType.IsVariadic {
		if len(args) > len(funcType.ParamTypes) {
			inferredType := b.inferConcreteNumericType(args[len(funcType.ParamTypes):])
			if inferredType != nil {
				resultType = inferredType
			}
		}
	}
	
	return &FunctionCall{
		Name:       funcName,
		Args:       args,
		ResultType: resultType,
		ArgCasts:   argCasts,
		IsBuiltin:  false, // 没有内置函数，都是用户注册的
		IsVariadic: funcType.IsVariadic,
	}
}

// Helper methods

func (b *Builder) addError(ctx antlr.ParserRuleContext, message string) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	b.errors = append(b.errors, fmt.Sprintf("IR building error at line %d, column %d: %s", line, column, message))
}

func (b *Builder) GetErrors() []string {
	return b.errors
}

func (b *Builder) isNumericType(t semantic.Type) bool {
	if bt, ok := t.(*semantic.BasicType); ok {
		return bt.Name == "int" || bt.Name == "float"
	}
	return false
}

func (b *Builder) isNumericOrNumericBase(t semantic.Type) bool {
	if bt, ok := t.(*semantic.BasicType); ok {
		return bt.Name == "int" || bt.Name == "float"
	}
	if nt, ok := t.(*semantic.NumericType); ok {
		return nt.Name == "numeric"
	}
	return false
}

func (b *Builder) promoteNumericTypes(left, right semantic.Type) semantic.Type {
	if !b.isNumericOrNumericBase(left) || !b.isNumericOrNumericBase(right) {
		return nil
	}
	
	// Handle NumericType (from functions that return numeric)
	if _, ok := left.(*semantic.NumericType); ok {
		left = semantic.IntType // Default numeric to int
	}
	if _, ok := right.(*semantic.NumericType); ok {
		right = semantic.IntType // Default numeric to int
	}
	
	leftBasic := left.(*semantic.BasicType)
	rightBasic := right.(*semantic.BasicType)
	
	// If either operand is float, result is float
	if leftBasic.Name == "float" || rightBasic.Name == "float" {
		return semantic.FloatType
	}
	
	// Both are int
	return semantic.IntType
}

func (b *Builder) promoteTypes(type1, type2 semantic.Type) semantic.Type {
	if type1.Equals(type2) {
		return type1
	}
	
	// Try numeric promotion
	if promoted := b.promoteNumericTypes(type1, type2); promoted != nil {
		return promoted
	}
	
	// No promotion possible
	return nil
}

func (b *Builder) canCastTo(from, to semantic.Type) bool {
	if from.Equals(to) {
		return true
	}
	
	// Numeric types can be cast to each other
	if b.isNumericType(from) && b.isNumericType(to) {
		return true
	}
	
	// Check if 'to' is a compatible supertype
	return from.IsCompatibleWith(to)
}

// inferConcreteNumericType 推断具体的数值类型
func (b *Builder) inferConcreteNumericType(args []IRExpr) semantic.Type {
	if len(args) == 0 {
		return semantic.IntType // 默认为 int
	}
	
	hasFloat := false
	for _, arg := range args {
		argType := arg.Type()
		if argType.Equals(semantic.FloatType) {
			hasFloat = true
		} else if !argType.Equals(semantic.IntType) {
			// 如果有非数值类型，返回nil表示无法推断
			return nil
		}
	}
	
	// 如果有任何一个是float，则统一提升到float
	if hasFloat {
		return semantic.FloatType
	}
	
	// 否则保持int
	return semantic.IntType
}

func (b *Builder) buildFieldAccess(ctx *parser.FieldAccessContext) IRExpr {
	// 构建左侧表达式（对象）
	object := b.BuildExpression(ctx.Expr())
	
	// 获取字段名
	fieldName := ctx.ID().GetText()
	
	// 获取对象类型
	objectType := object.Type()
	
	// 确保是结构体类型
	structType, ok := objectType.(*semantic.StructType)
	if !ok {
		b.addError(ctx, fmt.Sprintf("field access on non-struct type: %s", objectType))
		return &IntLiteral{Value: 0} // fallback
	}
	
	// 获取字段类型
	fieldType, exists := structType.GetFieldType(fieldName)
	if !exists {
		b.addError(ctx, fmt.Sprintf("field '%s' does not exist in struct type '%s'", fieldName, structType.Name))
		return &IntLiteral{Value: 0} // fallback
	}
	
	return &FieldAccess{
		Object:    object,
		FieldName: fieldName,
		FieldType: fieldType,
	}
}

func (b *Builder) buildIndexAccess(ctx *parser.IndexAccessContext) IRExpr {
	// 构建被索引的对象表达式
	object := b.BuildExpression(ctx.Expr(0))
	
	// 构建索引表达式
	index := b.BuildExpression(ctx.Expr(1))
	
	// 获取对象和索引的类型
	objectType := object.Type()
	indexType := index.Type()
	
	var resultType semantic.Type
	var indexCast semantic.Type
	
	// 根据对象类型确定结果类型和可能的索引类型转换
	switch objType := objectType.(type) {
	case *semantic.ArrayType:
		// 数组索引必须是整数
		if !indexType.Equals(semantic.IntType) {
			if b.canCastTo(indexType, semantic.IntType) {
				indexCast = semantic.IntType
			} else {
				b.addError(ctx, fmt.Sprintf("array index must be int, got %s", indexType))
				return &IntLiteral{Value: 0} // fallback
			}
		}
		resultType = objType.ElementType
		
	case *semantic.DictType:
		// 字典索引必须与键类型匹配
		if !indexType.Equals(objType.KeyType) {
			if b.canCastTo(indexType, objType.KeyType) {
				indexCast = objType.KeyType
			} else {
				b.addError(ctx, fmt.Sprintf("dictionary index type mismatch: expected %s, got %s", objType.KeyType, indexType))
				return &IntLiteral{Value: 0} // fallback
			}
		}
		resultType = objType.ValueType
		
	default:
		b.addError(ctx, fmt.Sprintf("index access not supported on type %s", objectType))
		return &IntLiteral{Value: 0} // fallback
	}
	
	return &IndexAccess{
		Object:     object,
		Index:      index,
		ResultType: resultType,
		IndexCast:  indexCast,
	}
}

