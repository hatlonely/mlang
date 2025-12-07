package semantic_test

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hatlonely/mlang/go/semantic"
	parser "github.com/hatlonely/mlang/gen/go"
)

func TestDemo(t *testing.T) {
	fmt.Println("=== mlang 语义分析器演示 ===")

	// 创建一个自定义验证器来演示注册功能
	validator := semantic.NewPureValidator()

	// 注册自定义函数
	fmt.Println("注册自定义函数:")
	validator.RegisterFunction("sqrt", []semantic.Type{semantic.NumericBaseType}, semantic.NumericBaseType)
	validator.RegisterFunction("isEmpty", []semantic.Type{semantic.AnyType}, semantic.BooleanType)
	
	// 注册之前的内置函数
	validator.RegisterFunction("len", []semantic.Type{semantic.AnyType}, semantic.IntType)
	validator.RegisterFunction("abs", []semantic.Type{semantic.NumericBaseType}, semantic.NumericBaseType)
	validator.RegisterVariadicFunction("max", []semantic.Type{semantic.NumericBaseType}, semantic.NumericBaseType, semantic.NumericBaseType)
	validator.RegisterVariadicFunction("min", []semantic.Type{semantic.NumericBaseType}, semantic.NumericBaseType, semantic.NumericBaseType)
	validator.RegisterVariadicFunction("sum", []semantic.Type{}, semantic.NumericBaseType, semantic.NumericBaseType)
	validator.RegisterVariadicFunction("concat", []semantic.Type{}, semantic.StringType, semantic.StringType)

	// 注册可变参数函数
	fmt.Println("注册可变参数函数:")
	validator.RegisterVariadicFunction("add", []semantic.Type{}, semantic.NumericBaseType, semantic.NumericBaseType)
	validator.RegisterVariadicFunction("multiply", []semantic.Type{semantic.NumericBaseType}, semantic.NumericBaseType, semantic.NumericBaseType) // 至少1个参数

	// 注册自定义比较运算符
	validator.RegisterBinaryOp("contains", semantic.StringType, semantic.StringType, semantic.BooleanType)
	validator.RegisterBinaryOp("startsWith", semantic.StringType, semantic.StringType, semantic.BooleanType)

	// 列出所有注册的函数
	functions := validator.GetSymbolTable().ListFunctions()
	for name, funcType := range functions {
		fmt.Printf("  %s: %s\n", name, funcType.String())
	}

	fmt.Println("\n=== 基础测试用例 ===")
	basicTestCases := []string{
		"1 + 2",           // int + int -> int
		"1.5 + 2.3",       // float + float -> float
		"1 + 2.5",         // int + float -> float
		"3.14 * 2",        // float * int -> float
		"[1, 2, 3]",       // int array
		"[1.1, 2.2, 3.3]", // float array
		"[1, \"hello\"]",  // 类型不一致，应该报错
		"{\"name\": \"test\", \"age\": 25}",
		"len([1, 2, 3])",
	}

	runTestCases(basicTestCases, validator)

	fmt.Println("\n=== 内置可变参数函数测试 ===")
	variadicTestCases := []string{
		"max(1)",                  // 单个参数 - int
		"max(1.5)",                // 单个参数 - float
		"max(1, 2)",               // 两个参数 - int
		"max(1.1, 2.2)",           // 两个参数 - float
		"max(1, 2.5)",             // 混合数值类型
		"max(1, 2, 3, 4)",         // 多个参数 - int
		"min(5.5, 3.1, 8.9, 1.2)", // min函数 - float
		"sum(1, 2, 3, 4, 5)",      // sum函数 - int
		"sum(1.1, 2.2, 3.3)",      // sum函数 - float
		"concat(\"hello\", \" \", \"world\", \"!\")", // 字符串拼接
		"max(1, \"2\")",      // 类型错误
		"sum()",              // 空参数（应该允许）
		"concat(\"a\", 123)", // 混合类型错误
	}

	runTestCases(variadicTestCases, validator)

	fmt.Println("\n=== 自定义可变参数函数测试 ===")
	customVariadicTestCases := []string{
		"add(1, 2, 3)",      // 自定义add函数
		"add()",             // 无参数（应该允许）
		"multiply(2)",       // 最少1个参数
		"multiply(2, 3, 4)", // 多个参数
		"multiply()",        // 参数不足（应该报错）
		"add(1, \"2\")",     // 类型错误
	}

	runTestCases(customVariadicTestCases, validator)

	fmt.Println("\n=== 固定参数函数测试 ===")
	fixedFuncTestCases := []string{
		"sqrt(16)",          // sqrt with int
		"sqrt(16.5)",        // sqrt with float
		"abs(5)",            // abs with int (改为正数)
		"abs(3.14)",         // abs with float (改为正数)
		"sqrt(\"invalid\")", // 类型错误
		"isEmpty([1, 2, 3])",
		"isEmpty(\"\")",
	}

	runTestCases(fixedFuncTestCases, validator)

	fmt.Println("\n=== 自定义比较运算符测试 ===")
	customCompareTestCases := []string{
		"\"hello world\" contains \"world\"",
		"\"test\" startsWith \"te\"",
		"\"hello\" contains 123", // 类型错误
		"42 startsWith \"4\"",    // 类型错误
	}

	runTestCases(customCompareTestCases, validator)

	fmt.Println("\n=== 可变参数函数管理演示 ===")
	fmt.Println("注册一个更复杂的可变参数函数 (average):")
	validator.RegisterVariadicFunction("average", []semantic.Type{}, semantic.NumericBaseType, semantic.NumericBaseType)

	fmt.Println("测试 average 函数:")
	averageTests := []string{
		"average(1, 2, 3, 4, 5)",
		"average(10)",
		"average()",
		"average(1, 2, \"3\")", // 类型错误
	}
	for _, test := range averageTests {
		errors, resultType := analyzeCodeWithValidator(test, validator)
		if len(errors) > 0 {
			fmt.Printf("  %s -> 错误: %s\n", test, errors[0].Error())
		} else {
			fmt.Printf("  %s -> 成功: %s\n", test, resultType.String())
		}
	}

	fmt.Println("\n删除可变参数函数 add:")
	validator.GetSymbolTable().UnregisterFunction("add")

	errors, _ := analyzeCodeWithValidator("add(1, 2, 3)", validator)
	if len(errors) > 0 {
		fmt.Printf("  预期错误: %s\n", errors[0].Error())
	}

	fmt.Println("\n=== 最终函数列表 ===")
	finalFunctions := validator.GetSymbolTable().ListFunctions()
	for name, funcType := range finalFunctions {
		fmt.Printf("  %s: %s\n", name, funcType.String())
	}
}

