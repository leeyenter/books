package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/leeyenter/books/backend/auth"
	"github.com/leeyenter/books/backend/utils"
	"net/http"
	"time"
)

func checkLogin(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getNumCredentials(c *gin.Context) {
	var user auth.User
	c.JSON(http.StatusOK, gin.H{"num": len(user.WebAuthnCredentials())})
}

func beginRegistration(c *gin.Context) {
	var user auth.User

	// if user > 1 credential, and not auth, return early
	if len(user.WebAuthnCredentials()) > 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	options, session, err := auth.Get().BeginRegistration(user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := auth.SaveSession(c.Request.Context(), user.WebAuthnID(), session); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, options)
}

func finishRegistration(c *gin.Context) {
	response, err := protocol.ParseCredentialCreationResponseBody(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user auth.User
	session, err := auth.GetSession(c.Request.Context(), user.WebAuthnID())
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	credential, err := auth.Get().CreateCredential(user, session, response)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := auth.AddCredential(c.Request.Context(), *credential); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	remoteIp := utils.GetRemoteIp(c.Request)
	token, err := auth.CreateJWT(remoteIp, time.Hour)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func beginLogin(c *gin.Context) {
	var user auth.User

	options, session, err := auth.Get().BeginLogin(user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := auth.SaveSession(c.Request.Context(), user.WebAuthnID(), session); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, options)
}

func finishLogin(c *gin.Context) {
	response, err := protocol.ParseCredentialRequestResponseBody(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user auth.User
	session, err := auth.GetSession(c.Request.Context(), user.WebAuthnID())
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = auth.Get().ValidateLogin(user, session, response)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	fmt.Println("remote:", c.Request.RemoteAddr)
	fmt.Println("for:", c.Request.Header.Get("X-Forwarded-For"))

	remoteIp := utils.GetRemoteIp(c.Request)
	token, err := auth.CreateJWT(remoteIp, time.Hour)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
