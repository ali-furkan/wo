// This package will rewrite for bad code in the future
// TODO (IDEA):
//	[] package of funcs maybe will move to struct
// 	[] move work of logical statements to this package
package workspace

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"

	"github.com/MakeNowJust/heredoc"
)

func childRun(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmdOut, err := cmd.Output()
	if err != nil {
		fmt.Println(cmdOut[:])
		return err
	}

	return nil
}

func CreateWork(w Work) error {
	_, err := os.Stat(w.Path)
	if !os.IsNotExist(err) {
		errStr := fmt.Sprintf("%s: folder exists", w.Path)
		return errors.New(errStr)
	}

	err = os.MkdirAll(w.Path, 0755)
	if err != nil {
		return err
	}

	return InitWork(w)
}

func InitWork(w Work) error {
	if w.Template != "" {
		url, err := url.ParseRequestURI(w.Template)
		if err != nil {
			return childRun("git", "clone", url.String(), w.Path)
		}
	}

	if w.InitGit {
		err := childRun("git", "init", w.Path)
		if err != nil {
			return err
		}
	}

	if w.InitReadme {
		err := createDefReadme(w.Name, w.Path)
		if err != nil {
			return err
		}
	}

	if w.License != "" {
		err := createLicense(w.Name, w.License, w.Path)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveWork(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}

	return nil
}

func PrintTinyStat(w Work) {
	stat := heredoc.Docf(`
		Created '%s' work by WO CLI

		Name: %s
		Init Git: %s
		Init Readme: %s
		Path: %s
		License: %s
	`, w.Name, w.Name, fmt.Sprint(w.InitGit), fmt.Sprint(w.InitReadme), w.Path, w.License)

	fmt.Println(stat)
}
