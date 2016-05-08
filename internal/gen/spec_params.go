package gen

import (
	"errors"
	"regexp"
)

var (
	reValidParams   = regexp.MustCompile(`\?(.+?)\b`)
	reInvalidParams = regexp.MustCompile(`\?(\d|\W|$)`)
)

var (
	// ErrInvalidParams is error for invalid query parameters definition.
	ErrInvalidParams = errors.New("query parameters must have a name and the name must start with letter")
)

// CreateParamsSpec returns a new ParamsSpec from an SQL statement.
func CreateParamsSpec(s string) (ParamsSpec, error) {
	if hasInvalidParams(s) {
		return ParamsSpec{}, ErrInvalidParams
	}

	pos := createParamsUsage(s)

	return ParamsSpec{
		Declaration: createParamsDeclaration(pos),
		Usage:       pos,
	}, nil
}

func createParamsUsage(s string) []string {
	all := reValidParams.FindAllStringSubmatch(s, -1)
	var arr []string
	for _, a := range all {
		arr = append(arr, a[1])
	}
	return arr
}

func hasInvalidParams(s string) bool {
	return reInvalidParams.MatchString(s)
}

func createParamsDeclaration(a []string) []string {
	var (
		m = make(map[string]bool)
		r []string
	)

	for _, e := range a {
		_, ok := m[e]
		if ok {
			continue
		}

		m[e] = true
		r = append(r, e)
	}
	return r
}

func removeParamsName(s string) string {
	return reValidParams.ReplaceAllString(s, "?")
}
