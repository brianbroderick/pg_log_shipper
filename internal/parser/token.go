package parser

import "strings"

type Token int

var keywords map[string]Token

const (
	ILLEGAL Token = iota
	EOF
	WS // whitespace
	NIL

	literalBeg

	// Literals
	IDENT
	STRING
	SENTENCE
	NUMBER
	INTEGER
	BADSTRING
	BADESCAPE

	literalEnd

	// Misc characters
	LPAREN      // (
	RPAREN      // )
	COMMA       // ,
	COLON       // :  // used
	APOSTROPHE  // '
	DOUBLECOLON // ::
	SEMICOLON   // ;
	DOT         // .
	GT          // >
	LT          // <
	SLASH       // /
	STAR        // *

	keywordBeg
	BIND          // bind
	PARSE         // parse
	EXECUTE       // execute
	UPDATE        // update_waiting
	DURATION      // query
	MS            // milliseconds
	APPLICATION   // query comment
	CONTROLLER    // query comment
	ACTION        // query comment
	LINE          // query comment
	AUTOMATIC     // analyze, vacuum
	REPLICATION   // authentication. repl_connection
	CHECKPOINT    // checkpoint_complete, checkpoint_starting
	CONNECTION    // connection_received
	COULD         // connection_reset
	DISCONNECTION // disconnection
	TEMPORARY     // temp_table

	keywordEnd
)

// These are how a string is mapped to the token
var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",
	NIL:     "NIL",

	IDENT:     "IDENT",
	NUMBER:    "NUMBER",
	STRING:    "STRING",
	SENTENCE:  "SENTENCE",
	BADSTRING: "BADSTRING",
	BADESCAPE: "BADESCAPE",

	LPAREN:      "(",
	RPAREN:      ")",
	COMMA:       ",",
	COLON:       ":",
	DOUBLECOLON: "::",
	SEMICOLON:   ";",
	DOT:         ".",
	APOSTROPHE:  "'",
	GT:          ">",
	LT:          "<",
	SLASH:       "/",
	STAR:        "*",

	BIND:          "bind",
	PARSE:         "parse",
	EXECUTE:       "execute",
	UPDATE:        "update",
	DURATION:      "duration",
	MS:            "ms",
	APPLICATION:   "application",
	CONTROLLER:    "controller",
	ACTION:        "action",
	LINE:          "line",
	AUTOMATIC:     "automatic",
	REPLICATION:   "replication",
	CHECKPOINT:    "checkpoint",
	CONNECTION:    "connection",
	COULD:         "could",
	DISCONNECTION: "disconnection",
	TEMPORARY:     "temporary",
}

func init() {
	keywords = make(map[string]Token)
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
	}
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return ""
}

// tokstr returns a literal if provided, otherwise returns the token string.
func tokstr(tok Token, lit string) string {
	if lit != "" {
		return lit
	}
	return tok.String()
}

// Lookup returns the token associated with a given string.
func Lookup(ident string) Token {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}

	return IDENT
}

// Pos specifies the line and character position of a token.
// The Char and Line are both zero-based indexes.
type Pos struct {
	Line int
	Char int
}
