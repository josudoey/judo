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
