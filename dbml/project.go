package dbml

import (
	"strings"

	"github.com/josudoey/judo/lit"
)

const PostgreSQL string = "PostgreSQL"

type Project struct {
	implementedElement

	Name         string
	DatabaseType string
	Note         string
}

func (p *Project) String() string {
	name := p.Name
	if strings.Contains(name, "-") {
		name = lit.Quote(name).String()
	}

	l := lit.Add(
		"Project ", name, " {\n",
		"  database_type: ", lit.SingleQuote(p.DatabaseType), "\n",
	)

	if p.Note != "" {
		l.Add("  Note: ", lit.SingleQuote(p.Note), "\n")
	}

	return l.Add("}\n").String()
}
