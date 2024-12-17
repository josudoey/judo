package dbml

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("TableSetting", func(result fmt.Stringer, expected string) {
	Expect(result.String()).To(Equal(expected))
},
	Entry(`NewHeaderColor`, Headercolor("#3498DB"), "headercolor: #3498DB"),
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
)
