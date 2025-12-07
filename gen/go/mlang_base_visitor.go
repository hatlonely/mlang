// Code generated from mlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mlang

import "github.com/antlr4-go/antlr/v4"

type BasemlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemlangVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitDictionary(ctx *DictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitCompareSymbol(ctx *CompareSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitCompareFuncInfix(ctx *CompareFuncInfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitFieldAccess(ctx *FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitArrayElements(ctx *ArrayElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitDictElements(ctx *DictElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitDictPair(ctx *DictPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitBinaryOp(ctx *BinaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemlangVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}
