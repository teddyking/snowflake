package snowgauge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSnowgauge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Snowgauge Suite")
}
