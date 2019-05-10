package emerge

import (
	"os"
	"os/exec"

	"fmt"
)

// Executor will pass the command
// given in parameter to emerge
// and redirect output to Std{in,out,err}
func Executor(s string) {
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
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
