package dbml

import (
	"github.com/josudoey/judo/lit"
)

// ref https://dbml.dbdiagram.io/docs/#table-definition

type Table struct {
	sealedElement

	Name          string
	TableSettings []TableSetting
	Columns       []*Column
	TableIndexes  []*TableIndex
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

	if len(t.TableIndexes) == 0 {
		return l.Add(lit.CurlyBracket(
			lit.Add(
				"\n",
				lit.Join(columns, "\n"),
				"\n",
			),
		), "\n").String()
	}

	tableIndexes := lit.Slice(t.TableIndexes)
	for i, tableIndex := range tableIndexes {
		tableIndexes[i] = lit.IndentSpace(tableIndex, 4)
	}

	return l.Add(lit.CurlyBracket(
		lit.Add(
			"\n",
			lit.Join(columns, "\n"),
			"\n\n",
			"  indexes ", lit.CurlyBracket(
				lit.Add(
					"\n",
					lit.Join(tableIndexes, "\n"),
					"\n  "),
			), "\n",
		),
	), "\n").String()
}
