# Generated from mlang.g4 by ANTLR 4.13.2
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO

def serializedATN():
    return [
        4,1,24,110,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,2,8,7,8,1,0,4,0,20,8,0,11,0,12,0,21,1,0,1,0,1,1,1,1,1,
        1,1,1,3,1,30,8,1,1,2,1,2,1,2,1,2,3,2,36,8,2,1,2,1,2,1,2,1,2,3,2,
        42,8,2,1,2,1,2,1,2,3,2,47,8,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,
        2,3,2,58,8,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,
        5,2,73,8,2,10,2,12,2,76,9,2,1,3,1,3,1,3,5,3,81,8,3,10,3,12,3,84,
        9,3,1,4,1,4,1,4,5,4,89,8,4,10,4,12,4,92,9,4,1,5,1,5,1,5,5,5,97,8,
        5,10,5,12,5,100,9,5,1,6,1,6,1,6,1,6,1,7,1,7,1,8,1,8,1,8,0,1,4,9,
        0,2,4,6,8,10,12,14,16,0,3,1,0,1,2,1,0,3,4,1,0,5,10,119,0,19,1,0,
        0,0,2,29,1,0,0,0,4,57,1,0,0,0,6,77,1,0,0,0,8,85,1,0,0,0,10,93,1,
        0,0,0,12,101,1,0,0,0,14,105,1,0,0,0,16,107,1,0,0,0,18,20,3,2,1,0,
        19,18,1,0,0,0,20,21,1,0,0,0,21,19,1,0,0,0,21,22,1,0,0,0,22,23,1,
        0,0,0,23,24,5,0,0,1,24,1,1,0,0,0,25,26,3,4,2,0,26,27,5,23,0,0,27,
        30,1,0,0,0,28,30,5,23,0,0,29,25,1,0,0,0,29,28,1,0,0,0,30,3,1,0,0,
        0,31,32,6,2,-1,0,32,33,3,16,8,0,33,35,5,11,0,0,34,36,3,6,3,0,35,
        34,1,0,0,0,35,36,1,0,0,0,36,37,1,0,0,0,37,38,5,12,0,0,38,58,1,0,
        0,0,39,41,5,13,0,0,40,42,3,8,4,0,41,40,1,0,0,0,41,42,1,0,0,0,42,
        43,1,0,0,0,43,58,5,14,0,0,44,46,5,15,0,0,45,47,3,10,5,0,46,45,1,
        0,0,0,46,47,1,0,0,0,47,48,1,0,0,0,48,58,5,16,0,0,49,50,5,11,0,0,
        50,51,3,4,2,0,51,52,5,12,0,0,52,58,1,0,0,0,53,58,5,21,0,0,54,58,
        5,22,0,0,55,58,5,19,0,0,56,58,5,20,0,0,57,31,1,0,0,0,57,39,1,0,0,
        0,57,44,1,0,0,0,57,49,1,0,0,0,57,53,1,0,0,0,57,54,1,0,0,0,57,55,
        1,0,0,0,57,56,1,0,0,0,58,74,1,0,0,0,59,60,10,12,0,0,60,61,7,0,0,
        0,61,73,3,4,2,13,62,63,10,11,0,0,63,64,7,1,0,0,64,73,3,4,2,12,65,
        66,10,10,0,0,66,67,7,2,0,0,67,73,3,4,2,11,68,69,10,9,0,0,69,70,3,
        14,7,0,70,71,3,4,2,10,71,73,1,0,0,0,72,59,1,0,0,0,72,62,1,0,0,0,
        72,65,1,0,0,0,72,68,1,0,0,0,73,76,1,0,0,0,74,72,1,0,0,0,74,75,1,
        0,0,0,75,5,1,0,0,0,76,74,1,0,0,0,77,82,3,4,2,0,78,79,5,17,0,0,79,
        81,3,4,2,0,80,78,1,0,0,0,81,84,1,0,0,0,82,80,1,0,0,0,82,83,1,0,0,
        0,83,7,1,0,0,0,84,82,1,0,0,0,85,90,3,4,2,0,86,87,5,17,0,0,87,89,
        3,4,2,0,88,86,1,0,0,0,89,92,1,0,0,0,90,88,1,0,0,0,90,91,1,0,0,0,
        91,9,1,0,0,0,92,90,1,0,0,0,93,98,3,12,6,0,94,95,5,17,0,0,95,97,3,
        12,6,0,96,94,1,0,0,0,97,100,1,0,0,0,98,96,1,0,0,0,98,99,1,0,0,0,
        99,11,1,0,0,0,100,98,1,0,0,0,101,102,3,4,2,0,102,103,5,18,0,0,103,
        104,3,4,2,0,104,13,1,0,0,0,105,106,5,20,0,0,106,15,1,0,0,0,107,108,
        5,20,0,0,108,17,1,0,0,0,11,21,29,35,41,46,57,72,74,82,90,98
    ]

