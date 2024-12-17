package dbml

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("ColumnSetting", func(result fmt.Stringer, expected string) {
	Expect(result.String()).To(Equal(expected))
},
	Entry(`Note`, ColumnNote("string to add notes"), `note: 'string to add notes'`),
	Entry(`PrimaryKey`, PrimaryKey(), "primary key"),
	Entry(`PK`, PK(), "pk"),
	Entry(`Null`, Null(), "null"),
	Entry(`NotNull`, NotNull(), "not null"),
	Entry(`DefaultInt`, DefaultInt(123), "default: 123"),
	Entry(`DefaultFloat`, DefaultFloat(123.456), "default: 123.456"),
	Entry(`DefaultString`, DefaultString("some string value"), `default: "some string value"`),
	Entry(`DefaultExpression`, DefaultExpression(`now() - interval '5 days'`), "default: `now() - interval '5 days'`"),
	Entry(`DefaultBoolean true`, DefaultBoolean(true), "default: true"),
	Entry(`DefaultBoolean false`, DefaultBoolean(false), "default: false"),
	Entry(`ColumnOneToMany`, ColumnOneToMany("posts", "user_id"), "ref: < posts.user_id"),
	Entry(`ColumnManyToOne`, ColumnManyToOne("users", "id"), "ref: > users.id"),
)
