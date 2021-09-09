package workspace

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type License struct {
	Name       string
	Identifier string
}

const spdxLicenseURI = "https://spdx.org/licenses/"

var Licenses = []License{
	{Name: "Academic Free License 3.0 ", Identifier: "AFL-3.0"},
	{Name: "Adaptive Public License", Identifier: "APL-1.0"},
	{Name: "Apache License 2.0", Identifier: "Apache-2.0"},
	{Name: "Apple Public Source License 2.0 ", Identifier: "APSL-2.0"},
	{Name: "Artistic License 2.0", Identifier: "Artistic-2.0"},
	{Name: "BSD Zero Clause License", Identifier: "0BSD"},
	{Name: "BSD 1-Clause License", Identifier: "BSD-1-Clause"},
	{Name: "BSD 2-Clause License", Identifier: "BSD-2-Clause"},
	{Name: "BSD 3-Clause License", Identifier: "BSD-3-Clause"},
	{Name: "Educational Community License 2.0", Identifier: "ECL-2.0"},
	{Name: "Eclipse Public License 2.0", Identifier: "EPL-2.0"},
	{Name: "GNU Affero General Public License 3.0", Identifier: "AGPL-3.0"},
	{Name: "GNU General Public License 2.0", Identifier: "GPL-2.0-only"},
	{Name: "GNU General Public License 3.0", Identifier: "GPL-3.0-only"},
	{Name: "GNU Lesser General Public License 2.1", Identifier: "LGPL-2.1-only"},
	{Name: "GNU Lesser General Public License 3.0", Identifier: "LGPL-3.0-only"},
	{Name: "IBM Public License 1.0", Identifier: "IPL-1.0"},
	{Name: "ISC License", Identifier: "ISC"},
	{Name: "Microsoft Public License", Identifier: "MS-PL"},
	{Name: "MIT License", Identifier: "MIT"},
	{Name: "Mozilla Public License 2.0", Identifier: "MPL-2.0"},
	{Name: "Open Software License 3.0", Identifier: "OSL-3.0"},
	{Name: "University of Illinois/NCSA Open Source License", Identifier: "	NCSA"},
	{Name: "The Unlicense", Identifier: "Unlicense"},
	{Name: "zLib License", Identifier: "ZLib"},
}

func getLicense(name string) (string, error) {
	for _, license := range Licenses {
		if license.Name == name {
			lisenceURI := fmt.Sprintf("%s/%s.txt", spdxLicenseURI, name)
			resp, err := http.Get(lisenceURI)
			if err != nil {
				return "", err
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}

			return fmt.Sprint(data), nil
		}
	}

	return "", errors.New("license not found")
}

func compileLicense(name string, cnt string) string {
	t := time.Now()
	c := ""

	c = strings.ReplaceAll(cnt, "<name>", name)
	c = strings.ReplaceAll(c, "<year>", strconv.Itoa(t.Year()))

	return c
}
