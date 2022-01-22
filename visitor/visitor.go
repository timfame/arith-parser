package visitor

import (
	"fmt"
	"github.com/timfame/arith-parser.git/tokenizer"
)

type Visitor interface {
	// Visit ... in Go we cannot declare multiple functions with same name,
	// so I have one common method for all tokens,
	// in which I check with type-switch what token passed as argument
	Visit(token tokenizer.Token) error
	visitOperation(op tokenizer.Operation) error
	visitNumber(n tokenizer.Number) error
	visitBracket(b tokenizer.Bracket) error
}

type baseVisitor struct {
	Visitor
}

func (b *baseVisitor) Visit(token tokenizer.Token) error {
	switch v := token.(type) {
	case tokenizer.Operation:
		return b.visitOperation(v)
	case tokenizer.Number:
		return b.visitNumber(v)
	case tokenizer.Bracket:
		return b.visitBracket(v)
	}
	return fmt.Errorf("unsupported token type: %T", token)
}
