package books

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/leeyenter/books/backend/db"
	"google.golang.org/api/iterator"
	"gopkg.in/guregu/null.v4"
	"time"
)

func GetAllRaw(ctx context.Context) ([]*firestore.DocumentSnapshot, error) {
	iter := db.Get().Books().Documents(ctx)
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

func GetAll(ctx context.Context) ([]Book, error) {
	docs, err := GetAllRaw(ctx)
	if err != nil {
		return nil, err
	}

	books := make([]Book, len(docs))
	for i, doc := range docs {
		var book Book
		err = doc.DataTo(&book)
		if err != nil {
			return nil, err
		}

		book.ID = doc.Ref.ID
		books[i] = book
	}

	return books, nil
}

func Add(ctx context.Context, book *Book) error {
	ref, _, err := db.Get().Books().Add(ctx, book)
	book.ID = ref.ID
	return err
}

func MarkBought(ctx context.Context, id, boughtType string) error {
	_, err := db.Get().Books().Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "boughtType",
			Value: boughtType,
		},
		{
			Path:  "boughtDate",
			Value: null.TimeFrom(time.Now()),
		},
	})

	return err
}

func SetBookPrice(ctx context.Context, id, source string, price int) error {
	_, err := db.Get().Books().Doc(id).Set(ctx, map[string]interface{}{
		"prices": map[string]interface{}{
			source: price,
		},
	}, firestore.MergeAll)

	return err
}

func RemoveBookPrice(ctx context.Context, id, source string) error {
	_, err := db.Get().Books().Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "prices." + source,
			Value: firestore.Delete,
		},
	})

	return err
}

func SetBookStatus(ctx context.Context, id, status string) error {
	_, err := db.Get().Books().Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "readStatus",
			Value: status,
		},
	})

	return err
}
