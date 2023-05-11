package auth

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/base64"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/leeyenter/books/backend/db"
	"google.golang.org/api/iterator"
)

func SaveSession(ctx context.Context, userIdBytes []byte, session *webauthn.SessionData) error {
	userIdString := base64.StdEncoding.EncodeToString(userIdBytes)
	_, err := db.Get().AuthSession().Doc(userIdString).Set(ctx, session)
	return err
}

func GetSession(ctx context.Context, userIdBytes []byte) (webauthn.SessionData, error) {
	var resp webauthn.SessionData
	userIdString := base64.StdEncoding.EncodeToString(userIdBytes)
	doc, err := db.Get().AuthSession().Doc(userIdString).Get(ctx)
	if err != nil {
		return resp, err
	}

	err = doc.DataTo(&resp)
	return resp, err
}

func AddCredential(ctx context.Context, credential webauthn.Credential) error {
	_, _, err := db.Get().AuthCredential().Add(ctx, credential)
	return err
}

func getCredentialDocs(ctx context.Context) ([]*firestore.DocumentSnapshot, error) {
	iter := db.Get().AuthCredential().Documents(ctx)
	docs := make([]*firestore.DocumentSnapshot, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}

func GetCredentials(ctx context.Context) ([]webauthn.Credential, error) {
	docs, err := getCredentialDocs(ctx)
	if err != nil {
		return nil, err
	}

	creds := make([]webauthn.Credential, len(docs))
	for i, doc := range docs {
		var cred webauthn.Credential
		if err := doc.DataTo(&cred); err != nil {
			return nil, err
		}

		creds[i] = cred
	}

	return creds, nil
}
