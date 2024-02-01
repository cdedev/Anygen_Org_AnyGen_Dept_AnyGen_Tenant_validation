// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796744001/Rule3.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule3Listener is a complete listener for a parse tree produced by Rule3Parser.
type Rule3Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField0 is called when entering the field0 production.
	EnterField0(c *Field0Context)

	// EnterField1 is called when entering the field1 production.
	EnterField1(c *Field1Context)

	// EnterField2 is called when entering the field2 production.
	EnterField2(c *Field2Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField0 is called when exiting the field0 production.
	ExitField0(c *Field0Context)

	// ExitField1 is called when exiting the field1 production.
	ExitField1(c *Field1Context)

	// ExitField2 is called when exiting the field2 production.
	ExitField2(c *Field2Context)
}
