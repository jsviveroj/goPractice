package main

import (
	"fmt"
	"goPractice/src/variables"
)

func main() {
	state, text := variables.ConvertToInt(5)
	fmt.Println(state)
	fmt.Println(text)
}
