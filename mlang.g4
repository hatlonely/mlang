grammar Expr;

// 解析规则
prog: stat+ EOF;

stat: expr NEWLINE | NEWLINE;

expr:
	expr op = ('*' | '/') expr									# MulDiv
	| expr op = ('+' | '-') expr								# AddSub
	| expr op = ('>' | '<' | '>=' | '<=' | '==' | '!=') expr	# CompareSymbol
	| expr compareOp expr										# CompareFuncInfix
	| func '(' exprList? ')'									# FunctionCall
	| '(' expr ')'												# Parens
	| NUMBER													# Number
	| ID														# Id;

exprList: expr (',' expr)*;

compareOp: ID;
func: ID;

ID: [a-zA-Z_][a-zA-Z_0-9]*;
NUMBER: [0-9]+ ('.' [0-9]+)?;
NEWLINE: '\r'? '\n';
WS: [ \t]+ -> skip;