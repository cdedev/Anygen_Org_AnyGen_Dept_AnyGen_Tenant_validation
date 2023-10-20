// Code generated from /usr/local/lib/Go/1bac2f9d-4c54-49b6-8ca4-15f54b9b250f/1697784827775/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule1Listener is a complete listener for a parse tree produced by Rule1Parser.
type BaseRule1Listener struct{}

var _ Rule1Listener = &BaseRule1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule1Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule1Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterField1 is called when production field1 is entered.
func (s *BaseRule1Listener) EnterField1(ctx *Field1Context) {}

// ExitField1 is called when production field1 is exited.
func (s *BaseRule1Listener) ExitField1(ctx *Field1Context) {}
