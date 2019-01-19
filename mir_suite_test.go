package mir_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMir(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mir Suite")
}
