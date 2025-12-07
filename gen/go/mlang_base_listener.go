// Code generated from mlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mlang

import "github.com/antlr4-go/antlr/v4"

// BasemlangListener is a complete listener for a parse tree produced by mlangParser.
type BasemlangListener struct{}

var _ mlangListener = &BasemlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasemlangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasemlangListener) ExitProg(ctx *ProgContext) {}

// EnterStat is called when production stat is entered.
func (s *BasemlangListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BasemlangListener) ExitStat(ctx *StatContext) {}

// EnterDictionary is called when production Dictionary is entered.
func (s *BasemlangListener) EnterDictionary(ctx *DictionaryContext) {}

// ExitDictionary is called when production Dictionary is exited.
func (s *BasemlangListener) ExitDictionary(ctx *DictionaryContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BasemlangListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BasemlangListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasemlangListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasemlangListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BasemlangListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BasemlangListener) ExitParens(ctx *ParensContext) {}

// EnterCompareSymbol is called when production CompareSymbol is entered.
func (s *BasemlangListener) EnterCompareSymbol(ctx *CompareSymbolContext) {}

// ExitCompareSymbol is called when production CompareSymbol is exited.
func (s *BasemlangListener) ExitCompareSymbol(ctx *CompareSymbolContext) {}

// EnterString is called when production String is entered.
func (s *BasemlangListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BasemlangListener) ExitString(ctx *StringContext) {}

// EnterCompareFuncInfix is called when production CompareFuncInfix is entered.
func (s *BasemlangListener) EnterCompareFuncInfix(ctx *CompareFuncInfixContext) {}

// ExitCompareFuncInfix is called when production CompareFuncInfix is exited.
func (s *BasemlangListener) ExitCompareFuncInfix(ctx *CompareFuncInfixContext) {}

// EnterArray is called when production Array is entered.
func (s *BasemlangListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production Array is exited.
func (s *BasemlangListener) ExitArray(ctx *ArrayContext) {}

// EnterNumber is called when production Number is entered.
func (s *BasemlangListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BasemlangListener) ExitNumber(ctx *NumberContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BasemlangListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BasemlangListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterId is called when production Id is entered.
func (s *BasemlangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production Id is exited.
func (s *BasemlangListener) ExitId(ctx *IdContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BasemlangListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BasemlangListener) ExitBoolean(ctx *BooleanContext) {}

// EnterFieldAccess is called when production FieldAccess is entered.
func (s *BasemlangListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production FieldAccess is exited.
func (s *BasemlangListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BasemlangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BasemlangListener) ExitExprList(ctx *ExprListContext) {}

// EnterArrayElements is called when production arrayElements is entered.
func (s *BasemlangListener) EnterArrayElements(ctx *ArrayElementsContext) {}

// ExitArrayElements is called when production arrayElements is exited.
func (s *BasemlangListener) ExitArrayElements(ctx *ArrayElementsContext) {}

// EnterDictElements is called when production dictElements is entered.
func (s *BasemlangListener) EnterDictElements(ctx *DictElementsContext) {}

// ExitDictElements is called when production dictElements is exited.
func (s *BasemlangListener) ExitDictElements(ctx *DictElementsContext) {}

// EnterDictPair is called when production dictPair is entered.
func (s *BasemlangListener) EnterDictPair(ctx *DictPairContext) {}

// ExitDictPair is called when production dictPair is exited.
func (s *BasemlangListener) ExitDictPair(ctx *DictPairContext) {}

// EnterBinaryOp is called when production binaryOp is entered.
func (s *BasemlangListener) EnterBinaryOp(ctx *BinaryOpContext) {}

// ExitBinaryOp is called when production binaryOp is exited.
func (s *BasemlangListener) ExitBinaryOp(ctx *BinaryOpContext) {}

// EnterFunc is called when production func is entered.
func (s *BasemlangListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BasemlangListener) ExitFunc(ctx *FuncContext) {}
