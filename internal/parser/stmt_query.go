package parser

import (
	"bytes"
)

//******
// QUERY
//******

// QueryStatement allows you to converse with others in the same room
type QueryStatement struct {
	Duration     string
	PreparedStep Token
	PreparedName string
	Query        string
}

// TODO: Finish writing options
func (s *QueryStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("DURATION")

	return buf.String()
}

//******
// Example: duration: 0.051 ms  execute <unnamed>: select * from servers where id IN ('1', '2', '3') and name = 'localhost'
//******

// parseQueryStatement parses a query and returns a Statement AST object.
// This function assumes the DURATION token has already been consumed.
func (p *Parser) parseQueryStatement() (*QueryStatement, error) {
	stmt := &QueryStatement{}

	// colon
	tok, pos, lit := p.Scan()
	if tok != COLON {
		return nil, newParseError(tokstr(tok, lit), []string{":"}, pos)
	}

	tok, _, lit = p.ScanIgnoreWhitespace()
	if tok != NUMBER {
		return nil, newParseError(tokstr(tok, lit), []string{"NUMBER"}, pos)
	}
	stmt.Duration = lit

	tok, _, _ = p.ScanIgnoreWhitespace()
	if tok != MS {
		return nil, newParseError(tokstr(tok, lit), []string{"MS"}, pos)
	}

	tok, _, lit = p.ScanIgnoreWhitespace()
	stmt.PreparedStep = tok

	// Loop over the Prepared Name and return the identity.
	// Sometimes they are surrounded by <>, like <unnamed>, but ignore < & >.
preparedLoop:
	for {
		tok, _, lit = p.ScanIgnoreWhitespace()
		switch tok {
		case IDENT:
			stmt.PreparedName = lit
		case COLON:
			break preparedLoop
		case EOF:
			return stmt, nil
		}
	}

	// Go past the space after the colon
	_, _, _ = p.Scan()

	// The rest is the query
	_, _, stmt.Query = p.ScanSentence()

	return stmt, nil
}
