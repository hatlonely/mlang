"""
测试集合类型的解析
包括数组和字典的各种用法
"""
import unittest
from test_base import MlangTestBase


class TestCollections(MlangTestBase):
    """集合类型测试"""
    
    def test_empty_array(self):
        """测试空数组"""
        self.assert_parse_success("[]", "空数组应该解析成功")
    
    def test_single_element_arrays(self):
        """测试单元素数组"""
        valid_arrays = [
            "[1]",
            "[42.5]", 
            '["hello"]',
            "[TRUE]",
            "[variable]"
        ]
        
        for array in valid_arrays:
            with self.subTest(array=array):
                self.assert_parse_success(array, f"单元素数组 {array} 应该解析成功")
    
    def test_multi_element_arrays(self):
        """测试多元素数组"""
        valid_arrays = [
            "[1, 2, 3]",
            "[1.1, 2.2, 3.3]",
            '["a", "b", "c"]',
            "[TRUE, FALSE, TRUE]",
            "[x, y, z]",
            "[1, 2.5, TRUE, variable]",  # 混合类型
            "[1,2,3]",  # 无空格
            "[ 1 , 2 , 3 ]"  # 多空格
        ]
        
        for array in valid_arrays:
            with self.subTest(array=array):
                self.assert_parse_success(array, f"多元素数组 {array} 应该解析成功")
    
    def test_empty_dictionary(self):
        """测试空字典"""
        self.assert_parse_success("{}", "空字典应该解析成功")
    
    def test_single_pair_dictionaries(self):
        """测试单键值对字典"""
        valid_dicts = [
            '{"key": "value"}',
            "{1: 2}",
            "{TRUE: FALSE}",
            "{variable: 42}",
            '{"name": variable}'
        ]
        
        for dictionary in valid_dicts:
            with self.subTest(dictionary=dictionary):
                self.assert_parse_success(dictionary, f"单键值对字典 {dictionary} 应该解析成功")
    
    def test_multi_pair_dictionaries(self):
        """测试多键值对字典"""
        valid_dicts = [
            '{"key1": "value1", "key2": "value2"}',
            "{1: 2, 3: 4}",
            '{"name": "test", "age": 25, "active": TRUE}',
            '{x: y, a: b}',  # 变量键值
            '{"a":1,"b":2}',  # 无空格
            '{ "a" : 1 , "b" : 2 }'  # 多空格
        ]
        
        for dictionary in valid_dicts:
            with self.subTest(dictionary=dictionary):
                self.assert_parse_success(dictionary, f"多键值对字典 {dictionary} 应该解析成功")
    
    def test_nested_arrays(self):
        """测试嵌套数组"""
        valid_nested = [
            "[[1, 2], [3, 4]]",
            "[[], [1], [2, 3]]",
            "[[1, 2, 3], [4, 5, 6], [7, 8, 9]]",
            '[[1, "a"], [2, "b"]]'
        ]
        
        for nested in valid_nested:
            with self.subTest(nested=nested):
                self.assert_parse_success(nested, f"嵌套数组 {nested} 应该解析成功")
    
    def test_nested_dictionaries(self):
        """测试嵌套字典"""
        valid_nested = [
            '{"outer": {"inner": "value"}}',
            '{"a": {"b": 1, "c": 2}, "d": {"e": 3}}',
            '{1: {2: 3}}',
            '{"empty": {}}'
        ]
        
        for nested in valid_nested:
            with self.subTest(nested=nested):
                self.assert_parse_success(nested, f"嵌套字典 {nested} 应该解析成功")
    
    def test_mixed_collections(self):
        """测试混合集合类型"""
        valid_mixed = [
            '[{"key": "value"}]',  # 数组包含字典
            '{"key": [1, 2, 3]}',  # 字典包含数组
            '[{"a": 1}, {"b": 2}]',  # 数组包含多个字典
            '{"arrays": [[1, 2], [3, 4]], "dicts": [{"x": 1}, {"y": 2}]}'  # 复杂嵌套
        ]
        
        for mixed in valid_mixed:
            with self.subTest(mixed=mixed):
                self.assert_parse_success(mixed, f"混合集合 {mixed} 应该解析成功")
    
    def test_collections_with_expressions(self):
        """测试包含表达式的集合"""
        valid_expr_collections = [
            "[1 + 2, 3 * 4]",
            '{"sum": 1 + 2, "product": 3 * 4}',
            "[func(x), func(y)]",
            '{"result": func(42)}'
        ]
        
        for expr_collection in valid_expr_collections:
            with self.subTest(expr_collection=expr_collection):
                self.assert_parse_success(expr_collection, f"表达式集合 {expr_collection} 应该解析成功")
    
    def test_invalid_arrays(self):
        """测试无效的数组格式"""
        invalid_arrays = [
            "[1, 2,]",     # 尾随逗号
            "[,1, 2]",     # 前置逗号
            "[1,, 2]",     # 双逗号
            "[1 2]",       # 缺少逗号
            "[",           # 未闭合
            "]",           # 仅右括号
            "[[1, 2]",     # 嵌套未闭合
        ]
        
        for invalid_array in invalid_arrays:
            with self.subTest(invalid_array=invalid_array):
                self.assert_parse_failure(invalid_array, f"无效数组 {invalid_array} 应该解析失败")
    
    def test_invalid_dictionaries(self):
        """测试无效的字典格式"""
        invalid_dicts = [
            '{"key": "value",}',   # 尾随逗号
            '{,"key": "value"}',   # 前置逗号
            '{"key":, "value"}',   # 缺少值
            '{"key" "value"}',     # 缺少冒号
            '{"key": "value" "key2": "value2"}',  # 缺少逗号
            '{',                   # 未闭合
            '}',                   # 仅右括号
            '{"key": }',           # 缺少值
            '{: "value"}',         # 缺少键
        ]
        
        for invalid_dict in invalid_dicts:
            with self.subTest(invalid_dict=invalid_dict):
                self.assert_parse_failure(invalid_dict, f"无效字典 {invalid_dict} 应该解析失败")


if __name__ == '__main__':
    unittest.main()