package calculator

import (
	"fmt"
	"github.com/timfame/arith-parser.git/tokenizer"
	"github.com/timfame/arith-parser.git/visitor"
)

type Calculator struct {
	parser *visitor.ParserVisitor
	calc   *visitor.CalcVisitor
	print  *visitor.PrintVisitor
}

func NewCalculator() *Calculator {
	return &Calculator{
		parser: visitor.NewParser(),
		calc:   visitor.NewCalc(),
		print:  visitor.NewPrint(),
	}
}

func (c *Calculator) Calculate(expression string) error {
	tokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		return err
	}

	parsedTokens, err := c.parser.ConvertToPolishNotation(tokens)
	if err != nil {
		return err
	}

	if err := c.print.Print(tokens); err != nil {
		return err
	}
	if err := c.print.Print(parsedTokens); err != nil {
		return err
	}

	result, err := c.calc.Calculate(parsedTokens)
	if err != nil {
		return err
	}

	fmt.Println("Result:", result)

	return nil
}
