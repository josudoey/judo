package dbml

import (
	"fmt"

	"github.com/josudoey/judo/lit"
)

type Note interface {
	fmt.Stringer
}

type noteSetting struct {
	sealedColumnSetting
	sealedTableIndexSetting
	text string
}

func (c *noteSetting) String() string {
	return lit.Add("note: ", lit.SingleQuote(c.text)).String()
}

func ColumnNote(text string) ColumnSetting {
	return &noteSetting{text: text}
}

func TableIndexNote(text string) TableIndexSetting {
	return &noteSetting{text: text}
}
