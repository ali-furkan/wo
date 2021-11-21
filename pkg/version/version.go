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

func NewVersion(v string) (*Version, error) {
	reSyntax := regexp.MustCompile("^v([0-9]+).([0-9]+).([0-9]+)(-[a-z]+)*?$")
	reCatchNumber := regexp.MustCompile("([0-9]+).([0-9]+).([0-9]+)")
	reCatchName := regexp.MustCompile("(-[a-z]+)")

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

	versionName := strings.Trim(reCatchName.FindString(v), "-")

	version := &Version{
		Major:       major,
		Minor:       minor,
		Patch:       patch,
		VersionName: versionName,
	}

	return version, nil
}

func (v *Version) String() string {
	baseVersion := fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)

	if v.VersionName != "" {
		baseVersion += fmt.Sprintf("-%s", v.VersionName)
	}

	return baseVersion
}

func (v *Version) IsGreaterThan(oVersion Version) bool {
	aVersionNumbers := [3]int{v.Major, v.Minor, v.Patch}
	bVersionNumbers := [3]int{oVersion.Major, oVersion.Minor, oVersion.Patch}

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

func (v *Version) IsEqualTo(oVersion Version, strict bool) bool {
	if strict {
		return v.String() == oVersion.String()
	}

	if v.Major != oVersion.Major {
		return false
	}

	if v.Minor != oVersion.Minor {
		return false
	}

	if v.Patch != oVersion.Patch {
		return false
	}

	return true
}

func CompareVersion(av, bv string) bool {
	aVersion, aErr := NewVersion(av)
	bVersion, bErr := NewVersion(bv)

	if aErr != nil || bErr != nil {
		return false
	}

	return aVersion.IsGreaterThan(*bVersion)
}
