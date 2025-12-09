// Code generated from mlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mlang

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by mlangParser.
type mlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mlangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by mlangParser#Dictionary.
	VisitDictionary(ctx *DictionaryContext) interface{}

	// Visit a parse tree produced by mlangParser#OrOp.
	VisitOrOp(ctx *OrOpContext) interface{}

	// Visit a parse tree produced by mlangParser#RelationalSymbol.
	VisitRelationalSymbol(ctx *RelationalSymbolContext) interface{}

	// Visit a parse tree produced by mlangParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by mlangParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by mlangParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by mlangParser#IndexAccess.
	VisitIndexAccess(ctx *IndexAccessContext) interface{}

	// Visit a parse tree produced by mlangParser#EqualitySymbol.
	VisitEqualitySymbol(ctx *EqualitySymbolContext) interface{}

	// Visit a parse tree produced by mlangParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by mlangParser#CompareFuncInfix.
	VisitCompareFuncInfix(ctx *CompareFuncInfixContext) interface{}

	// Visit a parse tree produced by mlangParser#Array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by mlangParser#NotCompareFuncInfix.
	VisitNotCompareFuncInfix(ctx *NotCompareFuncInfixContext) interface{}

	// Visit a parse tree produced by mlangParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by mlangParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by mlangParser#Id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by mlangParser#Boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by mlangParser#FieldAccess.
	VisitFieldAccess(ctx *FieldAccessContext) interface{}

	// Visit a parse tree produced by mlangParser#AndOp.
	VisitAndOp(ctx *AndOpContext) interface{}

	// Visit a parse tree produced by mlangParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by mlangParser#arrayElements.
	VisitArrayElements(ctx *ArrayElementsContext) interface{}

	// Visit a parse tree produced by mlangParser#dictElements.
	VisitDictElements(ctx *DictElementsContext) interface{}

	// Visit a parse tree produced by mlangParser#dictPair.
	VisitDictPair(ctx *DictPairContext) interface{}

	// Visit a parse tree produced by mlangParser#binaryOp.
	VisitBinaryOp(ctx *BinaryOpContext) interface{}

	// Visit a parse tree produced by mlangParser#func.
	VisitFunc(ctx *FuncContext) interface{}
}
