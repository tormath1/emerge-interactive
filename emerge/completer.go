package emerge

import "github.com/c-bata/go-prompt"

// Completer will provide suggestion in order to
// help the user when he's building his command
func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--info", Description: "Produces a list of information to include in bug reports which aids the  developers  when  fixing  the  reported  problem."},
		{Text: "--pretend", Description: "Instead of actually performing the merge, simply display what *would* have been installed if --pretend weren't used."},
		{Text: "-p", Description: "Instead of actually performing the merge, simply display what *would* have been installed if --pretend weren't used."},
		{Text: "--verbose", Description: "Tell  emerge  to  run in verbose mode."},
		{Text: "-v", Description: "Tell  emerge  to  run in verbose mode."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
