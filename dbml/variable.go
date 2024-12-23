package dbml

import (
	"regexp"

	"github.com/josudoey/judo/lit"
)

var basicRegexp = regexp.MustCompile("^[_0-9a-zA-Z]+$")

func Variable(name string) lit.Literal {
	if basicRegexp.Match([]byte(name)) {
		return lit.Add(name)
	}

	return lit.Quote(name)
}

func Variables(names []string) []lit.Literal {
	result := make([]lit.Literal, len(names))
	for i, name := range names {
		result[i] = Variable(name)
	}

	return result
}
