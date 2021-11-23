package version

import (
	"fmt"

	"github.com/ali-furkan/wo/pkg/version"
)

const (
	CurVersionStr  = "0.0.1"
	CurVersionName = "alpha"
)

var (
	CurVersionFormat = fmt.Sprintf("v%s-%s", CurVersionStr, CurVersionName)
	CurVersion, _    = version.NewVersion(CurVersionFormat)
)
