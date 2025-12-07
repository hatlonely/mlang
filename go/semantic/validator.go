package semantic

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hatlonely/mlang/gen/go"
)

// SemanticError represents a semantic analysis error
type SemanticError struct {
	Line    int
	Column  int
	Message string
}

func (e *SemanticError) Error() string {
	return fmt.Sprintf("semantic error at line %d, column %d: %s", e.Line, e.Column, e.Message)
}

// PureValidator performs only semantic validation without making implementation decisions
type PureValidator struct {
	symbolTable *SymbolTable
	errors      []*SemanticError
}

// NewPureValidator creates a new pure semantic validator
func NewPureValidator() *PureValidator {
	return &PureValidator{
		symbolTable: NewSymbolTable(),
		errors:      make([]*SemanticError, 0),
	}
}

// GetSymbolTable returns the symbol table
func (v *PureValidator) GetSymbolTable() *SymbolTable {
	return v.symbolTable
}

// AddError adds a semantic error
func (v *PureValidator) AddError(ctx antlr.ParserRuleContext, message string) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	v.errors = append(v.errors, &SemanticError{
		Line:    line,
		Column:  column,
		Message: message,
	})
}

// GetErrors returns all validation errors
func (v *PureValidator) GetErrors() []*SemanticError {
	return v.errors
}

// HasErrors returns true if there are validation errors
func (v *PureValidator) HasErrors() bool {
	return len(v.errors) > 0
}

// ClearErrors clears all errors
func (v *PureValidator) ClearErrors() {
	v.errors = make([]*SemanticError, 0)
}

// InferType infers the type of an expression (public interface)
func (v *PureValidator) InferType(ctx parser.IExprContext) Type {
	return v.inferExpressionType(ctx)
}

// RegisterFunction registers a function in the symbol table
func (v *PureValidator) RegisterFunction(name string, paramTypes []Type, returnType Type) error {
	return v.symbolTable.RegisterFunction(name, paramTypes, returnType)
}

// RegisterVariadicFunction registers a variadic function
func (v *PureValidator) RegisterVariadicFunction(name string, paramTypes []Type, variadicType Type, returnType Type) error {
	return v.symbolTable.RegisterVariadicFunction(name, paramTypes, variadicType, returnType)
}

// RegisterBinaryOp registers a binary operator function
func (v *PureValidator) RegisterBinaryOp(name string, leftType, rightType, returnType Type) error {
	return v.symbolTable.RegisterBinaryOp(name, leftType, rightType, returnType)
}

// RegisterVariable registers a variable in the symbol table
func (v *PureValidator) RegisterVariable(name string, varType Type) error {
	return v.symbolTable.RegisterVariable(name, varType)
}

// ValidateProgram validates the entire program
func (v *PureValidator) ValidateProgram(ctx *parser.ProgContext) bool {
	valid := true
	
	for _, statCtx := range ctx.AllStat() {
		if exprCtx := statCtx.Expr(); exprCtx != nil {
			if !v.ValidateExpression(exprCtx) {
				valid = false
			}
		}
	}
	
	return valid
}

