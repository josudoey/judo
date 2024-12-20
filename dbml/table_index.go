package dbml

import "github.com/josudoey/judo/lit"

// https://dbml.dbdiagram.io/docs/#index-definition
type TableIndex struct {
	ColumnNames []string
	Settings    []TableIndexSetting
}

func (i *TableIndex) String() string {
	if len(i.ColumnNames) == 0 {
		return ""
	}

	name := i.ColumnNames[0]
	if len(i.ColumnNames) > 1 {
		name = lit.RoundBracket(lit.Join(lit.Slice(i.ColumnNames), ", ")).String()
	}
	if len(i.Settings) == 0 {
		return name
	}

	return lit.Add(name, " ", lit.SquareBracket(lit.Join(lit.Slice(i.Settings), ", "))).String()
}
