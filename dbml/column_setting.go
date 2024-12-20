package dbml

import (
	"fmt"
	"strconv"
)

type ColumnSetting interface {
	fmt.Stringer
	columnSetting()
}

type sealedColumnSetting struct{}

func (s sealedColumnSetting) columnSetting() {}

type null struct{ sealedColumnSetting }

func (s *null) String() string {
	return "null"
}

func Null() ColumnSetting {
	return &null{}
}

type notNull struct{ sealedColumnSetting }

func (s *notNull) String() string {
	return "not null"
}

func NotNull() ColumnSetting {
	return &notNull{}
}

type defaultString struct {
	sealedColumnSetting
	value string
}

func (s *defaultString) String() string {
	return "default: " + strconv.Quote(s.value)
}

func DefaultString(value string) ColumnSetting {
	return &defaultString{value: value}
}

type defaultExpression struct {
	sealedColumnSetting
	value string
}

func (s *defaultExpression) String() string {
	return "default: `" + s.value + "`"
}

func DefaultExpression(value string) ColumnSetting {
	return &defaultExpression{value: value}
}

type defaultBoolean struct {
	sealedColumnSetting
	value bool
}

func (s *defaultBoolean) String() string {
	if s.value {
		return "default: true"
	}

	return "default: false"
}

func DefaultBoolean(value bool) ColumnSetting {
	return &defaultBoolean{value: value}
}

type defaultNull struct{ sealedColumnSetting }

func (s *defaultNull) String() string {
	return "default: null"
}

func DefaultNull() ColumnSetting {
	return &defaultNull{}
}

type defaultInt struct {
	sealedColumnSetting
	value int64
}

func (s *defaultInt) String() string {
	return "default: " + strconv.FormatInt(s.value, 10)
}

func DefaultInt(value int64) ColumnSetting {
	return &defaultInt{value: value}
}

type defaultFloat struct {
	sealedColumnSetting
	value float64
}

func (s *defaultFloat) String() string {
	return "default: " + strconv.FormatFloat(s.value, 'f', -1, 64)
}

func DefaultFloat(value float64) ColumnSetting {
	return &defaultFloat{value: value}
}

type columnOneToMany struct {
	sealedColumnSetting

	TableName  string
	ColumnName string
}

func (c *columnOneToMany) String() string {
	return "ref: < " + c.TableName + "." + c.ColumnName
}

func ColumnOneToMany(tableName string, columnName string) ColumnSetting {
	return &columnOneToMany{TableName: tableName, ColumnName: columnName}
}

type columnManyToOne struct {
	sealedColumnSetting

	TableName  string
	ColumnName string
}

func (c *columnManyToOne) String() string {
	return "ref: > " + c.TableName + "." + c.ColumnName
}

func ColumnManyToOne(tableName string, columnName string) ColumnSetting {
	return &columnManyToOne{TableName: tableName, ColumnName: columnName}
}

type columnOneToOne struct {
	sealedColumnSetting

	TableName  string
	ColumnName string
}

func (c *columnOneToOne) String() string {
	return "ref: - " + c.TableName + "." + c.ColumnName
}

func ColumnOneToOne(tableName string, columnName string) ColumnSetting {
	return &columnOneToOne{TableName: tableName, ColumnName: columnName}
}

type columnManyToMany struct {
	sealedColumnSetting

	TableName  string
	ColumnName string
}

func (c *columnManyToMany) String() string {
	return "ref: <> " + c.TableName + "." + c.ColumnName
}

func ColumnManyToMany(tableName string, columnName string) ColumnSetting {
	return &columnManyToMany{TableName: tableName, ColumnName: columnName}
}
