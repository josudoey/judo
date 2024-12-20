package lit

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Compose", func(items []Composable, expected string) {
	Expect(render(Compose(items))).To(Equal(expected))
},
	Entry("empty", nil, ""),
	Entry("1 items",
		[]Composable{func(s *strings.Builder, next func()) {
			s.WriteString("a")
		}},
		"a"),
	Entry("2 items",
		[]Composable{
			func(s *strings.Builder, next func()) {
				s.WriteString("a")
			},
			func(s *strings.Builder, next func()) {
				s.WriteString("b")
			}},
		"ab"),
	Entry("test next",
		[]Composable{
			func(s *strings.Builder, next func()) {
				s.WriteString("[")
				next()
				s.WriteString("]")
				next()
			},
			func(s *strings.Builder, next func()) {
				s.WriteString("a")
			}, func(s *strings.Builder, next func()) {
				s.WriteString("b")
			}},
		"[ab]"),
)
