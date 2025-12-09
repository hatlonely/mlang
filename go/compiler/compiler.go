package compiler

import (
	"fmt"
	"strings"
	
	"github.com/antlr4-go/antlr/v4"
	"github.com/hatlonely/mlang/go/codegen"
	"github.com/hatlonely/mlang/go/ir"
	"github.com/hatlonely/mlang/go/semantic"
	parser "github.com/hatlonely/mlang/gen/go"
)

// SyntaxError represents a syntax error with position information
type SyntaxError struct {
	Line    int
	Column  int
	Message string
	Length  int
}

// CustomErrorListener handles syntax errors with better formatting
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	errors     []SyntaxError
	sourceCode string
}

func NewCustomErrorListener(sourceCode string) *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		errors:               make([]SyntaxError, 0),
		sourceCode:           sourceCode,
	}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// 计算错误token的长度
	length := 1
	if token, ok := offendingSymbol.(antlr.Token); ok && token != nil {
		text := token.GetText()
		if text != "" {
			length = len(text)
		}
	}
	
	c.errors = append(c.errors, SyntaxError{
		Line:    line,
		Column:  column,
		Message: msg,
		Length:  length,
	})
}

func (c *CustomErrorListener) GetErrors() []SyntaxError {
	return c.errors
}

func (comp *Compiler) formatSyntaxError(err SyntaxError, sourceCode string) string {
	// 构建错误信息格式
	result := fmt.Sprintf("syntax error at line %d, column %d: %s\n", err.Line, err.Column, err.Message)
	
	// 分割源代码为行
	lines := strings.Split(sourceCode, "\n")
	if err.Line > 0 && err.Line <= len(lines) {
		sourceLine := lines[err.Line-1]
		result += fmt.Sprintf("  %s\n", sourceLine)
		
		// 添加指示箭头
		if err.Column >= 0 {
			spaces := strings.Repeat(" ", err.Column+2) // +2 for the "  " indent
			length := err.Length
			if length <= 0 {
				length = 1 // 默认长度
			}
			arrows := strings.Repeat("^", length)
			result += fmt.Sprintf("%s%s", spaces, arrows)
		}
	}
	
	return result
}

// CompileResult contains the result of compilation
type CompileResult struct {
	GoCode     string        // Go expression
	Errors     []string
	Warnings   []string
	ResultType semantic.Type
}

// Compiler compiles mlang source code to target languages
type Compiler struct {
	validator *semantic.PureValidator
}

// NewCompiler creates a new compiler instance
func NewCompiler() *Compiler {
	return &Compiler{
		validator: semantic.NewPureValidator(),
	}
}

// CompileToGo compiles mlang source code to Go expression
func (c *Compiler) CompileToGo(input string) *CompileResult {
	result := &CompileResult{
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}
	
	// Step 1: Parse the input with custom error handling
	lexer := parser.NewmlangLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewmlangParser(stream)
	
	// 添加自定义错误监听器
	errorListener := NewCustomErrorListener(input)
	lexer.RemoveErrorListeners() // 移除词法器默认错误监听器
	lexer.AddErrorListener(errorListener)
	p.RemoveErrorListeners() // 移除解析器默认错误监听器
	p.AddErrorListener(errorListener)
	
	// Parse the program
	tree := p.Prog()
	
	// 检查语法错误
	if len(errorListener.GetErrors()) > 0 {
		for _, err := range errorListener.GetErrors() {
			formattedError := c.formatSyntaxError(err, input)
			result.Errors = append(result.Errors, formattedError)
		}
		return result
	}
	
	// Step 2: Semantic validation
	c.validator.ClearErrors()
	valid := c.validator.ValidateProgram(tree.(*parser.ProgContext))
	
	if !valid {
		for _, err := range c.validator.GetErrors() {
			formattedError := c.formatError(err, input)
			result.Errors = append(result.Errors, formattedError)
		}
		return result
	}
	
	// Step 3: Build IR
	irBuilder := ir.NewBuilder(c.validator.GetSymbolTable())
	program := irBuilder.BuildProgram(tree.(*parser.ProgContext))
	
	if errors := irBuilder.GetErrors(); len(errors) > 0 {
		result.Errors = append(result.Errors, errors...)
		return result
	}
	
	result.ResultType = program.ResultType
	
	// Step 4: Generate Go expression
	goGen := codegen.NewGoGenerator()
	goCode, err := goGen.GenerateProgram(program)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("Code generation error: %v", err))
		return result
	}
	
	result.GoCode = goCode
	return result
}

