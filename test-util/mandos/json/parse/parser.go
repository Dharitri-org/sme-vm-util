package mandosjsonparse

import (
	fr "github.com/Dharitri-org/sme-vm-util/test-util/mandos/json/fileresolver"
	vi "github.com/Dharitri-org/sme-vm-util/test-util/mandos/json/valueinterpreter"
)

// Parser performs parsing of both json tests (older) and scenarios (new).
type Parser struct {
	ValueInterpreter vi.ValueInterpreter
}

// NewParser provides a new Parser instance.
func NewParser(fileResolver fr.FileResolver) Parser {
	return Parser{
		ValueInterpreter: vi.ValueInterpreter{
			FileResolver: fileResolver,
		},
	}
}
