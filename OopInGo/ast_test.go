package OopInGo

import "testing"

func TestAST(t *testing.T) {
	ast := Or{And{Val{"A"}, Val{"B"}}, Val{"C"}}
	values := map[string]bool{
		"A": true,
		"B": false,
		"C": true,
	}

	result := ast.eval(values)
	if result != true {
		t.Error("Expression should return true but returned false!")
	}
}
