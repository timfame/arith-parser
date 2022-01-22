package visitor

import (
	"fmt"
	"github.com/timfame/arith-parser.git/tokenizer"
)

type PrintVisitor struct {
	*baseVisitor
}

func NewPrint() *PrintVisitor {
	b := &baseVisitor{}
	r := &PrintVisitor{
		baseVisitor: b,
	}
	b.Visitor = r
	return r
}

func (p *PrintVisitor) Print(tokens tokenizer.Tokens) error {
	for _, token := range tokens {
		if err := p.Visit(token); err != nil {
			return err
		}
	}
	fmt.Println()
	return nil
}

func (p *PrintVisitor) visitOperation(op tokenizer.Operation) error {
	fmt.Print(op.String())
	fmt.Print(" ")
	return nil
}

func (p *PrintVisitor) visitNumber(n tokenizer.Number) error {
	fmt.Print(n.String())
	fmt.Print(" ")
	return nil
}

func (p *PrintVisitor) visitBracket(b tokenizer.Bracket) error {
	fmt.Print(b.String())
	fmt.Print(" ")
	return nil
}
