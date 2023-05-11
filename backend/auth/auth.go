package auth

import (
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"math/rand"
	"sync"
	"time"
)

var w *webauthn.WebAuthn
var once sync.Once

func Get() *webauthn.WebAuthn {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
		_ = initWebAuth()
	})

	return w
}

func initWebAuth() error {
	wconfig := &webauthn.Config{
		RPDisplayName: "Books",                                                        // Display Name for your site
		RPID:          "localhost",                                                    // Generally the FQDN for your site
		RPOrigins:     []string{"http://localhost:5173", "https://ytbooks.pages.dev"}, // The origin URLs allowed for WebAuthn requests
	}

	var err error
	if w, err = webauthn.New(wconfig); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
