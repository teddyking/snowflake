package flaker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFlaker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flaker Service Suite")
}
