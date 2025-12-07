package compiler

import (
	"fmt"
	
	"github.com/antlr4-go/antlr/v4"
	"github.com/hatlonely/mlang/go/codegen"
	"github.com/hatlonely/mlang/go/ir"
	"github.com/hatlonely/mlang/go/semantic"
	parser "github.com/hatlonely/mlang/gen/go"
)

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
	
	// Step 1: Parse the input
	lexer := parser.NewmlangLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewmlangParser(stream)
	
	// Parse the program
	tree := p.Prog()
	
	// Step 2: Semantic validation
	c.validator.ClearErrors()
	valid := c.validator.ValidateProgram(tree.(*parser.ProgContext))
	
	if !valid {
		for _, err := range c.validator.GetErrors() {
			result.Errors = append(result.Errors, err.Error())
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