package ParserCombinators

func parseA(input ParserInput) ParserResult {
	if input.CurrentCodePoint() == 'A' {
		return ParserResult{'A', input.RemainingInput()}
	} else {
		return ParserResult{nil, input.RemainingInput()}
	}
}
