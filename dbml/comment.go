package dbml

type sealedComment struct{}

func (s *sealedComment) comment() {}

type singleLineComment struct {
	sealedComment
	text string
}

type Comment interface {
	comment()
	String() string
}

func (s *singleLineComment) String() string {
	return "// " + s.text
}

func SingleLineComment(text string) Comment {
	return &singleLineComment{text: text}
}
