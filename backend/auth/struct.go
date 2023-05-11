package auth

import (
	"context"
	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct{}

var _ webauthn.User = (*User)(nil)

func (u User) WebAuthnID() []byte {
	return []byte("yenter")
}

func (u User) WebAuthnName() string {
	return "yenter"
}

func (u User) WebAuthnDisplayName() string {
	return "Yen Ter"
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	creds, _ := GetCredentials(context.Background())
	return creds
}

func (u User) WebAuthnIcon() string {
	return "no-icon"
}
