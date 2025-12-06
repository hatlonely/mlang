"""
mlang 文法测试基础类
提供通用的解析器测试方法
"""
import sys
import os
import unittest
from antlr4 import *
from antlr4.error.ErrorListener import ErrorListener

# 添加gen/py到Python路径
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'gen', 'py'))

from mlangLexer import mlangLexer
from mlangParser import mlangParser


class ParseErrorListener(ErrorListener):
    """自定义错误监听器，收集解析错误"""
    
    def __init__(self):
        super().__init__()
        self.errors = []
    
    def syntaxError(self, recognizer, offendingSymbol, line, column, msg, e):
        self.errors.append(f"Line {line}:{column} - {msg}")


class MlangTestBase(unittest.TestCase):
    """mlang 文法测试基础类"""
    
    def parse_expression(self, text):
        """
        解析表达式并返回解析树和错误信息
        
        Args:
            text: 要解析的文本
            
        Returns:
            tuple: (parse_tree, errors_list)
        """
        # 创建输入流
        input_stream = InputStream(text + "\n")
        
        # 创建词法分析器
        lexer = mlangLexer(input_stream)
        
        # 创建错误监听器
        error_listener = ParseErrorListener()
        lexer.removeErrorListeners()
        lexer.addErrorListener(error_listener)
        
        # 创建token流
        token_stream = CommonTokenStream(lexer)
        
        # 创建语法分析器
        parser = mlangParser(token_stream)
        parser.removeErrorListeners()
        parser.addErrorListener(error_listener)
        
        # 解析
        try:
            tree = parser.prog()
            return tree, error_listener.errors
        except Exception as e:
            error_listener.errors.append(f"Parse exception: {str(e)}")
            return None, error_listener.errors
    
    def assert_parse_success(self, text, msg=None):
        """断言表达式解析成功"""
        tree, errors = self.parse_expression(text)
        if errors:
            error_msg = f"Failed to parse '{text}': {'; '.join(errors)}"
            if msg:
                error_msg = f"{msg}: {error_msg}"
            self.fail(error_msg)
        self.assertIsNotNone(tree, f"Parse tree is None for: {text}")
        return tree
    
    def assert_parse_failure(self, text, msg=None):
        """断言表达式解析失败"""
        tree, errors = self.parse_expression(text)
        if not errors:
            error_msg = f"Expected parsing to fail for '{text}', but it succeeded"
            if msg:
                error_msg = f"{msg}: {error_msg}"
            self.fail(error_msg)
        return errors
    
    def get_parse_tree_text(self, tree, parser):
        """获取解析树的文本表示"""
        return tree.toStringTree(recog=parser)