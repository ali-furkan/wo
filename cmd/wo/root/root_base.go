package root

import "github.com/MakeNowJust/heredoc"

const (
	// Base fields for Root Cmd
	CmdUsage     = "wo <command> [subcommand] [flags]"
	CmdShortDesc = "WO CLI"
	CmdLongDesc  = "Easily manage your work via command line"
)

var (
	CmdExample = heredoc.Doc(`
		$ wo init hello-world
		$ wo open hello-world
		$ wo remove -r hello-world
	`)
)
