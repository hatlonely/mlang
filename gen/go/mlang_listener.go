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

	// EnterIndexLvalue is called when entering the IndexLvalue production.
	EnterIndexLvalue(c *IndexLvalueContext)

	// EnterSimpleLvalue is called when entering the SimpleLvalue production.
	EnterSimpleLvalue(c *SimpleLvalueContext)

	// EnterFieldLvalue is called when entering the FieldLvalue production.
	EnterFieldLvalue(c *FieldLvalueContext)

	// EnterDictionary is called when entering the Dictionary production.
	EnterDictionary(c *DictionaryContext)

	// EnterOrOp is called when entering the OrOp production.
	EnterOrOp(c *OrOpContext)

	// EnterRelationalSymbol is called when entering the RelationalSymbol production.
	EnterRelationalSymbol(c *RelationalSymbolContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterIndexAccess is called when entering the IndexAccess production.
	EnterIndexAccess(c *IndexAccessContext)

	// EnterEqualitySymbol is called when entering the EqualitySymbol production.
	EnterEqualitySymbol(c *EqualitySymbolContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterCompareFuncInfix is called when entering the CompareFuncInfix production.
	EnterCompareFuncInfix(c *CompareFuncInfixContext)

	// EnterArray is called when entering the Array production.
	EnterArray(c *ArrayContext)

	// EnterNotCompareFuncInfix is called when entering the NotCompareFuncInfix production.
	EnterNotCompareFuncInfix(c *NotCompareFuncInfixContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterFieldAccess is called when entering the FieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterAndOp is called when entering the AndOp production.
	EnterAndOp(c *AndOpContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterArrayElements is called when entering the arrayElements production.
	EnterArrayElements(c *ArrayElementsContext)

	// EnterDictElements is called when entering the dictElements production.
	EnterDictElements(c *DictElementsContext)

	// EnterDictPair is called when entering the dictPair production.
	EnterDictPair(c *DictPairContext)

	// EnterBinaryOp is called when entering the binaryOp production.
	EnterBinaryOp(c *BinaryOpContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitIndexLvalue is called when exiting the IndexLvalue production.
	ExitIndexLvalue(c *IndexLvalueContext)

	// ExitSimpleLvalue is called when exiting the SimpleLvalue production.
	ExitSimpleLvalue(c *SimpleLvalueContext)

	// ExitFieldLvalue is called when exiting the FieldLvalue production.
	ExitFieldLvalue(c *FieldLvalueContext)

	// ExitDictionary is called when exiting the Dictionary production.
	ExitDictionary(c *DictionaryContext)

	// ExitOrOp is called when exiting the OrOp production.
	ExitOrOp(c *OrOpContext)

	// ExitRelationalSymbol is called when exiting the RelationalSymbol production.
	ExitRelationalSymbol(c *RelationalSymbolContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitIndexAccess is called when exiting the IndexAccess production.
	ExitIndexAccess(c *IndexAccessContext)

	// ExitEqualitySymbol is called when exiting the EqualitySymbol production.
	ExitEqualitySymbol(c *EqualitySymbolContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitCompareFuncInfix is called when exiting the CompareFuncInfix production.
	ExitCompareFuncInfix(c *CompareFuncInfixContext)

	// ExitArray is called when exiting the Array production.
	ExitArray(c *ArrayContext)

	// ExitNotCompareFuncInfix is called when exiting the NotCompareFuncInfix production.
	ExitNotCompareFuncInfix(c *NotCompareFuncInfixContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitFieldAccess is called when exiting the FieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitAndOp is called when exiting the AndOp production.
	ExitAndOp(c *AndOpContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitArrayElements is called when exiting the arrayElements production.
	ExitArrayElements(c *ArrayElementsContext)

	// ExitDictElements is called when exiting the dictElements production.
	ExitDictElements(c *DictElementsContext)

	// ExitDictPair is called when exiting the dictPair production.
	ExitDictPair(c *DictPairContext)

	// ExitBinaryOp is called when exiting the binaryOp production.
	ExitBinaryOp(c *BinaryOpContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)
}
