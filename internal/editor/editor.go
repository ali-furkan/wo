package editor

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

const ErrNotFoundEditor = "editor not found on machine"

func Scan() ([]Editor, error) {
	var s int
	e := []Editor{}

	switch runtime.GOOS {
	case "windows":
		s = 0
	case "darwin":
		s = 1
	default:
		s = 2
	}

	for name, execEditor := range editors {
		execPath := execEditor[0]
		if len(execEditor) > 1 {
			execPath = execEditor[s]
		}
		path, err := exec.LookPath(execPath)
		if err != nil {
			continue
		}
		e = append(e, Editor{Name: name, Exec: path})
	}

	if len(e) == 0 {
		return e, errors.New(ErrNotFoundEditor)
	}

	return e, nil
}

func OpenEditor(e Editor, path string) error {
	cmd := exec.Command(e.Exec, path)

	data, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Printf("Running: %s", string(data))

	return nil
}
