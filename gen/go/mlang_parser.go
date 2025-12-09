// Code generated from mlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mlang

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type mlangParser struct {
	*antlr.BaseParser
}

var MlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mlangParserInit() {
	staticData := &MlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'['", "']'", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'>'", "'<'",
		"'>='", "'<='", "'=='", "'!='", "'NOT'", "'not'", "'&&'", "'||'", "'{'",
		"'}'", "','", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "DOT", "BOOLEAN", "ID", "NUMBER", "STRING",
		"NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "expr", "exprList", "arrayElements", "dictElements", "dictPair",
		"binaryOp", "func",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 120, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 24, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 30, 8, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 46, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 83, 8, 1, 10, 1, 12, 1, 86, 9, 1, 1, 2, 1, 2, 1,
		2, 5, 2, 91, 8, 2, 10, 2, 12, 2, 94, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 99,
		8, 3, 10, 3, 12, 3, 102, 9, 3, 1, 4, 1, 4, 1, 4, 5, 4, 107, 8, 4, 10, 4,
		12, 4, 110, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		0, 1, 2, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 5, 1, 0, 5, 6, 1, 0, 7, 8, 1,
		0, 9, 12, 1, 0, 13, 14, 1, 0, 15, 16, 134, 0, 16, 1, 0, 0, 0, 2, 45, 1,
		0, 0, 0, 4, 87, 1, 0, 0, 0, 6, 95, 1, 0, 0, 0, 8, 103, 1, 0, 0, 0, 10,
		111, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 17, 3, 2,
		1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 20, 6, 1, -1, 0, 20, 21,
		3, 14, 7, 0, 21, 23, 5, 3, 0, 0, 22, 24, 3, 4, 2, 0, 23, 22, 1, 0, 0, 0,
		23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5, 4, 0, 0, 26, 46, 1,
		0, 0, 0, 27, 29, 5, 1, 0, 0, 28, 30, 3, 6, 3, 0, 29, 28, 1, 0, 0, 0, 29,
		30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 46, 5, 2, 0, 0, 32, 34, 5, 19,
		0, 0, 33, 35, 3, 8, 4, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 46, 5, 20, 0, 0, 37, 38, 5, 3, 0, 0, 38, 39, 3, 2, 1, 0,
		39, 40, 5, 4, 0, 0, 40, 46, 1, 0, 0, 0, 41, 46, 5, 26, 0, 0, 42, 46, 5,
		27, 0, 0, 43, 46, 5, 24, 0, 0, 44, 46, 5, 25, 0, 0, 45, 19, 1, 0, 0, 0,
		45, 27, 1, 0, 0, 0, 45, 32, 1, 0, 0, 0, 45, 37, 1, 0, 0, 0, 45, 41, 1,
		0, 0, 0, 45, 42, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46,
		84, 1, 0, 0, 0, 47, 48, 10, 15, 0, 0, 48, 49, 7, 0, 0, 0, 49, 83, 3, 2,
		1, 16, 50, 51, 10, 14, 0, 0, 51, 52, 7, 1, 0, 0, 52, 83, 3, 2, 1, 15, 53,
		54, 10, 13, 0, 0, 54, 55, 7, 2, 0, 0, 55, 83, 3, 2, 1, 14, 56, 57, 10,
		12, 0, 0, 57, 58, 7, 3, 0, 0, 58, 83, 3, 2, 1, 13, 59, 60, 10, 11, 0, 0,
		60, 61, 7, 4, 0, 0, 61, 62, 3, 12, 6, 0, 62, 63, 3, 2, 1, 12, 63, 83, 1,
		0, 0, 0, 64, 65, 10, 10, 0, 0, 65, 66, 3, 12, 6, 0, 66, 67, 3, 2, 1, 11,
		67, 83, 1, 0, 0, 0, 68, 69, 10, 9, 0, 0, 69, 70, 5, 17, 0, 0, 70, 83, 3,
		2, 1, 10, 71, 72, 10, 8, 0, 0, 72, 73, 5, 18, 0, 0, 73, 83, 3, 2, 1, 9,
		74, 75, 10, 18, 0, 0, 75, 76, 5, 1, 0, 0, 76, 77, 3, 2, 1, 0, 77, 78, 5,
		2, 0, 0, 78, 83, 1, 0, 0, 0, 79, 80, 10, 17, 0, 0, 80, 81, 5, 23, 0, 0,
		81, 83, 5, 25, 0, 0, 82, 47, 1, 0, 0, 0, 82, 50, 1, 0, 0, 0, 82, 53, 1,
		0, 0, 0, 82, 56, 1, 0, 0, 0, 82, 59, 1, 0, 0, 0, 82, 64, 1, 0, 0, 0, 82,
		68, 1, 0, 0, 0, 82, 71, 1, 0, 0, 0, 82, 74, 1, 0, 0, 0, 82, 79, 1, 0, 0,
		0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 3, 1,
		0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 92, 3, 2, 1, 0, 88, 89, 5, 21, 0, 0, 89,
		91, 3, 2, 1, 0, 90, 88, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0,
		0, 92, 93, 1, 0, 0, 0, 93, 5, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 100,
		3, 2, 1, 0, 96, 97, 5, 21, 0, 0, 97, 99, 3, 2, 1, 0, 98, 96, 1, 0, 0, 0,
		99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 7,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 108, 3, 10, 5, 0, 104, 105, 5, 21,
		0, 0, 105, 107, 3, 10, 5, 0, 106, 104, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0,
		108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 9, 1, 0, 0, 0, 110, 108,
		1, 0, 0, 0, 111, 112, 3, 2, 1, 0, 112, 113, 5, 22, 0, 0, 113, 114, 3, 2,
		1, 0, 114, 11, 1, 0, 0, 0, 115, 116, 5, 25, 0, 0, 116, 13, 1, 0, 0, 0,
		117, 118, 5, 25, 0, 0, 118, 15, 1, 0, 0, 0, 9, 23, 29, 34, 45, 82, 84,
		92, 100, 108,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// mlangParserInit initializes any static state used to implement mlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewmlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MlangParserInit() {
	staticData := &MlangParserStaticData
	staticData.once.Do(mlangParserInit)
}

// NewmlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewmlangParser(input antlr.TokenStream) *mlangParser {
	MlangParserInit()
	this := new(mlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "mlang.g4"

	return this
}

// mlangParser tokens.
const (
	mlangParserEOF     = antlr.TokenEOF
	mlangParserT__0    = 1
	mlangParserT__1    = 2
	mlangParserT__2    = 3
	mlangParserT__3    = 4
	mlangParserT__4    = 5
	mlangParserT__5    = 6
	mlangParserT__6    = 7
	mlangParserT__7    = 8
	mlangParserT__8    = 9
	mlangParserT__9    = 10
	mlangParserT__10   = 11
	mlangParserT__11   = 12
	mlangParserT__12   = 13
	mlangParserT__13   = 14
	mlangParserT__14   = 15
	mlangParserT__15   = 16
	mlangParserT__16   = 17
	mlangParserT__17   = 18
	mlangParserT__18   = 19
	mlangParserT__19   = 20
	mlangParserT__20   = 21
	mlangParserT__21   = 22
	mlangParserDOT     = 23
	mlangParserBOOLEAN = 24
	mlangParserID      = 25
	mlangParserNUMBER  = 26
	mlangParserSTRING  = 27
	mlangParserNEWLINE = 28
	mlangParserWS      = 29
)

// mlangParser rules.
const (
	mlangParserRULE_prog          = 0
	mlangParserRULE_expr          = 1
	mlangParserRULE_exprList      = 2
	mlangParserRULE_arrayElements = 3
	mlangParserRULE_dictElements  = 4
	mlangParserRULE_dictPair      = 5
	mlangParserRULE_binaryOp      = 6
	mlangParserRULE_func          = 7
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(mlangParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mlangParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expr(0)
	}
	{
		p.SetState(17)
		p.Match(mlangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DictionaryContext struct {
	ExprContext
}

func NewDictionaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DictionaryContext {
	var p = new(DictionaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DictionaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryContext) DictElements() IDictElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictElementsContext)
}

func (s *DictionaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterDictionary(s)
	}
}

func (s *DictionaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitDictionary(s)
	}
}

