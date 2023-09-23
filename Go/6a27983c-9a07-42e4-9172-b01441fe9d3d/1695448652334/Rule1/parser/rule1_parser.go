// Code generated from /usr/local/lib/Go/6a27983c-9a07-42e4-9172-b01441fe9d3d/1695448652334/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

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
		"", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD0", "FIELD1", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "field0", "field1",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 15, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 0, 11, 0, 6, 1, 0, 0,
		0, 2, 10, 1, 0, 0, 0, 4, 12, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 8, 3, 4,
		2, 0, 8, 9, 5, 0, 0, 1, 9, 1, 1, 0, 0, 0, 10, 11, 5, 1, 0, 0, 11, 3, 1,
		0, 0, 0, 12, 13, 5, 2, 0, 0, 13, 5, 1, 0, 0, 0, 0,
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
	Rule1ParserFIELD0   = 1
	Rule1ParserFIELD1   = 2
	Rule1ParserOWNKEY   = 3
	Rule1ParserSPLITKEY = 4
	Rule1ParserWS       = 5
)

// Rule1Parser rules.
const (
	Rule1ParserRULE_expression = 0
	Rule1ParserRULE_field0     = 1
	Rule1ParserRULE_field1     = 2
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

func (s *ExpressionContext) Field0() IField0Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField0Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField0Context)
}

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
		p.SetState(6)
		p.Field0()
	}
	{
		p.SetState(7)
		p.Field1()
	}
	{
		p.SetState(8)
		p.Match(Rule1ParserEOF)
	}

	return localctx
}

// IField0Context is an interface to support dynamic dispatch.
type IField0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField0Context differentiates from other interfaces.
	IsField0Context()
}

type Field0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField0Context() *Field0Context {
	var p = new(Field0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule1ParserRULE_field0
	return p
}

func (*Field0Context) IsField0Context() {}

func NewField0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field0Context {
	var p = new(Field0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_field0

	return p
}

func (s *Field0Context) GetParser() antlr.Parser { return s.parser }

func (s *Field0Context) FIELD0() antlr.TerminalNode {
	return s.GetToken(Rule1ParserFIELD0, 0)
}

func (s *Field0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterField0(s)
	}
}

func (s *Field0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitField0(s)
	}
}

func (p *Rule1Parser) Field0() (localctx IField0Context) {
	this := p
	_ = this

	localctx = NewField0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Rule1ParserRULE_field0)

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
		p.SetState(10)
		p.Match(Rule1ParserFIELD0)
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
	p.EnterRule(localctx, 4, Rule1ParserRULE_field1)

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
		p.SetState(12)
		p.Match(Rule1ParserFIELD1)
	}

	return localctx
}
