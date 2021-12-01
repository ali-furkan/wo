package space

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/MakeNowJust/heredoc"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var WorkspaceNameValidationRules = []validation.Rule{validation.Required, validation.Length(2, 64), is.PrintableASCII}

func childRun(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmdOut, err := cmd.Output()
	if err != nil {
		fmt.Println(cmdOut[:])
		return err
	}

	return nil
}

func CreateWorkspace(w Workspace, opts Options) error {
	_, err := os.Stat(w.Path)
	if !os.IsNotExist(err) {
		errStr := fmt.Sprintf("%s: folder exists", w.Path)
		return errors.New(errStr)
	}

	err = os.MkdirAll(w.Path, 0755)
	if err != nil {
		return err
	}

	return InitWorkspace(w, opts)
}

func InitWorkspace(w Workspace, opts Options) error {
	_, err := os.Stat(w.Path)
	if os.IsNotExist(err) {
		errStr := fmt.Sprintf("%s: folder doesn't exists", w.Path)
		return errors.New(errStr)
	}

	if opts.Git == "enabled" {
		err := childRun("git", "init", w.Path)
		if err != nil {
			return err
		}
	}

	if opts.Readme == "enabled" {
		err := createDefReadme(w.Name, w.Path)
		if err != nil {
			return err
		}
	}

	if opts.License != "" {
		err := createLicense(w.Name, opts.License, w.Path)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveWorkspace(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}

	return nil
}

func PrintTinyStat(w Workspace) {
	stat := heredoc.Docf(`
		Success, Created '%s' workspace at the '%s' by Wo

		Name: %s
		Path: %s
	`, w.Name, w.Name)

	fmt.Println(stat)
}

func MoveWorkspace(w map[string]string, newpath string) error {
	p := newpath
	if !filepath.IsAbs(newpath) {
		d, err := os.Getwd()
		if err != nil {
			return err
		}

		p = filepath.Join(d, newpath)
	}

	p = filepath.Clean(p)

	return os.Rename(w["path"], p)
}
