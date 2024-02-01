// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796744001/Rule3.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Rule3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rule3lexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rule3lexerLexerInit() {
	staticData := &rule3lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD0", "FIELD1", "FIELD2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"FIELD0", "FIELD1", "FIELD2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 65, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 5, 0, 17, 8, 0, 10, 0, 12, 0, 20,
		9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 4, 5, 60, 8, 5, 11, 5, 12, 5, 61, 1,
		5, 1, 5, 0, 0, 6, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 1, 0, 4, 3, 0, 32,
		32, 65, 90, 97, 122, 3, 0, 33, 64, 91, 95, 123, 126, 1, 0, 48, 57, 3, 0,
		9, 10, 13, 13, 32, 32, 68, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 1, 13,
		1, 0, 0, 0, 3, 21, 1, 0, 0, 0, 5, 30, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9,
		44, 1, 0, 0, 0, 11, 59, 1, 0, 0, 0, 13, 14, 3, 9, 4, 0, 14, 18, 3, 7, 3,
		0, 15, 17, 7, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17, 20, 1, 0, 0, 0, 18, 16,
		1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 2, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0,
		21, 22, 3, 9, 4, 0, 22, 23, 3, 7, 3, 0, 23, 27, 3, 7, 3, 0, 24, 26, 7,
		1, 0, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27,
		28, 1, 0, 0, 0, 28, 4, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 3, 9, 4,
		0, 31, 32, 3, 7, 3, 0, 32, 33, 3, 7, 3, 0, 33, 37, 3, 7, 3, 0, 34, 36,
		7, 2, 0, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 6, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 36,
		0, 0, 41, 42, 5, 126, 0, 0, 42, 43, 5, 36, 0, 0, 43, 8, 1, 0, 0, 0, 44,
		45, 5, 32, 0, 0, 45, 46, 5, 33, 0, 0, 46, 47, 5, 36, 0, 0, 47, 48, 5, 126,
		0, 0, 48, 49, 5, 36, 0, 0, 49, 50, 5, 35, 0, 0, 50, 51, 5, 37, 0, 0, 51,
		52, 5, 35, 0, 0, 52, 53, 5, 38, 0, 0, 53, 54, 5, 36, 0, 0, 54, 55, 5, 38,
		0, 0, 55, 56, 5, 33, 0, 0, 56, 57, 5, 32, 0, 0, 57, 10, 1, 0, 0, 0, 58,
		60, 7, 3, 0, 0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0,
		0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 6, 5, 0, 0, 64, 12,
		1, 0, 0, 0, 5, 0, 18, 27, 37, 61, 1, 6, 0, 0,
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

// Rule3LexerInit initializes any static state used to implement Rule3Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRule3Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule3LexerInit() {
	staticData := &rule3lexerLexerStaticData
	staticData.once.Do(rule3lexerLexerInit)
}

// NewRule3Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRule3Lexer(input antlr.CharStream) *Rule3Lexer {
	Rule3LexerInit()
	l := new(Rule3Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rule3lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rule3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Rule3Lexer tokens.
const (
	Rule3LexerFIELD0   = 1
	Rule3LexerFIELD1   = 2
	Rule3LexerFIELD2   = 3
	Rule3LexerOWNKEY   = 4
	Rule3LexerSPLITKEY = 5
	Rule3LexerWS       = 6
)
