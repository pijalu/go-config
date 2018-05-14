package parser

import "github.com/pijalu/go-config/source"

// Store actual call to parserz
type MockParserCall struct {
	Source string
	Data   interface{}
}

// MockParser is a mocked parsers
type MockParser struct {
	ParseCalls []MockParserCall  // ParseCalls saves the call lists
	RError     error             // RError defines returned errors
	RChangeSet *source.ChangeSet // RChangeSet defines the returned changeset
}

func (m *MockParser) Parse(source string, i interface{}) (*source.ChangeSet, error) {
	m.ParseCalls = append(m.ParseCalls, MockParserCall{
		Source: source,
		Data:   i,
	})

	return m.RChangeSet, m.RError
}
