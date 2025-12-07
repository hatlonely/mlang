package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	
	"github.com/hatlonely/mlang/go/compiler"
	"github.com/hatlonely/mlang/go/semantic"
)

func main() {
	var (
		inputFile  = flag.String("i", "", "Input mlang file (default: stdin)")
		outputFile = flag.String("o", "", "Output Go expression file (default: stdout)")
		expression = flag.String("e", "", "Evaluate expression directly")
		showHelp   = flag.Bool("h", false, "Show help")
	)
	flag.Parse()

	if *showHelp {
		printUsage()
		return
	}

	// Get input
	var input string
	var err error

	if *expression != "" {
		input = *expression
	} else {
		input, err = readInput(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	}

	// Compile
	comp := compiler.NewCompiler()
	
	// Register some common custom functions for demo
	registerCommonFunctions(comp)
	
	// Register some common variables for demo
	registerCommonVariables(comp)
	
	// Register struct types and variables
	registerStructTypes(comp)
	
	result := comp.CompileToGo(input)

	// Handle errors
	if len(result.Errors) > 0 {
		fmt.Fprintf(os.Stderr, "Compilation errors:\n")
		for _, err := range result.Errors {
			fmt.Fprintf(os.Stderr, "  %s\n", err)
		}
		os.Exit(1)
	}

	// Handle warnings
	if len(result.Warnings) > 0 {
		fmt.Fprintf(os.Stderr, "Warnings:\n")
		for _, warning := range result.Warnings {
			fmt.Fprintf(os.Stderr, "  %s\n", warning)
		}
	}

	// Write output
	err = writeOutput(*outputFile, result.GoCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	// Print summary to stderr if writing to file
	if *outputFile != "" {
		fmt.Fprintf(os.Stderr, "Successfully compiled to Go expression\n")
		fmt.Fprintf(os.Stderr, "Result type: %s\n", result.ResultType)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "mlang - Compiler for the mlang language\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s [options]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  %s -e '1 + 2 * 3'\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -i input.mlang -o output.go\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -e 'max(1, 2, 3)' | go run\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -e 'userAge + 10'\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -e 'pi * 2'\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nSupported expressions:\n")
	fmt.Fprintf(os.Stderr, "  - Arithmetic: 1 + 2 * 3, 1.5 + 2.0\n")
	fmt.Fprintf(os.Stderr, "  - Comparisons: 1 < 2, \"a\" == \"b\"\n")
	fmt.Fprintf(os.Stderr, "  - Arrays: [1, 2, 3], [1.5, 2.5]\n")
	fmt.Fprintf(os.Stderr, "  - Dictionaries: {\"key\": \"value\"}\n")
	fmt.Fprintf(os.Stderr, "  - Functions: len([1, 2]), max(1, 2, 3)\n")
	fmt.Fprintf(os.Stderr, "  - Variables: pi, e, userAge, userName\n")
	fmt.Fprintf(os.Stderr, "  - Built-ins: abs, max, min, sum, concat\n")
}

func readInput(filename string) (string, error) {
	var reader io.Reader
	if filename == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return "", err
		}
		defer file.Close()
		reader = file
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func writeOutput(filename, content string) error {
	var writer io.Writer
	if filename == "" {
		writer = os.Stdout
	} else {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		writer = file
	}

	_, err := writer.Write([]byte(content))
	return err
}

func registerCommonFunctions(comp *compiler.Compiler) {
	// Register previously built-in functions
	comp.RegisterFunction("len", []semantic.Type{semantic.AnyType}, semantic.IntType)
	
	// Register abs function for both int and float
	comp.RegisterFunction("abs", []semantic.Type{semantic.IntType}, semantic.IntType)
	comp.RegisterFunction("abs", []semantic.Type{semantic.FloatType}, semantic.FloatType)
	
	// Register max function with specific type overloads
	comp.RegisterVariadicFunction("max", []semantic.Type{}, semantic.IntType, semantic.IntType)
	comp.RegisterVariadicFunction("max", []semantic.Type{}, semantic.FloatType, semantic.FloatType)
	
	// Register min function with specific type overloads
	comp.RegisterVariadicFunction("min", []semantic.Type{}, semantic.IntType, semantic.IntType)
	comp.RegisterVariadicFunction("min", []semantic.Type{}, semantic.FloatType, semantic.FloatType)
	
	// Register sum function with specific type overloads
	comp.RegisterVariadicFunction("sum", []semantic.Type{}, semantic.IntType, semantic.IntType)
	comp.RegisterVariadicFunction("sum", []semantic.Type{}, semantic.FloatType, semantic.FloatType)
	
	// Register concat function
	comp.RegisterVariadicFunction("concat", []semantic.Type{}, semantic.StringType, semantic.StringType)
	
	// Register some useful mathematical functions
	comp.RegisterFunction("sqrt", []semantic.Type{semantic.IntType}, semantic.FloatType)
	comp.RegisterFunction("sqrt", []semantic.Type{semantic.FloatType}, semantic.FloatType)
	comp.RegisterFunction("pow", []semantic.Type{semantic.IntType, semantic.IntType}, semantic.IntType)
	comp.RegisterFunction("pow", []semantic.Type{semantic.FloatType, semantic.FloatType}, semantic.FloatType)
	comp.RegisterFunction("pow", []semantic.Type{semantic.IntType, semantic.FloatType}, semantic.FloatType)
	comp.RegisterFunction("pow", []semantic.Type{semantic.FloatType, semantic.IntType}, semantic.FloatType)
	comp.RegisterFunction("round", []semantic.Type{semantic.FloatType}, semantic.IntType)
	
	// Register string functions
	comp.RegisterBinaryOp("contains", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("startsWith", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("endsWith", semantic.StringType, semantic.StringType, semantic.BooleanType)
	
	// Register comparison operators for numbers
	comp.RegisterBinaryOp("gt", semantic.IntType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("gt", semantic.FloatType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("gt", semantic.IntType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("gt", semantic.FloatType, semantic.IntType, semantic.BooleanType)
	
	comp.RegisterBinaryOp("lt", semantic.IntType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("lt", semantic.FloatType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("lt", semantic.IntType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("lt", semantic.FloatType, semantic.IntType, semantic.BooleanType)
	
	comp.RegisterBinaryOp("ge", semantic.IntType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("ge", semantic.FloatType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("ge", semantic.IntType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("ge", semantic.FloatType, semantic.IntType, semantic.BooleanType)
	
	comp.RegisterBinaryOp("le", semantic.IntType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("le", semantic.FloatType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("le", semantic.IntType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("le", semantic.FloatType, semantic.IntType, semantic.BooleanType)
	
	comp.RegisterBinaryOp("eq", semantic.IntType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("eq", semantic.FloatType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("eq", semantic.IntType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("eq", semantic.FloatType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("eq", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("eq", semantic.BooleanType, semantic.BooleanType, semantic.BooleanType)
	
	comp.RegisterBinaryOp("ne", semantic.IntType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("ne", semantic.FloatType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("ne", semantic.IntType, semantic.FloatType, semantic.BooleanType)
	comp.RegisterBinaryOp("ne", semantic.FloatType, semantic.IntType, semantic.BooleanType)
	comp.RegisterBinaryOp("ne", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("ne", semantic.BooleanType, semantic.BooleanType, semantic.BooleanType)
	
	// Register string comparison operators
	comp.RegisterBinaryOp("gt", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("lt", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("ge", semantic.StringType, semantic.StringType, semantic.BooleanType)
	comp.RegisterBinaryOp("le", semantic.StringType, semantic.StringType, semantic.BooleanType)
	
	// Register utility functions
	comp.RegisterFunction("isEmpty", []semantic.Type{semantic.AnyType}, semantic.BooleanType)
	comp.RegisterFunction("toString", []semantic.Type{semantic.AnyType}, semantic.StringType)
}

func registerCommonVariables(comp *compiler.Compiler) {
	// Register mathematical constants
	comp.RegisterVariable("pi", semantic.FloatType)
	comp.RegisterVariable("e", semantic.FloatType)
	
	// Register system variables
	comp.RegisterVariable("version", semantic.StringType)
	comp.RegisterVariable("debug", semantic.BooleanType)
	
	// Register configuration variables
	comp.RegisterVariable("maxRetries", semantic.IntType)
	comp.RegisterVariable("timeout", semantic.FloatType)
	comp.RegisterVariable("apiKey", semantic.StringType)
	
	// Register user variables
	comp.RegisterVariable("userAge", semantic.IntType)
	comp.RegisterVariable("userName", semantic.StringType)
	comp.RegisterVariable("isActive", semantic.BooleanType)
}

func registerStructTypes(comp *compiler.Compiler) {
	// Register User struct type
	userFields := map[string]semantic.Type{
		"id":   semantic.IntType,
		"name": semantic.StringType,
		"age":  semantic.IntType,
	}
	comp.RegisterStructType("User", userFields)
	
	// Register Point struct type
	pointFields := map[string]semantic.Type{
		"x": semantic.FloatType,
		"y": semantic.FloatType,
	}
	comp.RegisterStructType("Point", pointFields)
	
	// Register some struct variables for demo
	comp.RegisterStructVariable("user", "User")
	comp.RegisterStructVariable("point", "Point")
}