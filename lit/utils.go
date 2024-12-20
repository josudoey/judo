package lit

import (
	"fmt"
	"strconv"
	"strings"
)

func RoundBracket(v any) Literal {
	return Add("(", v, ")")
}

func CurlyBracket(v any) Literal {
	return Add("{", v, "}")
}

func SquareBracket(v any) Literal {
	return Add("[", v, "]")
}

func Join(items []any, sep string) Literal {
	if len(items) == 0 {
		return Add()
	}

	l := Add(items[0])
	for _, item := range items[1:] {
		l.Add(sep, item)
	}
	return l
}

func IndentSpace(v any, numberOfSpace int) Literal {
	return Add(strings.Repeat(" ", numberOfSpace), v)
}

func Format(format string, a ...any) Literal {
	var c Composable = func(b *strings.Builder, next func()) {
		b.WriteString(Wrap(fmt.Sprintf(format, a...)).String())
	}
	return Wrap(c)
}

func Quote(v any) Literal {
	var c Composable = func(b *strings.Builder, next func()) {
		b.WriteString(strconv.Quote(Wrap(v).String()))
	}
	return Wrap(c)
}
func SingleQuote(v any) Literal {
	var c Composable = func(b *strings.Builder, next func()) {
		b.WriteString("'" + strings.ReplaceAll(Wrap(v).String(), "'", "\\'") + "'")
	}
	return Wrap(c)
}
