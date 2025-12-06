package semantic

import (
	"fmt"

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

// Analyzer performs semantic analysis on mlang AST
type Analyzer struct {
	symbolTable *SymbolTable
	errors      []*SemanticError
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		symbolTable: NewSymbolTable(),
		errors:      make([]*SemanticError, 0),
	}
}

func (a *Analyzer) AddError(ctx antlr.ParserRuleContext, message string) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	a.errors = append(a.errors, &SemanticError{
		Line:    line,
		Column:  column,
		Message: message,
	})
}

func (a *Analyzer) GetErrors() []*SemanticError {
	return a.errors
}

func (a *Analyzer) HasErrors() bool {
	return len(a.errors) > 0
}

// RegisterFunction 向分析器注册一个新函数
func (a *Analyzer) RegisterFunction(name string, paramTypes []Type, returnType Type) error {
	return a.symbolTable.RegisterFunction(name, paramTypes, returnType)
}

// RegisterVariadicFunction 向分析器注册一个新的可变参数函数
func (a *Analyzer) RegisterVariadicFunction(name string, paramTypes []Type, variadicType Type, returnType Type) error {
	return a.symbolTable.RegisterVariadicFunction(name, paramTypes, variadicType, returnType)
}

// RegisterBinaryOp 向分析器注册一个比较运算符
func (a *Analyzer) RegisterBinaryOp(name string, leftType, rightType, returnType Type) error {
	return a.symbolTable.RegisterBinaryOp(name, leftType, rightType, returnType)
}

// UnregisterFunction 从分析器中删除一个函数
func (a *Analyzer) UnregisterFunction(name string) {
	a.symbolTable.UnregisterFunction(name)
}

// ListFunctions 列出所有已注册的函数
func (a *Analyzer) ListFunctions() map[string]*FunctionType {
	return a.symbolTable.ListFunctions()
}

// GetSymbolTable 获取符号表（用于高级操作）
func (a *Analyzer) GetSymbolTable() *SymbolTable {
	return a.symbolTable
}

// AnalyzeProgram analyzes the entire program
func (a *Analyzer) AnalyzeProgram(ctx *parser.ProgContext) Type {
	var lastType Type = VoidType

	for _, statCtx := range ctx.AllStat() {
		if exprCtx := statCtx.Expr(); exprCtx != nil {
			lastType = a.AnalyzeExpression(exprCtx)
		}
	}

	return lastType
}

