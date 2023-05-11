package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/leeyenter/books/backend/utils"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	var recorder *httptest.ResponseRecorder
	var c *gin.Context
	var validJwt string

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(recorder)
		validJwt, _ = CreateJWT(utils.SaltedRemoteIPForTests, time.Minute)
	})

	It("blocks incorrect auth values logins", func() {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("auth", "invalid-value")
		c.Request = req
		Middleware(c)
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
	})

	It("blocks logins with an incorrect address", func() {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("auth", validJwt)
		req.RemoteAddr = "incorrect-ip-address"
		c.Request = req
		Middleware(c)
		Expect(recorder.Code).To(Equal(http.StatusUnprocessableEntity))
	})

	It("allows valid user logins", func() {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		utils.WrapRequestForTest(req, validJwt)
		c.Request = req
		Middleware(c)
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})
})
