package lit

import (
	"fmt"
)

type Literal interface {
	fmt.Stringer
	Add(v ...any) Literal
}

type literal struct {
	elems []any
}

func (b *literal) Add(elems ...any) Literal {
	b.elems = append(b.elems, elems...)
	return b
}

func (b *literal) String() string {
	return concat(b.elems).String()
}

func Add(elems ...any) Literal {
	return &literal{elems: elems}
}

func Slice[T any](a []T) []any {
	items := make([]any, len(a))
	for i, item := range a {
		items[i] = Wrap(item)
	}

	return items
}
