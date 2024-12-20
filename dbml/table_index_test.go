package dbml

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("TableIndex", func(c *TableIndex, expected string) {
	Expect(c.String()).To(Equal(expected))
},
	Entry("composite primary key",
		&TableIndex{
			ColumnNames: []string{"id", "country"},
			Settings:    []TableIndexSetting{PK()},
		},
		"(id, country) [pk]"),
	Entry("created_at",
		&TableIndex{
			ColumnNames: []string{"created_at"},
			Settings:    []TableIndexSetting{TableIndexName("created_at_index"), TableIndexNote("Date")},
		},
		"created_at [name: 'created_at_index', note: 'Date']"),
	Entry("only name",
		&TableIndex{
			ColumnNames: []string{"booking_date"},
		},
		"booking_date"),
	Entry("unique",
		&TableIndex{
			ColumnNames: []string{"country", "booking_date"},
			Settings:    []TableIndexSetting{Unique()},
		},
		"(country, booking_date) [unique]"),
	Entry("type",
		&TableIndex{
			ColumnNames: []string{"booking_date"},
			Settings:    []TableIndexSetting{TableIndexType("hash")},
		},
		"booking_date [type: hash]"),
)
