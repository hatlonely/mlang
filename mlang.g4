grammar mlang;

// 解析规则
prog: stat+ EOF;

stat: expr NEWLINE? | NEWLINE;

expr:
	expr op = ('*' | '/') expr									# MulDiv
	| expr op = ('+' | '-') expr								# AddSub
	| expr op = ('>' | '<' | '>=' | '<=' | '==' | '!=') expr	# CompareSymbol
	| expr binaryOp expr										# CompareFuncInfix
	| expr DOT ID												# FieldAccess
	| func '(' exprList? ')'									# FunctionCall
	| '[' arrayElements? ']'									# Array
	| '{' dictElements? '}'										# Dictionary
	| '(' expr ')'												# Parens
	| NUMBER													# Number
	| STRING													# String
	| BOOLEAN													# Boolean
	| ID														# Id;

exprList: expr (',' expr)*;

arrayElements: expr (',' expr)*;

dictElements: dictPair (',' dictPair)*;

dictPair: expr ':' expr;

binaryOp: ID;
func: ID;

DOT: '.';
BOOLEAN: 'TRUE' | 'FALSE' | 'true' | 'false';
ID: [a-zA-Z_][a-zA-Z_0-9]*;
NUMBER: [0-9]+ ('.' [0-9]+)?;
STRING:
	'"' (~["\\] | '\\' .)* '"' // 双引号字符串，支持转义
	| '`' (~[`])* '`'; // 反引号字符串，不支持转义
NEWLINE: '\r'? '\n';
WS: [ \t]+ -> skip;