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
	analyzer.RegisterFunction("concat", []semantic.Type{semantic.StringType, semantic.StringType}, semantic.StringType)
	analyzer.RegisterFunction("isEmpty", []semantic.Type{semantic.AnyType}, semantic.BooleanType)
	
	// 注册自定义比较运算符
	analyzer.RegisterCompareOp("contains", semantic.StringType, semantic.StringType, semantic.BooleanType)
	analyzer.RegisterCompareOp("startsWith", semantic.StringType, semantic.StringType, semantic.BooleanType)

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
		"max(1, 2)",
		"max(1, \"2\")", // 参数类型错误
	}

	runTestCases(basicTestCases, analyzer)

	fmt.Println("\n=== 自定义函数测试 ===")
	customFuncTestCases := []string{
		"sqrt(16)",
		"sqrt(\"invalid\")", // 类型错误
		"concat(\"hello\", \" world\")",
		"concat(1, 2)", // 类型错误
		"isEmpty([1, 2, 3])",
		"isEmpty(\"\")",
	}

	runTestCases(customFuncTestCases, analyzer)

	fmt.Println("\n=== 自定义比较运算符测试 ===")
	customCompareTestCases := []string{
		"\"hello world\" contains \"world\"",
		"\"test\" startsWith \"te\"",
		"\"hello\" contains 123", // 类型错误
		"42 startsWith \"4\"", // 类型错误
	}

	runTestCases(customCompareTestCases, analyzer)

	fmt.Println("\n=== 函数管理演示 ===")
	fmt.Println("删除 sqrt 函数:")
	analyzer.UnregisterFunction("sqrt")
	
	errors, _ := semantic.AnalyzeCodeWithAnalyzer("sqrt(16)", analyzer)
	if len(errors) > 0 {
		fmt.Printf("  预期错误: %s\n", errors[0].Error())
	}

	fmt.Println("\n重新注册 sqrt 函数 (支持字符串参数):")
	analyzer.RegisterFunction("sqrt", []semantic.Type{semantic.AnyType}, semantic.NumberType)
	
	errors, resultType := semantic.AnalyzeCodeWithAnalyzer("sqrt(\"16\")", analyzer)
	if len(errors) == 0 {
		fmt.Printf("  成功: 结果类型 %s\n", resultType.String())
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