package parser

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("PLATFORM_ENV", "test")
}

func TestParserParseStatement(t *testing.T) {
	var tests = []struct {
		s   string
		obj Statement
		p   string
		err string
	}{
		// Single field statement
		{
			s: `duration: 0.051 ms  execute <unnamed>: select * from servers where id IN ('1', '2', '3') and name = 'localhost'`,
			obj: &QueryStatement{
				Duration:     "0.051",
				PreparedStep: EXECUTE,
				PreparedName: "unnamed",
				Query:        `select * from servers where id IN ('1', '2', '3') and name = 'localhost'`},
			p: `DURATION`,
		},
	}

	for _, tt := range tests {
		obj, err := NewParser(strings.NewReader(tt.s)).ParseStatement()
		assert.NoError(t, err)
		assert.Equal(t, tt.obj, obj)
		assert.Equal(t, tt.p, tt.obj.String())
	}
}
