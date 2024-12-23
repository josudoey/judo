package dbml

import "fmt"

type implementedComment struct{}

func (s *implementedComment) comment() {}

type singleLineComment struct {
	implementedComment
	implementedElement

	text string
}

type Comment interface {
	fmt.Stringer
	Element

	comment()
}

func (s *singleLineComment) String() string {
	return "// " + s.text
}

func SingleLineComment(text string) Comment {
	return &singleLineComment{text: text}
}
