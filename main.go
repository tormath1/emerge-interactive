package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"

	"os"
	"os/exec"
	"strings"
)

func executor(s string) {
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	cmd := exec.Command("/bin/sh", "-c", "emerge "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("unable to run command: %v\n", err)
	}
	return
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--info", Description: "Produces a list of information to include in bug reports which aids the  developers  when  fixing  the  reported  problem."},
		{Text: "--pretend", Description: "Instead of actually performing the merge, simply display what *would* have been installed if --pretend weren't used."},
		{Text: "-p", Description: "Instead of actually performing the merge, simply display what *would* have been installed if --pretend weren't used."},
		{Text: "--verbose", Description: "Tell  emerge  to  run in verbose mode."},
		{Text: "-v", Description: "Tell  emerge  to  run in verbose mode."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("emerge-interactive v0.0.1")
	fmt.Println("Run `Ctrl+d` or `exit` to quit")
	p := prompt.New(executor, completer, prompt.OptionPrefix("> "))
	p.Run()
}
