package borscht_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBorscht(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Borscht Suite")
}

func envMustHave(key string) string {
	val := os.Getenv(key)
	Expect(val).NotTo(BeEmpty(), fmt.Sprintf("must set env var %s", key))
	return val
}
