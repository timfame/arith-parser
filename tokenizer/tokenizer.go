package tokenizer

func Tokenize(expr string) (tokens Tokens, err error) {
	var current state = &arithmeticState{}
	for i := 0; i < len(expr); i++ {
		current, err = current.checkRune(rune(expr[i]), &tokens)
		if err != nil {
			return nil, err
		}
	}
	if err := current.checkEnd(&tokens); err != nil {
		return nil, err
	}
	return
}
