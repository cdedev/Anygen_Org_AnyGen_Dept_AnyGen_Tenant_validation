// Code generated from /usr/local/lib/Go/2f0e2722-5549-4bd6-90a9-d1d89a67a55c/1700029269676/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule2Listener is a complete listener for a parse tree produced by Rule2Parser.
type BaseRule2Listener struct{}

var _ Rule2Listener = &BaseRule2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterQuestions is called when production questions is entered.
func (s *BaseRule2Listener) EnterQuestions(ctx *QuestionsContext) {}

// ExitQuestions is called when production questions is exited.
func (s *BaseRule2Listener) ExitQuestions(ctx *QuestionsContext) {}
