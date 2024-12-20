package dbml

type Setting interface {
	ColumnSetting
	TableIndexSetting
}

type pK struct {
	sealedColumnSetting
	sealedTableIndexSetting
}

func (s *pK) String() string {
	return "pk"
}

func PK() Setting {
	return &pK{}
}

type primaryKey struct {
	sealedColumnSetting
	sealedTableIndexSetting
}

func (s *primaryKey) String() string {
	return "primary key"
}

func PrimaryKey() Setting {
	return &primaryKey{}
}

type unique struct {
	sealedColumnSetting
	sealedTableIndexSetting
}

func (s *unique) String() string {
	return "unique"
}

func Unique() Setting {
	return &unique{}
}
