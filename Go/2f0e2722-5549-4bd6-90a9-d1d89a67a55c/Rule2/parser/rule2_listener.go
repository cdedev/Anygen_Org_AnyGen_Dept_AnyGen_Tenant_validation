// Code generated from /usr/local/lib/Go/2f0e2722-5549-4bd6-90a9-d1d89a67a55c/1700029269676/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule2Listener is a complete listener for a parse tree produced by Rule2Parser.
type Rule2Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterQuestions is called when entering the questions production.
	EnterQuestions(c *QuestionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitQuestions is called when exiting the questions production.
	ExitQuestions(c *QuestionsContext)
}
