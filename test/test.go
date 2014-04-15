package main

import (
	"fmt"
	"github.com/elsonwu/random"
)

func main() {
	fmt.Println("char: ", random.Char())
	fmt.Println("letter: ", random.Letter())
	fmt.Println("letters: ", random.Letters(10))
	fmt.Println("String: ", random.String(10))
	fmt.Println("int: ", random.Int(0, 100))
	fmt.Println("symbol: ", random.Symbol())
	fmt.Println("symbols: ", random.Symbols(3))
}
