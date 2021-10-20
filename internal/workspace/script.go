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

	if script.Name == "" {
		script.Name = "unknown"
	}

	color.HiBlack(heredoc.Docf(BeforeRunScriptFormat, script.Name, script.Name))

	strArgs := strings.Join(script.Args, " ")
	env := []string{}

	for _, childScript := range strings.Split(script.Run, "\n") {
		if strings.TrimSpace(childScript) == "" {
			continue
		}

		if script.Workingdir != "" {
			childScript = fmt.Sprintf("cd %s && %s", script.Workingdir, childScript)
		}

		color.HiBlack(ScriptShowFormat, childScript)

		cmd := exec.Command(shell, "-c", strings.ReplaceAll(childScript, "@args", strArgs))

		if len(env) > 0 {
			cmd.Env = env
		} else {
			cmd.Env = os.Environ()
		}
		cmd.Env = append(cmd.Env, script.Env...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}

		env = cmd.Env
	}

	return nil
}
