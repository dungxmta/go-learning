package version

import (
	"fmt"
	"testing"
)

func TestSemVer(t *testing.T) {
	var str = `Valid Semantic Versions

0.0.4
1.2.3
10.20.30
1.1.2-prerelease+meta
1.1.2+meta
1.1.2+meta-valid
1.0.0-alpha
1.0.0-beta
1.0.0-alpha.beta
1.0.0-alpha.beta.1
1.0.0-alpha.1
1.0.0-alpha0.valid
1.0.0-alpha.0valid
1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay
1.0.0-rc.1+build.1
2.0.0-rc.1+build.123
1.2.3-beta
10.2.3-DEV-SNAPSHOT
1.2.3-SNAPSHOT-123
1.0.0
2.0.0
1.1.7
2.0.0+build.1848
2.0.1-alpha.1227
1.0.0-alpha+beta
1.2.3----RC-SNAPSHOT.12.9.1--.12+788
1.2.3----R-S.12.9.1--.12+meta
1.2.3----RC-SNAPSHOT.12.9.1--.12
1.0.0+0.build.1-rc.10000aaa-kk-0.1
99999999999999999999999.999999999999999999.99999999999999999
1.0.0-0A.is.legal


Invalid Semantic Versions

1
1.2
1.2.3-0123
1.2.3-0123.0123
1.1.2+.123
+invalid
-invalid
-invalid+invalid
-invalid.01
alpha
alpha.beta
alpha.beta.1
alpha.1
alpha+beta
alpha_beta
alpha.
alpha..
beta
1.0.0-alpha_beta
-alpha.
1.0.0-alpha..
1.0.0-alpha..1
1.0.0-alpha...1
1.0.0-alpha....1
1.0.0-alpha.....1
1.0.0-alpha......1
1.0.0-alpha.......1
01.1.1
1.01.1
1.1.01
1.2
1.2.3.DEV
1.2-SNAPSHOT
1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788
1.2-RC-SNAPSHOT
-1.0.3-gamma+b7718
+justmeta
9.8.7+meta+meta
9.8.7-whatever+meta+meta
99999999999999999999999.999999999999999999.99999999999999999----RC-SNAPSHOT.12.09.1--------------------------------..12`

	for i, match := range reS.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}

func TestParse(t *testing.T) {
	inps := []string{
		"NT 10.0.2004 Build number 2004",
		"10.0.9600",
		"1.2.3",
		"1.2.3.4",
		"1.2.3.4.5",
		"1.2.3a",
		"1.2b.3a",
		"1c.2b.3a",
		"1a.2",
		"ignore 1a.2b ok 1.2.3",
		"1a",
		"1.2.3aa",
		"1.2ba.3a",
		"1.2b.3a ok",
		"1.2b.3a-not",
		"1ca.2b.3a",
		"1c.2b.3a.4d.5ff",
		"1c.2b.3a.4d.5f.6g",
		"9.8.7-whatever+meta",
	}

	for _, i := range inps {
		v, err := Parse(i)
		fmt.Printf("%30v | %20v | %v\n", i, v, err)
	}
}

func TestExtractNum(t *testing.T) {
	inps := []string{
		"NT 10.0.2004 Build number 2004",
		"ABC 10.0.9600 DEF",
		"1.2.3",
		"1.2.3.4",
		"1.2.3.4.5",
		"1.2.3a",
		"1.2b.3a",
		"1c.2b.3a",
		"1c.2b.3a.4d.5f",
		"1.2",
		"1a.2",
		"ok 1a.2b ignore 1.2.3",

		// invalid
		"1a",
		"1.2.3aa",
		"1.2ba.3a",
		"1ca.2b.3a",
		"1c.2b.3a.4d.5ff",
		"1c.2b.3a.4d.5f.6g",
		"9.8.7-whatever+meta",
	}

	for _, i := range inps {
		v, err := ExtractNum(i)
		fmt.Printf("%30v | %v | %v\n", i, v, err)
	}
}
