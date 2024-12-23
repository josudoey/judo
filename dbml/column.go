package dbml

import (
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
	if len(c.Settings) == 0 {
		return lit.Add(c.Name, " ", Variable(c.Type))
	}

	return lit.Add(Variable(c.Name), " ", Variable(c.Type), " ", lit.SquareBracket(lit.Join(c.Settings, ", ")))
}