// RegisterFunction registers a custom function
func (c *Compiler) RegisterFunction(name string, paramTypes []semantic.Type, returnType semantic.Type) error {
	return c.validator.RegisterFunction(name, paramTypes, returnType)
}

// RegisterVariadicFunction registers a custom variadic function
func (c *Compiler) RegisterVariadicFunction(name string, paramTypes []semantic.Type, variadicType semantic.Type, returnType semantic.Type) error {
	return c.validator.RegisterVariadicFunction(name, paramTypes, variadicType, returnType)
}

// RegisterBinaryOp registers a custom binary operator
func (c *Compiler) RegisterBinaryOp(name string, leftType, rightType, returnType semantic.Type) error {
	return c.validator.RegisterBinaryOp(name, leftType, rightType, returnType)
}

// RegisterVariable registers a variable with its type
func (c *Compiler) RegisterVariable(name string, varType semantic.Type) error {
	return c.validator.RegisterVariable(name, varType)
}

// RegisterStructType registers a custom struct type
func (c *Compiler) RegisterStructType(name string, fields map[string]semantic.Type) error {
	return c.validator.GetSymbolTable().RegisterStructType(name, fields)
}

// RegisterStructVariable registers a struct variable
func (c *Compiler) RegisterStructVariable(varName string, structTypeName string) error {
	return c.validator.GetSymbolTable().RegisterStructVariable(varName, structTypeName)
}

// GetValidator returns the semantic validator (for advanced usage)
func (c *Compiler) GetValidator() *semantic.PureValidator {
	return c.validator
}

// CompileExample demonstrates the usage
func CompileExample() {
	compiler := NewCompiler()
	
	// Register custom functions
	compiler.RegisterFunction("sqrt", []semantic.Type{semantic.NumericBaseType}, semantic.NumericBaseType)
	compiler.RegisterBinaryOp("contains", semantic.StringType, semantic.StringType, semantic.BooleanType)
	
	// Compile examples
	examples := []string{
		"1 + 2 * 3",
		"[1, 2, 3.5, 4]",
		`{"name": "test", "age": 25}`,
		"max(1, 2, 3, 4, 5)",
		"len([1, 2, 3])",
		`"hello" contains "ell"`,
	}
	
	for i, example := range examples {
		fmt.Printf("=== Example %d: %s ===\n", i+1, example)
		
		result := compiler.CompileToGo(example)
		
		if len(result.Errors) > 0 {
			fmt.Println("Errors:")
			for _, err := range result.Errors {
				fmt.Printf("  %s\n", err)
			}
		} else {
			fmt.Println("Generated Go Expression:")
			fmt.Printf("  %s\n", result.GoCode)
			fmt.Printf("Result Type: %s\n", result.ResultType)
		}
		
		fmt.Println()
	}
}

// formatError formats a semantic error with source code and position highlighting
func (c *Compiler) formatError(err *semantic.SemanticError, sourceCode string) string {
	// 构建错误信息格式
	result := fmt.Sprintf("semantic error at line %d, column %d: %s\n", err.Line, err.Column, err.Message)
	
	// 分割源代码为行
	lines := strings.Split(sourceCode, "\n")
	if err.Line > 0 && err.Line <= len(lines) {
		sourceLine := lines[err.Line-1]
		result += fmt.Sprintf("  %s\n", sourceLine)
		
		// 添加指示箭头
		if err.Column >= 0 {
			spaces := strings.Repeat(" ", err.Column+2) // +2 for the "  " indent
			length := err.Length
			if length <= 0 {
				length = 1 // 默认长度
			}
			arrows := strings.Repeat("^", length)
			result += fmt.Sprintf("%s%s", spaces, arrows)
		}
	}
	
	return result
}