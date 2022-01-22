package visitor

import (
	"fmt"
	"github.com/timfame/arith-parser.git/tokenizer"
)

type ParserVisitor struct {
	*baseVisitor
	numberStack tokenizer.Tokens
	symbolStack tokenizer.Tokens
}

func NewParser() *ParserVisitor {
	b := &baseVisitor{}
	r := &ParserVisitor{
		baseVisitor: b,
	}
	b.Visitor = r
	return r
}

func (p *ParserVisitor) ConvertToPolishNotation(tokens tokenizer.Tokens) (tokenizer.Tokens, error) {
	p.numberStack = make([]tokenizer.Token, 0, len(tokens))
	p.symbolStack = make([]tokenizer.Token, 0, len(tokens))
	for _, token := range tokens {
		if err := p.Visit(token); err != nil {
			return nil, err
		}
	}
	for len(p.symbolStack) > 0 {
		p.numberStack = append(p.numberStack, p.symbolStack[len(p.symbolStack) - 1])
		p.symbolStack = p.symbolStack[:len(p.symbolStack) - 1]
	}
	return p.numberStack, nil
}

func (p *ParserVisitor) visitOperation(op tokenizer.Operation) error {
	switch op {
	case tokenizer.Plus, tokenizer.Minus:
		if d := len(p.symbolStack); d > 0 {
			if l := p.symbolStack[d - 1]; l != tokenizer.Open {
				p.numberStack = append(p.numberStack, l)
				p.symbolStack = p.symbolStack[:d - 1]
			}
		}
	case tokenizer.Mul, tokenizer.Div:
		if d := len(p.symbolStack); d > 0 {
			if l := p.symbolStack[d - 1]; l == tokenizer.Mul || l == tokenizer.Div {
				p.numberStack = append(p.numberStack, l)
				p.symbolStack = p.symbolStack[:d - 1]
			}
		}
	}

	p.symbolStack = append(p.symbolStack, op)

	return nil
}

func (p *ParserVisitor) visitNumber(n tokenizer.Number) error {
	p.numberStack = append(p.numberStack, n)
	return nil
}

func (p *ParserVisitor) visitBracket(b tokenizer.Bracket) error {
	if b == tokenizer.Open {
		p.symbolStack = append(p.symbolStack, b)
		return nil
	}
	for len(p.symbolStack) > 0 && p.symbolStack[len(p.symbolStack) - 1] != tokenizer.Open {
		p.numberStack = append(p.numberStack, p.symbolStack[len(p.symbolStack) - 1])
		p.symbolStack = p.symbolStack[:len(p.symbolStack) - 1]
	}
	if len(p.symbolStack) == 0 {
		return fmt.Errorf("bad expression, left bracket is missed")
	}
	p.symbolStack = p.symbolStack[:len(p.symbolStack) - 1]
	return nil
}
