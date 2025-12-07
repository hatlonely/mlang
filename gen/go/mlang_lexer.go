// Code generated from mlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type mlangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MlangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mlanglexerLexerInit() {
	staticData := &MlangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'*'", "'/'", "'+'", "'-'", "'>'", "'<'", "'>='", "'<='", "'=='",
		"'!='", "'['", "']'", "'('", "')'", "'{'", "'}'", "','", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "DOT", "BOOLEAN", "ID", "NUMBER", "STRING", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "DOT", "BOOLEAN", "ID", "NUMBER", "STRING", "NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 165, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 112, 8, 19, 1, 20, 1, 20, 5, 20, 116, 8, 20, 10,
		20, 12, 20, 119, 9, 20, 1, 21, 4, 21, 122, 8, 21, 11, 21, 12, 21, 123,
		1, 21, 1, 21, 4, 21, 128, 8, 21, 11, 21, 12, 21, 129, 3, 21, 132, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 138, 8, 22, 10, 22, 12, 22, 141, 9,
		22, 1, 22, 1, 22, 1, 22, 5, 22, 146, 8, 22, 10, 22, 12, 22, 149, 9, 22,
		1, 22, 3, 22, 152, 8, 22, 1, 23, 3, 23, 155, 8, 23, 1, 23, 1, 23, 1, 24,
		4, 24, 160, 8, 24, 11, 24, 12, 24, 161, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34, 34, 92,
		92, 1, 0, 96, 96, 2, 0, 9, 9, 32, 32, 177, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 1, 51, 1, 0, 0, 0, 3, 53, 1, 0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0,
		0, 0, 9, 59, 1, 0, 0, 0, 11, 61, 1, 0, 0, 0, 13, 63, 1, 0, 0, 0, 15, 66,
		1, 0, 0, 0, 17, 69, 1, 0, 0, 0, 19, 72, 1, 0, 0, 0, 21, 75, 1, 0, 0, 0,
		23, 77, 1, 0, 0, 0, 25, 79, 1, 0, 0, 0, 27, 81, 1, 0, 0, 0, 29, 83, 1,
		0, 0, 0, 31, 85, 1, 0, 0, 0, 33, 87, 1, 0, 0, 0, 35, 89, 1, 0, 0, 0, 37,
		91, 1, 0, 0, 0, 39, 111, 1, 0, 0, 0, 41, 113, 1, 0, 0, 0, 43, 121, 1, 0,
		0, 0, 45, 151, 1, 0, 0, 0, 47, 154, 1, 0, 0, 0, 49, 159, 1, 0, 0, 0, 51,
		52, 5, 42, 0, 0, 52, 2, 1, 0, 0, 0, 53, 54, 5, 47, 0, 0, 54, 4, 1, 0, 0,
		0, 55, 56, 5, 43, 0, 0, 56, 6, 1, 0, 0, 0, 57, 58, 5, 45, 0, 0, 58, 8,
		1, 0, 0, 0, 59, 60, 5, 62, 0, 0, 60, 10, 1, 0, 0, 0, 61, 62, 5, 60, 0,
		0, 62, 12, 1, 0, 0, 0, 63, 64, 5, 62, 0, 0, 64, 65, 5, 61, 0, 0, 65, 14,
		1, 0, 0, 0, 66, 67, 5, 60, 0, 0, 67, 68, 5, 61, 0, 0, 68, 16, 1, 0, 0,
		0, 69, 70, 5, 61, 0, 0, 70, 71, 5, 61, 0, 0, 71, 18, 1, 0, 0, 0, 72, 73,
		5, 33, 0, 0, 73, 74, 5, 61, 0, 0, 74, 20, 1, 0, 0, 0, 75, 76, 5, 91, 0,
		0, 76, 22, 1, 0, 0, 0, 77, 78, 5, 93, 0, 0, 78, 24, 1, 0, 0, 0, 79, 80,
		5, 40, 0, 0, 80, 26, 1, 0, 0, 0, 81, 82, 5, 41, 0, 0, 82, 28, 1, 0, 0,
		0, 83, 84, 5, 123, 0, 0, 84, 30, 1, 0, 0, 0, 85, 86, 5, 125, 0, 0, 86,
		32, 1, 0, 0, 0, 87, 88, 5, 44, 0, 0, 88, 34, 1, 0, 0, 0, 89, 90, 5, 58,
		0, 0, 90, 36, 1, 0, 0, 0, 91, 92, 5, 46, 0, 0, 92, 38, 1, 0, 0, 0, 93,
		94, 5, 84, 0, 0, 94, 95, 5, 82, 0, 0, 95, 96, 5, 85, 0, 0, 96, 112, 5,
		69, 0, 0, 97, 98, 5, 70, 0, 0, 98, 99, 5, 65, 0, 0, 99, 100, 5, 76, 0,
		0, 100, 101, 5, 83, 0, 0, 101, 112, 5, 69, 0, 0, 102, 103, 5, 116, 0, 0,
		103, 104, 5, 114, 0, 0, 104, 105, 5, 117, 0, 0, 105, 112, 5, 101, 0, 0,
		106, 107, 5, 102, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5, 108, 0, 0,
		109, 110, 5, 115, 0, 0, 110, 112, 5, 101, 0, 0, 111, 93, 1, 0, 0, 0, 111,
		97, 1, 0, 0, 0, 111, 102, 1, 0, 0, 0, 111, 106, 1, 0, 0, 0, 112, 40, 1,
		0, 0, 0, 113, 117, 7, 0, 0, 0, 114, 116, 7, 1, 0, 0, 115, 114, 1, 0, 0,
		0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		42, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 122, 7, 2, 0, 0, 121, 120, 1,
		0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0,
		0, 124, 131, 1, 0, 0, 0, 125, 127, 5, 46, 0, 0, 126, 128, 7, 2, 0, 0, 127,
		126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 125, 1, 0, 0, 0, 131, 132, 1, 0,
		0, 0, 132, 44, 1, 0, 0, 0, 133, 139, 5, 34, 0, 0, 134, 138, 8, 3, 0, 0,
		135, 136, 5, 92, 0, 0, 136, 138, 9, 0, 0, 0, 137, 134, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 152, 5, 34,
		0, 0, 143, 147, 5, 96, 0, 0, 144, 146, 8, 4, 0, 0, 145, 144, 1, 0, 0, 0,
		146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 152, 5, 96, 0, 0, 151, 133,
		1, 0, 0, 0, 151, 143, 1, 0, 0, 0, 152, 46, 1, 0, 0, 0, 153, 155, 5, 13,
		0, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 157, 5, 10, 0, 0, 157, 48, 1, 0, 0, 0, 158, 160, 7, 5, 0, 0, 159,
		158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 6, 24, 0, 0, 164, 50, 1, 0,
		0, 0, 12, 0, 111, 117, 123, 129, 131, 137, 139, 147, 151, 154, 161, 1,
		6, 0, 0,
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

