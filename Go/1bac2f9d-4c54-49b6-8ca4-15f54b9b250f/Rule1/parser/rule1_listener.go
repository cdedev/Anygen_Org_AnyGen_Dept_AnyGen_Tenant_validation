// Code generated from /usr/local/lib/Go/1bac2f9d-4c54-49b6-8ca4-15f54b9b250f/1697784827775/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule1Listener is a complete listener for a parse tree produced by Rule1Parser.
type Rule1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField1 is called when entering the field1 production.
	EnterField1(c *Field1Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField1 is called when exiting the field1 production.
	ExitField1(c *Field1Context)
}
