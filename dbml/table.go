package dbml

import (
	"github.com/josudoey/judo/lit"
)

// ref https://dbml.dbdiagram.io/docs/#table-definition

type TableSetting interface {
	tableSetting()
	String() string
}

type sealedTableSetting struct{}

func (s sealedTableSetting) tableSetting() {}

type headercolor struct {
	sealedTableSetting
	code string
}

func (s *headercolor) String() string {
	return "headercolor: " + s.code
}

func Headercolor(code string) TableSetting {
	return &headercolor{code: code}
}

type Table struct {
	sealedElement

	Name          string
	TableSettings []TableSetting
	Columns       []*Column
}

func (t *Table) String() string {
	l := lit.Add("Table ", t.Name, " ")
	if len(t.TableSettings) > 0 {
		l.Add(lit.SquareBracket(lit.Slice(t.TableSettings)), " ")
	}

	columns := lit.Slice(t.Columns)
	for i, column := range columns {
		columns[i] = lit.IndentSpace(column, 2)
	}

	return l.Add(lit.CurlyBracket(
		lit.Add(
			"\n",
			lit.Join(columns, "\n"),
			"\n",
		),
	), "\n").String()
}
