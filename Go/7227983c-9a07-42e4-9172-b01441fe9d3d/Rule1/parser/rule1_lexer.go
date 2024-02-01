// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796480666/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type Rule1Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rule1lexerLexerStaticData struct {
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

func rule1lexerLexerInit() {
	staticData := &rule1lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD1", "FIELD2", "FIELD3", "FIELD4", "FIELD5", "OWNKEY", "SPLITKEY",
		"WS",
	}
	staticData.ruleNames = []string{
		"FIELD1", "FIELD2", "FIELD3", "FIELD4", "FIELD5", "OWNKEY", "SPLITKEY",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 92, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 5, 0, 21,
		8, 0, 10, 0, 12, 0, 24, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 30, 8, 1, 10,
		1, 12, 1, 33, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 40, 8, 2, 10, 2,
		12, 2, 43, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 51, 8, 3, 10,
		3, 12, 3, 54, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 63,
		8, 4, 10, 4, 12, 4, 66, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		4, 7, 87, 8, 7, 11, 7, 12, 7, 88, 1, 7, 1, 7, 0, 0, 8, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 6, 1, 0, 48, 57, 4, 0, 32, 32, 48,
		57, 65, 90, 97, 122, 4, 0, 32, 32, 46, 46, 65, 90, 97, 122, 2, 0, 46, 46,
		48, 57, 5, 0, 32, 32, 46, 46, 48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 97, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5, 34, 1, 0, 0,
		0, 7, 44, 1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 67, 1, 0, 0, 0, 13, 71, 1,
		0, 0, 0, 15, 86, 1, 0, 0, 0, 17, 18, 3, 13, 6, 0, 18, 22, 3, 11, 5, 0,
		19, 21, 7, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 22, 20, 1,
		0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 2, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 25,
		26, 3, 13, 6, 0, 26, 27, 3, 11, 5, 0, 27, 31, 3, 11, 5, 0, 28, 30, 7, 1,
		0, 0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32,
		1, 0, 0, 0, 32, 4, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 3, 13, 6, 0,
		35, 36, 3, 11, 5, 0, 36, 37, 3, 11, 5, 0, 37, 41, 3, 11, 5, 0, 38, 40,
		7, 2, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0,
		41, 42, 1, 0, 0, 0, 42, 6, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 3, 13,
		6, 0, 45, 46, 3, 11, 5, 0, 46, 47, 3, 11, 5, 0, 47, 48, 3, 11, 5, 0, 48,
		52, 3, 11, 5, 0, 49, 51, 7, 3, 0, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0,
		0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 8, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 55, 56, 3, 13, 6, 0, 56, 57, 3, 11, 5, 0, 57, 58, 3, 11, 5,
		0, 58, 59, 3, 11, 5, 0, 59, 60, 3, 11, 5, 0, 60, 64, 3, 11, 5, 0, 61, 63,
		7, 4, 0, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 10, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 5,
		36, 0, 0, 68, 69, 5, 126, 0, 0, 69, 70, 5, 36, 0, 0, 70, 12, 1, 0, 0, 0,
		71, 72, 5, 32, 0, 0, 72, 73, 5, 33, 0, 0, 73, 74, 5, 36, 0, 0, 74, 75,
		5, 126, 0, 0, 75, 76, 5, 36, 0, 0, 76, 77, 5, 35, 0, 0, 77, 78, 5, 37,
		0, 0, 78, 79, 5, 35, 0, 0, 79, 80, 5, 38, 0, 0, 80, 81, 5, 36, 0, 0, 81,
		82, 5, 38, 0, 0, 82, 83, 5, 33, 0, 0, 83, 84, 5, 32, 0, 0, 84, 14, 1, 0,
		0, 0, 85, 87, 7, 5, 0, 0, 86, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86,
		1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 6, 7, 0, 0,
		91, 16, 1, 0, 0, 0, 7, 0, 22, 31, 41, 52, 64, 88, 1, 6, 0, 0,
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

// Rule1LexerInit initializes any static state used to implement Rule1Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRule1Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule1LexerInit() {
	staticData := &rule1lexerLexerStaticData
	staticData.once.Do(rule1lexerLexerInit)
}

// NewRule1Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRule1Lexer(input antlr.CharStream) *Rule1Lexer {
	Rule1LexerInit()
	l := new(Rule1Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rule1lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rule1.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Rule1Lexer tokens.
const (
	Rule1LexerFIELD1   = 1
	Rule1LexerFIELD2   = 2
	Rule1LexerFIELD3   = 3
	Rule1LexerFIELD4   = 4
	Rule1LexerFIELD5   = 5
	Rule1LexerOWNKEY   = 6
	Rule1LexerSPLITKEY = 7
	Rule1LexerWS       = 8
)
