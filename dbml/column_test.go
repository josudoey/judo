package dbml

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Column", func(c *Column, expected string) {
	Expect(c.String()).To(Equal(expected))
},
	Entry("Column address",
		&Column{
			Name:     "address",
			Type:     "varchar(255)",
			Settings: []ColumnSetting{Unique(), NotNull(), ColumnNote("to include unit number")},
		},
		`address "varchar(255)" [unique, not null, note: 'to include unit number']`),
	Entry("Column id",
		&Column{
			Name:     "id",
			Type:     "integer",
			Settings: []ColumnSetting{PK(), Unique(), DefaultInt(123), ColumnNote("Number")},
		},
		`id integer [pk, unique, default: 123, note: 'Number']`),
)
