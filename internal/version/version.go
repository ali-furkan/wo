package version

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int

	VersionName string
}

func GetVersion() string {
	if CurVersion == "" {
		return "unknown version"
	}

	version := fmt.Sprintf("v%s", CurVersion)

	if CurVersionName != "" {
		version = fmt.Sprintf("v%s-%s", CurVersion, CurVersionName)
	}

	return version
}

func (v *Version) String() string {
	if CurVersion == "" {
		return "unknown version"
	}

	version := fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)

	if CurVersionName != "" {
		version = fmt.Sprintf("v%s-%s", CurVersion, CurVersionName)
	}

	return version
}

func NewVersion(v string) (*Version, error) {
	reSyntax := regexp.MustCompile("^v([0-9]+).([0-9]+).([0-9]+)(-[a-z]+)*?$")
	reCatchNumber := regexp.MustCompile("([0-9]+).([0-9]+).([0-9]+)")
	reCatchName := regexp.MustCompile("([a-z]+)")

	match := reSyntax.MatchString(v)
	if !match {
		return nil, errors.New("v is not version")
	}

	versionCatchNumbers := strings.Split(reCatchNumber.FindString(v), ".")

	major, majorErr := strconv.Atoi(versionCatchNumbers[0])
	minor, minorErr := strconv.Atoi(versionCatchNumbers[1])
	patch, patchErr := strconv.Atoi(versionCatchNumbers[2])

	if majorErr != nil || minorErr != nil || patchErr != nil {
		return nil, errors.New("v could not parse")
	}

	versionName := reCatchName.FindString(v)

	version := &Version{
		Major:       major,
		Minor:       minor,
		Patch:       patch,
		VersionName: versionName,
	}

	return version, nil
}

func IsGreaterThan(av, bv string) bool {

	aVersion, aErr := NewVersion(av)
	bVersion, bErr := NewVersion(bv)

	if aErr != nil || bErr != nil {
		return false
	}

	aVersionNumbers := [3]int{aVersion.Major, aVersion.Minor, aVersion.Patch}
	bVersionNumbers := [3]int{bVersion.Major, bVersion.Minor, bVersion.Patch}

	for i := 0; i < 3; i++ {
		if aVersionNumbers[i] > bVersionNumbers[i] {
			return true
		}

		if aVersionNumbers[i] < bVersionNumbers[i] {
			return false
		}
	}

	return false
}
