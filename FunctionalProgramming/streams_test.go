package FunctionalProgramming

import (
	"fmt"
	"strings"
	"testing"
)

func TestStreams(t *testing.T) {
	myStream := []Any{"Hello", "your", "royal", "majesty"}
	mapper := func(element Any) Any {
		return strings.ToUpper(element.(string))
	}
	predicate := func(element Any) bool {
		return element.(string) != "ROYAL"
	}
	accumulator := func(elementOne, elementTwo Any) Any {
		return elementOne.(string) + " " + elementTwo.(string)
	}
	result := ToStream(myStream).Map(mapper).Filter(predicate).Reduce(accumulator)

	if result != "HELLO YOUR MAJESTY" {
		t.Errorf("Result sould be \"HELLO YOUR MAJESTY\", but was %s", result)
	}
}

func TestWordCount(t *testing.T) {
	mySlice := []Any{"A", "B", "C", "A", "B", "A"}
	mapper := func(element Any) Any {
		return []Pair{{element, 1}}
	}
	accumulator := func(elementOne, elementTwo Any) Any {
		elementOneContainsElementTwo := false
		for index, element := range elementOne.([]Pair) {
			if element.first == elementTwo.([]Pair)[0].first { // element two always contains only one entry
				elementOne.([]Pair)[index].second = elementOne.([]Pair)[index].second.(int) + 1
				elementOneContainsElementTwo = true
				break
			}
		}
		if !elementOneContainsElementTwo {
			elementOne = append(elementOne.([]Pair), elementTwo.([]Pair)[0])
		}
		return elementOne
	}

	result := ToStream(mySlice).Map(mapper).Reduce(accumulator)
	fmt.Printf("%v", result)
}