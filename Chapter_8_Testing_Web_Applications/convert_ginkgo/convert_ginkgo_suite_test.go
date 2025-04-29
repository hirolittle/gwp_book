package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConvertGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertGinkgo Suite")
}
