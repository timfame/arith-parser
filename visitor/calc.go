package visitor

import (
	"fmt"
	"github.com/timfame/arith-parser.git/tokenizer"
)

type CalcVisitor struct {
	*baseVisitor
	stack []int
}

func NewCalc() *CalcVisitor {
	b := &baseVisitor{}
	r := &CalcVisitor{
		baseVisitor: b,
	}
	b.Visitor = r
	return r
}

func (c *CalcVisitor) Calculate(tokens tokenizer.Tokens) (int, error) {
	c.stack = nil
	for _, token := range tokens {
		if err := c.Visit(token); err != nil {
			return 0, err
		}
	}
	if len(c.stack) != 1 {
		return 0, fmt.Errorf("wrong reverse polish notation")
	}
	return c.stack[0], nil
}

func (c *CalcVisitor) visitOperation(op tokenizer.Operation) error {
	d := len(c.stack)
	if d < 2 {
		return fmt.Errorf("wrong reverse polish notation")
	}
	x, y := c.stack[d - 2], c.stack[d - 1]
	c.stack = c.stack[:d - 2]
	switch op {
	case tokenizer.Plus:
		c.stack = append(c.stack, x + y)
	case tokenizer.Minus:
		c.stack = append(c.stack, x - y)
	case tokenizer.Mul:
		c.stack = append(c.stack, x * y)
	case tokenizer.Div:
		c.stack = append(c.stack, x / y)
	default:
		return fmt.Errorf("wrong operation token: %v", op)
	}
	return nil
}

func (c *CalcVisitor) visitNumber(n tokenizer.Number) error {
	c.stack = append(c.stack, n.GetValue())
	return nil
}

func (c *CalcVisitor) visitBracket(b tokenizer.Bracket) error {
	return fmt.Errorf("wrong token in reverse polish notation: %v", b)
}
