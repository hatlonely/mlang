"""
测试复杂表达式和边界案例
包括深度嵌套、运算符优先级、综合测试等
"""
import unittest
from test_base import MlangTestBase


class TestComplexCases(MlangTestBase):
    """复杂案例测试"""
    
    def test_operator_precedence(self):
        """测试运算符优先级"""
        expressions_with_precedence = [
            ("2 + 3 * 4", "应该是 2 + (3 * 4)"),
            ("10 / 2 + 3", "应该是 (10 / 2) + 3"),
            ("1 + 2 * 3 + 4", "应该是 1 + (2 * 3) + 4"),
            ("(1 + 2) * (3 + 4)", "显式括号优先级"),
            ("2 * 3 + 4 * 5", "应该是 (2 * 3) + (4 * 5)"),
            ("a + b * c / d - e", "复杂的优先级组合")
        ]
        
        for expr, description in expressions_with_precedence:
            with self.subTest(expr=expr, desc=description):
                self.assert_parse_success(expr, f"优先级测试: {description}")
    
    def test_deeply_nested_expressions(self):
        """测试深度嵌套表达式"""
        nested_expressions = [
            "((((1))))",
            "func(func(func(x)))",
            "(((a + b) * c) + d)",
            "add(mul(a, b), div(c, d))",
            "compare(func1(x, y), func2(a, b))",
            "calculate((x + y) * (z - w), (a / b) + c)"
        ]
        
        for expr in nested_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"深度嵌套表达式 {expr} 应该解析成功")
    
    def test_complex_function_calls(self):
        """测试复杂的函数调用"""
        complex_calls = [
            'process(1, "hello", TRUE, 3.14)',
            'func(add(1, 2), sub(5, 3), "result")',
            'calculate((x + y), (a > b), func(z))',
            'nested(outer(inner(value)), "test")',
            'transform(data, filter(items, predicate), "output")',
            'api_call("endpoint", params, headers, TRUE)'
        ]
        
        for call in complex_calls:
            with self.subTest(call=call):
                self.assert_parse_success(call, f"复杂函数调用 {call} 应该解析成功")
    
    def test_comparison_chains(self):
        """测试比较链"""
        comparison_expressions = [
            "a > b",
            "x <= y",
            "value == expected",
            "result != NULL",
            "score >= threshold",
            "count < limit"
        ]
        
        for expr in comparison_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"比较表达式 {expr} 应该解析成功")
    
    def test_mixed_complex_expressions(self):
        """测试混合复杂表达式"""
        complex_mixed = [
            'func(1 + 2, "test", x > y)',
            'calculate((a + b) * c, result != "")',
            'process(data, filter == TRUE, count >= 10)',
            'transform(input, (x * 2) + offset, "mode")',
            'validate(user, age >= 18, name != "")',
            'query(table, id == key, active == TRUE)'
        ]
        
        for expr in complex_mixed:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"混合复杂表达式 {expr} 应该解析成功")
    
    def test_edge_cases(self):
        """测试边界案例"""
        edge_cases = [
            '""',           # 空字符串
            "``",           # 空原始字符串
            "func()",       # 无参函数
            "a",            # 单个标识符
            "0",            # 零
            "0.0"           # 零浮点数
        ]
        
        for case in edge_cases:
            with self.subTest(case=case):
                self.assert_parse_success(case, f"边界案例 {case} 应该解析成功")
    
    def test_whitespace_tolerance(self):
        """测试空白字符容忍度"""
        expressions_with_spaces = [
            "1+2",          # 无空格
            "1 + 2",        # 正常空格
            "1  +  2",      # 多个空格
            "func()",       # 函数调用无空格
            "func( )",      # 函数调用有空格
            "func( 1 , 2 )", # 参数间有空格
            "( 1 + 2 )",    # 括号内有空格
            "a\t+\tb"       # 制表符
        ]
        
        for expr in expressions_with_spaces:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"空白字符测试 {repr(expr)} 应该解析成功")
    
    def test_syntax_error_recovery(self):
        """测试语法错误"""
        syntax_errors = [
            "1 + + 2",      # 连续运算符
            "func(",        # 不完整的函数调用
            "(1 + 2",       # 不匹配的括号
            "1 2 3",        # 缺少运算符
            "func(1, , 2)", # 空的函数参数
            "1 + * 2",      # 运算符错误
            "...",          # 无效的符号
            "1 + 2 +",      # 不完整的表达式
            "= 1",          # 运算符开头
        ]
        
        for error_case in syntax_errors:
            with self.subTest(error_case=error_case):
                self.assert_parse_failure(error_case, f"语法错误 {repr(error_case)} 应该解析失败")
    
    def test_examples_from_file(self):
        """测试examples.mlang文件中的有效表达式"""
        valid_examples = [
            # 数字
            "42", "3.14", "123.456",
            
            # 字符串
            '"hello world"',
            '`原始字符串`',
            
            # 布尔值
            "TRUE", "FALSE", "true", "false",
            
            # 标识符
            "variable_name", "myFunction", "_private_var", "counter123",
            
            # 算术表达式
            "1 + 2", "3.14 * 2", "10 / 3", "15 - 7", "2 + 3 * 4", "(2 + 3) * 4",
            
            # 比较表达式
            "5 > 3", "10 <= 20", "42 == 42", "3.14 >= 3.0", "x < y",
            
            # 函数调用
            'print("hello world")', "add(1, 2)", "max(10, 20, 30)", "calculate()",
            
            # 复合表达式
            "add(1 + 2, 3 * 4)", "func1(func2(x), y + z)",
            
            # 中缀比较
            "x equals y", "a contains b", "str matches pattern"
        ]
        
        for example in valid_examples:
            with self.subTest(example=example):
                self.assert_parse_success(example, f"示例 {example} 应该解析成功")


if __name__ == '__main__':
    unittest.main()