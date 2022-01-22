package main

import (
	"fmt"
	"github.com/timfame/arith-parser.git/calculator"
)

func main() {
	calc := calculator.NewCalculator()
	run(calc, "2 + 3 - 5")
	run(calc, "2 + 3 * 5")
	run(calc, "(2 + 3) * 5")
}

func run(calc *calculator.Calculator, expression string) {
	fmt.Println(expression)

	if err := calc.Calculate(expression); err != nil {
		panic(err)
	}

	fmt.Println()
}