func runTestCases(testCases []string, validator *semantic.PureValidator) {
	for i, testCase := range testCases {
		fmt.Printf("\nTest %d: %s\n", i+1, testCase)

		errors, resultType := analyzeCodeWithValidator(testCase, validator)

		if len(errors) > 0 {
			fmt.Println("  语义错误:")
			for _, err := range errors {
				fmt.Printf("    %s\n", err.Error())
			}
		} else {
			fmt.Println("  ✓ 无语义错误")
		}

		fmt.Printf("  结果类型: %s\n", resultType.String())
	}
}

// Helper function to analyze code with PureValidator
func analyzeCodeWithValidator(input string, validator *semantic.PureValidator) ([]*semantic.SemanticError, semantic.Type) {
	lexer := parser.NewmlangLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewmlangParser(stream)

	// Parse the input
	tree := p.Prog()

	// Clear previous errors
	validator.ClearErrors()

	// Validate with the provided validator
	valid := validator.ValidateProgram(tree.(*parser.ProgContext))
	if !valid {
		return validator.GetErrors(), semantic.VoidType
	}

	// Get result type from last statement
	progCtx := tree.(*parser.ProgContext)
	var lastType semantic.Type = semantic.VoidType
	for _, statCtx := range progCtx.AllStat() {
		if exprCtx := statCtx.Expr(); exprCtx != nil {
			lastType = validator.InferType(exprCtx)
		}
	}

	return validator.GetErrors(), lastType
}
