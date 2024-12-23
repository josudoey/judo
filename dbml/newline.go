package dbml

type newLine struct {
	implementedElement
}

func (s *newLine) String() string {
	return "\n"
}

func NewLine() Element {
	return &newLine{}
}
