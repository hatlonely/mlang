package main

import (
	"fmt"

	"github.com/hatlonely/mlang/go/semantic"
)

func main() {
	fmt.Println("=== mlang 语义分析器演示 ===\n")

	// 创建一个自定义分析器来演示注册功能
	analyzer := semantic.NewAnalyzer()

	// 注册自定义函数
	fmt.Println("注册自定义函数:")
	analyzer.RegisterFunction("sqrt", []semantic.Type{semantic.NumberType}, semantic.NumberType)
	analyzer.RegisterFunction("isEmpty", []semantic.Type{semantic.AnyType}, semantic.BooleanType)
	
	// 注册可变参数函数
	fmt.Println("注册可变参数函数:")
	analyzer.RegisterVariadicFunction("add", []semantic.Type{}, semantic.NumberType, semantic.NumberType)
	analyzer.RegisterVariadicFunction("multiply", []semantic.Type{semantic.NumberType}, semantic.NumberType, semantic.NumberType) // 至少1个参数

	// 注册自定义比较运算符
	analyzer.RegisterBinaryOp("contains", semantic.StringType, semantic.StringType, semantic.BooleanType)
	analyzer.RegisterBinaryOp("startsWith", semantic.StringType, semantic.StringType, semantic.BooleanType)

	// 列出所有注册的函数
	functions := analyzer.ListFunctions()
	for name, funcType := range functions {
		fmt.Printf("  %s: %s\n", name, funcType.String())
	}

	fmt.Println("\n=== 基础测试用例 ===")
	basicTestCases := []string{
		"1 + 2",
		"[1, 2, 3]",
		"[1, \"hello\"]", // 类型不一致，应该报错
		"{\"name\": \"test\", \"age\": 25}",
		"len([1, 2, 3])",
	}

	runTestCases(basicTestCases, analyzer)

	fmt.Println("\n=== 内置可变参数函数测试 ===")
	variadicTestCases := []string{
		"max(1)",           // 单个参数
		"max(1, 2)",        // 两个参数
		"max(1, 2, 3, 4)",  // 多个参数
		"min(5, 3, 8, 1)",  // min函数
		"sum(1, 2, 3, 4, 5)", // sum函数
		"concat(\"hello\", \" \", \"world\", \"!\")", // 字符串拼接
		"max(1, \"2\")",     // 类型错误
		"sum()",            // 空参数（应该允许）
		"concat(\"a\", 123)", // 混合类型错误
	}

	runTestCases(variadicTestCases, analyzer)

	fmt.Println("\n=== 自定义可变参数函数测试 ===")
	customVariadicTestCases := []string{
		"add(1, 2, 3)",      // 自定义add函数
		"add()",             // 无参数（应该允许）
		"multiply(2)",       // 最少1个参数
		"multiply(2, 3, 4)", // 多个参数
		"multiply()",        // 参数不足（应该报错）
		"add(1, \"2\")",     // 类型错误
	}

	runTestCases(customVariadicTestCases, analyzer)

	fmt.Println("\n=== 固定参数函数测试 ===")
	fixedFuncTestCases := []string{
		"sqrt(16)",
		"sqrt(\"invalid\")", // 类型错误
		"isEmpty([1, 2, 3])",
		"isEmpty(\"\")",
	}

	runTestCases(fixedFuncTestCases, analyzer)

	fmt.Println("\n=== 自定义比较运算符测试 ===")
	customCompareTestCases := []string{
		"\"hello world\" contains \"world\"",
		"\"test\" startsWith \"te\"",
		"\"hello\" contains 123", // 类型错误
		"42 startsWith \"4\"",    // 类型错误
	}

	runTestCases(customCompareTestCases, analyzer)

	fmt.Println("\n=== 可变参数函数管理演示 ===")
	fmt.Println("注册一个更复杂的可变参数函数 (average):")
	analyzer.RegisterVariadicFunction("average", []semantic.Type{}, semantic.NumberType, semantic.NumberType)
	
	fmt.Println("测试 average 函数:")
	averageTests := []string{
		"average(1, 2, 3, 4, 5)",
		"average(10)",
		"average()",
		"average(1, 2, \"3\")", // 类型错误
	}
	for _, test := range averageTests {
		errors, resultType := semantic.AnalyzeCodeWithAnalyzer(test, analyzer)
		if len(errors) > 0 {
			fmt.Printf("  %s -> 错误: %s\n", test, errors[0].Error())
		} else {
			fmt.Printf("  %s -> 成功: %s\n", test, resultType.String())
		}
	}

	fmt.Println("\n删除可变参数函数 add:")
	analyzer.UnregisterFunction("add")

	errors, _ := semantic.AnalyzeCodeWithAnalyzer("add(1, 2, 3)", analyzer)
	if len(errors) > 0 {
		fmt.Printf("  预期错误: %s\n", errors[0].Error())
	}

	fmt.Println("\n=== 最终函数列表 ===")
	finalFunctions := analyzer.ListFunctions()
	for name, funcType := range finalFunctions {
		fmt.Printf("  %s: %s\n", name, funcType.String())
	}
}

func runTestCases(testCases []string, analyzer *semantic.Analyzer) {
	for i, testCase := range testCases {
		fmt.Printf("\nTest %d: %s\n", i+1, testCase)

		errors, resultType := semantic.AnalyzeCodeWithAnalyzer(testCase, analyzer)

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
