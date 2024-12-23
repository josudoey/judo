package dbml

import (
	"fmt"

	"github.com/josudoey/judo/lit"
)

type TableIndexSetting interface {
	fmt.Stringer

	tableIndexSetting()
}

type implementedTableIndexSetting struct{}

func (s implementedTableIndexSetting) tableIndexSetting() {}

type tableIndexName struct {
	implementedTableIndexSetting

	text string
}

func (c *tableIndexName) String() string {
	return lit.Add("name: ", lit.SingleQuote(c.text)).String()
}

func TableIndexName(text string) TableIndexSetting {
	return &tableIndexName{text: text}
}

type tableIndexType struct {
	implementedTableIndexSetting

	text string
}

func (c *tableIndexType) String() string {
	return lit.Add("type: ", c.text).String()
}

func TableIndexType(text string) TableIndexSetting {
	return &tableIndexType{text: text}
}
