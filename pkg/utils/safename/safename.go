package safename

import (
	"fmt"
	"regexp"
)

type safeName struct {
	Pattern         string
	CompiledPattern regexp.Regexp
}

type SafeNameClient interface {
	Check(stringToCheck string) bool
}

func UsePattern(pattern string) (SafeNameClient, error) {
	compiledRegExp, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to compile pattern: %s", err)
	}

	return safeName{
		CompiledPattern: *compiledRegExp,
		Pattern:         pattern,
	}, nil
}

// returns no error
func MustUsePattern(pattern string) SafeNameClient {
	safeNameClient, _ := UsePattern(pattern)
	return safeNameClient
}

func (sn safeName) Check(stringToCheck string) bool {
	return sn.CompiledPattern.MatchString(stringToCheck)
}
