package dbml

import (
	"fmt"

	"github.com/josudoey/judo/lit"
)

// ref https://dbml.dbdiagram.io/docs/#syntax-consistency
type document struct {
	children []Element
}

type Document interface {
	fmt.Stringer
	Append(e Element) Document
}

func (d *document) Append(e Element) Document {
	d.children = append(d.children, e)
	return d
}

func (d *document) String() string {
	return lit.Join(lit.Slice(d.children), "\n").String()
}