class mlangParser ( Parser ):

    grammarFileName = "mlang.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'*'", "'/'", "'+'", "'-'", "'>'", "'<'", 
                     "'>='", "'<='", "'=='", "'!='", "'('", "')'", "'['", 
                     "']'", "'{'", "'}'", "','", "':'" ]

    symbolicNames = [ "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "BOOLEAN", 
                      "ID", "NUMBER", "STRING", "NEWLINE", "WS" ]

    RULE_prog = 0
    RULE_stat = 1
    RULE_expr = 2
    RULE_exprList = 3
    RULE_arrayElements = 4
    RULE_dictElements = 5
    RULE_dictPair = 6
    RULE_compareOp = 7
    RULE_func = 8

    ruleNames =  [ "prog", "stat", "expr", "exprList", "arrayElements", 
                   "dictElements", "dictPair", "compareOp", "func" ]

    EOF = Token.EOF
    T__0=1
    T__1=2
    T__2=3
    T__3=4
    T__4=5
    T__5=6
    T__6=7
    T__7=8
    T__8=9
    T__9=10
    T__10=11
    T__11=12
    T__12=13
    T__13=14
    T__14=15
    T__15=16
    T__16=17
    T__17=18
    BOOLEAN=19
    ID=20
    NUMBER=21
    STRING=22
    NEWLINE=23
    WS=24

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.13.2")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class ProgContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def EOF(self):
            return self.getToken(mlangParser.EOF, 0)

        def stat(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.StatContext)
            else:
                return self.getTypedRuleContext(mlangParser.StatContext,i)


        def getRuleIndex(self):
            return mlangParser.RULE_prog

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterProg" ):
                listener.enterProg(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitProg" ):
                listener.exitProg(self)




    def prog(self):

        localctx = mlangParser.ProgContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_prog)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 19 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 18
                self.stat()
                self.state = 21 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not ((((_la) & ~0x3f) == 0 and ((1 << _la) & 16295936) != 0)):
                    break

            self.state = 23
            self.match(mlangParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StatContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expr(self):
            return self.getTypedRuleContext(mlangParser.ExprContext,0)


        def NEWLINE(self):
            return self.getToken(mlangParser.NEWLINE, 0)

        def getRuleIndex(self):
            return mlangParser.RULE_stat

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterStat" ):
                listener.enterStat(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitStat" ):
                listener.exitStat(self)




    def stat(self):

        localctx = mlangParser.StatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_stat)
        try:
            self.state = 29
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [11, 13, 15, 19, 20, 21, 22]:
                self.enterOuterAlt(localctx, 1)
                self.state = 25
                self.expr(0)
                self.state = 26
                self.match(mlangParser.NEWLINE)
                pass
            elif token in [23]:
                self.enterOuterAlt(localctx, 2)
                self.state = 28
                self.match(mlangParser.NEWLINE)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ExprContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return mlangParser.RULE_expr

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)


    class ArrayContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def arrayElements(self):
            return self.getTypedRuleContext(mlangParser.ArrayElementsContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterArray" ):
                listener.enterArray(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitArray" ):
                listener.exitArray(self)


    class DictionaryContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def dictElements(self):
            return self.getTypedRuleContext(mlangParser.DictElementsContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDictionary" ):
                listener.enterDictionary(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDictionary" ):
                listener.exitDictionary(self)


    class NumberContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def NUMBER(self):
            return self.getToken(mlangParser.NUMBER, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNumber" ):
                listener.enterNumber(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNumber" ):
                listener.exitNumber(self)


    class MulDivContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMulDiv" ):
                listener.enterMulDiv(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMulDiv" ):
                listener.exitMulDiv(self)


    class AddSubContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAddSub" ):
                listener.enterAddSub(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAddSub" ):
                listener.exitAddSub(self)


    class ParensContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def expr(self):
            return self.getTypedRuleContext(mlangParser.ExprContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterParens" ):
                listener.enterParens(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitParens" ):
                listener.exitParens(self)


    class CompareSymbolContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCompareSymbol" ):
                listener.enterCompareSymbol(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCompareSymbol" ):
                listener.exitCompareSymbol(self)


    class FunctionCallContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def func(self):
            return self.getTypedRuleContext(mlangParser.FuncContext,0)

        def exprList(self):
            return self.getTypedRuleContext(mlangParser.ExprListContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFunctionCall" ):
                listener.enterFunctionCall(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFunctionCall" ):
                listener.exitFunctionCall(self)


    class StringContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def STRING(self):
            return self.getToken(mlangParser.STRING, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterString" ):
                listener.enterString(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitString" ):
                listener.exitString(self)


    class IdContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def ID(self):
            return self.getToken(mlangParser.ID, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterId" ):
                listener.enterId(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitId" ):
                listener.exitId(self)


    class BooleanContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def BOOLEAN(self):
            return self.getToken(mlangParser.BOOLEAN, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBoolean" ):
                listener.enterBoolean(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBoolean" ):
                listener.exitBoolean(self)


    class CompareFuncInfixContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a mlangParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)

        def compareOp(self):
            return self.getTypedRuleContext(mlangParser.CompareOpContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCompareFuncInfix" ):
                listener.enterCompareFuncInfix(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCompareFuncInfix" ):
                listener.exitCompareFuncInfix(self)



    def expr(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = mlangParser.ExprContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 4
        self.enterRecursionRule(localctx, 4, self.RULE_expr, _p)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 57
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,5,self._ctx)
            if la_ == 1:
                localctx = mlangParser.FunctionCallContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 32
                self.func()
                self.state = 33
                self.match(mlangParser.T__10)
                self.state = 35
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if (((_la) & ~0x3f) == 0 and ((1 << _la) & 7907328) != 0):
                    self.state = 34
                    self.exprList()


                self.state = 37
                self.match(mlangParser.T__11)
                pass

            elif la_ == 2:
                localctx = mlangParser.ArrayContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 39
                self.match(mlangParser.T__12)
                self.state = 41
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if (((_la) & ~0x3f) == 0 and ((1 << _la) & 7907328) != 0):
                    self.state = 40
                    self.arrayElements()


                self.state = 43
                self.match(mlangParser.T__13)
                pass

            elif la_ == 3:
                localctx = mlangParser.DictionaryContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 44
                self.match(mlangParser.T__14)
                self.state = 46
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if (((_la) & ~0x3f) == 0 and ((1 << _la) & 7907328) != 0):
                    self.state = 45
                    self.dictElements()


                self.state = 48
                self.match(mlangParser.T__15)
                pass

            elif la_ == 4:
                localctx = mlangParser.ParensContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 49
                self.match(mlangParser.T__10)
                self.state = 50
                self.expr(0)
                self.state = 51
                self.match(mlangParser.T__11)
                pass

            elif la_ == 5:
                localctx = mlangParser.NumberContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 53
                self.match(mlangParser.NUMBER)
                pass

            elif la_ == 6:
                localctx = mlangParser.StringContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 54
                self.match(mlangParser.STRING)
                pass

            elif la_ == 7:
                localctx = mlangParser.BooleanContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 55
                self.match(mlangParser.BOOLEAN)
                pass

            elif la_ == 8:
                localctx = mlangParser.IdContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 56
                self.match(mlangParser.ID)
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 74
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,7,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    self.state = 72
                    self._errHandler.sync(self)
                    la_ = self._interp.adaptivePredict(self._input,6,self._ctx)
                    if la_ == 1:
                        localctx = mlangParser.MulDivContext(self, mlangParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 59
                        if not self.precpred(self._ctx, 12):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 12)")
                        self.state = 60
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not(_la==1 or _la==2):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 61
                        self.expr(13)
                        pass

                    elif la_ == 2:
                        localctx = mlangParser.AddSubContext(self, mlangParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 62
                        if not self.precpred(self._ctx, 11):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 11)")
                        self.state = 63
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not(_la==3 or _la==4):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 64
                        self.expr(12)
                        pass

                    elif la_ == 3:
                        localctx = mlangParser.CompareSymbolContext(self, mlangParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 65
                        if not self.precpred(self._ctx, 10):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 10)")
                        self.state = 66
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 2016) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 67
                        self.expr(11)
                        pass

                    elif la_ == 4:
                        localctx = mlangParser.CompareFuncInfixContext(self, mlangParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 68
                        if not self.precpred(self._ctx, 9):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 9)")
                        self.state = 69
                        self.compareOp()
                        self.state = 70
                        self.expr(10)
                        pass

             
                self.state = 76
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,7,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class ExprListContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)


        def getRuleIndex(self):
            return mlangParser.RULE_exprList

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExprList" ):
                listener.enterExprList(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExprList" ):
                listener.exitExprList(self)




    def exprList(self):

        localctx = mlangParser.ExprListContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_exprList)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 77
            self.expr(0)
            self.state = 82
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==17:
                self.state = 78
                self.match(mlangParser.T__16)
                self.state = 79
                self.expr(0)
                self.state = 84
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ArrayElementsContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)


        def getRuleIndex(self):
            return mlangParser.RULE_arrayElements

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterArrayElements" ):
                listener.enterArrayElements(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitArrayElements" ):
                listener.exitArrayElements(self)




    def arrayElements(self):

        localctx = mlangParser.ArrayElementsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_arrayElements)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 85
            self.expr(0)
            self.state = 90
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==17:
                self.state = 86
                self.match(mlangParser.T__16)
                self.state = 87
                self.expr(0)
                self.state = 92
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class DictElementsContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def dictPair(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.DictPairContext)
            else:
                return self.getTypedRuleContext(mlangParser.DictPairContext,i)


        def getRuleIndex(self):
            return mlangParser.RULE_dictElements

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDictElements" ):
                listener.enterDictElements(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDictElements" ):
                listener.exitDictElements(self)




    def dictElements(self):

        localctx = mlangParser.DictElementsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_dictElements)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 93
            self.dictPair()
            self.state = 98
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==17:
                self.state = 94
                self.match(mlangParser.T__16)
                self.state = 95
                self.dictPair()
                self.state = 100
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class DictPairContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(mlangParser.ExprContext)
            else:
                return self.getTypedRuleContext(mlangParser.ExprContext,i)


        def getRuleIndex(self):
            return mlangParser.RULE_dictPair

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDictPair" ):
                listener.enterDictPair(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDictPair" ):
                listener.exitDictPair(self)




    def dictPair(self):

        localctx = mlangParser.DictPairContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_dictPair)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 101
            self.expr(0)
            self.state = 102
            self.match(mlangParser.T__17)
            self.state = 103
            self.expr(0)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class CompareOpContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(mlangParser.ID, 0)

        def getRuleIndex(self):
            return mlangParser.RULE_compareOp

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCompareOp" ):
                listener.enterCompareOp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCompareOp" ):
                listener.exitCompareOp(self)




    def compareOp(self):

        localctx = mlangParser.CompareOpContext(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_compareOp)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 105
            self.match(mlangParser.ID)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class FuncContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(mlangParser.ID, 0)

        def getRuleIndex(self):
            return mlangParser.RULE_func

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFunc" ):
                listener.enterFunc(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFunc" ):
                listener.exitFunc(self)




    def func(self):

        localctx = mlangParser.FuncContext(self, self._ctx, self.state)
        self.enterRule(localctx, 16, self.RULE_func)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 107
            self.match(mlangParser.ID)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx



    def sempred(self, localctx:RuleContext, ruleIndex:int, predIndex:int):
        if self._predicates == None:
            self._predicates = dict()
        self._predicates[2] = self.expr_sempred
        pred = self._predicates.get(ruleIndex, None)
        if pred is None:
            raise Exception("No predicate with index:" + str(ruleIndex))
        else:
            return pred(localctx, predIndex)

    def expr_sempred(self, localctx:ExprContext, predIndex:int):
            if predIndex == 0:
                return self.precpred(self._ctx, 12)
         

            if predIndex == 1:
                return self.precpred(self._ctx, 11)
         

            if predIndex == 2:
                return self.precpred(self._ctx, 10)
         

            if predIndex == 3:
                return self.precpred(self._ctx, 9)
         




