# MLang - A Go Expression Compiler

MLang is a domain-specific language for mathematical and data manipulation expressions that compiles to Go expressions. It starts with **no built-in functions** - all functionality is added through user registration.

## Architecture

The compiler follows a clean layered architecture:

```
mlang source → Lexer/Parser → Semantic Validation → IR Generation → Code Generation
```

### Key Components

1. **Semantic Validation** (`go/semantic/validator.go`): Pure semantic validation without implementation decisions
2. **IR Layer** (`go/ir/`): Intermediate representation with type promotion and optimization information  
3. **Code Generation** (`go/codegen/`): Target-language specific code generators
4. **Compiler** (`go/compiler/`): Unified compilation pipeline

## Features

### Data Types
- **Numbers**: `42`, `3.14` (automatic int/float promotion)
- **Strings**: `"hello"`, `\`world\`` (with/without escape sequences)
- **Booleans**: `true`, `false`
- **Arrays**: `[1, 2, 3]`, `[1.5, 2.5]` (with type promotion)
- **Dictionaries**: `{"key": "value", "num": 42}`

### Operations
- **Arithmetic**: `+`, `-`, `*`, `/` (with automatic type promotion)
- **Comparisons**: `<`, `>`, `<=`, `>=`, `==`, `!=`
- **Custom binary operators**: Extensible through function registration

### No Built-in Functions
MLang starts with **no built-in functions**. All functions must be registered by the user:
- Functions via `RegisterFunction()`
- Binary operators via `RegisterBinaryOp()`
- Variadic functions via `RegisterVariadicFunction()`

### Extensibility
- Custom functions via `RegisterFunction()`
- Custom binary operators via `RegisterBinaryOp()`
- Variadic functions via `RegisterVariadicFunction()`

## Usage

### Command Line
```bash
# Generate Go expression (basic arithmetic works without registration)
./mlang -e "1 + 2 * 3"

# Arrays and dictionaries work without registration
./mlang -e "[1, 2, 3, 4]"
./mlang -e '{"key": "value"}'

# Functions need to be registered in code (see programmatic API)
./mlang -i input.mlang -o output.go
```

### Programmatic API
```go
compiler := compiler.NewCompiler()

// Register user functions (no built-ins exist)
compiler.RegisterFunction("max", []semantic.Type{semantic.IntType, semantic.IntType}, semantic.IntType)
compiler.RegisterFunction("len", []semantic.Type{semantic.AnyType}, semantic.IntType)
compiler.RegisterBinaryOp("contains", semantic.StringType, semantic.StringType, semantic.BooleanType)

// Compile to Go expression
result := compiler.CompileToGo("max(1, 2) + len([3, 4, 5])")
if len(result.Errors) == 0 {
    fmt.Printf("Go Expression: %s\n", result.GoCode)
    fmt.Printf("Result Type: %s\n", result.ResultType)
}
// Output: max(1, 2) + len([]int{3, 4, 5})
```

## Examples

### Basic Arithmetic
```mlang
1 + 2 * 3          // → 7
1.5 + 2.0          // → 3.5 (float arithmetic)
1 + 2.5            // → 3.5 (int promoted to float)
```

### Data Structures
```mlang
[1, 2, 3, 4]                    // → []int{1, 2, 3, 4}
[1, 2.5, 3]                     // → []float64{1, 2.5, 3} (promotion)
{"name": "mlang", "version": 1} // → map[string]interface{}{"name": "mlang", "version": 1}
```

### User-Registered Functions
```mlang
// After registering max, len, concat functions:
max(5, 3)                     // → max(5, 3)
len([1, 2, 3])               // → len([]int{1, 2, 3})
concat("Hello", "World")     // → concat("Hello", "World")
"text" contains "ext"        // → contains("text", "ext")
```

### Generated Go Expressions

| Input | Output |
|-------|--------|
| `1 + 2.5 * 3` | `(float64(1) + (2.5 * float64(3)))` |
| `[1, 2, 3, 4]` | `[]int{1, 2, 3, 4}` |
| `{"key": "value"}` | `map[string]string{"key": "value"}` |
| `max(5, 3)` | `max(5, 3)` (after registration) |
| `len([1, 2, 3])` | `len([]int{1, 2, 3})` (after registration) |

## Architecture Benefits

### Clean Separation of Concerns
- **Semantic Validation**: Only validates correctness, no implementation decisions
- **IR Layer**: Handles type promotion, optimization decisions  
- **Code Generation**: Target-language specific implementations

### Extensibility
- Easy to add new target languages
- Pluggable function registration
- Type system extensibility

### Type Safety
- Static type checking and inference
- Automatic numeric type promotion
- Compile-time error detection

### Pure Expression Generation
- Generates only Go expressions, not full programs
- Easy to embed in existing Go codebases
- No assumptions about runtime environment

## Building and Testing

```bash
# Build the compiler
go build -o mlang ./cmd/main.go

# Run tests
go test ./go/compiler -v
go test ./go/semantic -v
go test ./go/ir -v

# Run examples  
./mlang -e "1 + 2 * 3"
./mlang -e "[1, 2, 3]"
./mlang -e '{"name": "mlang"}'
```

## Future Extensions

- **More Target Languages**: Rust, Python, JavaScript, C++ expression generation
- **Negative Numbers**: Support for unary minus operator (`-42`)
- **Control Flow**: `if`, `for`, `while` constructs
- **User-Defined Functions**: Function definition syntax in MLang itself
- **Modules**: Import/export system for reusable function libraries
- **Optimizations**: Constant folding, dead code elimination

The architecture is designed to support these extensions with minimal changes to the core semantic validation and IR layers.

## Expression Embedding Examples

The generated Go expressions can be embedded directly in Go programs:

```go
package main

import "fmt"

// User-provided functions that MLang expressions can call
func max(a, b int) int {
    if a > b { return a }
    return b
}

func len(arr []int) int {
    return len(arr)
}

func main() {
    // MLang generates these expressions:
    arithmetic := (float64(1) + (2.5 * float64(3)))
    array := []int{1, 2, 3, 4}
    maxResult := max(5, 3)
    arrayLen := len([]int{1, 2, 3})
    
    fmt.Printf("Arithmetic: %g\n", arithmetic)   // 8.5
    fmt.Printf("Array: %v\n", array)             // [1 2 3 4]  
    fmt.Printf("Max: %d\n", maxResult)           // 5
    fmt.Printf("Length: %d\n", arrayLen)         // 3
}
```

This makes MLang perfect for:
- **Configuration files** with computed values
- **Template systems** with embedded expressions  
- **DSL implementations** for domain-specific calculations
- **Code generation** pipelines
- **Mathematical expression** compilation