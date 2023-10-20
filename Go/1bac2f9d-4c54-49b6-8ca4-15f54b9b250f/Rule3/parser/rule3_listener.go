// Code generated from /usr/local/lib/Go/1bac2f9d-4c54-49b6-8ca4-15f54b9b250f/1697785004374/Rule3.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule3Listener is a complete listener for a parse tree produced by Rule3Parser.
type Rule3Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField0 is called when entering the field0 production.
	EnterField0(c *Field0Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField0 is called when exiting the field0 production.
	ExitField0(c *Field0Context)
}
