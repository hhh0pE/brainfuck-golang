package main

import (
	"flag"
	"fmt"
)

var storage []int
var brackets []int

func createStorageIfNotExist(index int) {
	if len(storage)==index {
		storage = append(storage, 0)
	}
}

func bracketPush(position int) int {
	brackets = append(brackets, position)
	return len(brackets)-1
}

func bracketPop() {
	brackets = brackets[:len(brackets)-1]
}

func main() {
	input := flag.String("code", "", "code of programm")
	flag.Parse()

	var i, bi int
	storage = make([]int, 1)
	brackets = make([]int, 1)

	for si:=0; si<len(*input); si++ {
		switch (*input)[si] {
		case '+':
			storage[i]++
		case '-':
			storage[i]--
		case '.':
			fmt.Print(string(storage[i]))
		case ',':
			fmt.Scanf("%d", storage[i])
		case '[':
			bi = bracketPush(si)
		case ']':
			if storage[i]>0 {
				si = brackets[bi]
			} else {
				bracketPop()
				bi--
			}
		case '>':
			i++
			createStorageIfNotExist(i)
		case '<':
			i--
		}
	}
}