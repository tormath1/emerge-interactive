package emerge

import "github.com/c-bata/go-prompt"

// Completer will provide suggestion in order to
// help the user when he's building his command
func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--sync", Description: "Updates repositories, for which auto-sync, sync-type and sync-uri attributes are set in repos.con."},
		{Text: "--info", Description: "Produces a list of information to include in bug reports which aids the  developers  when  fixing  the  reported  problem."},
		{Text: "--search", Description: "Searches  for matches of the supplied strin  repository."},
		{Text: "--help", Description: "Displays help information for emerge."},
		{Text: "--version", Description: "Displays the version number of emerge."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
