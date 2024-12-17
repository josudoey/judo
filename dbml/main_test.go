package dbml

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDBML(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DBML Suite")
}
