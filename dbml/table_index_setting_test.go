package dbml

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("TableIndexSetting", func(result fmt.Stringer, expected string) {
	Expect(result.String()).To(Equal(expected))
},
	Entry("Note", TableIndexNote("Date"), `note: 'Date'`),
	Entry("PrimaryKey", PrimaryKey(), "primary key"),
	Entry("PK", PK(), "pk"),
	Entry("TableIndexName", TableIndexName("created_at_index"), `name: 'created_at_index'`),
	Entry("TableIndexType", TableIndexType("hash"), "type: hash"),
)
