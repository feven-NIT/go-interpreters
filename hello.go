package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		runPrompt()
	} else {
		runFile(os.Args[1])
	}
}

func runPrompt() {
	fmt.Println("<")

}

func runFile(path string) {
	fmt.Println(path)
}
