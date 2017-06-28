package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfMysqlExampleApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfMysqlExampleApp Suite")
}
