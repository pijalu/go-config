package parser

import "github.com/pijalu/go-config/changeset"

// Store actual call to parserz
type MockParserCall struct {
	Source string
	Data   interface{}
}

// MockParser is a mocked parsers
type MockParser struct {
	ParseCalls []MockParserCall     // ParseCalls saves the call lists
	RError     error                // RError defines returned errors
	RChangeSet *changeset.ChangeSet // RChangeSet defines the returned changeset
}

func (m *MockParser) Parse(source string, i interface{}) (*changeset.ChangeSet, error) {
	m.ParseCalls = append(m.ParseCalls, MockParserCall{
		Source: source,
		Data:   i,
	})

	return m.RChangeSet, m.RError
}
