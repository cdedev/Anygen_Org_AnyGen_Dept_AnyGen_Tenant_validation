// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796649792/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule2

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

type Rule2Parser struct {
	*antlr.BaseParser
}

var rule2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rule2ParserInit() {
	staticData := &rule2ParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD0", "FIELD1", "FIELD2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "field0", "field1", "field2",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 20, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0,
		2, 4, 6, 0, 0, 15, 0, 8, 1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 15, 1, 0, 0,
		0, 6, 17, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 3, 4, 2, 0, 10, 11, 3, 6,
		3, 0, 11, 12, 5, 0, 0, 1, 12, 1, 1, 0, 0, 0, 13, 14, 5, 1, 0, 0, 14, 3,
		1, 0, 0, 0, 15, 16, 5, 2, 0, 0, 16, 5, 1, 0, 0, 0, 17, 18, 5, 3, 0, 0,
		18, 7, 1, 0, 0, 0, 0,
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

// Rule2ParserInit initializes any static state used to implement Rule2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRule2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule2ParserInit() {
	staticData := &rule2ParserStaticData
	staticData.once.Do(rule2ParserInit)
}

// NewRule2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewRule2Parser(input antlr.TokenStream) *Rule2Parser {
	Rule2ParserInit()
	this := new(Rule2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rule2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rule2.g4"

	return this
}

// Rule2Parser tokens.
const (
	Rule2ParserEOF      = antlr.TokenEOF
	Rule2ParserFIELD0   = 1
	Rule2ParserFIELD1   = 2
	Rule2ParserFIELD2   = 3
	Rule2ParserOWNKEY   = 4
	Rule2ParserSPLITKEY = 5
	Rule2ParserWS       = 6
)

// Rule2Parser rules.
const (
	Rule2ParserRULE_expression = 0
	Rule2ParserRULE_field0     = 1
	Rule2ParserRULE_field1     = 2
	Rule2ParserRULE_field2     = 3
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
	p.RuleIndex = Rule2ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_expression

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

func (s *ExpressionContext) Field2() IField2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField2Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Rule2ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Rule2Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Rule2ParserRULE_expression)

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
		p.SetState(8)
		p.Field0()
	}
	{
		p.SetState(9)
		p.Field1()
	}
	{
		p.SetState(10)
		p.Field2()
	}
	{
		p.SetState(11)
		p.Match(Rule2ParserEOF)
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
	p.RuleIndex = Rule2ParserRULE_field0
	return p
}

func (*Field0Context) IsField0Context() {}

func NewField0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field0Context {
	var p = new(Field0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_field0

	return p
}

func (s *Field0Context) GetParser() antlr.Parser { return s.parser }

func (s *Field0Context) FIELD0() antlr.TerminalNode {
	return s.GetToken(Rule2ParserFIELD0, 0)
}

func (s *Field0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterField0(s)
	}
}

func (s *Field0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitField0(s)
	}
}

func (p *Rule2Parser) Field0() (localctx IField0Context) {
	this := p
	_ = this

	localctx = NewField0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Rule2ParserRULE_field0)

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
		p.SetState(13)
		p.Match(Rule2ParserFIELD0)
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
	p.RuleIndex = Rule2ParserRULE_field1
	return p
}

func (*Field1Context) IsField1Context() {}

func NewField1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field1Context {
	var p = new(Field1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_field1

	return p
}

func (s *Field1Context) GetParser() antlr.Parser { return s.parser }

func (s *Field1Context) FIELD1() antlr.TerminalNode {
	return s.GetToken(Rule2ParserFIELD1, 0)
}

func (s *Field1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterField1(s)
	}
}

func (s *Field1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitField1(s)
	}
}

func (p *Rule2Parser) Field1() (localctx IField1Context) {
	this := p
	_ = this

	localctx = NewField1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Rule2ParserRULE_field1)

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
		p.SetState(15)
		p.Match(Rule2ParserFIELD1)
	}

	return localctx
}

// IField2Context is an interface to support dynamic dispatch.
type IField2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField2Context differentiates from other interfaces.
	IsField2Context()
}

type Field2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField2Context() *Field2Context {
	var p = new(Field2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule2ParserRULE_field2
	return p
}

func (*Field2Context) IsField2Context() {}

func NewField2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field2Context {
	var p = new(Field2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_field2

	return p
}

func (s *Field2Context) GetParser() antlr.Parser { return s.parser }

func (s *Field2Context) FIELD2() antlr.TerminalNode {
	return s.GetToken(Rule2ParserFIELD2, 0)
}

func (s *Field2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterField2(s)
	}
}

func (s *Field2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitField2(s)
	}
}

func (p *Rule2Parser) Field2() (localctx IField2Context) {
	this := p
	_ = this

	localctx = NewField2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Rule2ParserRULE_field2)

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
		p.SetState(17)
		p.Match(Rule2ParserFIELD2)
	}

	return localctx
}
