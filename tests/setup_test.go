package tests

import (
	"testing"
	"os"
	"github.com/smbody/common-server/tests/mocks"
	f "github.com/smbody/common-server/container"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	f.Init(&mocks.TestFactory{})
}

func tearDown() {

}