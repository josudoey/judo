package dbml

import "fmt"

type Element interface {
	fmt.Stringer
	element()
}

type sealedElement struct{}

func (s sealedElement) element() {}
