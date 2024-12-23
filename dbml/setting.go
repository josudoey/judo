package dbml

type Setting interface {
	ColumnSetting
	TableIndexSetting
}

type pK struct {
	implementedColumnSetting
	implementedTableIndexSetting
}

func (s *pK) String() string {
	return "pk"
}

func PK() Setting {
	return &pK{}
}

type primaryKey struct {
	implementedColumnSetting
	implementedTableIndexSetting
}

func (s *primaryKey) String() string {
	return "primary key"
}

func PrimaryKey() Setting {
	return &primaryKey{}
}

type unique struct {
	implementedColumnSetting
	implementedTableIndexSetting
}

func (s *unique) String() string {
	return "unique"
}

func Unique() Setting {
	return &unique{}
}
