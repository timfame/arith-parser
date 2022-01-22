package tokenizer

import "fmt"

type Token interface {
	String() string
	//Accept()
}

type Tokens []Token

type Operation string

func (p Operation) String() string {
	return string(p)
}

var (
	Plus  Operation = "PLUS"
	Minus Operation = "MINUS"
	Mul   Operation = "MUL"
	Div   Operation = "DIV"
)

type Bracket string

func (b Bracket) String() string {
	return string(b)
}

var (
	Open  Bracket = "OPEN"
	Close Bracket = "CLOSE"
)

type Number struct {
	x int
}

func (n Number) String() string {
	return fmt.Sprintf("NUMBER(%d)", n.x)
}

func (n Number) GetValue() int {
	return n.x
}
