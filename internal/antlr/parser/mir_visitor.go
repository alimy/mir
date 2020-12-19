// Code generated from Mir.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Mir

import "github.com/alimy/antlr4-go"

// A complete Visitor for a parse tree produced by MirParser.
type MirVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MirParser#r.
	VisitR(ctx *RContext) interface{}
}
