package dbml

import "fmt"

type implementedComment struct{}

func (s *implementedComment) comment() {}

type singleLineComment struct {
	implementedComment

	text string
}

type Comment interface {
	fmt.Stringer

	comment()
}

func (s *singleLineComment) String() string {
	return "// " + s.text
}

func SingleLineComment(text string) Comment {
	return &singleLineComment{text: text}
}
