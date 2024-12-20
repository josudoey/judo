package dbml

import "fmt"

type TableSetting interface {
	fmt.Stringer
	tableSetting()
}

type sealedTableSetting struct{}

func (s sealedTableSetting) tableSetting() {}

type headercolor struct {
	sealedTableSetting
	code string
}

func (s *headercolor) String() string {
	return "headercolor: " + s.code
}

func Headercolor(code string) TableSetting {
	return &headercolor{code: code}
}