func (s *DictionaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitDictionary(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrOpContext struct {
	ExprContext
	op antlr.Token
}

func NewOrOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrOpContext {
	var p = new(OrOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrOpContext) GetOp() antlr.Token { return s.op }

func (s *OrOpContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterOrOp(s)
	}
}

func (s *OrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitOrOp(s)
	}
}

func (s *OrOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitOrOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalSymbolContext struct {
	ExprContext
	op antlr.Token
}

func NewRelationalSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalSymbolContext {
	var p = new(RelationalSymbolContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelationalSymbolContext) GetOp() antlr.Token { return s.op }

func (s *RelationalSymbolContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalSymbolContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalSymbolContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RelationalSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterRelationalSymbol(s)
	}
}

func (s *RelationalSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitRelationalSymbol(s)
	}
}

func (s *RelationalSymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitRelationalSymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexAccessContext struct {
	ExprContext
}

func NewIndexAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAccessContext {
	var p = new(IndexAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IndexAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAccessContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IndexAccessContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterIndexAccess(s)
	}
}

func (s *IndexAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitIndexAccess(s)
	}
}

func (s *IndexAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitIndexAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualitySymbolContext struct {
	ExprContext
	op antlr.Token
}

func NewEqualitySymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualitySymbolContext {
	var p = new(EqualitySymbolContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualitySymbolContext) GetOp() antlr.Token { return s.op }

func (s *EqualitySymbolContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualitySymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualitySymbolContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualitySymbolContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualitySymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterEqualitySymbol(s)
	}
}

func (s *EqualitySymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitEqualitySymbol(s)
	}
}

func (s *EqualitySymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitEqualitySymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(mlangParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareFuncInfixContext struct {
	ExprContext
}

func NewCompareFuncInfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareFuncInfixContext {
	var p = new(CompareFuncInfixContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CompareFuncInfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareFuncInfixContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CompareFuncInfixContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompareFuncInfixContext) BinaryOp() IBinaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
}

func (s *CompareFuncInfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterCompareFuncInfix(s)
	}
}

func (s *CompareFuncInfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitCompareFuncInfix(s)
	}
}

func (s *CompareFuncInfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitCompareFuncInfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayContext struct {
	ExprContext
}

func NewArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayContext {
	var p = new(ArrayContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ArrayElements() IArrayElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayElementsContext)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotCompareFuncInfixContext struct {
	ExprContext
}

func NewNotCompareFuncInfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotCompareFuncInfixContext {
	var p = new(NotCompareFuncInfixContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotCompareFuncInfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotCompareFuncInfixContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NotCompareFuncInfixContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotCompareFuncInfixContext) BinaryOp() IBinaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
}

func (s *NotCompareFuncInfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterNotCompareFuncInfix(s)
	}
}

func (s *NotCompareFuncInfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitNotCompareFuncInfix(s)
	}
}

func (s *NotCompareFuncInfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitNotCompareFuncInfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(mlangParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExprContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FunctionCallContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(mlangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ExprContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(mlangParserBOOLEAN, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldAccessContext struct {
	ExprContext
}

func NewFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAccessContext {
	var p = new(FieldAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(mlangParserDOT, 0)
}

func (s *FieldAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(mlangParserID, 0)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (s *FieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitFieldAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndOpContext struct {
	ExprContext
	op antlr.Token
}

func NewAndOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndOpContext {
	var p = new(AndOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndOpContext) GetOp() antlr.Token { return s.op }

func (s *AndOpContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterAndOp(s)
	}
}

func (s *AndOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitAndOp(s)
	}
}

func (s *AndOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitAndOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *mlangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, mlangParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(20)
			p.Func_()
		}
		{
			p.SetState(21)
			p.Match(mlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252182538) != 0 {
			{
				p.SetState(22)
				p.ExprList()
			}

		}
		{
			p.SetState(25)
			p.Match(mlangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewArrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.Match(mlangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252182538) != 0 {
			{
				p.SetState(28)
				p.ArrayElements()
			}

		}
		{
			p.SetState(31)
			p.Match(mlangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewDictionaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Match(mlangParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252182538) != 0 {
			{
				p.SetState(33)
				p.DictElements()
			}

		}
		{
			p.SetState(36)
			p.Match(mlangParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Match(mlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.expr(0)
		}
		{
			p.SetState(39)
			p.Match(mlangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.Match(mlangParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Match(mlangParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(mlangParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(mlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(48)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mlangParserT__4 || _la == mlangParserT__5) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(49)
					p.expr(16)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(51)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mlangParserT__6 || _la == mlangParserT__7) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(52)
					p.expr(15)
				}

			case 3:
				localctx = NewRelationalSymbolContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(54)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalSymbolContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7680) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalSymbolContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(55)
					p.expr(14)
				}

			case 4:
				localctx = NewEqualitySymbolContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(57)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualitySymbolContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mlangParserT__12 || _la == mlangParserT__13) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualitySymbolContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)
					p.expr(13)
				}

			case 5:
				localctx = NewNotCompareFuncInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(60)
					_la = p.GetTokenStream().LA(1)

					if !(_la == mlangParserT__14 || _la == mlangParserT__15) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.BinaryOp()
				}
				{
					p.SetState(62)
					p.expr(12)
				}

			case 6:
				localctx = NewCompareFuncInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(65)
					p.BinaryOp()
				}
				{
					p.SetState(66)
					p.expr(11)
				}

			case 7:
				localctx = NewAndOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(69)

					var _m = p.Match(mlangParserT__16)

					localctx.(*AndOpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)
					p.expr(10)
				}

			case 8:
				localctx = NewOrOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(72)

					var _m = p.Match(mlangParserT__17)

					localctx.(*OrOpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(73)
					p.expr(9)
				}

			case 9:
				localctx = NewIndexAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(75)
					p.Match(mlangParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)
					p.expr(0)
				}
				{
					p.SetState(77)
					p.Match(mlangParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, mlangParserRULE_expr)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(80)
					p.Match(mlangParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(81)
					p.Match(mlangParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mlangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.expr(0)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == mlangParserT__20 {
		{
			p.SetState(88)
			p.Match(mlangParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.expr(0)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayElementsContext is an interface to support dynamic dispatch.
type IArrayElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArrayElementsContext differentiates from other interfaces.
	IsArrayElementsContext()
}

type ArrayElementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElementsContext() *ArrayElementsContext {
	var p = new(ArrayElementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_arrayElements
	return p
}

func InitEmptyArrayElementsContext(p *ArrayElementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_arrayElements
}

func (*ArrayElementsContext) IsArrayElementsContext() {}

func NewArrayElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElementsContext {
	var p = new(ArrayElementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_arrayElements

	return p
}

func (s *ArrayElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElementsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayElementsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterArrayElements(s)
	}
}

func (s *ArrayElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitArrayElements(s)
	}
}

func (s *ArrayElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitArrayElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) ArrayElements() (localctx IArrayElementsContext) {
	localctx = NewArrayElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mlangParserRULE_arrayElements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.expr(0)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == mlangParserT__20 {
		{
			p.SetState(96)
			p.Match(mlangParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.expr(0)
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictElementsContext is an interface to support dynamic dispatch.
type IDictElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDictPair() []IDictPairContext
	DictPair(i int) IDictPairContext

	// IsDictElementsContext differentiates from other interfaces.
	IsDictElementsContext()
}

type DictElementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictElementsContext() *DictElementsContext {
	var p = new(DictElementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_dictElements
	return p
}

func InitEmptyDictElementsContext(p *DictElementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_dictElements
}

func (*DictElementsContext) IsDictElementsContext() {}

func NewDictElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictElementsContext {
	var p = new(DictElementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_dictElements

	return p
}

func (s *DictElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *DictElementsContext) AllDictPair() []IDictPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictPairContext); ok {
			len++
		}
	}

	tst := make([]IDictPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictPairContext); ok {
			tst[i] = t.(IDictPairContext)
			i++
		}
	}

	return tst
}

func (s *DictElementsContext) DictPair(i int) IDictPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictPairContext)
}

func (s *DictElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterDictElements(s)
	}
}

func (s *DictElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitDictElements(s)
	}
}

func (s *DictElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitDictElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) DictElements() (localctx IDictElementsContext) {
	localctx = NewDictElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mlangParserRULE_dictElements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.DictPair()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == mlangParserT__20 {
		{
			p.SetState(104)
			p.Match(mlangParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.DictPair()
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictPairContext is an interface to support dynamic dispatch.
type IDictPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsDictPairContext differentiates from other interfaces.
	IsDictPairContext()
}

type DictPairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictPairContext() *DictPairContext {
	var p = new(DictPairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_dictPair
	return p
}

func InitEmptyDictPairContext(p *DictPairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_dictPair
}

func (*DictPairContext) IsDictPairContext() {}

func NewDictPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictPairContext {
	var p = new(DictPairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_dictPair

	return p
}

func (s *DictPairContext) GetParser() antlr.Parser { return s.parser }

func (s *DictPairContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DictPairContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DictPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterDictPair(s)
	}
}

func (s *DictPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitDictPair(s)
	}
}

func (s *DictPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitDictPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) DictPair() (localctx IDictPairContext) {
	localctx = NewDictPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mlangParserRULE_dictPair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.expr(0)
	}
	{
		p.SetState(112)
		p.Match(mlangParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOpContext is an interface to support dynamic dispatch.
type IBinaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsBinaryOpContext differentiates from other interfaces.
	IsBinaryOpContext()
}

type BinaryOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpContext() *BinaryOpContext {
	var p = new(BinaryOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_binaryOp
	return p
}

func InitEmptyBinaryOpContext(p *BinaryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_binaryOp
}

func (*BinaryOpContext) IsBinaryOpContext() {}

func NewBinaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpContext {
	var p = new(BinaryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_binaryOp

	return p
}

func (s *BinaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpContext) ID() antlr.TerminalNode {
	return s.GetToken(mlangParserID, 0)
}

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterBinaryOp(s)
	}
}

func (s *BinaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitBinaryOp(s)
	}
}

func (s *BinaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitBinaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) BinaryOp() (localctx IBinaryOpContext) {
	localctx = NewBinaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mlangParserRULE_binaryOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(mlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mlangParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mlangParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(mlangParserID, 0)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mlangListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mlangVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mlangParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mlangParserRULE_func)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(mlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *mlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *mlangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
