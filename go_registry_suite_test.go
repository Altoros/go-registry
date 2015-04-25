package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoRegistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoRegistry Suite")
}
