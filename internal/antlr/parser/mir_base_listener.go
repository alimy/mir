// Code generated from Mir.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Mir

import "github.com/alimy/antlr4-go"

// BaseMirListener is a complete listener for a parse tree produced by MirParser.
type BaseMirListener struct{}

var _ MirListener = &BaseMirListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMirListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMirListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMirListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMirListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterR is called when production r is entered.
func (s *BaseMirListener) EnterR(ctx *RContext) {}

// ExitR is called when production r is exited.
func (s *BaseMirListener) ExitR(ctx *RContext) {}
