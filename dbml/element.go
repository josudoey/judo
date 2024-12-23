package dbml

import "fmt"

type Element interface {
	fmt.Stringer

	element()
}

type implementedElement struct{}

func (s implementedElement) element() {}