// ValidateExpression validates an expression and returns true if valid
func (v *PureValidator) ValidateExpression(ctx parser.IExprContext) bool {
	switch expr := ctx.(type) {
	case *parser.NumberContext:
		return v.validateNumber(expr)
	case *parser.StringContext:
		return v.validateString(expr)
	case *parser.BooleanContext:
		return v.validateBoolean(expr)
	case *parser.IdContext:
		return v.validateIdentifier(expr)
	case *parser.ArrayContext:
		return v.validateArray(expr)
	case *parser.DictionaryContext:
		return v.validateDictionary(expr)
	case *parser.ParensContext:
		return v.ValidateExpression(expr.Expr())
	case *parser.MulDivContext:
		return v.validateBinaryOperation(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.AddSubContext:
		return v.validateBinaryOperation(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.CompareSymbolContext:
		return v.validateBinaryOperation(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.CompareFuncInfixContext:
		return v.validateCustomBinaryOp(expr)
	case *parser.FunctionCallContext:
		return v.validateFunctionCall(expr)
	case *parser.FieldAccessContext:
		return v.validateFieldAccess(expr)
	case *parser.IndexAccessContext:
		return v.validateIndexAccess(expr)
	default:
		v.AddError(ctx, fmt.Sprintf("unsupported expression type: %T", expr))
		return false
	}
}

func (v *PureValidator) validateNumber(ctx *parser.NumberContext) bool {
	text := ctx.GetText()
	if strings.Contains(text, ".") {
		// Validate float format
		if !isValidFloat(text) {
			v.AddError(ctx, "invalid float format: "+text)
			return false
		}
	} else {
		// Validate integer format
		if !isValidInteger(text) {
			v.AddError(ctx, "invalid integer format: "+text)
			return false
		}
	}
	return true
}

func (v *PureValidator) validateString(ctx *parser.StringContext) bool {
	text := ctx.GetText()
	if len(text) < 2 {
		v.AddError(ctx, "invalid string format: "+text)
		return false
	}
	
	// Check string delimiters
	if (text[0] == '"' && text[len(text)-1] == '"') || 
	   (text[0] == '`' && text[len(text)-1] == '`') {
		return true
	}
	
	v.AddError(ctx, "invalid string format: "+text)
	return false
}

func (v *PureValidator) validateBoolean(ctx *parser.BooleanContext) bool {
	text := strings.ToLower(ctx.GetText())
	if text == "true" || text == "false" {
		return true
	}
	v.AddError(ctx, "invalid boolean value: "+ctx.GetText())
	return false
}

func (v *PureValidator) validateIdentifier(ctx *parser.IdContext) bool {
	name := ctx.GetText()
	_, exists := v.symbolTable.Lookup(name)
	if !exists {
		v.AddError(ctx, "undefined identifier: "+name)
		return false
	}
	return true
}

func (v *PureValidator) validateArray(ctx *parser.ArrayContext) bool {
	if ctx.ArrayElements() == nil {
		return true // Empty array is valid
	}

	elements := ctx.ArrayElements().AllExpr()
	if len(elements) == 0 {
		return true
	}

	valid := true
	
	// Validate all elements
	for _, elemCtx := range elements {
		if !v.ValidateExpression(elemCtx) {
			valid = false
		}
	}
	
	// Check type consistency (but don't decide on promotion)
	firstType := v.inferExpressionType(elements[0])
	for i := 1; i < len(elements); i++ {
		elemType := v.inferExpressionType(elements[i])
		if !v.areTypesCompatibleInArray(firstType, elemType) {
			v.AddError(elements[i], fmt.Sprintf("array element type incompatible: %s vs %s", firstType, elemType))
			valid = false
		}
	}

	return valid
}

func (v *PureValidator) validateDictionary(ctx *parser.DictionaryContext) bool {
	if ctx.DictElements() == nil {
		return true // Empty dictionary is valid
	}

	pairs := ctx.DictElements().AllDictPair()
	if len(pairs) == 0 {
		return true
	}

	valid := true
	
	// Validate all key-value pairs
	for _, pair := range pairs {
		if !v.ValidateExpression(pair.Expr(0)) {
			valid = false
		}
		if !v.ValidateExpression(pair.Expr(1)) {
			valid = false
		}
	}
	
	// Check type consistency
	firstPair := pairs[0]
	firstKeyType := v.inferExpressionType(firstPair.Expr(0))
	firstValueType := v.inferExpressionType(firstPair.Expr(1))
	
	for i := 1; i < len(pairs); i++ {
		pair := pairs[i]
		keyType := v.inferExpressionType(pair.Expr(0))
		valueType := v.inferExpressionType(pair.Expr(1))
		
		if !v.areTypesCompatibleInDict(firstKeyType, keyType) {
			v.AddError(pair.Expr(0), fmt.Sprintf("dictionary key type incompatible: %s vs %s", firstKeyType, keyType))
			valid = false
		}
		if !v.areTypesCompatibleInDict(firstValueType, valueType) {
			v.AddError(pair.Expr(1), fmt.Sprintf("dictionary value type incompatible: %s vs %s", firstValueType, valueType))
			valid = false
		}
	}

	return valid
}

func (v *PureValidator) validateBinaryOperation(leftCtx, rightCtx parser.IExprContext, op string) bool {
	leftValid := v.ValidateExpression(leftCtx)
	rightValid := v.ValidateExpression(rightCtx)
	
	if !leftValid || !rightValid {
		return false
	}
	
	leftType := v.inferExpressionType(leftCtx)
	rightType := v.inferExpressionType(rightCtx)
	
	return v.validateBinaryOpTypes(leftCtx, rightCtx, leftType, rightType, op)
}

func (v *PureValidator) validateBinaryOpTypes(leftCtx, rightCtx parser.IExprContext, leftType, rightType Type, op string) bool {
	switch op {
	case "+", "-", "*", "/":
		// Arithmetic operations require numeric types
		if !v.isNumericOrNumericBase(leftType) {
			v.AddError(leftCtx, fmt.Sprintf("arithmetic operation requires numeric type, got %s", leftType))
			return false
		}
		if !v.isNumericOrNumericBase(rightType) {
			v.AddError(rightCtx, fmt.Sprintf("arithmetic operation requires numeric type, got %s", rightType))
			return false
		}
		return true
		
	case ">", "<", ">=", "<=":
		// Comparison operations: both numeric or both same type
		if v.isNumericType(leftType) && v.isNumericType(rightType) {
			return true
		}
		if !leftType.Equals(rightType) {
			v.AddError(rightCtx, fmt.Sprintf("comparison requires compatible types: %s vs %s", leftType, rightType))
			return false
		}
		return true
		
	case "==", "!=":
		// Equality: both numeric or both same type
		if v.isNumericType(leftType) && v.isNumericType(rightType) {
			return true
		}
		if !leftType.Equals(rightType) {
			v.AddError(rightCtx, fmt.Sprintf("equality comparison requires compatible types: %s vs %s", leftType, rightType))
			return false
		}
		return true
		
	default:
		v.AddError(leftCtx, "unsupported binary operator: "+op)
		return false
	}
}

func (v *PureValidator) validateCustomBinaryOp(ctx *parser.CompareFuncInfixContext) bool {
	funcName := ctx.BinaryOp().GetText()
	
	leftValid := v.ValidateExpression(ctx.Expr(0))
	rightValid := v.ValidateExpression(ctx.Expr(1))
	
	if !leftValid || !rightValid {
		return false
	}
	
	// Get argument types
	leftType := v.inferExpressionType(ctx.Expr(0))
	rightType := v.inferExpressionType(ctx.Expr(1))
	
	// Try to resolve binary operator overload
	symbol, err := v.symbolTable.ResolveBinaryOpOverload(funcName, leftType, rightType)
	if err != nil {
		v.AddError(ctx, fmt.Sprintf("cannot resolve binary operator %s for types (%s, %s): %v", funcName, leftType, rightType, err))
		return false
	}
	
	binaryOpType, ok := symbol.Type.(*BinaryOpType)
	if !ok {
		v.AddError(ctx, funcName+" is not a binary operator")
		return false
	}
	
	// Store the resolved type for code generation
	_ = binaryOpType // Use the resolved binary operator type
	
	return true
}

func (v *PureValidator) validateFunctionCall(ctx *parser.FunctionCallContext) bool {
	funcName := ctx.Func_().GetText()
	
	// Validate arguments
	valid := true
	var argTypes []Type
	if ctx.ExprList() != nil {
		for _, exprCtx := range ctx.ExprList().AllExpr() {
			if !v.ValidateExpression(exprCtx) {
				valid = false
			} else {
				argTypes = append(argTypes, v.inferExpressionType(exprCtx))
			}
		}
	}
	
	if !valid {
		return false
	}
	
	// Try to resolve function overload first
	symbol, err := v.symbolTable.ResolveFunctionOverload(funcName, argTypes)
	overloadResolved := err == nil
	
	if err != nil {
		// Fallback to old lookup method for backward compatibility
		var exists bool
		symbol, exists = v.symbolTable.Lookup(funcName)
		if !exists {
			v.AddError(ctx, "undefined function: "+funcName)
			return false
		}
	}
	
	funcType, ok := symbol.Type.(*FunctionType)
	if !ok {
		v.AddError(ctx, funcName+" is not a function")
		return false
	}
	
	// Check argument count
	if !v.isArgCountValid(funcType, len(argTypes)) {
		if funcType.IsVariadic {
			minArgs := len(funcType.ParamTypes)
			v.AddError(ctx, fmt.Sprintf("variadic function %s expects at least %d arguments, got %d",
				funcName, minArgs, len(argTypes)))
		} else {
			v.AddError(ctx, fmt.Sprintf("function %s expects %d arguments, got %d",
				funcName, len(funcType.ParamTypes), len(argTypes)))
		}
		return false
	}
	
	// Check argument type compatibility
	// If overload was resolved, we're more lenient with implicit conversions
	return v.validateFunctionArguments(ctx, funcType, argTypes, overloadResolved)
}

func (v *PureValidator) validateFunctionArguments(ctx *parser.FunctionCallContext, funcType *FunctionType, argTypes []Type, overloadResolved bool) bool {
	valid := true
	
	// Check fixed parameters
	fixedParamCount := len(funcType.ParamTypes)
	for i := 0; i < fixedParamCount && i < len(argTypes); i++ {
		expectedType := funcType.ParamTypes[i]
		argType := argTypes[i]
		
		// If overload was resolved, allow implicit conversions
		compatible := false
		if overloadResolved {
			compatible = argType.IsCompatibleWith(expectedType) || v.canImplicitlyCast(argType, expectedType)
		} else {
			compatible = argType.IsCompatibleWith(expectedType)
		}
		
		if !compatible {
			v.AddError(ctx.ExprList().Expr(i),
				fmt.Sprintf("argument %d type mismatch: expected %s, got %s", i+1, expectedType, argType))
			valid = false
		}
	}
	
	// Check variadic parameters
	if funcType.IsVariadic && len(argTypes) > fixedParamCount {
		expectedType := funcType.VariadicType
		for i := fixedParamCount; i < len(argTypes); i++ {
			argType := argTypes[i]
			
			// If overload was resolved, allow implicit conversions
			compatible := false
			if overloadResolved {
				compatible = argType.IsCompatibleWith(expectedType) || v.canImplicitlyCast(argType, expectedType)
			} else {
				compatible = argType.IsCompatibleWith(expectedType)
			}
			
			if !compatible {
				v.AddError(ctx.ExprList().Expr(i),
					fmt.Sprintf("variadic argument %d type mismatch: expected %s, got %s",
						i+1, expectedType, argType))
				valid = false
			}
		}
	}
	
	return valid
}

func (v *PureValidator) validateFieldAccess(ctx *parser.FieldAccessContext) bool {
	// 验证左侧表达式
	if !v.ValidateExpression(ctx.Expr()) {
		return false
	}
	
	// 获取左侧表达式的类型
	leftType := v.inferExpressionType(ctx.Expr())
	
	// 检查是否为结构体类型
	structType, ok := leftType.(*StructType)
	if !ok {
		v.AddError(ctx, fmt.Sprintf("field access on non-struct type: %s", leftType))
		return false
	}
	
	// 获取字段名
	fieldName := ctx.ID().GetText()
	
	// 检查字段是否存在
	_, exists := structType.GetFieldType(fieldName)
	if !exists {
		v.AddError(ctx, fmt.Sprintf("field '%s' does not exist in struct type '%s'", fieldName, structType.Name))
		return false
	}
	
	return true
}

func (v *PureValidator) validateIndexAccess(ctx *parser.IndexAccessContext) bool {
	// 验证左侧表达式（被索引的对象）
	if !v.ValidateExpression(ctx.Expr(0)) {
		return false
	}
	
	// 验证右侧表达式（索引）
	if !v.ValidateExpression(ctx.Expr(1)) {
		return false
	}
	
	// 获取对象和索引的类型
	objectType := v.inferExpressionType(ctx.Expr(0))
	indexType := v.inferExpressionType(ctx.Expr(1))
	
	// 检查对象类型是否支持索引访问
	switch objType := objectType.(type) {
	case *ArrayType:
		// 数组索引必须是整数
		if !indexType.Equals(IntType) {
			v.AddError(ctx.Expr(1), fmt.Sprintf("array index must be int, got %s", indexType))
			return false
		}
		return true
	case *DictType:
		// 字典索引必须与键类型匹配
		if !indexType.Equals(objType.KeyType) && !v.canImplicitlyCast(indexType, objType.KeyType) {
			v.AddError(ctx.Expr(1), fmt.Sprintf("dictionary index type mismatch: expected %s, got %s", objType.KeyType, indexType))
			return false
		}
		return true
	default:
		v.AddError(ctx.Expr(0), fmt.Sprintf("index access not supported on type %s", objectType))
		return false
	}
}

// Helper methods

func (v *PureValidator) inferExpressionType(ctx parser.IExprContext) Type {
	switch expr := ctx.(type) {
	case *parser.NumberContext:
		text := expr.GetText()
		if strings.Contains(text, ".") {
			return FloatType
		}
		return IntType
	case *parser.StringContext:
		return StringType
	case *parser.BooleanContext:
		return BooleanType
	case *parser.IdContext:
		name := expr.GetText()
		if symbol, exists := v.symbolTable.Lookup(name); exists {
			return symbol.Type
		}
		return AnyType
	case *parser.ArrayContext:
		if expr.ArrayElements() == nil || len(expr.ArrayElements().AllExpr()) == 0 {
			return &ArrayType{ElementType: AnyType}
		}
		firstElemType := v.inferExpressionType(expr.ArrayElements().AllExpr()[0])
		return &ArrayType{ElementType: firstElemType}
	case *parser.DictionaryContext:
		if expr.DictElements() == nil || len(expr.DictElements().AllDictPair()) == 0 {
			return &DictType{KeyType: AnyType, ValueType: AnyType}
		}
		firstPair := expr.DictElements().AllDictPair()[0]
		keyType := v.inferExpressionType(firstPair.Expr(0))
		valueType := v.inferExpressionType(firstPair.Expr(1))
		return &DictType{KeyType: keyType, ValueType: valueType}
	case *parser.ParensContext:
		return v.inferExpressionType(expr.Expr())
	case *parser.MulDivContext:
		leftType := v.inferExpressionType(expr.Expr(0))
		rightType := v.inferExpressionType(expr.Expr(1))
		if v.isNumericType(leftType) && v.isNumericType(rightType) {
			// Promote to float if either is float
			if v.isFloatType(leftType) || v.isFloatType(rightType) {
				return FloatType
			}
			return IntType
		}
		return AnyType
	case *parser.AddSubContext:
		leftType := v.inferExpressionType(expr.Expr(0))
		rightType := v.inferExpressionType(expr.Expr(1))
		if v.isNumericType(leftType) && v.isNumericType(rightType) {
			// Promote to float if either is float
			if v.isFloatType(leftType) || v.isFloatType(rightType) {
				return FloatType
			}
			return IntType
		}
		return AnyType
	case *parser.CompareSymbolContext, *parser.CompareFuncInfixContext:
		return BooleanType
	case *parser.FunctionCallContext:
		funcName := expr.Func_().GetText()
		if symbol, exists := v.symbolTable.Lookup(funcName); exists {
			if funcType, ok := symbol.Type.(*FunctionType); ok {
				return funcType.ReturnType
			}
		}
		return AnyType
	case *parser.FieldAccessContext:
		// 获取左侧表达式的类型
		leftType := v.inferExpressionType(expr.Expr())
		
		// 如果是结构体类型，返回字段类型
		if structType, ok := leftType.(*StructType); ok {
			fieldName := expr.ID().GetText()
			if fieldType, exists := structType.GetFieldType(fieldName); exists {
				return fieldType
			}
		}
		return AnyType
	case *parser.IndexAccessContext:
		// 获取被索引对象的类型
		objectType := v.inferExpressionType(expr.Expr(0))
		
		// 根据对象类型推断索引访问的结果类型
		switch objType := objectType.(type) {
		case *ArrayType:
			return objType.ElementType
		case *DictType:
			return objType.ValueType
		default:
			return AnyType
		}
	default:
		return AnyType
	}
}

func (v *PureValidator) isNumericType(t Type) bool {
	if bt, ok := t.(*BasicType); ok {
		return bt.Name == "int" || bt.Name == "float"
	}
	return false
}

func (v *PureValidator) isFloatType(t Type) bool {
	if bt, ok := t.(*BasicType); ok {
		return bt.Name == "float"
	}
	return false
}

func (v *PureValidator) isNumericOrNumericBase(t Type) bool {
	if bt, ok := t.(*BasicType); ok {
		return bt.Name == "int" || bt.Name == "float"
	}
	if nt, ok := t.(*NumericType); ok {
		return nt.Name == "numeric"
	}
	return false
}

func (v *PureValidator) areTypesCompatibleInArray(type1, type2 Type) bool {
	if type1.Equals(type2) {
		return true
	}
	// Allow numeric type mixing in arrays (let IR decide promotion)
	return v.isNumericType(type1) && v.isNumericType(type2)
}

func (v *PureValidator) areTypesCompatibleInDict(type1, type2 Type) bool {
	if type1.Equals(type2) {
		return true
	}
	// Allow numeric type mixing in dictionaries (let IR decide promotion)
	return v.isNumericType(type1) && v.isNumericType(type2)
}

func (v *PureValidator) isArgCountValid(funcType *FunctionType, argCount int) bool {
	if funcType.IsVariadic {
		return argCount >= len(funcType.ParamTypes)
	}
	return argCount == len(funcType.ParamTypes)
}

// canImplicitlyCast 检查是否可以进行隐式类型转换
func (v *PureValidator) canImplicitlyCast(from, to Type) bool {
	if from.Equals(to) {
		return true
	}
	
	// int -> float 隐式转换
	if from.Equals(IntType) && to.Equals(FloatType) {
		return true
	}
	
	// 数值类型与NumericBaseType的兼容性
	if to.Equals(NumericBaseType) {
		return from.Equals(IntType) || from.Equals(FloatType)
	}
	
	// 检查类型兼容性
	return from.IsCompatibleWith(to)
}

// Utility functions for format validation

func isValidFloat(s string) bool {
	// Basic float validation - could be more sophisticated
	return strings.Count(s, ".") <= 1
}

func isValidInteger(s string) bool {
	// Basic integer validation - could be more sophisticated
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

// ValidateCodePure is the main entry point for pure semantic validation
func ValidateCodePure(input string) ([]*SemanticError, *PureValidator) {
	lexer := parser.NewmlangLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewmlangParser(stream)

	// Parse the input
	tree := p.Prog()

	// Create validator and validate
	validator := NewPureValidator()
	validator.ValidateProgram(tree.(*parser.ProgContext))

	return validator.GetErrors(), validator
}