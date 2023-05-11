package main

import (
	"github.com/leeyenter/books/backend/auth"
	"github.com/leeyenter/books/backend/utils"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth Handlers", func() {
	var validJwt string
	BeforeEach(func() {
		validJwt, _ = auth.CreateJWT(utils.SaltedRemoteIPForTests, time.Minute)
	})

	Describe("login check handler", func() {
		It("returns 401 when not authenticated", func() {
			r, err := http.NewRequest("GET", "/auth/", nil)
			Expect(err).To(BeNil())

			router := setupRouter()
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})

		It("returns 200 when authenticated", func() {
			r, err := http.NewRequest("GET", "/auth/", nil)
			Expect(err).To(BeNil())

			utils.WrapRequestForTest(r, validJwt)
			router := setupRouter()
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})
