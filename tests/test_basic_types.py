"""
测试基本类型的解析
包括数字、字符串、布尔值和标识符
"""
import unittest
from test_base import MlangTestBase


class TestBasicTypes(MlangTestBase):
    """基本类型测试"""
    
    def test_integers(self):
        """测试整数字面量"""
        valid_integers = [
            "42",
            "0", 
            "123",
            "999"
        ]
        
        for integer in valid_integers:
            with self.subTest(integer=integer):
                self.assert_parse_success(integer, f"整数 {integer} 应该解析成功")
    
    def test_floats(self):
        """测试浮点数字面量"""
        valid_floats = [
            "3.14",
            "0.0",
            "123.456",
            "99.99"
        ]
        
        for float_num in valid_floats:
            with self.subTest(float_num=float_num):
                self.assert_parse_success(float_num, f"浮点数 {float_num} 应该解析成功")
    
    def test_invalid_numbers(self):
        """测试无效的数字格式"""
        invalid_numbers = [
            "42.",      # 以点结尾
            ".14",      # 以点开头
            "42.14.5",  # 多个小数点
            "1.2.3.4"   # 多个小数点
        ]
        
        for invalid_num in invalid_numbers:
            with self.subTest(invalid_num=invalid_num):
                self.assert_parse_failure(invalid_num, f"无效数字 {invalid_num} 应该解析失败")
    
    def test_double_quote_strings(self):
        """测试双引号字符串"""
        valid_strings = [
            '"hello"',
            '"hello world"',
            '"含有中文的字符串"',
            '""',  # 空字符串
            '"包含\\"转义\\"字符"',
            '"包含\\n换行符"'
        ]
        
        for string in valid_strings:
            with self.subTest(string=string):
                self.assert_parse_success(string, f"双引号字符串 {string} 应该解析成功")
    
    def test_backtick_strings(self):
        """测试反引号字符串"""
        valid_strings = [
            '`hello`',
            '`raw string with "quotes"`',
            '``',  # 空字符串
            '`包含 \\ 反斜杠的原始字符串`'
        ]
        
        for string in valid_strings:
            with self.subTest(string=string):
                self.assert_parse_success(string, f"反引号字符串 {string} 应该解析成功")
    
    def test_invalid_strings(self):
        """测试无效的字符串格式"""
        invalid_strings = [
            '"未闭合字符串',        # 未闭合的双引号
            "'单引号字符串'",       # 不支持单引号
            '`未闭合反引号字符串',   # 未闭合的反引号
        ]
        
        for invalid_str in invalid_strings:
            with self.subTest(invalid_str=invalid_str):
                self.assert_parse_failure(invalid_str, f"无效字符串 {invalid_str} 应该解析失败")
    
    def test_boolean_literals(self):
        """测试布尔字面量"""
        valid_booleans = [
            "TRUE",
            "FALSE", 
            "true",
            "false"
        ]
        
        for boolean in valid_booleans:
            with self.subTest(boolean=boolean):
                self.assert_parse_success(boolean, f"布尔值 {boolean} 应该解析成功")
    
    def test_identifiers(self):
        """测试标识符"""
        valid_identifiers = [
            "variable",
            "myFunction",
            "_private_var",
            "counter123",
            "_",
            "a",
            "CamelCase",
            "snake_case_name"
        ]
        
        for identifier in valid_identifiers:
            with self.subTest(identifier=identifier):
                self.assert_parse_success(identifier, f"标识符 {identifier} 应该解析成功")
    
    def test_invalid_identifiers(self):
        """测试无效的标识符"""
        invalid_identifiers = [
            "123abc",     # 以数字开头
            "func#name",   # 包含特殊字符  
            "42variable"   # 以数字开头
        ]
        
        for invalid_id in invalid_identifiers:
            with self.subTest(invalid_id=invalid_id):
                self.assert_parse_failure(invalid_id, f"无效标识符 {invalid_id} 应该解析失败")


if __name__ == '__main__':
    unittest.main()