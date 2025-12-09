#!/bin/bash

echo "=== mlang 赋值语句功能测试 ==="

echo
echo "1. 简单变量赋值："
./mlang -e "userAge = 25"
./mlang -e "userName = \"Bob\""
./mlang -e "isActive = true"

echo
echo "2. 数组元素赋值："
./mlang -e "numbers[0] = 100"
./mlang -e "names[1] = \"Alice\""

echo
echo "3. 字典赋值："
./mlang -e 'configs["timeout"] = "30s"'
./mlang -e 'lookup["key"] = 42'

echo
echo "4. 结构体字段赋值："
./mlang -e "user.name = \"Charlie\""
./mlang -e "point.x = 3.14"

echo
echo "5. 嵌套赋值："
./mlang -e "users[0].age = 30"
./mlang -e "company.address.city = \"Beijing\""

echo
echo "6. 类型转换赋值（int to float）："
./mlang -e "timeout = 42"  # int 42 赋值给 float 类型的 timeout

echo
echo "=== 验证表达式功能仍然正常 ==="
./mlang -e "userAge + 10"
./mlang -e "max(1, 2, 3)"
./mlang -e "[1, 2, 3]"

echo
echo "=== 错误情况测试 ==="
echo "未定义变量赋值错误："
./mlang -e "undefined = 25" 2>&1 || true

echo
echo "类型不兼容赋值错误："
./mlang -e 'userAge = "not a number"' 2>&1 || true

echo
echo "数组索引类型错误："
./mlang -e 'numbers["invalid"] = 100' 2>&1 || true

echo
echo "字典索引类型错误："
./mlang -e 'lookup[123.45] = "error"' 2>&1 || true

echo
echo "=== 测试完成 ==="