// AnalyzeExpression analyzes expressions and returns their type
func (a *Analyzer) AnalyzeExpression(ctx parser.IExprContext) Type {
	switch expr := ctx.(type) {
	case *parser.NumberContext:
		return a.analyzeNumber(expr)
	case *parser.StringContext:
		return a.analyzeString(expr)
	case *parser.BooleanContext:
		return a.analyzeBoolean(expr)
	case *parser.IdContext:
		return a.analyzeId(expr)
	case *parser.ArrayContext:
		return a.analyzeArray(expr)
	case *parser.DictionaryContext:
		return a.analyzeDictionary(expr)
	case *parser.ParensContext:
		return a.analyzeParens(expr)
	case *parser.MulDivContext:
		return a.analyzeBinaryOp(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.AddSubContext:
		return a.analyzeBinaryOp(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.CompareSymbolContext:
		return a.analyzeBinaryOp(expr.Expr(0), expr.Expr(1), expr.GetOp().GetText())
	case *parser.CompareFuncInfixContext:
		return a.analyzeCompareFuncInfix(expr)
	case *parser.FunctionCallContext:
		return a.analyzeFunctionCall(expr)
	default:
		a.AddError(ctx, fmt.Sprintf("unsupported expression type: %T", expr))
		return AnyType
	}
}

func (a *Analyzer) analyzeNumber(_ *parser.NumberContext) Type {
	return NumberType
}

func (a *Analyzer) analyzeString(_ *parser.StringContext) Type {
	return StringType
}

func (a *Analyzer) analyzeBoolean(_ *parser.BooleanContext) Type {
	return BooleanType
}

func (a *Analyzer) analyzeId(ctx *parser.IdContext) Type {
	name := ctx.GetText()
	symbol, exists := a.symbolTable.Lookup(name)
	if !exists {
		a.AddError(ctx, fmt.Sprintf("undefined identifier: %s", name))
		return AnyType
	}
	return symbol.Type
}

func (a *Analyzer) analyzeArray(ctx *parser.ArrayContext) Type {
	if ctx.ArrayElements() == nil {
		// 空数组，推断为 any 类型数组
		return &ArrayType{ElementType: AnyType}
	}

	elements := ctx.ArrayElements().AllExpr()
	if len(elements) == 0 {
		return &ArrayType{ElementType: AnyType}
	}

	// 分析第一个元素确定数组类型
	elementType := a.AnalyzeExpression(elements[0])

	// 检查所有元素类型是否一致
	for i := 1; i < len(elements); i++ {
		exprType := a.AnalyzeExpression(elements[i])
		if !elementType.Equals(exprType) {
			a.AddError(elements[i], fmt.Sprintf("array element type mismatch: expected %s, got %s", elementType, exprType))
		}
	}

	return &ArrayType{ElementType: elementType}
}

func (a *Analyzer) analyzeDictionary(ctx *parser.DictionaryContext) Type {
	if ctx.DictElements() == nil {
		// 空字典，推断为 any:any 类型
		return &DictType{KeyType: AnyType, ValueType: AnyType}
	}

	pairs := ctx.DictElements().AllDictPair()
	if len(pairs) == 0 {
		return &DictType{KeyType: AnyType, ValueType: AnyType}
	}

	// 分析第一个键值对确定字典类型
	firstPair := pairs[0]
	keyType := a.AnalyzeExpression(firstPair.Expr(0))
	valueType := a.AnalyzeExpression(firstPair.Expr(1))

	// 检查所有键值对类型是否一致
	for i := 1; i < len(pairs); i++ {
		pair := pairs[i]
		kType := a.AnalyzeExpression(pair.Expr(0))
		vType := a.AnalyzeExpression(pair.Expr(1))

		if !keyType.Equals(kType) {
			a.AddError(pair.Expr(0), fmt.Sprintf("dictionary key type mismatch: expected %s, got %s", keyType, kType))
		}
		if !valueType.Equals(vType) {
			a.AddError(pair.Expr(1), fmt.Sprintf("dictionary value type mismatch: expected %s, got %s", valueType, vType))
		}
	}

	return &DictType{KeyType: keyType, ValueType: valueType}
}

func (a *Analyzer) analyzeParens(ctx *parser.ParensContext) Type {
	return a.AnalyzeExpression(ctx.Expr())
}

func (a *Analyzer) analyzeBinaryOp(leftCtx, rightCtx parser.IExprContext, op string) Type {
	leftType := a.AnalyzeExpression(leftCtx)
	rightType := a.AnalyzeExpression(rightCtx)

	switch op {
	case "+", "-", "*", "/":
		// 算术运算符需要两个数字类型
		if !leftType.Equals(NumberType) {
			a.AddError(leftCtx, fmt.Sprintf("arithmetic operation requires number type, got %s", leftType))
		}
		if !rightType.Equals(NumberType) {
			a.AddError(rightCtx, fmt.Sprintf("arithmetic operation requires number type, got %s", rightType))
		}
		return NumberType
	case ">", "<", ">=", "<=", "==", "!=":
		// 比较运算符需要相同类型
		if !leftType.Equals(rightType) {
			a.AddError(rightCtx, fmt.Sprintf("comparison requires same types: %s vs %s", leftType, rightType))
		}
		return BooleanType
	default:
		a.AddError(leftCtx, fmt.Sprintf("unsupported binary operator: %s", op))
		return AnyType
	}
}

func (a *Analyzer) analyzeCompareFuncInfix(ctx *parser.CompareFuncInfixContext) Type {
	funcName := ctx.BinaryOp().GetText()
	leftType := a.AnalyzeExpression(ctx.Expr(0))
	rightType := a.AnalyzeExpression(ctx.Expr(1))

	// 查找比较函数
	symbol, exists := a.symbolTable.Lookup(funcName)
	if !exists {
		a.AddError(ctx, fmt.Sprintf("undefined comparison function: %s", funcName))
		return BooleanType // 假设比较函数返回布尔值
	}

	if funcType, ok := symbol.Type.(*FunctionType); ok {
		// 检查参数类型
		if len(funcType.ParamTypes) != 2 {
			a.AddError(ctx, fmt.Sprintf("comparison function %s expects 2 parameters", funcName))
		} else {
			if !funcType.ParamTypes[0].Equals(leftType) {
				a.AddError(ctx.Expr(0), fmt.Sprintf("parameter type mismatch: expected %s, got %s", funcType.ParamTypes[0], leftType))
			}
			if !funcType.ParamTypes[1].Equals(rightType) {
				a.AddError(ctx.Expr(1), fmt.Sprintf("parameter type mismatch: expected %s, got %s", funcType.ParamTypes[1], rightType))
			}
		}
		return funcType.ReturnType
	}

	a.AddError(ctx, fmt.Sprintf("%s is not a function", funcName))
	return BooleanType
}

func (a *Analyzer) analyzeFunctionCall(ctx *parser.FunctionCallContext) Type {
	funcName := ctx.Func_().GetText()

	// 查找函数符号
	symbol, exists := a.symbolTable.Lookup(funcName)
	if !exists {
		a.AddError(ctx, fmt.Sprintf("undefined function: %s", funcName))
		return AnyType
	}

	funcType, ok := symbol.Type.(*FunctionType)
	if !ok {
		a.AddError(ctx, fmt.Sprintf("%s is not a function", funcName))
		return AnyType
	}

	// 分析参数
	var argTypes []Type
	if ctx.ExprList() != nil {
		for _, exprCtx := range ctx.ExprList().AllExpr() {
			argTypes = append(argTypes, a.AnalyzeExpression(exprCtx))
		}
	}

	// 检查参数数量
	if funcType.IsVariadic {
		// 可变参数函数：至少要有固定参数的数量
		minArgs := len(funcType.ParamTypes)
		if len(argTypes) < minArgs {
			a.AddError(ctx, fmt.Sprintf("variadic function %s expects at least %d arguments, got %d",
				funcName, minArgs, len(argTypes)))
			return funcType.ReturnType
		}
	} else {
		// 固定参数函数：参数数量必须完全匹配
		if len(argTypes) != len(funcType.ParamTypes) {
			a.AddError(ctx, fmt.Sprintf("function %s expects %d arguments, got %d",
				funcName, len(funcType.ParamTypes), len(argTypes)))
			return funcType.ReturnType
		}
	}

	// 检查固定参数类型
	fixedParamCount := len(funcType.ParamTypes)
	for i := 0; i < fixedParamCount && i < len(argTypes); i++ {
		expectedType := funcType.ParamTypes[i]
		argType := argTypes[i]
		if !expectedType.Equals(AnyType) && !argType.Equals(expectedType) {
			a.AddError(ctx.ExprList().Expr(i),
				fmt.Sprintf("argument %d type mismatch: expected %s, got %s", i+1, expectedType, argType))
		}
	}

	// 检查可变参数类型
	if funcType.IsVariadic && len(argTypes) > fixedParamCount {
		for i := fixedParamCount; i < len(argTypes); i++ {
			argType := argTypes[i]
			if !funcType.VariadicType.Equals(AnyType) && !argType.Equals(funcType.VariadicType) {
				a.AddError(ctx.ExprList().Expr(i),
					fmt.Sprintf("variadic argument %d type mismatch: expected %s, got %s",
						i+1, funcType.VariadicType, argType))
			}
		}
	}

	return funcType.ReturnType
}

// AnalyzeCode is the main entry point for semantic analysis
func AnalyzeCode(input string) ([]*SemanticError, Type) {
	lexer := parser.NewmlangLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewmlangParser(stream)

	// Parse the input
	tree := p.Prog()

	// Create analyzer and analyze
	analyzer := NewAnalyzer()
	resultType := analyzer.AnalyzeProgram(tree.(*parser.ProgContext))

	return analyzer.GetErrors(), resultType
}

// AnalyzeCodeWithAnalyzer 使用提供的分析器分析代码
func AnalyzeCodeWithAnalyzer(input string, analyzer *Analyzer) ([]*SemanticError, Type) {
	lexer := parser.NewmlangLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewmlangParser(stream)

	// Parse the input
	tree := p.Prog()

	// Clear previous errors
	analyzer.errors = make([]*SemanticError, 0)

	// Analyze with the provided analyzer
	resultType := analyzer.AnalyzeProgram(tree.(*parser.ProgContext))

	return analyzer.GetErrors(), resultType
}
