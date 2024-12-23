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

	name := Variable(i.ColumnNames[0])
	if len(i.ColumnNames) > 1 {
		name = lit.RoundBracket(lit.Join(Variables(i.ColumnNames), ", "))
	}

	if len(i.Settings) == 0 {
		return name.String()
	}

	return lit.Add(name, " ", lit.SquareBracket(lit.Join(i.Settings, ", "))).String()
}
