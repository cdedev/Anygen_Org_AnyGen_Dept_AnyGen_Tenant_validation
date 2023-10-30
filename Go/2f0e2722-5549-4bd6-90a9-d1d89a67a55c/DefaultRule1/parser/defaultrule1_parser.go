// Code generated from /usr/local/lib/Go/2f0e2722-5549-4bd6-90a9-d1d89a67a55c/1698641154656/DefaultRule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DefaultRule1

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DefaultRule1Parser struct {
	*antlr.BaseParser
}

var defaultrule1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func defaultrule1ParserInit() {
	staticData := &defaultrule1ParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TEXTTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "texttype",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 10, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 7, 0, 4, 1, 0, 0, 0, 2, 7, 1, 0, 0, 0, 4, 5, 3, 2,
		1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 5, 1, 0, 0, 8, 3, 1, 0,
		0, 0, 0,
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

// DefaultRule1ParserInit initializes any static state used to implement DefaultRule1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDefaultRule1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DefaultRule1ParserInit() {
	staticData := &defaultrule1ParserStaticData
	staticData.once.Do(defaultrule1ParserInit)
}

// NewDefaultRule1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewDefaultRule1Parser(input antlr.TokenStream) *DefaultRule1Parser {
	DefaultRule1ParserInit()
	this := new(DefaultRule1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &defaultrule1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DefaultRule1.g4"

	return this
}

// DefaultRule1Parser tokens.
const (
	DefaultRule1ParserEOF      = antlr.TokenEOF
	DefaultRule1ParserTEXTTYPE = 1
	DefaultRule1ParserOWNKEY   = 2
	DefaultRule1ParserSPLITKEY = 3
	DefaultRule1ParserWS       = 4
)

// DefaultRule1Parser rules.
const (
	DefaultRule1ParserRULE_expression = 0
	DefaultRule1ParserRULE_texttype   = 1
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefaultRule1ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultRule1ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Texttype() ITexttypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITexttypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITexttypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DefaultRule1ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultRule1Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultRule1Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DefaultRule1Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DefaultRule1ParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Texttype()
	}
	{
		p.SetState(5)
		p.Match(DefaultRule1ParserEOF)
	}

	return localctx
}

// ITexttypeContext is an interface to support dynamic dispatch.
type ITexttypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTexttypeContext differentiates from other interfaces.
	IsTexttypeContext()
}

type TexttypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTexttypeContext() *TexttypeContext {
	var p = new(TexttypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefaultRule1ParserRULE_texttype
	return p
}

func (*TexttypeContext) IsTexttypeContext() {}

func NewTexttypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TexttypeContext {
	var p = new(TexttypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultRule1ParserRULE_texttype

	return p
}

func (s *TexttypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TexttypeContext) TEXTTYPE() antlr.TerminalNode {
	return s.GetToken(DefaultRule1ParserTEXTTYPE, 0)
}

func (s *TexttypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TexttypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TexttypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultRule1Listener); ok {
		listenerT.EnterTexttype(s)
	}
}

func (s *TexttypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultRule1Listener); ok {
		listenerT.ExitTexttype(s)
	}
}

func (p *DefaultRule1Parser) Texttype() (localctx ITexttypeContext) {
	this := p
	_ = this

	localctx = NewTexttypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DefaultRule1ParserRULE_texttype)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)
		p.Match(DefaultRule1ParserTEXTTYPE)
	}

	return localctx
}
