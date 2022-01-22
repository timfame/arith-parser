package tokenizer

import "fmt"

type state interface {
	checkRune(r rune, tokens *Tokens) (state, error)
	checkEnd(tokens *Tokens) error
}

type arithmeticState struct {}

func (as *arithmeticState) checkRune(r rune, tokens *Tokens) (state, error) {
	switch r {
	case '+':
		*tokens = append(*tokens, Plus)
	case '-':
		*tokens = append(*tokens, Minus)
	case '*':
		*tokens = append(*tokens, Mul)
	case '/':
		*tokens = append(*tokens, Div)
	case '(':
		*tokens = append(*tokens, Open)
	case ')':
		*tokens = append(*tokens, Close)
	default:
		if '0' <= r && r <= '9' {
			newState := &numberState{current: 0}
			return newState.checkRune(r, tokens)
		}
		if r != ' ' && r != '\t' && r != '\n' {
			return nil, fmt.Errorf("bad expression with \"%c\"", r)
		}
	}
	return as, nil
}

func (as *arithmeticState) checkEnd(tokens *Tokens) error {
	return nil
}

type numberState struct {
	current int
}

func (ns *numberState) checkRune(r rune, tokens *Tokens) (state, error) {
	if '0' <= r && r <= '9' {
		ns.current *= 10
		ns.current += int(r - '0')
		return ns, nil
	}
	*tokens = append(*tokens, Number{x: ns.current})
	newState := &arithmeticState{}
	return newState.checkRune(r, tokens)
}

func (ns *numberState) checkEnd(tokens *Tokens) error {
	*tokens = append(*tokens, Number{x: ns.current})
	return nil
}
