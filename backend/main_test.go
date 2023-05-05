package main

import (
	"github.com/leeyenter/books/backend/db"

	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("ENV", "test")
	clearBooks()
	code := m.Run()
	db.Get().Close()
	os.Exit(code)
}

func TestBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backend Suite")
}
