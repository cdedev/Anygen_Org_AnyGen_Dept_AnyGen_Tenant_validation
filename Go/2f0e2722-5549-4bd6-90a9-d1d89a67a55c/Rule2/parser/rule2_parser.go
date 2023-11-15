// Code generated from /usr/local/lib/Go/2f0e2722-5549-4bd6-90a9-d1d89a67a55c/1700029269676/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

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
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "QUESTIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "questions",
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
	Rule2ParserEOF       = antlr.TokenEOF
	Rule2ParserQUESTIONS = 1
	Rule2ParserOWNKEY    = 2
	Rule2ParserSPLITKEY  = 3
	Rule2ParserWS        = 4
)

// Rule2Parser rules.
const (
	Rule2ParserRULE_expression = 0
	Rule2ParserRULE_questions  = 1
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

func (s *ExpressionContext) Questions() IQuestionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuestionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuestionsContext)
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
		p.SetState(4)
		p.Questions()
	}
	{
		p.SetState(5)
		p.Match(Rule2ParserEOF)
	}

	return localctx
}

// IQuestionsContext is an interface to support dynamic dispatch.
type IQuestionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuestionsContext differentiates from other interfaces.
	IsQuestionsContext()
}

type QuestionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuestionsContext() *QuestionsContext {
	var p = new(QuestionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule2ParserRULE_questions
	return p
}

func (*QuestionsContext) IsQuestionsContext() {}

func NewQuestionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionsContext {
	var p = new(QuestionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_questions

	return p
}

func (s *QuestionsContext) GetParser() antlr.Parser { return s.parser }

func (s *QuestionsContext) QUESTIONS() antlr.TerminalNode {
	return s.GetToken(Rule2ParserQUESTIONS, 0)
}

func (s *QuestionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuestionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterQuestions(s)
	}
}

func (s *QuestionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitQuestions(s)
	}
}

func (p *Rule2Parser) Questions() (localctx IQuestionsContext) {
	this := p
	_ = this

	localctx = NewQuestionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Rule2ParserRULE_questions)

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
		p.Match(Rule2ParserQUESTIONS)
	}

	return localctx
}
