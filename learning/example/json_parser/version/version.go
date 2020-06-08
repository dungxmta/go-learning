package version

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"regexp"
	"strconv"
	"strings"
)

// semantic version: MAJOR.MINOR.PATCH - https://semver.org/
const (
	// true semantic version
	PatternSemVer = `(?m)^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	// contains: extend sem ver to 5 nums
	PatternSem5C = `(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(\.(?P<rev1>0|[1-9]\d*))?(\.(?P<rev2>0|[1-9]\d*))?(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`

	PatternNum = `^(?P<v>[0-9]+)(?P<rev>[a-zA-Z])?$`
)

var (
	reS = regexp.MustCompile(PatternSemVer)

	reSC = regexp.MustCompile(PatternSem5C)

	reNum = regexp.MustCompile(PatternNum)
)

var (
	reVersExactly = []*regexp.Regexp{
		reS,
	}
	reVersContain = []*regexp.Regexp{
		reSC,
	}
)

func Parse(s string) (string, error) {
	var nums []string

	// check exactly
	for _, r := range reVersExactly {
		gr := r.FindStringSubmatch(s)
		if len(gr) > 0 {
			// log.Println("exactly:", gr)
			out, ok := SemVer(gr[0])
			if ok {
				for _, i := range out {
					nums = append(nums, strconv.Itoa(i))
				}
				// break
				return JoinVersion(nums), nil
			}
		}
	}

	// check contains
	for _, r := range reVersContain {
		gr := r.FindStringSubmatch(s)
		if len(gr) > 0 {
			// log.Println("contain:", gr)
			out, ok := SemVer(gr[0])
			if ok {
				for _, i := range out {
					nums = append(nums, strconv.Itoa(i))
				}
				// break
				return JoinVersion(nums), nil
			}
		}
	}

	// v, err := version.NewVersion(s)
	// if err != nil {
	// 	return "", err
	// }

	nums, ok := ExtractNum(s)
	if !ok {
		return s, fmt.Errorf("invalid version %v", s)
	}

	// var ver string
	// // nums := v.Segments()
	// for _, i := range nums {
	// 	ver = fmt.Sprintf("%v%10v", ver, i)
	// }

	return JoinVersion(nums), nil
}

func JoinVersion(nums []string) string {
	var ver string
	// nums := v.Segments()
	for _, i := range nums {
		ver = fmt.Sprintf("%v%10v", ver, i)
	}

	return ver
}

func SemVer(s string) ([]int, bool) {
	v, err := version.NewVersion(s)
	if err != nil {
		return nil, false
	}

	return v.Segments(), true
}

// split by space " "                                           =/> A B C
// check each part contains at least 1 dot "." and max dot is 4 =/> 1a 2b 3c 4d 5e
// check format each dot: <number><one_char>                    =/> 1 2 3 4 5
// if all fail, fallback to check format "1a"                   =/> 1 0 0
func ExtractNum(s string) ([]string, bool) {
	var out []string
	var ok bool

	// split " "
	parts := strings.Split(s, " ")

	for _, part := range parts {
		ok = true
		nums := strings.Split(part, ".")
		// check .
		if len(nums) <= 1 || len(nums) > 5 {
			ok = false
			out = []string{}
			continue
		}

		for _, n := range nums {
			gr := reNum.FindStringSubmatch(n)
			if len(gr) > 0 {
				// log.Println("-> ", gr)
				out = append(out, gr[1])
				ok = true
				continue
			}
			// break if one num is invalid
			ok = false
			out = []string{}
			break
		}

		// one part is ok
		if ok {
			break
		}
	}

	// fallback to check exactly "1a"
	if !ok && len(parts) == 1 {
		gr := reNum.FindStringSubmatch(s)
		if len(gr) > 0 {
			out = append(out, gr[1])
			ok = true
		}
	}

	if ok {
		// min 3 nums
		if len(out) < 3 {
			for i := 0; i <= 3-len(out); i++ {
				out = append(out, "0")
			}
		}
		return out, ok
	}

	return nil, false
}
