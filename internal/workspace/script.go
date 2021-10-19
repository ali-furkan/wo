package workspace

import (
	"os"
	"os/exec"
	"strings"

	"github.com/riywo/loginshell"
)

const (
	RunScriptFormat = `
		$ wo run %s
		%s
		%s

		%s
	`
)

func RunScript(script Script) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return err
	}

	cmd := exec.Command(shell, "-c", script.Run, strings.Join(script.Args, " "))

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, script.Env...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
