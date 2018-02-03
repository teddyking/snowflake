package snowflake_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSnowflake(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Snowflake Suite")
}
