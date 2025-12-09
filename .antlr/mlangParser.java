// Generated from /Users/mvbj0483/hatlonely/github.com/hatlonely/mlang/mlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class mlangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, DOT=24, BOOLEAN=25, 
		ID=26, NUMBER=27, STRING=28, NEWLINE=29, WS=30;
	public static final int
		RULE_prog = 0, RULE_stat = 1, RULE_lvalue = 2, RULE_expr = 3, RULE_exprList = 4, 
		RULE_arrayElements = 5, RULE_dictElements = 6, RULE_dictPair = 7, RULE_binaryOp = 8, 
		RULE_func = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "stat", "lvalue", "expr", "exprList", "arrayElements", "dictElements", 
			"dictPair", "binaryOp", "func"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'['", "']'", "'('", "')'", "'*'", "'/'", "'+'", "'-'", 
			"'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'NOT'", "'not'", "'&&'", 
			"'||'", "'{'", "'}'", "','", "':'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"DOT", "BOOLEAN", "ID", "NUMBER", "STRING", "NEWLINE", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "mlang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public mlangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgContext extends ParserRuleContext {
		public StatContext stat() {
			return getRuleContext(StatContext.class,0);
		}
		public TerminalNode EOF() { return getToken(mlangParser.EOF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterProg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitProg(this);
		}
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		try {
			setState(26);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(20);
				stat();
				setState(21);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(23);
				expr(0);
				setState(24);
				match(EOF);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatContext extends ParserRuleContext {
		public LvalueContext lvalue() {
			return getRuleContext(LvalueContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stat; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterStat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitStat(this);
		}
	}

	public final StatContext stat() throws RecognitionException {
		StatContext _localctx = new StatContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(28);
			lvalue(0);
			setState(29);
			match(T__0);
			setState(30);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LvalueContext extends ParserRuleContext {
		public LvalueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lvalue; }
	 
		public LvalueContext() { }
		public void copyFrom(LvalueContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexLvalueContext extends LvalueContext {
		public LvalueContext lvalue() {
			return getRuleContext(LvalueContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IndexLvalueContext(LvalueContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterIndexLvalue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitIndexLvalue(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SimpleLvalueContext extends LvalueContext {
		public TerminalNode ID() { return getToken(mlangParser.ID, 0); }
		public SimpleLvalueContext(LvalueContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterSimpleLvalue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitSimpleLvalue(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FieldLvalueContext extends LvalueContext {
		public LvalueContext lvalue() {
			return getRuleContext(LvalueContext.class,0);
		}
		public TerminalNode DOT() { return getToken(mlangParser.DOT, 0); }
		public TerminalNode ID() { return getToken(mlangParser.ID, 0); }
		public FieldLvalueContext(LvalueContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterFieldLvalue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitFieldLvalue(this);
		}
	}

	public final LvalueContext lvalue() throws RecognitionException {
		return lvalue(0);
	}

	private LvalueContext lvalue(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		LvalueContext _localctx = new LvalueContext(_ctx, _parentState);
		LvalueContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_lvalue, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new SimpleLvalueContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(33);
			match(ID);
			}
			_ctx.stop = _input.LT(-1);
			setState(45);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(43);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new IndexLvalueContext(new LvalueContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_lvalue);
						setState(35);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(36);
						match(T__1);
						setState(37);
						expr(0);
						setState(38);
						match(T__2);
						}
						break;
					case 2:
						{
						_localctx = new FieldLvalueContext(new LvalueContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_lvalue);
						setState(40);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(41);
						match(DOT);
						setState(42);
						match(ID);
						}
						break;
					}
					} 
				}
				setState(47);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DictionaryContext extends ExprContext {
		public DictElementsContext dictElements() {
			return getRuleContext(DictElementsContext.class,0);
		}
		public DictionaryContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterDictionary(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitDictionary(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OrOpContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public OrOpContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterOrOp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitOrOp(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RelationalSymbolContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public RelationalSymbolContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterRelationalSymbol(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitRelationalSymbol(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MulDivContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MulDivContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterMulDiv(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitMulDiv(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AddSubContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterAddSub(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitAddSub(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParensContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ParensContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterParens(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitParens(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexAccessContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public IndexAccessContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterIndexAccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitIndexAccess(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EqualitySymbolContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public EqualitySymbolContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterEqualitySymbol(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitEqualitySymbol(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ExprContext {
		public TerminalNode STRING() { return getToken(mlangParser.STRING, 0); }
		public StringContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitString(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CompareFuncInfixContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public BinaryOpContext binaryOp() {
			return getRuleContext(BinaryOpContext.class,0);
		}
		public CompareFuncInfixContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterCompareFuncInfix(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitCompareFuncInfix(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayContext extends ExprContext {
		public ArrayElementsContext arrayElements() {
			return getRuleContext(ArrayElementsContext.class,0);
		}
		public ArrayContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterArray(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitArray(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NotCompareFuncInfixContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public BinaryOpContext binaryOp() {
			return getRuleContext(BinaryOpContext.class,0);
		}
		public NotCompareFuncInfixContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterNotCompareFuncInfix(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitNotCompareFuncInfix(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NumberContext extends ExprContext {
		public TerminalNode NUMBER() { return getToken(mlangParser.NUMBER, 0); }
		public NumberContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterNumber(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitNumber(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallContext extends ExprContext {
		public FuncContext func() {
			return getRuleContext(FuncContext.class,0);
		}
		public ExprListContext exprList() {
			return getRuleContext(ExprListContext.class,0);
		}
		public FunctionCallContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterFunctionCall(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitFunctionCall(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ExprContext {
		public TerminalNode ID() { return getToken(mlangParser.ID, 0); }
		public IdContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterId(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitId(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanContext extends ExprContext {
		public TerminalNode BOOLEAN() { return getToken(mlangParser.BOOLEAN, 0); }
		public BooleanContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterBoolean(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitBoolean(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FieldAccessContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOT() { return getToken(mlangParser.DOT, 0); }
		public TerminalNode ID() { return getToken(mlangParser.ID, 0); }
		public FieldAccessContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterFieldAccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitFieldAccess(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AndOpContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AndOpContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterAndOp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitAndOp(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				_localctx = new FunctionCallContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(49);
				func();
				setState(50);
				match(T__3);
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 504365076L) != 0)) {
					{
					setState(51);
					exprList();
					}
				}

				setState(54);
				match(T__4);
				}
				break;
			case 2:
				{
				_localctx = new ArrayContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(56);
				match(T__1);
				setState(58);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 504365076L) != 0)) {
					{
					setState(57);
					arrayElements();
					}
				}

				setState(60);
				match(T__2);
				}
				break;
			case 3:
				{
				_localctx = new DictionaryContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(61);
				match(T__19);
				setState(63);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 504365076L) != 0)) {
					{
					setState(62);
					dictElements();
					}
				}

				setState(65);
				match(T__20);
				}
				break;
			case 4:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(66);
				match(T__3);
				setState(67);
				expr(0);
				setState(68);
				match(T__4);
				}
				break;
			case 5:
				{
				_localctx = new NumberContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(70);
				match(NUMBER);
				}
				break;
			case 6:
				{
				_localctx = new StringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(71);
				match(STRING);
				}
				break;
			case 7:
				{
				_localctx = new BooleanContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(72);
				match(BOOLEAN);
				}
				break;
			case 8:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(73);
				match(ID);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(113);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(111);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(76);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(77);
						((MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__5 || _la==T__6) ) {
							((MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(78);
						expr(16);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(79);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(80);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__7 || _la==T__8) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(81);
						expr(15);
						}
						break;
					case 3:
						{
						_localctx = new RelationalSymbolContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(82);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(83);
						((RelationalSymbolContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 15360L) != 0)) ) {
							((RelationalSymbolContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(84);
						expr(14);
						}
						break;
					case 4:
						{
						_localctx = new EqualitySymbolContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(85);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(86);
						((EqualitySymbolContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__13 || _la==T__14) ) {
							((EqualitySymbolContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(87);
						expr(13);
						}
						break;
					case 5:
						{
						_localctx = new NotCompareFuncInfixContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(88);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(89);
						_la = _input.LA(1);
						if ( !(_la==T__15 || _la==T__16) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(90);
						binaryOp();
						setState(91);
						expr(12);
						}
						break;
					case 6:
						{
						_localctx = new CompareFuncInfixContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(93);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(94);
						binaryOp();
						setState(95);
						expr(11);
						}
						break;
					case 7:
						{
						_localctx = new AndOpContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(97);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(98);
						((AndOpContext)_localctx).op = match(T__17);
						setState(99);
						expr(10);
						}
						break;
					case 8:
						{
						_localctx = new OrOpContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(100);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(101);
						((OrOpContext)_localctx).op = match(T__18);
						setState(102);
						expr(9);
						}
						break;
					case 9:
						{
						_localctx = new IndexAccessContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(103);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(104);
						match(T__1);
						setState(105);
						expr(0);
						setState(106);
						match(T__2);
						}
						break;
					case 10:
						{
						_localctx = new FieldAccessContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(108);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(109);
						match(DOT);
						setState(110);
						match(ID);
						}
						break;
					}
					} 
				}
				setState(115);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprListContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ExprListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterExprList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitExprList(this);
		}
	}

	public final ExprListContext exprList() throws RecognitionException {
		ExprListContext _localctx = new ExprListContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_exprList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			expr(0);
			setState(121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__21) {
				{
				{
				setState(117);
				match(T__21);
				setState(118);
				expr(0);
				}
				}
				setState(123);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayElementsContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ArrayElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayElements; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterArrayElements(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitArrayElements(this);
		}
	}

	public final ArrayElementsContext arrayElements() throws RecognitionException {
		ArrayElementsContext _localctx = new ArrayElementsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_arrayElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			expr(0);
			setState(129);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__21) {
				{
				{
				setState(125);
				match(T__21);
				setState(126);
				expr(0);
				}
				}
				setState(131);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DictElementsContext extends ParserRuleContext {
		public List<DictPairContext> dictPair() {
			return getRuleContexts(DictPairContext.class);
		}
		public DictPairContext dictPair(int i) {
			return getRuleContext(DictPairContext.class,i);
		}
		public DictElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictElements; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterDictElements(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitDictElements(this);
		}
	}

	public final DictElementsContext dictElements() throws RecognitionException {
		DictElementsContext _localctx = new DictElementsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_dictElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			dictPair();
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__21) {
				{
				{
				setState(133);
				match(T__21);
				setState(134);
				dictPair();
				}
				}
				setState(139);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DictPairContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public DictPairContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dictPair; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterDictPair(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitDictPair(this);
		}
	}

	public final DictPairContext dictPair() throws RecognitionException {
		DictPairContext _localctx = new DictPairContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_dictPair);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			expr(0);
			setState(141);
			match(T__22);
			setState(142);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryOpContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(mlangParser.ID, 0); }
		public BinaryOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryOp; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterBinaryOp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitBinaryOp(this);
		}
	}

	public final BinaryOpContext binaryOp() throws RecognitionException {
		BinaryOpContext _localctx = new BinaryOpContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_binaryOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(mlangParser.ID, 0); }
		public FuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).enterFunc(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof mlangListener ) ((mlangListener)listener).exitFunc(this);
		}
	}

	public final FuncContext func() throws RecognitionException {
		FuncContext _localctx = new FuncContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_func);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 2:
			return lvalue_sempred((LvalueContext)_localctx, predIndex);
		case 3:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean lvalue_sempred(LvalueContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		case 5:
			return precpred(_ctx, 12);
		case 6:
			return precpred(_ctx, 11);
		case 7:
			return precpred(_ctx, 10);
		case 8:
			return precpred(_ctx, 9);
		case 9:
			return precpred(_ctx, 8);
		case 10:
			return precpred(_ctx, 18);
		case 11:
			return precpred(_ctx, 17);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u001e\u0095\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0001\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0003\u0000\u001b\b\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0005\u0002,\b\u0002\n\u0002\f\u0002/\t\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u00035\b\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003;\b\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0003\u0003@\b\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0003\u0003K\b\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0005\u0003p\b\u0003\n\u0003\f\u0003s\t"+
		"\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004x\b\u0004\n\u0004"+
		"\f\u0004{\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005\u0080"+
		"\b\u0005\n\u0005\f\u0005\u0083\t\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0005\u0006\u0088\b\u0006\n\u0006\f\u0006\u008b\t\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t"+
		"\u0000\u0002\u0004\u0006\n\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0000\u0005\u0001\u0000\u0006\u0007\u0001\u0000\b\t\u0001\u0000\n\r\u0001"+
		"\u0000\u000e\u000f\u0001\u0000\u0010\u0011\u00a4\u0000\u001a\u0001\u0000"+
		"\u0000\u0000\u0002\u001c\u0001\u0000\u0000\u0000\u0004 \u0001\u0000\u0000"+
		"\u0000\u0006J\u0001\u0000\u0000\u0000\bt\u0001\u0000\u0000\u0000\n|\u0001"+
		"\u0000\u0000\u0000\f\u0084\u0001\u0000\u0000\u0000\u000e\u008c\u0001\u0000"+
		"\u0000\u0000\u0010\u0090\u0001\u0000\u0000\u0000\u0012\u0092\u0001\u0000"+
		"\u0000\u0000\u0014\u0015\u0003\u0002\u0001\u0000\u0015\u0016\u0005\u0000"+
		"\u0000\u0001\u0016\u001b\u0001\u0000\u0000\u0000\u0017\u0018\u0003\u0006"+
		"\u0003\u0000\u0018\u0019\u0005\u0000\u0000\u0001\u0019\u001b\u0001\u0000"+
		"\u0000\u0000\u001a\u0014\u0001\u0000\u0000\u0000\u001a\u0017\u0001\u0000"+
		"\u0000\u0000\u001b\u0001\u0001\u0000\u0000\u0000\u001c\u001d\u0003\u0004"+
		"\u0002\u0000\u001d\u001e\u0005\u0001\u0000\u0000\u001e\u001f\u0003\u0006"+
		"\u0003\u0000\u001f\u0003\u0001\u0000\u0000\u0000 !\u0006\u0002\uffff\uffff"+
		"\u0000!\"\u0005\u001a\u0000\u0000\"-\u0001\u0000\u0000\u0000#$\n\u0002"+
		"\u0000\u0000$%\u0005\u0002\u0000\u0000%&\u0003\u0006\u0003\u0000&\'\u0005"+
		"\u0003\u0000\u0000\',\u0001\u0000\u0000\u0000()\n\u0001\u0000\u0000)*"+
		"\u0005\u0018\u0000\u0000*,\u0005\u001a\u0000\u0000+#\u0001\u0000\u0000"+
		"\u0000+(\u0001\u0000\u0000\u0000,/\u0001\u0000\u0000\u0000-+\u0001\u0000"+
		"\u0000\u0000-.\u0001\u0000\u0000\u0000.\u0005\u0001\u0000\u0000\u0000"+
		"/-\u0001\u0000\u0000\u000001\u0006\u0003\uffff\uffff\u000012\u0003\u0012"+
		"\t\u000024\u0005\u0004\u0000\u000035\u0003\b\u0004\u000043\u0001\u0000"+
		"\u0000\u000045\u0001\u0000\u0000\u000056\u0001\u0000\u0000\u000067\u0005"+
		"\u0005\u0000\u00007K\u0001\u0000\u0000\u00008:\u0005\u0002\u0000\u0000"+
		"9;\u0003\n\u0005\u0000:9\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000"+
		";<\u0001\u0000\u0000\u0000<K\u0005\u0003\u0000\u0000=?\u0005\u0014\u0000"+
		"\u0000>@\u0003\f\u0006\u0000?>\u0001\u0000\u0000\u0000?@\u0001\u0000\u0000"+
		"\u0000@A\u0001\u0000\u0000\u0000AK\u0005\u0015\u0000\u0000BC\u0005\u0004"+
		"\u0000\u0000CD\u0003\u0006\u0003\u0000DE\u0005\u0005\u0000\u0000EK\u0001"+
		"\u0000\u0000\u0000FK\u0005\u001b\u0000\u0000GK\u0005\u001c\u0000\u0000"+
		"HK\u0005\u0019\u0000\u0000IK\u0005\u001a\u0000\u0000J0\u0001\u0000\u0000"+
		"\u0000J8\u0001\u0000\u0000\u0000J=\u0001\u0000\u0000\u0000JB\u0001\u0000"+
		"\u0000\u0000JF\u0001\u0000\u0000\u0000JG\u0001\u0000\u0000\u0000JH\u0001"+
		"\u0000\u0000\u0000JI\u0001\u0000\u0000\u0000Kq\u0001\u0000\u0000\u0000"+
		"LM\n\u000f\u0000\u0000MN\u0007\u0000\u0000\u0000Np\u0003\u0006\u0003\u0010"+
		"OP\n\u000e\u0000\u0000PQ\u0007\u0001\u0000\u0000Qp\u0003\u0006\u0003\u000f"+
		"RS\n\r\u0000\u0000ST\u0007\u0002\u0000\u0000Tp\u0003\u0006\u0003\u000e"+
		"UV\n\f\u0000\u0000VW\u0007\u0003\u0000\u0000Wp\u0003\u0006\u0003\rXY\n"+
		"\u000b\u0000\u0000YZ\u0007\u0004\u0000\u0000Z[\u0003\u0010\b\u0000[\\"+
		"\u0003\u0006\u0003\f\\p\u0001\u0000\u0000\u0000]^\n\n\u0000\u0000^_\u0003"+
		"\u0010\b\u0000_`\u0003\u0006\u0003\u000b`p\u0001\u0000\u0000\u0000ab\n"+
		"\t\u0000\u0000bc\u0005\u0012\u0000\u0000cp\u0003\u0006\u0003\nde\n\b\u0000"+
		"\u0000ef\u0005\u0013\u0000\u0000fp\u0003\u0006\u0003\tgh\n\u0012\u0000"+
		"\u0000hi\u0005\u0002\u0000\u0000ij\u0003\u0006\u0003\u0000jk\u0005\u0003"+
		"\u0000\u0000kp\u0001\u0000\u0000\u0000lm\n\u0011\u0000\u0000mn\u0005\u0018"+
		"\u0000\u0000np\u0005\u001a\u0000\u0000oL\u0001\u0000\u0000\u0000oO\u0001"+
		"\u0000\u0000\u0000oR\u0001\u0000\u0000\u0000oU\u0001\u0000\u0000\u0000"+
		"oX\u0001\u0000\u0000\u0000o]\u0001\u0000\u0000\u0000oa\u0001\u0000\u0000"+
		"\u0000od\u0001\u0000\u0000\u0000og\u0001\u0000\u0000\u0000ol\u0001\u0000"+
		"\u0000\u0000ps\u0001\u0000\u0000\u0000qo\u0001\u0000\u0000\u0000qr\u0001"+
		"\u0000\u0000\u0000r\u0007\u0001\u0000\u0000\u0000sq\u0001\u0000\u0000"+
		"\u0000ty\u0003\u0006\u0003\u0000uv\u0005\u0016\u0000\u0000vx\u0003\u0006"+
		"\u0003\u0000wu\u0001\u0000\u0000\u0000x{\u0001\u0000\u0000\u0000yw\u0001"+
		"\u0000\u0000\u0000yz\u0001\u0000\u0000\u0000z\t\u0001\u0000\u0000\u0000"+
		"{y\u0001\u0000\u0000\u0000|\u0081\u0003\u0006\u0003\u0000}~\u0005\u0016"+
		"\u0000\u0000~\u0080\u0003\u0006\u0003\u0000\u007f}\u0001\u0000\u0000\u0000"+
		"\u0080\u0083\u0001\u0000\u0000\u0000\u0081\u007f\u0001\u0000\u0000\u0000"+
		"\u0081\u0082\u0001\u0000\u0000\u0000\u0082\u000b\u0001\u0000\u0000\u0000"+
		"\u0083\u0081\u0001\u0000\u0000\u0000\u0084\u0089\u0003\u000e\u0007\u0000"+
		"\u0085\u0086\u0005\u0016\u0000\u0000\u0086\u0088\u0003\u000e\u0007\u0000"+
		"\u0087\u0085\u0001\u0000\u0000\u0000\u0088\u008b\u0001\u0000\u0000\u0000"+
		"\u0089\u0087\u0001\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000\u0000"+
		"\u008a\r\u0001\u0000\u0000\u0000\u008b\u0089\u0001\u0000\u0000\u0000\u008c"+
		"\u008d\u0003\u0006\u0003\u0000\u008d\u008e\u0005\u0017\u0000\u0000\u008e"+
		"\u008f\u0003\u0006\u0003\u0000\u008f\u000f\u0001\u0000\u0000\u0000\u0090"+
		"\u0091\u0005\u001a\u0000\u0000\u0091\u0011\u0001\u0000\u0000\u0000\u0092"+
		"\u0093\u0005\u001a\u0000\u0000\u0093\u0013\u0001\u0000\u0000\u0000\f\u001a"+
		"+-4:?Joqy\u0081\u0089";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}