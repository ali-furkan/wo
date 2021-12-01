package auth

import (
	"fmt"
	"net/http"

	"github.com/cli/oauth/api"
	"github.com/cli/oauth/device"
	"github.com/pkg/browser"
)

const (
	ClientID  = "826db533af5eefa85c71"
	DeviceURI = "https://github.com/login/device/code"
	PollURL   = "https://github.com/login/oauth/access_token"
)

type GitHub struct {
	HttpClient *http.Client
	Scopes     []string
}

func (gh *GitHub) RequestCodeAndPollToken() (*api.AccessToken, error) {
	code, err := device.RequestCode(gh.HttpClient, DeviceURI, ClientID, gh.Scopes)

	if err != nil {
		return nil, err
	}

	fmt.Printf("First copy your one-time code: %s\n", code.UserCode)
	fmt.Printf("then open: %s\n", code.VerificationURI)
	err = browser.OpenURL(code.VerificationURI)
	if err != nil {
		return nil, err
	}

	accessToken, err := device.PollToken(gh.HttpClient, PollURL, ClientID, code)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
