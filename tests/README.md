# mlang 文法测试

这个目录包含了 mlang.g4 文法的 Python 测试套件。

## 测试结构

```
tests/
├── __init__.py              # Python包初始化
├── test_base.py             # 测试基础类
├── test_basic_types.py      # 基本类型测试
├── test_expressions.py      # 表达式测试  
├── test_complex_cases.py    # 复杂案例测试
├── run_tests.py            # 测试运行器
└── README.md               # 本文件
```

## 前置要求

1. **Python 3.6+**
2. **ANTLR4 Python运行时**:
   ```bash
   pip install antlr4-python3-runtime
   ```
3. **生成的解析器代码**: 
   确保在 `gen/py/` 目录中有ANTLR生成的Python代码

## 运行测试

### 运行所有测试
```bash
python tests/run_tests.py
```

### 运行单个测试文件
```bash
cd tests
python test_basic_types.py
python test_expressions.py  
python test_complex_cases.py
```

### 使用unittest运行
```bash
python -m unittest tests.test_basic_types
python -m unittest tests.test_expressions
python -m unittest tests.test_complex_cases
```

## 测试内容

### test_basic_types.py
- 整数和浮点数字面量
- 双引号和反引号字符串
- 布尔值 (TRUE/FALSE/true/false)
- 标识符
- 无效格式的错误案例

### test_expressions.py  
- 算术表达式 (+, -, *, /)
- 比较表达式 (>, <, ==, !=, >=, <=)
- 函数调用
- 中缀比较函数
- 括号表达式
- 嵌套表达式

### test_complex_cases.py
- 运算符优先级
- 深度嵌套表达式
- 复杂函数调用
- 混合类型表达式
- 边界案例
- 空白字符容忍度
- 语法错误恢复

## 测试输出

测试运行器会显示：
- ANTLR4运行时版本检查
- 生成文件检查
- 详细的测试结果
- 成功/失败/错误统计
- 失败测试的详细信息