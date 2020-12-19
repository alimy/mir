// Code generated from Mir.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/alimy/antlr4-go"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 27, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 3, 6, 3, 17, 10, 3, 13, 3, 14, 3, 18, 3, 4, 6, 4, 22, 10, 4, 13,
	4, 14, 4, 23, 3, 4, 3, 4, 2, 2, 5, 3, 3, 5, 4, 7, 5, 3, 2, 4, 3, 2, 99,
	124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 28, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 3, 9, 3, 2, 2, 2, 5, 16, 3, 2, 2, 2, 7, 21, 3,
	2, 2, 2, 9, 10, 7, 106, 2, 2, 10, 11, 7, 103, 2, 2, 11, 12, 7, 110, 2,
	2, 12, 13, 7, 110, 2, 2, 13, 14, 7, 113, 2, 2, 14, 4, 3, 2, 2, 2, 15, 17,
	9, 2, 2, 2, 16, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2,
	18, 19, 3, 2, 2, 2, 19, 6, 3, 2, 2, 2, 20, 22, 9, 3, 2, 2, 21, 20, 3, 2,
	2, 2, 22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25,
	3, 2, 2, 2, 25, 26, 8, 4, 2, 2, 26, 8, 3, 2, 2, 2, 5, 2, 18, 23, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'hello'",
}

var lexerSymbolicNames = []string{
	"", "", "ID", "WS",
}

var lexerRuleNames = []string{
	"T__0", "ID", "WS",
}

type MirLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewMirLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *MirLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMirLexer(input antlr.CharStream) *MirLexer {
	l := new(MirLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Mir.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MirLexer tokens.
const (
	MirLexerT__0 = 1
	MirLexerID   = 2
	MirLexerWS   = 3
)
