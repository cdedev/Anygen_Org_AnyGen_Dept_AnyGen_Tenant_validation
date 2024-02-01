// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796480666/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

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
		"", "", "", "", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD1", "FIELD2", "FIELD3", "FIELD4", "FIELD5", "OWNKEY", "SPLITKEY",
		"WS",
	}
	staticData.ruleNames = []string{
		"expression", "field1", "field2", "field3", "field4", "field5",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 30, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8,
		10, 0, 0, 23, 0, 12, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 21, 1, 0, 0, 0,
		6, 23, 1, 0, 0, 0, 8, 25, 1, 0, 0, 0, 10, 27, 1, 0, 0, 0, 12, 13, 3, 2,
		1, 0, 13, 14, 3, 4, 2, 0, 14, 15, 3, 6, 3, 0, 15, 16, 3, 8, 4, 0, 16, 17,
		3, 10, 5, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 20, 5, 1, 0, 0,
		20, 3, 1, 0, 0, 0, 21, 22, 5, 2, 0, 0, 22, 5, 1, 0, 0, 0, 23, 24, 5, 3,
		0, 0, 24, 7, 1, 0, 0, 0, 25, 26, 5, 4, 0, 0, 26, 9, 1, 0, 0, 0, 27, 28,
		5, 5, 0, 0, 28, 11, 1, 0, 0, 0, 0,
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
	Rule1ParserFIELD2   = 2
	Rule1ParserFIELD3   = 3
	Rule1ParserFIELD4   = 4
	Rule1ParserFIELD5   = 5
	Rule1ParserOWNKEY   = 6
	Rule1ParserSPLITKEY = 7
	Rule1ParserWS       = 8
)

// Rule1Parser rules.
const (
	Rule1ParserRULE_expression = 0
	Rule1ParserRULE_field1     = 1
	Rule1ParserRULE_field2     = 2
	Rule1ParserRULE_field3     = 3
	Rule1ParserRULE_field4     = 4
	Rule1ParserRULE_field5     = 5
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

func (s *ExpressionContext) Field3() IField3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField3Context)
}

func (s *ExpressionContext) Field4() IField4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField4Context)
}

func (s *ExpressionContext) Field5() IField5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField5Context)
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
		p.SetState(12)
		p.Field1()
	}
	{
		p.SetState(13)
		p.Field2()
	}
	{
		p.SetState(14)
		p.Field3()
	}
	{
		p.SetState(15)
		p.Field4()
	}
	{
		p.SetState(16)
		p.Field5()
	}
	{
		p.SetState(17)
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
		p.SetState(19)
		p.Match(Rule1ParserFIELD1)
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
	p.RuleIndex = Rule1ParserRULE_field2
	return p
}

func (*Field2Context) IsField2Context() {}

func NewField2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field2Context {
	var p = new(Field2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_field2

	return p
}

func (s *Field2Context) GetParser() antlr.Parser { return s.parser }

func (s *Field2Context) FIELD2() antlr.TerminalNode {
	return s.GetToken(Rule1ParserFIELD2, 0)
}

func (s *Field2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterField2(s)
	}
}

func (s *Field2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitField2(s)
	}
}

func (p *Rule1Parser) Field2() (localctx IField2Context) {
	this := p
	_ = this

	localctx = NewField2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Rule1ParserRULE_field2)

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
		p.SetState(21)
		p.Match(Rule1ParserFIELD2)
	}

	return localctx
}

// IField3Context is an interface to support dynamic dispatch.
type IField3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField3Context differentiates from other interfaces.
	IsField3Context()
}

type Field3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField3Context() *Field3Context {
	var p = new(Field3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule1ParserRULE_field3
	return p
}

func (*Field3Context) IsField3Context() {}

func NewField3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field3Context {
	var p = new(Field3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_field3

	return p
}

func (s *Field3Context) GetParser() antlr.Parser { return s.parser }

func (s *Field3Context) FIELD3() antlr.TerminalNode {
	return s.GetToken(Rule1ParserFIELD3, 0)
}

func (s *Field3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterField3(s)
	}
}

func (s *Field3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitField3(s)
	}
}

func (p *Rule1Parser) Field3() (localctx IField3Context) {
	this := p
	_ = this

	localctx = NewField3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Rule1ParserRULE_field3)

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
		p.SetState(23)
		p.Match(Rule1ParserFIELD3)
	}

	return localctx
}

// IField4Context is an interface to support dynamic dispatch.
type IField4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField4Context differentiates from other interfaces.
	IsField4Context()
}

type Field4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField4Context() *Field4Context {
	var p = new(Field4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule1ParserRULE_field4
	return p
}

func (*Field4Context) IsField4Context() {}

func NewField4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field4Context {
	var p = new(Field4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_field4

	return p
}

func (s *Field4Context) GetParser() antlr.Parser { return s.parser }

func (s *Field4Context) FIELD4() antlr.TerminalNode {
	return s.GetToken(Rule1ParserFIELD4, 0)
}

func (s *Field4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterField4(s)
	}
}

func (s *Field4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitField4(s)
	}
}

func (p *Rule1Parser) Field4() (localctx IField4Context) {
	this := p
	_ = this

	localctx = NewField4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Rule1ParserRULE_field4)

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
		p.SetState(25)
		p.Match(Rule1ParserFIELD4)
	}

	return localctx
}

// IField5Context is an interface to support dynamic dispatch.
type IField5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField5Context differentiates from other interfaces.
	IsField5Context()
}

type Field5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField5Context() *Field5Context {
	var p = new(Field5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule1ParserRULE_field5
	return p
}

func (*Field5Context) IsField5Context() {}

func NewField5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field5Context {
	var p = new(Field5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule1ParserRULE_field5

	return p
}

func (s *Field5Context) GetParser() antlr.Parser { return s.parser }

func (s *Field5Context) FIELD5() antlr.TerminalNode {
	return s.GetToken(Rule1ParserFIELD5, 0)
}

func (s *Field5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.EnterField5(s)
	}
}

func (s *Field5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule1Listener); ok {
		listenerT.ExitField5(s)
	}
}

func (p *Rule1Parser) Field5() (localctx IField5Context) {
	this := p
	_ = this

	localctx = NewField5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Rule1ParserRULE_field5)

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
		p.SetState(27)
		p.Match(Rule1ParserFIELD5)
	}

	return localctx
}
