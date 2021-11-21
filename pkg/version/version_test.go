package version

import (
	"strings"
	"testing"
)

type args struct {
	av string
	bv string
}

type testCase struct {
	name string
	args args
	want bool
}

func TestCreateNewVersion(t *testing.T) {
}

// func TestBenchmarkCreateVersion() {

// }

// func TestBenchmarkFormatString() {

// }

func TestIsGreaterThan(t *testing.T) {
	testCases := []testCase{
		{"It should return false when it was got equal versions (with tags)", args{"v1.0.0-beta", "v1.0.0"}, false},
		{"It should return false when it was got equal versions (without tags) ", args{"v1.0.0", "v1.0.0"}, false},
		{"A should be greater than B version in major", args{"v2.0.0", "v1.0.0"}, true},
		{"A should be greater than B version in minor", args{"v1.1.0", "v1.0.0"}, true},
		{"A should be greater than B version in patch", args{"v1.0.1", "v1.0.0"}, true},
		{"A should be less than B version in major", args{"v1.0.0", "v2.0.0"}, false},
		{"A should be less than B version in minor", args{"v1.0.0", "v1.1.0"}, false},
		{"A should be less than B version in patch", args{"v1.0.0", "v1.0.1"}, false},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// CompareVersion does the same job as v.IsGreaterThan
			got := CompareVersion(tC.args.av, tC.args.bv)

			if got != tC.want {
				t.Errorf("v.IsGreaterThan() = %v, want %v", got, tC.want)
			}
		})
	}
}

// func TestBenchmarkIsGreaterThan() {

// }

func TestIsEqualTo(t *testing.T) {
	testCases := []testCase{
		{"It should return true when it was got equal versions but names are different (without strict)", args{"v1.0.0-beta", "v1.0.0"}, true},
		{"It should return false when it was got equal versions but names are different (with strict)", args{"v1.0.0-beta", "v1.0.0"}, false},
		{"It should return true when it was got equal versions (wiht strict)", args{"v1.0.0-abc", "v1.0.0-abc"}, true},
		{"It should return false when A version greater than B in major", args{"v2.0.0", "v1.0.0"}, false},
		{"It should return false when A version greater than B in minor", args{"v1.1.0", "v1.0.0"}, false},
		{"It should return false when A version greater than B in patch", args{"v1.0.1", "v1.0.0"}, false},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			aVersion, aErr := NewVersion(tC.args.av)
			bVersion, bErr := NewVersion(tC.args.bv)

			t.Logf("Versions: %s - %s\n", aVersion.String(), bVersion.String())

			if aErr != nil || bErr != nil {
				t.Errorf("Create NewVersion Error: %s %s", aErr.Error(), bErr.Error())
				return
			}

			isStrict := false
			if strings.HasSuffix(tC.name, "(with strict)") {
				isStrict = true
			}
			got := aVersion.IsEqualTo(*bVersion, isStrict)

			if got != tC.want {
				t.Errorf("v.IsEqualTo() = %v, want %v", got, tC.want)
			}
		})
	}
}
