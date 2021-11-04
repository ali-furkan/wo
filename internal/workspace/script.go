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
	ScriptShowFormat = "$ %s"
)

func RunScript(script Script, quiet bool) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return err
	}

	if script.Name == "" {
		script.Name = "unknown"
	}

	color.HiBlack(heredoc.Docf(BeforeRunScriptFormat, script.Name, script.Name))

	strArgs := strings.Join(script.Args, " ")
	strScript := ""

	if script.Workingdir != "" {
		strScript += fmt.Sprintf("cd %s &&", script.Workingdir)
	}

	scripts := strings.Split(strings.TrimSpace(script.Run), "\n")

	for n, childScript := range scripts {
		if !quiet {
			strScript += fmt.Sprintf("echo \"%s\" &&", color.HiBlackString(ScriptShowFormat, childScript))
		}
		strScript += strings.ReplaceAll(childScript, "@args", strArgs)
		if n != len(scripts)-1 {
			strScript += "&&"
		}
	}

	cmd := exec.Command(shell, "-c", strScript)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, script.Env...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
