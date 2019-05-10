package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"

	"github.com/tormath1/emerge-interactive/emerge"
)

func main() {
	defer fmt.Println("Bye !")

	fmt.Println("emerge-interactive v0.0.1")
	fmt.Println("Run `Ctrl+d` or `exit` to quit")
	p := prompt.New(emerge.Executor, emerge.Completer, prompt.OptionPrefix("> "))
	p.Run()
}
