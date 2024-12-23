package dbml

import "fmt"

type TableSetting interface {
	fmt.Stringer

	tableSetting()
}

type implementedTableSetting struct{}

func (s implementedTableSetting) tableSetting() {}

type headercolor struct {
	implementedTableSetting

	code string
}

func (s *headercolor) String() string {
	return "headercolor: " + s.code
}

func Headercolor(code string) TableSetting {
	return &headercolor{code: code}
}
