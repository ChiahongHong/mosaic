// Code generated from Policy.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 127,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 6, 2, 54, 10, 2, 13, 2,
	14, 2, 55, 3, 2, 3, 2, 3, 2, 7, 2, 61, 10, 2, 12, 2, 14, 2, 64, 11, 2,
	3, 3, 6, 3, 67, 10, 3, 13, 3, 14, 3, 68, 3, 4, 6, 4, 72, 10, 4, 13, 4,
	14, 4, 73, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 116,
	10, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25,
	3, 25, 2, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 3, 2, 4, 5, 2, 11, 12, 15,
	15, 34, 34, 6, 2, 37, 40, 47, 48, 165, 165, 169, 169, 2, 127, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 3, 53, 3, 2, 2, 2, 5, 66, 3, 2, 2, 2, 7, 71, 3, 2, 2,
	2, 9, 77, 3, 2, 2, 2, 11, 80, 3, 2, 2, 2, 13, 83, 3, 2, 2, 2, 15, 85, 3,
	2, 2, 2, 17, 88, 3, 2, 2, 2, 19, 90, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23,
	95, 3, 2, 2, 2, 25, 98, 3, 2, 2, 2, 27, 101, 3, 2, 2, 2, 29, 103, 3, 2,
	2, 2, 31, 105, 3, 2, 2, 2, 33, 107, 3, 2, 2, 2, 35, 109, 3, 2, 2, 2, 37,
	111, 3, 2, 2, 2, 39, 115, 3, 2, 2, 2, 41, 117, 3, 2, 2, 2, 43, 119, 3,
	2, 2, 2, 45, 121, 3, 2, 2, 2, 47, 123, 3, 2, 2, 2, 49, 125, 3, 2, 2, 2,
	51, 54, 5, 39, 20, 2, 52, 54, 5, 37, 19, 2, 53, 51, 3, 2, 2, 2, 53, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 62, 3, 2, 2, 2, 57, 61, 5, 39, 20, 2, 58, 61, 5, 45, 23, 2, 59, 61,
	5, 37, 19, 2, 60, 57, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2,
	2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 4, 3,
	2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 67, 5, 45, 23, 2, 66, 65, 3, 2, 2, 2,
	67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 6, 3, 2,
	2, 2, 70, 72, 9, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71,
	3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 8, 4, 2, 2,
	76, 8, 3, 2, 2, 2, 77, 78, 7, 63, 2, 2, 78, 79, 7, 63, 2, 2, 79, 10, 3,
	2, 2, 2, 80, 81, 7, 35, 2, 2, 81, 82, 7, 63, 2, 2, 82, 12, 3, 2, 2, 2,
	83, 84, 7, 62, 2, 2, 84, 14, 3, 2, 2, 2, 85, 86, 7, 62, 2, 2, 86, 87, 7,
	63, 2, 2, 87, 16, 3, 2, 2, 2, 88, 89, 7, 64, 2, 2, 89, 18, 3, 2, 2, 2,
	90, 91, 7, 64, 2, 2, 91, 92, 7, 63, 2, 2, 92, 20, 3, 2, 2, 2, 93, 94, 7,
	66, 2, 2, 94, 22, 3, 2, 2, 2, 95, 96, 7, 49, 2, 2, 96, 97, 7, 94, 2, 2,
	97, 24, 3, 2, 2, 2, 98, 99, 7, 94, 2, 2, 99, 100, 7, 49, 2, 2, 100, 26,
	3, 2, 2, 2, 101, 102, 7, 46, 2, 2, 102, 28, 3, 2, 2, 2, 103, 104, 7, 93,
	2, 2, 104, 30, 3, 2, 2, 2, 105, 106, 7, 95, 2, 2, 106, 32, 3, 2, 2, 2,
	107, 108, 7, 42, 2, 2, 108, 34, 3, 2, 2, 2, 109, 110, 7, 43, 2, 2, 110,
	36, 3, 2, 2, 2, 111, 112, 9, 3, 2, 2, 112, 38, 3, 2, 2, 2, 113, 116, 5,
	41, 21, 2, 114, 116, 5, 43, 22, 2, 115, 113, 3, 2, 2, 2, 115, 114, 3, 2,
	2, 2, 116, 40, 3, 2, 2, 2, 117, 118, 4, 67, 92, 2, 118, 42, 3, 2, 2, 2,
	119, 120, 4, 99, 124, 2, 120, 44, 3, 2, 2, 2, 121, 122, 4, 50, 59, 2, 122,
	46, 3, 2, 2, 2, 123, 124, 4, 51, 59, 2, 124, 48, 3, 2, 2, 2, 125, 126,
	7, 50, 2, 2, 126, 50, 3, 2, 2, 2, 10, 2, 53, 55, 60, 62, 68, 73, 115, 3,
	8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'@'", "'/\\'",
	"'\\/'", "','", "'['", "']'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "Name", "Number", "WS", "TK_EQ", "TK_NEQ", "TK_LT", "TK_LTEQ", "TK_GT",
	"TK_GTEQ", "TK_AT", "TK_AND", "TK_OR", "TK_COMMA", "TK_LBRACKET", "TK_RBRACKET",
	"TK_LPAREN", "TK_RPAREN",
}

var lexerRuleNames = []string{
	"Name", "Number", "WS", "TK_EQ", "TK_NEQ", "TK_LT", "TK_LTEQ", "TK_GT",
	"TK_GTEQ", "TK_AT", "TK_AND", "TK_OR", "TK_COMMA", "TK_LBRACKET", "TK_RBRACKET",
	"TK_LPAREN", "TK_RPAREN", "SpecialCharacter", "Letter", "LetterUppercase",
	"LetterLowercase", "Digit", "NonZeroDigit", "ZeroDigit",
}

type PolicyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPolicyLexer(input antlr.CharStream) *PolicyLexer {

	l := new(PolicyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Policy.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PolicyLexer tokens.
const (
	PolicyLexerName        = 1
	PolicyLexerNumber      = 2
	PolicyLexerWS          = 3
	PolicyLexerTK_EQ       = 4
	PolicyLexerTK_NEQ      = 5
	PolicyLexerTK_LT       = 6
	PolicyLexerTK_LTEQ     = 7
	PolicyLexerTK_GT       = 8
	PolicyLexerTK_GTEQ     = 9
	PolicyLexerTK_AT       = 10
	PolicyLexerTK_AND      = 11
	PolicyLexerTK_OR       = 12
	PolicyLexerTK_COMMA    = 13
	PolicyLexerTK_LBRACKET = 14
	PolicyLexerTK_RBRACKET = 15
	PolicyLexerTK_LPAREN   = 16
	PolicyLexerTK_RPAREN   = 17
)
