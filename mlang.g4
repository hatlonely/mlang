grammar Expr;

// 解析规则
prog: stat+ EOF;

stat: expr NEWLINE # printExpr | NEWLINE # blank;

expr:
	expr op = ('*' | '/') expr									# MulDiv
	| expr op = ('+' | '-') expr								# AddSub
	| expr op = ('>' | '<' | '>=' | '<=' | '==' | '!=') expr	# CompareSymbol
	| expr compareOp expr										# CompareFuncInfix
	| logicOp expr												# LogicUnary
	| expr logicOp expr											# LogicBinary
	| ID '(' exprList? ')'										# FunctionCall
	| '(' expr ')'												# Parens
	| NUMBER													# Number
	| ID														# Id;

exprList: expr (',' expr)*;

// 比较运算符(文本形式)
compareOp: GT | LT | GTE | LTE | EQ | NEQ;

// 逻辑运算符(文本形式)
logicOp: AND | OR | NOT;

// 词法规则 - 关键字必须在 ID 之前定义
GT: 'gt';
LT: 'lt';
GTE: 'gte';
LTE: 'lte';
EQ: 'eq';
NEQ: 'neq';
AND: 'and';
OR: 'or';
NOT: 'not';

ID: [a-zA-Z_][a-zA-Z_0-9]*;
NUMBER: [0-9]+ ('.' [0-9]+)?;
NEWLINE: '\r'? '\n';
WS: [ \t]+ -> skip;