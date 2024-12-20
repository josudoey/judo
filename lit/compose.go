package lit

import (
	"fmt"
	"strings"
	"sync"
)

type Composable func(b *strings.Builder, next func())

func (c Composable) String() string {
	return render(c)
}

func (c Composable) Add(elems ...any) Literal {
	return &literal{elems: elems}
}

func once(fn func()) func() {
	var one sync.Once
	return func() {
		one.Do(fn)
	}
}

func Nop()                                  {}
func Empty(s *strings.Builder, next func()) {}

func Compose(items []Composable) Composable {
	if len(items) == 0 {
		return Empty
	}

	return func(s *strings.Builder, next func()) {
		count := len(items)
		actions := make([]func(), count)
		var dispatch = func(i int) func() {
			return once(func() {
				items[i](s, actions[i+1])
				actions[i+1]()
			})
		}

		lastIndex := len(items) - 1
		for i := 0; i < lastIndex; i++ {
			actions[i] = dispatch(i)
		}

		actions[lastIndex] = once(func() {
			items[lastIndex](s, Nop)
		})

		actions[0]()
	}
}

func render(c Composable) string {
	result := new(strings.Builder)
	c(result, Nop)
	return result.String()
}

func concat[T any](items []T) Composable {
	result := make([]Composable, len(items))
	for i, item := range items {
		result[i] = Wrap(item)
	}

	return Compose(result)
}

func write(p []byte) Composable {
	return func(b *strings.Builder, next func()) {
		b.Write(p)
	}
}

func writeString(s string) Composable {
	return func(b *strings.Builder, next func()) {
		b.WriteString(s)
	}
}

func Wrap(a any) Composable {
	switch v := a.(type) {
	case Composable:
		return v
	case fmt.Stringer:
		return writeString(v.String())
	case string:
		return writeString(v)
	case []byte:
		return write(v)
	case []any:
		return concat(v)
	}

	return Empty
}
