package lit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("CurlyBracket", func(v any, expected string) {
	Expect(CurlyBracket(v).String()).To(Equal(expected))
},
	Entry("empty", Empty, "{}"),
	Entry("a", "a", "{a}"),
)

var _ = DescribeTable("SquareBracket", func(v any, expected string) {
	Expect(SquareBracket(v).String()).To(Equal(expected))
},
	Entry("empty", Empty, "[]"),
	Entry("a", "a", "[a]"),
)

var _ = DescribeTable("Join", func(items []any, sep string, expected string) {
	Expect(Join(items, sep).String()).To(Equal(expected))
},
	Entry("empty",
		[]any{}, "",
		""),
	Entry("1 items",
		[]any{"a"}, ", ",
		"a"),
	Entry("2 items",
		[]any{"a", "b"}, ", ",
		"a, b"),
)

var _ = DescribeTable("IndentSpace", func(v any, numberOfSpace int, expected string) {
	Expect(IndentSpace(v, numberOfSpace).String()).To(Equal(expected))
},
	Entry("empty",
		Empty, 0, ""),
	Entry("0 space",
		"a", 0, "a"),
	Entry("2 space",
		"a", 2, "  a"),
)

var _ = DescribeTable("Quote", func(v any, expected string) {
	Expect(Quote(v).String()).To(Equal(expected))
},
	Entry("empty", Empty, `""`),
	Entry("a", "a", `"a"`),
	Entry("new line", "\n", `"\n"`),
	Entry("single quote", "''", `"''"`),
	Entry("double quote", `""`, `"\"\""`),
)

var _ = DescribeTable("SingleQuote", func(v any, expected string) {
	Expect(SingleQuote(v).String()).To(Equal(expected))
},
	Entry("empty", Empty, "''"),
	Entry("a", "a", "'a'"),
	Entry("new line", "\n", "'\n'"),
	Entry("single quote", "''", `'\'\''`),
	Entry("double quote", `""`, `'""'`),
)
