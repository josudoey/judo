package dbml

type blankline struct {
	implementedElement
}

func (s *blankline) String() string {
	return ""
}

func Blankline() Element {
	return &blankline{}
}
