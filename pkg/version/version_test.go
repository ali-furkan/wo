package version

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	name string
	args interface{}
	want interface{}
}

func TestCreateNewVersion(t *testing.T) {
	const (
		major = 10
		minor = 11
		patch = 12
		tag   = "alpha"

		errNotMatchFieldsFormat = "version %s doesn't match the %s: want '%d' -> got '%d'"
	)

	verFormat := fmt.Sprintf("v%d.%d.%d-%s", major, minor, patch, tag)

	ver, err := NewVersion(verFormat)
	if err != nil {
		t.Fatal(err)
	}

	if ver.Major != major {
		t.Fatalf(errNotMatchFieldsFormat, "major", "major", major, ver.Major)
	}

	if ver.Minor != minor {
		t.Fatalf(errNotMatchFieldsFormat, "minor", "minor", minor, ver.Minor)
	}

	if ver.Patch != patch {
		t.Fatalf(errNotMatchFieldsFormat, "patch", "patch", patch, ver.Patch)
	}

	if ver.VersionName != tag {
		t.Fatalf("version tag doesn't match the tag: want '%s' -> got '%s'", tag, ver.VersionName)
	}

	if ver.String() != verFormat {
		t.Fatalf("ver convert string func works wrongly: got %s -> want %s", ver.String(), verFormat)
	}
}

func benchmarkCreateVersion(ver string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := NewVersion(ver)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkCreateVersionWithTag(b *testing.B) {
	benchmarkCreateVersion("v31.15.7-alpha", b)
}

func BenchmarkCreateVersionWithoutTag(b *testing.B) {
	benchmarkCreateVersion("v31.15.7", b)
}

func benchmarkFormatString(ver string, b *testing.B) {
	v, err := NewVersion(ver)
	if err != nil {
		b.Fail()
	}

	for n := 0; n < b.N; n++ {
		_ = v.String()
	}
}

func BenchmarkFormatStringWithTag(b *testing.B) {
	benchmarkFormatString("v31.15.7-alpha", b)
}

func BenchmarkFormatStringWithoutTag(b *testing.B) {
	benchmarkFormatString("v31.15.7", b)
}

func TestIsGreaterThan(t *testing.T) {
	type args struct {
		av string
		bv string
	}

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
			got := CompareVersion(tC.args.(args).av, tC.args.(args).bv)

			if got != tC.want.(bool) {
				t.Errorf("v.IsGreaterThan() = %v, want %v", got, tC.want.(bool))
			}
		})
	}
}

func BenchmarkIsGreaterThan(b *testing.B) {
	type args struct {
		av string
		bv string
	}

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

	for n := 0; n < b.N; n++ {
		tC := testCases[n%len(testCases)]

		got := CompareVersion(tC.args.(args).av, tC.args.(args).bv)

		if got != tC.want.(bool) {
			b.Errorf("Benchmark - v.IsGreaterThan() = %v, want %v", got, tC.want.(bool))
		}
	}
}

func TestIsEqualTo(t *testing.T) {
	type args struct {
		av string
		bv string
	}

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
			aVersion, aErr := NewVersion(tC.args.(args).av)
			bVersion, bErr := NewVersion(tC.args.(args).bv)

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

			if got != tC.want.(bool) {
				t.Errorf("v.IsEqualTo() = %v, want %v", got, tC.want.(bool))
			}
		})
	}
}
