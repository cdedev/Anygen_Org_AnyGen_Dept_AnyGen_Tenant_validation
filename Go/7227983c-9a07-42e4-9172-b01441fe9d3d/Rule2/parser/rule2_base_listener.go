// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796649792/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

// EnterField0 is called when production field0 is entered.
func (s *BaseRule2Listener) EnterField0(ctx *Field0Context) {}

// ExitField0 is called when production field0 is exited.
func (s *BaseRule2Listener) ExitField0(ctx *Field0Context) {}

// EnterField1 is called when production field1 is entered.
func (s *BaseRule2Listener) EnterField1(ctx *Field1Context) {}

// ExitField1 is called when production field1 is exited.
func (s *BaseRule2Listener) ExitField1(ctx *Field1Context) {}

// EnterField2 is called when production field2 is entered.
func (s *BaseRule2Listener) EnterField2(ctx *Field2Context) {}

// ExitField2 is called when production field2 is exited.
func (s *BaseRule2Listener) ExitField2(ctx *Field2Context) {}
