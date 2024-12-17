package dbml

import (
	"fmt"

	"github.com/josudoey/judo/lit"
)

type Note interface {
	fmt.Stringer
}

type columnNote struct {
	sealedColumnSetting
	text string
}

func (c *columnNote) String() string {
	return lit.Add("note: ", lit.SingleQuote(c.text)).String()
}

func ColumnNote(text string) ColumnSetting {
	return &columnNote{text: text}
}
