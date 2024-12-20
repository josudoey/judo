package dbml

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Table", func(t *Table, expected string) {
	Expect(t.String()).To(Equal(expected))
},
	Entry(`users`,
		&Table{
			Name: "users",
			Columns: []*Column{
				{Name: "id", Type: "integer"},
				{Name: "username", Type: "varchar"},
				{Name: "role", Type: "varchar"},
				{Name: "created_at", Type: "timestamp"},
			},
		},
		`Table users {
  id integer
  username varchar
  role varchar
  created_at timestamp
}
`),
	Entry(`posts`,
		&Table{
			Name: "posts",
			Columns: []*Column{
				{Name: "id", Type: "integer", Settings: []ColumnSetting{PrimaryKey()}},
				{Name: "title", Type: "varchar"},
				{Name: "body", Type: "text", Settings: []ColumnSetting{ColumnNote("Content of the post")}},
				{Name: "user_id", Type: "integer"},
				{Name: "created_at", Type: "timestamp"},
			},
		},
		`Table posts {
  id integer [primary key]
  title varchar
  body text [note: 'Content of the post']
  user_id integer
  created_at timestamp
}
`),
	Entry(`bookings`,
		&Table{
			Name: "bookings",
			Columns: []*Column{
				{Name: "id", Type: "integer"},
				{Name: "country", Type: "varchar"},
				{Name: "booking_date", Type: "date"},
				{Name: "created_at", Type: "timestamp"},
		},
			TableIndexes: []*TableIndex{
				{ColumnNames: []string{"id", "country"}, Settings: []TableIndexSetting{PK()}},
				{ColumnNames: []string{"created_at"}, Settings: []TableIndexSetting{TableIndexName("created_at_index"), TableIndexNote("Date")}},
				{ColumnNames: []string{"booking_date"}},
				{ColumnNames: []string{"country", "booking_date"}, Settings: []TableIndexSetting{Unique()}},
				{ColumnNames: []string{"booking_date"}, Settings: []TableIndexSetting{TableIndexType("hash")}},
			},
		},
		`Table bookings {
  id integer
  country varchar
  booking_date date
  created_at timestamp

  indexes {
    (id, country) [pk]
    created_at [name: 'created_at_index', note: 'Date']
    booking_date
    (country, booking_date) [unique]
    booking_date [type: hash]
  }
}
`),
)
