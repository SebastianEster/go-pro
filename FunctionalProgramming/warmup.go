package FunctionalProgramming

var myFunc = func (i int) int {
	return i*i
}

func apply (function func(interface{}) interface{}, parameter interface{}) interface{} {
	return function(parameter)
}

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}