// Code generated from mlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mlang

import "github.com/antlr4-go/antlr/v4"

// mlangListener is a complete listener for a parse tree produced by mlangParser.
type mlangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterArray is called when entering the Array production.
	EnterArray(c *ArrayContext)

	// EnterDictionary is called when entering the Dictionary production.
	EnterDictionary(c *DictionaryContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterCompareSymbol is called when entering the CompareSymbol production.
	EnterCompareSymbol(c *CompareSymbolContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterCompareFuncInfix is called when entering the CompareFuncInfix production.
	EnterCompareFuncInfix(c *CompareFuncInfixContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterArrayElements is called when entering the arrayElements production.
	EnterArrayElements(c *ArrayElementsContext)

	// EnterDictElements is called when entering the dictElements production.
	EnterDictElements(c *DictElementsContext)

	// EnterDictPair is called when entering the dictPair production.
	EnterDictPair(c *DictPairContext)

	// EnterCompareOp is called when entering the compareOp production.
	EnterCompareOp(c *CompareOpContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitArray is called when exiting the Array production.
	ExitArray(c *ArrayContext)

	// ExitDictionary is called when exiting the Dictionary production.
	ExitDictionary(c *DictionaryContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitCompareSymbol is called when exiting the CompareSymbol production.
	ExitCompareSymbol(c *CompareSymbolContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitCompareFuncInfix is called when exiting the CompareFuncInfix production.
	ExitCompareFuncInfix(c *CompareFuncInfixContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitArrayElements is called when exiting the arrayElements production.
	ExitArrayElements(c *ArrayElementsContext)

	// ExitDictElements is called when exiting the dictElements production.
	ExitDictElements(c *DictElementsContext)

	// ExitDictPair is called when exiting the dictPair production.
	ExitDictPair(c *DictPairContext)

	// ExitCompareOp is called when exiting the compareOp production.
	ExitCompareOp(c *CompareOpContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)
}
