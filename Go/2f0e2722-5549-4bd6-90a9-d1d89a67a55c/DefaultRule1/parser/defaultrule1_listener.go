// Code generated from /usr/local/lib/Go/2f0e2722-5549-4bd6-90a9-d1d89a67a55c/1698641154656/DefaultRule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DefaultRule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DefaultRule1Listener is a complete listener for a parse tree produced by DefaultRule1Parser.
type DefaultRule1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTexttype is called when entering the texttype production.
	EnterTexttype(c *TexttypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTexttype is called when exiting the texttype production.
	ExitTexttype(c *TexttypeContext)
}
