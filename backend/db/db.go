package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var db *DB
var once sync.Once

type DB struct {
	client *firestore.Client
	env    string
}

func Get() *DB {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
		db = &DB{}
		db.init()
	})

	return db
}

func (d *DB) init() {
	ctx := context.Background()
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	var err error

	if env == "test" || env == "dev" {
		d.client, err = firestore.NewClient(ctx, "ytbooks", option.WithCredentialsFile("ytbooks-firebase-adminsdk-yosds-19aba8652d.json"))
	} else {
		d.client, err = firestore.NewClient(ctx, "ytbooks")
	}

	if err != nil {
		log.Panicln(err)
	}

	d.env = env
}

func (d *DB) Books() *firestore.CollectionRef {
	return d.client.Collection("books-" + d.env)
}

func (d *DB) BulkWriter() *firestore.WriteBatch {
	return d.client.Batch()
}

func (d *DB) Close() {
	_ = d.client.Close()
}
