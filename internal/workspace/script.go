package workspace

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/fatih/color"
	"github.com/riywo/loginshell"
)

const (
	BeforeRunScriptFormat = `
		Running %s script
		$ wo run %s
	`
	ScriptShowFormat = "$ %s\n"
)

func RunScript(script Script) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return err
	}

	color.HiBlack(heredoc.Docf(BeforeRunScriptFormat, script.Name))

	for _, childScript := range strings.Split(script.Run, "\n") {
		if strings.TrimSpace(childScript) == "" {
			continue
		}

		if script.Workingdir != "" {
			childScript = fmt.Sprintf("cd %s && %s", script.Workingdir, childScript)
		}

		color.HiBlack(ScriptShowFormat, childScript)

		cmd := exec.Command(shell, "-c", childScript, strings.Join(script.Args, " "))

		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, script.Env...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
