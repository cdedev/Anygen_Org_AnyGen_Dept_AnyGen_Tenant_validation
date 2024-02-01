// Code generated from /usr/local/lib/Go/7227983c-9a07-42e4-9172-b01441fe9d3d/1706796480666/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule1Listener is a complete listener for a parse tree produced by Rule1Parser.
type Rule1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField1 is called when entering the field1 production.
	EnterField1(c *Field1Context)

	// EnterField2 is called when entering the field2 production.
	EnterField2(c *Field2Context)

	// EnterField3 is called when entering the field3 production.
	EnterField3(c *Field3Context)

	// EnterField4 is called when entering the field4 production.
	EnterField4(c *Field4Context)

	// EnterField5 is called when entering the field5 production.
	EnterField5(c *Field5Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField1 is called when exiting the field1 production.
	ExitField1(c *Field1Context)

	// ExitField2 is called when exiting the field2 production.
	ExitField2(c *Field2Context)

	// ExitField3 is called when exiting the field3 production.
	ExitField3(c *Field3Context)

	// ExitField4 is called when exiting the field4 production.
	ExitField4(c *Field4Context)

	// ExitField5 is called when exiting the field5 production.
	ExitField5(c *Field5Context)
}