// mlangLexerInit initializes any static state used to implement mlangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewmlangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MlangLexerInit() {
	staticData := &MlangLexerLexerStaticData
	staticData.once.Do(mlanglexerLexerInit)
}

// NewmlangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewmlangLexer(input antlr.CharStream) *mlangLexer {
	MlangLexerInit()
	l := new(mlangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MlangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "mlang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// mlangLexer tokens.
const (
	mlangLexerT__0    = 1
	mlangLexerT__1    = 2
	mlangLexerT__2    = 3
	mlangLexerT__3    = 4
	mlangLexerT__4    = 5
	mlangLexerT__5    = 6
	mlangLexerT__6    = 7
	mlangLexerT__7    = 8
	mlangLexerT__8    = 9
	mlangLexerT__9    = 10
	mlangLexerT__10   = 11
	mlangLexerT__11   = 12
	mlangLexerT__12   = 13
	mlangLexerT__13   = 14
	mlangLexerT__14   = 15
	mlangLexerT__15   = 16
	mlangLexerT__16   = 17
	mlangLexerT__17   = 18
	mlangLexerDOT     = 19
	mlangLexerBOOLEAN = 20
	mlangLexerID      = 21
	mlangLexerNUMBER  = 22
	mlangLexerSTRING  = 23
	mlangLexerNEWLINE = 24
	mlangLexerWS      = 25
)
