// Code generated from /usr/local/lib/Go/1bac2f9d-4c54-49b6-8ca4-15f54b9b250f/1697784827775/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1

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

type Rule1Parser struct {
	*antlr.BaseParser
}

var rule1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rule1ParserInit() {
	staticData := &rule1ParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD1", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "field1",
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

// Rule1ParserInit initializes any static state used to implement Rule1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRule1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule1ParserInit() {
	staticData := &rule1ParserStaticData
	staticData.once.Do(rule1ParserInit)
}

// NewRule1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewRule1Parser(input antlr.TokenStream) *Rule1Parser {
	Rule1ParserInit()
	this := new(Rule1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rule1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rule1.g4"

	return this
}

// Rule1Parser tokens.
const (
	Rule1ParserEOF      = antlr.TokenEOF
	Rule1ParserFIELD1   = 1
	Rule1ParserOWNKEY   = 2
	Rule1ParserSPLITKEY = 3
	Rule1ParserWS       = 4
)

// Rule1Parser rules.
const (
	Rule1ParserRULE_expression = 0
	Rule1ParserRULE_field1     = 1
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
	p.RuleIndex = Rule1ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Field1() IField1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField1Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Rule1ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Rule1Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Rule1ParserRULE_expression)

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
		p.Field1()
	}
	{
		p.SetState(5)
		p.Match(Rule1ParserEOF)
	}

	return localctx
}

// IField1Context is an interface to support dynamic dispatch.
type IField1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField1Context differentiates from other interfaces.
	IsField1Context()
}

type Field1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField1Context() *Field1Context {
	var p = new(Field1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule1ParserRULE_field1
	return p
}

func (*Field1Context) IsField1Context() {}

func NewField1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field1Context {
	var p = new(Field1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_field1

	return p
}

func (s *Field1Context) GetParser() antlr.Parser { return s.parser }

func (s *Field1Context) FIELD1() antlr.TerminalNode {
	return s.GetToken(Rule1ParserFIELD1, 0)
}

func (s *Field1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterField1(s)
	}
}

func (s *Field1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitField1(s)
	}
}

func (p *Rule1Parser) Field1() (localctx IField1Context) {
	this := p
	_ = this

	localctx = NewField1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Rule1ParserRULE_field1)

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
		p.Match(Rule1ParserFIELD1)
	}

	return localctx
}
