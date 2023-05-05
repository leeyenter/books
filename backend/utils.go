package main

import (
	"context"
	"fmt"
	"github.com/leeyenter/books/backend/books"
	"github.com/leeyenter/books/backend/db"
)

func clearBooks() {
	bulkWriter := db.Get().BulkWriter()
	docs, err := books.GetAllRaw(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(docs) == 0 {
		return
	}

	for _, doc := range docs {
		_ = bulkWriter.Delete(doc.Ref)
	}

	_, err = bulkWriter.Commit(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
