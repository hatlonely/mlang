// Generated from /Users/mvbj0483/hatlonely/github.com/hatlonely/mlang/mlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link mlangParser}.
 */
public interface mlangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link mlangParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(mlangParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(mlangParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterStat(mlangParser.StatContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitStat(mlangParser.StatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IndexLvalue}
	 * labeled alternative in {@link mlangParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void enterIndexLvalue(mlangParser.IndexLvalueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IndexLvalue}
	 * labeled alternative in {@link mlangParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void exitIndexLvalue(mlangParser.IndexLvalueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SimpleLvalue}
	 * labeled alternative in {@link mlangParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void enterSimpleLvalue(mlangParser.SimpleLvalueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SimpleLvalue}
	 * labeled alternative in {@link mlangParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void exitSimpleLvalue(mlangParser.SimpleLvalueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FieldLvalue}
	 * labeled alternative in {@link mlangParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void enterFieldLvalue(mlangParser.FieldLvalueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FieldLvalue}
	 * labeled alternative in {@link mlangParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void exitFieldLvalue(mlangParser.FieldLvalueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dictionary}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterDictionary(mlangParser.DictionaryContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dictionary}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitDictionary(mlangParser.DictionaryContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OrOp}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterOrOp(mlangParser.OrOpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OrOp}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitOrOp(mlangParser.OrOpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RelationalSymbol}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterRelationalSymbol(mlangParser.RelationalSymbolContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RelationalSymbol}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitRelationalSymbol(mlangParser.RelationalSymbolContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(mlangParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(mlangParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(mlangParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(mlangParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterParens(mlangParser.ParensContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitParens(mlangParser.ParensContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IndexAccess}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIndexAccess(mlangParser.IndexAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IndexAccess}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIndexAccess(mlangParser.IndexAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code EqualitySymbol}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterEqualitySymbol(mlangParser.EqualitySymbolContext ctx);
	/**
	 * Exit a parse tree produced by the {@code EqualitySymbol}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitEqualitySymbol(mlangParser.EqualitySymbolContext ctx);
	/**
	 * Enter a parse tree produced by the {@code String}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterString(mlangParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code String}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitString(mlangParser.StringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CompareFuncInfix}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterCompareFuncInfix(mlangParser.CompareFuncInfixContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CompareFuncInfix}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitCompareFuncInfix(mlangParser.CompareFuncInfixContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Array}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArray(mlangParser.ArrayContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Array}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArray(mlangParser.ArrayContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NotCompareFuncInfix}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNotCompareFuncInfix(mlangParser.NotCompareFuncInfixContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NotCompareFuncInfix}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNotCompareFuncInfix(mlangParser.NotCompareFuncInfixContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Number}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNumber(mlangParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Number}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNumber(mlangParser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FunctionCall}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(mlangParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FunctionCall}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(mlangParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Id}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterId(mlangParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Id}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitId(mlangParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Boolean}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterBoolean(mlangParser.BooleanContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Boolean}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitBoolean(mlangParser.BooleanContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FieldAccess}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFieldAccess(mlangParser.FieldAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FieldAccess}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFieldAccess(mlangParser.FieldAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AndOp}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAndOp(mlangParser.AndOpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AndOp}
	 * labeled alternative in {@link mlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAndOp(mlangParser.AndOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#exprList}.
	 * @param ctx the parse tree
	 */
	void enterExprList(mlangParser.ExprListContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#exprList}.
	 * @param ctx the parse tree
	 */
	void exitExprList(mlangParser.ExprListContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#arrayElements}.
	 * @param ctx the parse tree
	 */
	void enterArrayElements(mlangParser.ArrayElementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#arrayElements}.
	 * @param ctx the parse tree
	 */
	void exitArrayElements(mlangParser.ArrayElementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#dictElements}.
	 * @param ctx the parse tree
	 */
	void enterDictElements(mlangParser.DictElementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#dictElements}.
	 * @param ctx the parse tree
	 */
	void exitDictElements(mlangParser.DictElementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#dictPair}.
	 * @param ctx the parse tree
	 */
	void enterDictPair(mlangParser.DictPairContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#dictPair}.
	 * @param ctx the parse tree
	 */
	void exitDictPair(mlangParser.DictPairContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#binaryOp}.
	 * @param ctx the parse tree
	 */
	void enterBinaryOp(mlangParser.BinaryOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#binaryOp}.
	 * @param ctx the parse tree
	 */
	void exitBinaryOp(mlangParser.BinaryOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link mlangParser#func}.
	 * @param ctx the parse tree
	 */
	void enterFunc(mlangParser.FuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link mlangParser#func}.
	 * @param ctx the parse tree
	 */
	void exitFunc(mlangParser.FuncContext ctx);
}