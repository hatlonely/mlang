#!/usr/bin/env python3
"""
mlang 文法测试运行器
运行所有测试并生成报告
"""
import unittest
import sys
import os

# 添加当前目录到Python路径
sys.path.insert(0, os.path.dirname(__file__))

# 导入测试模块
from test_basic_types import TestBasicTypes
from test_expressions import TestExpressions
from test_complex_cases import TestComplexCases
from test_collections import TestCollections


def create_test_suite():
    """创建测试套件"""
    suite = unittest.TestSuite()
    
    # 添加基本类型测试
    suite.addTest(unittest.makeSuite(TestBasicTypes))
    
    # 添加表达式测试
    suite.addTest(unittest.makeSuite(TestExpressions))
    
    # 添加集合类型测试（数组和字典）
    suite.addTest(unittest.makeSuite(TestCollections))
    
    # 添加复杂案例测试
    suite.addTest(unittest.makeSuite(TestComplexCases))
    
    return suite


def run_tests():
    """运行所有测试"""
    print("=" * 60)
    print("mlang 文法解析器测试")
    print("=" * 60)
    
    # 检查ANTLR4是否安装
    try:
        import antlr4
        print("✓ ANTLR4 运行时已安装")
    except ImportError:
        print("✗ 错误: ANTLR4 Python运行时未安装")
        print("请运行: pip install antlr4-python3-runtime")
        sys.exit(1)
    
    # 检查生成的解析器文件
    gen_py_dir = os.path.join(os.path.dirname(__file__), '..', 'gen', 'py')
    required_files = ['mlangLexer.py', 'mlangParser.py']
    
    for filename in required_files:
        filepath = os.path.join(gen_py_dir, filename)
        if os.path.exists(filepath):
            print(f"✓ 找到生成文件: {filename}")
        else:
            print(f"✗ 缺少生成文件: {filename}")
            print("请先运行ANTLR生成解析器代码")
            sys.exit(1)
    
    print("\n" + "-" * 60)
    print("开始运行测试...")
    print("-" * 60)
    
    # 创建测试套件并运行
    suite = create_test_suite()
    runner = unittest.TextTestRunner(verbosity=2)
    result = runner.run(suite)
    
    # 输出结果摘要
    print("\n" + "=" * 60)
    print("测试结果摘要:")
    print(f"总测试数: {result.testsRun}")
    print(f"成功: {result.testsRun - len(result.failures) - len(result.errors)}")
    print(f"失败: {len(result.failures)}")
    print(f"错误: {len(result.errors)}")
    
    if result.failures:
        print(f"\n失败的测试:")
        for test, traceback in result.failures:
            print(f"  - {test}")
    
    if result.errors:
        print(f"\n错误的测试:")
        for test, traceback in result.errors:
            print(f"  - {test}")
    
    print("=" * 60)
    
    # 返回成功状态
    return len(result.failures) == 0 and len(result.errors) == 0


if __name__ == '__main__':
    success = run_tests()
    sys.exit(0 if success else 1)