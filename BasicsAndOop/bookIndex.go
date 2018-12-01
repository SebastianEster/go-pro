package main

import "fmt"

type Page []string

type Book []Page

type Index map[string][]int

func makeIndex(book Book) Index {
	var index = make(Index)
	for i, page := range book {
		for _, word := range page {
			pages := index[word]
			index[word] = append(pages, i)
		}
	}
	return index
}

func (index Index) String() string {
	retString := ""
	for key, value := range index {
		retString += key + ": " + fmt.Sprint(value) + "\n"
	}
	return retString
}

func main() {
	page1 := Page{"Hello", "Hello", "world"}
	page2 := Page{"Hello", "Peter"}
	page3 := Page{"Hello", "your", "majesty"}
	book := Book{page1, page2, page3}
	index := makeIndex(book)
	fmt.Println(index)
}