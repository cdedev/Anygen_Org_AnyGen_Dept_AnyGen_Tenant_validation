// Code generated from /usr/local/lib/Go/2f0e2722-5549-4bd6-90a9-d1d89a67a55c/1698641154656/DefaultRule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DefaultRule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDefaultRule1Listener is a complete listener for a parse tree produced by DefaultRule1Parser.
type BaseDefaultRule1Listener struct{}

var _ DefaultRule1Listener = &BaseDefaultRule1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDefaultRule1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDefaultRule1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDefaultRule1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDefaultRule1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDefaultRule1Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDefaultRule1Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterTexttype is called when production texttype is entered.
func (s *BaseDefaultRule1Listener) EnterTexttype(ctx *TexttypeContext) {}

// ExitTexttype is called when production texttype is exited.
func (s *BaseDefaultRule1Listener) ExitTexttype(ctx *TexttypeContext) {}
