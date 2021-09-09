package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ali-furkan/wo/internal/version"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

const (
	WoReleaseURI       = "https://api.github.com/repos/ali-furkan/wo/releases/latest"
	ErrReleaseNotFound = "release not found"
)

type ReleaseInfo struct {
	Version     string `json:"tag_name"`
	URL         string `json:"url"`
	InfoURL     string `json:"html_url"`
	PublishedAt string `json:"published_at"`
}

type UpdateState struct {
	LastCheckedAt time.Time   `yaml:"last_checked_at"`
	LatestRelease ReleaseInfo `yaml:"latest_release"`
}

func CheckForUpdate() (*ReleaseInfo, error) {
	statePath, err := getStatePath()
	if err != nil {
		return nil, err
	}

	state, err := getUpdateState(statePath)
	if err != nil {
		return nil, err
	}

	if state != nil && time.Since(state.LastCheckedAt).Hours() < 6 {
		return &state.LatestRelease, nil
	}

	releaseInfo, err := getLatestReleaseInfo()
	if err != nil {
		if err.Error() == ErrReleaseNotFound {
			return nil, nil
		}
		return nil, err
	}

	err = setUpdateState(statePath, time.Now(), releaseInfo)
	if err != nil {
		return nil, err
	}

	if version.IsGreaterThan(releaseInfo.Version, version.GetVersion()) {
		return &releaseInfo, nil
	}

	return nil, nil
}

func Update() error {
	release, err := CheckForUpdate()
	if err != nil {
		return err
	}

	if release == nil {
		fmt.Println("Already up to date!")
		return nil
	}

	uri := fmt.Sprintf("%s/wo_%s_%s", release.URL, runtime.GOOS, runtime.GOARCH)
	res, err := http.Get(uri)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(exePath, data, 0222)
	if err != nil {
		return err
	}

	statePath, err := getStatePath()
	if err != nil {
		return err
	}

	return setUpdateState(statePath, time.Now(), *release)
}

func getStatePath() (string, error) {
	woDir, err := homedir.Expand("~/.wo/")
	if err != nil {
		return "", err
	}

	statePath := filepath.Join(woDir, "state.yml")

	return statePath, nil
}

func getLatestReleaseInfo() (latestRelease ReleaseInfo, err error) {
	res, err := http.Get(WoReleaseURI)
	if err != nil {
		err = errors.New("fetch repo failure: please check your internet connection")
		return
	}

	if res.StatusCode == 404 {
		err = errors.New(ErrReleaseNotFound)
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &latestRelease)
	if err != nil {
		return
	}

	return
}

func getUpdateState(statePath string) (*UpdateState, error) {
	file, err := os.OpenFile(statePath, os.O_RDONLY, 0444)
	if err != nil {
		file, err = os.Create(statePath)

		if err != nil {
			return nil, err
		}
	}

	defer file.Close()

	state := &UpdateState{}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(state)

	if err != nil && err != io.EOF {
		return nil, err
	}

	return state, nil
}

func setUpdateState(statePath string, t time.Time, r ReleaseInfo) error {
	state := UpdateState{LastCheckedAt: t, LatestRelease: r}
	data, err := yaml.Marshal(state)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(statePath, data, 0600)

	return err
}
