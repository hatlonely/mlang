# Generated from mlang.g4 by ANTLR 4.13.2
from antlr4 import *
if "." in __name__:
    from .mlangParser import mlangParser
else:
    from mlangParser import mlangParser

# This class defines a complete listener for a parse tree produced by mlangParser.
class mlangListener(ParseTreeListener):

    # Enter a parse tree produced by mlangParser#prog.
    def enterProg(self, ctx:mlangParser.ProgContext):
        pass

    # Exit a parse tree produced by mlangParser#prog.
    def exitProg(self, ctx:mlangParser.ProgContext):
        pass


    # Enter a parse tree produced by mlangParser#stat.
    def enterStat(self, ctx:mlangParser.StatContext):
        pass

    # Exit a parse tree produced by mlangParser#stat.
    def exitStat(self, ctx:mlangParser.StatContext):
        pass


    # Enter a parse tree produced by mlangParser#Array.
    def enterArray(self, ctx:mlangParser.ArrayContext):
        pass

    # Exit a parse tree produced by mlangParser#Array.
    def exitArray(self, ctx:mlangParser.ArrayContext):
        pass


    # Enter a parse tree produced by mlangParser#Dictionary.
    def enterDictionary(self, ctx:mlangParser.DictionaryContext):
        pass

    # Exit a parse tree produced by mlangParser#Dictionary.
    def exitDictionary(self, ctx:mlangParser.DictionaryContext):
        pass


    # Enter a parse tree produced by mlangParser#Number.
    def enterNumber(self, ctx:mlangParser.NumberContext):
        pass

    # Exit a parse tree produced by mlangParser#Number.
    def exitNumber(self, ctx:mlangParser.NumberContext):
        pass


    # Enter a parse tree produced by mlangParser#MulDiv.
    def enterMulDiv(self, ctx:mlangParser.MulDivContext):
        pass

    # Exit a parse tree produced by mlangParser#MulDiv.
    def exitMulDiv(self, ctx:mlangParser.MulDivContext):
        pass


    # Enter a parse tree produced by mlangParser#AddSub.
    def enterAddSub(self, ctx:mlangParser.AddSubContext):
        pass

    # Exit a parse tree produced by mlangParser#AddSub.
    def exitAddSub(self, ctx:mlangParser.AddSubContext):
        pass


    # Enter a parse tree produced by mlangParser#Parens.
    def enterParens(self, ctx:mlangParser.ParensContext):
        pass

    # Exit a parse tree produced by mlangParser#Parens.
    def exitParens(self, ctx:mlangParser.ParensContext):
        pass


    # Enter a parse tree produced by mlangParser#CompareSymbol.
    def enterCompareSymbol(self, ctx:mlangParser.CompareSymbolContext):
        pass

    # Exit a parse tree produced by mlangParser#CompareSymbol.
    def exitCompareSymbol(self, ctx:mlangParser.CompareSymbolContext):
        pass


    # Enter a parse tree produced by mlangParser#FunctionCall.
    def enterFunctionCall(self, ctx:mlangParser.FunctionCallContext):
        pass

    # Exit a parse tree produced by mlangParser#FunctionCall.
    def exitFunctionCall(self, ctx:mlangParser.FunctionCallContext):
        pass


    # Enter a parse tree produced by mlangParser#String.
    def enterString(self, ctx:mlangParser.StringContext):
        pass

    # Exit a parse tree produced by mlangParser#String.
    def exitString(self, ctx:mlangParser.StringContext):
        pass


    # Enter a parse tree produced by mlangParser#Id.
    def enterId(self, ctx:mlangParser.IdContext):
        pass

    # Exit a parse tree produced by mlangParser#Id.
    def exitId(self, ctx:mlangParser.IdContext):
        pass


    # Enter a parse tree produced by mlangParser#Boolean.
    def enterBoolean(self, ctx:mlangParser.BooleanContext):
        pass

    # Exit a parse tree produced by mlangParser#Boolean.
    def exitBoolean(self, ctx:mlangParser.BooleanContext):
        pass


    # Enter a parse tree produced by mlangParser#CompareFuncInfix.
    def enterCompareFuncInfix(self, ctx:mlangParser.CompareFuncInfixContext):
        pass

    # Exit a parse tree produced by mlangParser#CompareFuncInfix.
    def exitCompareFuncInfix(self, ctx:mlangParser.CompareFuncInfixContext):
        pass


    # Enter a parse tree produced by mlangParser#exprList.
    def enterExprList(self, ctx:mlangParser.ExprListContext):
        pass

    # Exit a parse tree produced by mlangParser#exprList.
    def exitExprList(self, ctx:mlangParser.ExprListContext):
        pass


    # Enter a parse tree produced by mlangParser#arrayElements.
    def enterArrayElements(self, ctx:mlangParser.ArrayElementsContext):
        pass

    # Exit a parse tree produced by mlangParser#arrayElements.
    def exitArrayElements(self, ctx:mlangParser.ArrayElementsContext):
        pass


    # Enter a parse tree produced by mlangParser#dictElements.
    def enterDictElements(self, ctx:mlangParser.DictElementsContext):
        pass

    # Exit a parse tree produced by mlangParser#dictElements.
    def exitDictElements(self, ctx:mlangParser.DictElementsContext):
        pass


    # Enter a parse tree produced by mlangParser#dictPair.
    def enterDictPair(self, ctx:mlangParser.DictPairContext):
        pass

    # Exit a parse tree produced by mlangParser#dictPair.
    def exitDictPair(self, ctx:mlangParser.DictPairContext):
        pass


    # Enter a parse tree produced by mlangParser#compareOp.
    def enterCompareOp(self, ctx:mlangParser.CompareOpContext):
        pass

    # Exit a parse tree produced by mlangParser#compareOp.
    def exitCompareOp(self, ctx:mlangParser.CompareOpContext):
        pass


    # Enter a parse tree produced by mlangParser#func.
    def enterFunc(self, ctx:mlangParser.FuncContext):
        pass

    # Exit a parse tree produced by mlangParser#func.
    def exitFunc(self, ctx:mlangParser.FuncContext):
        pass



del mlangParser