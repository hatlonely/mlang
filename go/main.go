package main

import (
	"fmt"

	"github.com/hatlonely/mlang/go/semantic"
)

func main() {
	// 测试用例
	testCases := []string{
		"1 + 2",
		"true && false", // 这应该报错，因为语法中没有 && 运算符
		"[1, 2, 3]",
		"[1, \"hello\"]", // 类型不一致，应该报错
		"{\"name\": \"test\", \"age\": 25}",
		"{1: \"one\", \"two\": 2}", // 键类型不一致，应该报错
		"len([1, 2, 3])",
		"max(1, 2)",
		"max(1, \"2\")", // 参数类型错误
		"unknown_func()", // 未定义函数
		"1 > 2",
		"\"hello\" + \"world\"", // 字符串不支持加法
	}

	for i, testCase := range testCases {
		fmt.Printf("\n=== Test Case %d: %s ===\n", i+1, testCase)
		
		errors, resultType := semantic.AnalyzeCode(testCase)
		
		if len(errors) > 0 {
			fmt.Println("Semantic Errors:")
			for _, err := range errors {
				fmt.Printf("  %s\n", err.Error())
			}
		} else {
			fmt.Println("No semantic errors")
		}
		
		fmt.Printf("Result type: %s\n", resultType.String())
	}
}