package dbml

import (
	"strings"

	"github.com/josudoey/judo/lit"
)

// ref https://dbml.dbdiagram.io/docs/#column-definition
type Column struct {
	Name     string
	Type     string
	Settings []ColumnSetting
}

func (c *Column) String() string {
	return ColumnLiteral(c).String()
}

func ColumnLiteral(c *Column) lit.Literal {
	typ := c.Type
	if strings.Contains(typ, " ") {
		typ = lit.Quote(c.Type).String()
	}

	if len(c.Settings) == 0 {
		return lit.Add(c.Name, " ", typ)
	}

	return lit.Add(c.Name, " ", typ, " ", lit.SquareBracket(lit.Join(c.Settings, ", ")))
}
