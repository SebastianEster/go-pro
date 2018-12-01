package ParserCombinators

type ParserInput interface {
	CurrentCodePoint () rune
	RemainingInput () ParserInput
}

type ParserResult struct {
	Result interface{} // nil if parsing fails
	RemainingInput ParserInput
}


type ParserInputImpl struct {
	data []rune
}

func (input ParserInputImpl) CurrentCodePoint() rune {
	return input.data[0]
}

func (input ParserInputImpl) RemainingInput() ParserInput {
	if len(input.data) > 1 {
		return ParserInputImpl{input.data[1:]}
	} else {
		return nil
	}
}

func StringToInput(s string) ParserInput {
	return ParserInputImpl{[]rune(s)}
}