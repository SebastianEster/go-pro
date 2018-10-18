package OopInGo

type Node interface {
	eval(values map[string]bool) bool
}

type Or struct {
	leftChild Node
	rightChild Node
}

type And struct {
	leftChild Node
	rightChild Node
}

type Val struct {
	id string
}

func (or Or) eval(values map[string]bool) bool {
	return or.leftChild.eval(values) || or.rightChild.eval(values)
}

func (and And) eval(values map[string]bool) bool {
	return and.leftChild.eval(values) && and.rightChild.eval(values)
}

func (val Val) eval(values map[string]bool) bool {
	return values[val.id]
}