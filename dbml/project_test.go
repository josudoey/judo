package dbml

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Project", func(p *Project, expected string) {
	Expect(p.String()).To(Equal(expected))
},
	Entry(`PostgreSQL`, &Project{
		Name:         "project_name",
		DatabaseType: PostgreSQL,
		Note:         "Description of the project",
	}, `Project project_name {
  database_type: 'PostgreSQL'
  Note: 'Description of the project'
}
`))
