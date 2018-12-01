package FunctionalProgramming

type Any interface {}
type unaryFunction func(Any) Any

func compose(f, g unaryFunction) unaryFunction {
	return func(x Any) Any {
		return f(g(x))
	}
}