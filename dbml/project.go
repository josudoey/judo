package dbml

import (
	"github.com/josudoey/judo/lit"
)

const PostgreSQL string = "PostgreSQL"

type Project struct {
	sealedElement

	Name         string
	DatabaseType string
	Note         string
}

func (p *Project) String() string {
	l := lit.Add(
		"Project ", p.Name, " {\n",
		"  database_type: ", lit.SingleQuote(p.DatabaseType), "\n",
	)

	if p.Note != "" {
		l.Add("  Note: ", lit.SingleQuote(p.Note), "\n")
	}

	return l.Add("}\n").String()
}
