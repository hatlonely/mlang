"""
测试各种表达式的解析
包括算术表达式、比较表达式、函数调用等
"""
import unittest
from test_base import MlangTestBase


class TestExpressions(MlangTestBase):
    """表达式测试"""
    
    def test_arithmetic_expressions(self):
        """测试算术表达式"""
        valid_expressions = [
            "1 + 2",
            "3 - 4", 
            "5 * 6",
            "7 / 8",
            "2 + 3 * 4",    # 运算符优先级
            "(2 + 3) * 4",  # 括号
            "a + b",
            "x * y + z",
            "3.14 * 2",
            "10 / 3"
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"算术表达式 {expr} 应该解析成功")
    
    def test_comparison_expressions(self):
        """测试比较表达式"""
        valid_expressions = [
            "5 > 3",
            "10 <= 20", 
            "42 == 42",
            "x != y",
            "3.14 >= 3.0",
            "a < b",
            '"hello" == "world"',
            'name != ""',
            "TRUE == FALSE"
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"比较表达式 {expr} 应该解析成功")
    
    def test_function_calls(self):
        """测试函数调用"""
        valid_expressions = [
            "print()",
            'print("hello")',
            "add(1, 2)",
            "max(10, 20, 30)",
            "calculate()",
            'process("data", 123, true)',
            "func(x, y, z)"
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"函数调用 {expr} 应该解析成功")
    
    def test_infix_comparison_functions(self):
        """测试中缀比较函数"""
        valid_expressions = [
            "x equals y",
            "a contains b", 
            "str matches pattern",
            "data includes value"
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"中缀比较函数 {expr} 应该解析成功")
    
    def test_parenthesized_expressions(self):
        """测试括号表达式"""
        valid_expressions = [
            "(1)",
            "(1 + 2)",
            "((a + b))",
            "(func(x))",
            '("hello")',
            "(TRUE)",
            "(variable)"
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"括号表达式 {expr} 应该解析成功")
    
    def test_nested_expressions(self):
        """测试嵌套表达式"""
        valid_expressions = [
            "add(1 + 2, 3 * 4)",
            "func(func2(x), y + z)",
            "((x + y) * z) / (a - b)",
            'compare("hello", "world")',
            "calculate(3.14)",
            "func((x + y), (z * 2))"
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"嵌套表达式 {expr} 应该解析成功")
    
    def test_mixed_type_expressions(self):
        """测试混合类型表达式"""
        valid_expressions = [
            'age >= 18',
            'name != ""',
            'score > 90',
            'calculate(3.14, "pi")',
            'func(1, "test", TRUE)',
            'process(42, variable, "string")'
        ]
        
        for expr in valid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_success(expr, f"混合类型表达式 {expr} 应该解析成功")
    
    def test_invalid_expressions(self):
        """测试无效表达式"""
        invalid_expressions = [
            "1 +",           # 不完整的表达式
            "* 2",           # 运算符开头
            "1 + + 2",       # 连续运算符
            "func(",         # 未闭合的函数调用
            "func)",         # 语法错误
            "(1 + 2",        # 未闭合的括号
            "1 + 2)",        # 不匹配的括号
            "1 2",           # 缺少运算符
            "func(1, )",     # 函数参数错误
        ]
        
        for expr in invalid_expressions:
            with self.subTest(expr=expr):
                self.assert_parse_failure(expr, f"无效表达式 {expr} 应该解析失败")


if __name__ == '__main__':
    unittest.main()