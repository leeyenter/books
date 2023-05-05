package books

import (
	"github.com/brianvoe/gofakeit/v6"
	"gopkg.in/guregu/null.v4"
	"time"
)

type Book struct {
	ID         string         `json:"id" firestore:"-"`
	Title      string         `json:"title" firestore:"title"`
	Authors    []string       `json:"authors" firestore:"authors"`
	Prices     map[string]int `json:"prices" firestore:"prices,omitempty"`
	BoughtDate null.Time      `json:"boughtDate" firestore:"boughtDate,omitempty"`
	BoughtType string         `json:"boughtType" firestore:"boughtType,omitempty"`
	ReadStatus string         `json:"readStatus" firestore:"readStatus,omitempty"`
}

func CreateRandom(bought bool) Book {
	book := Book{
		Title:   gofakeit.Phrase(),
		Authors: []string{gofakeit.Name(), gofakeit.Name()},
		Prices: map[string]int{
			gofakeit.BeerName(): gofakeit.Number(100, 1000),
			gofakeit.Drink():    gofakeit.Number(2000, 6000),
			gofakeit.Dog():      gofakeit.Number(5000, 10000),
		},
		ReadStatus: gofakeit.RandomString([]string{"Unstarted", "In progress", "Completed"}),
	}

	if bought {
		book.BoughtType = gofakeit.RandomString([]string{"Kindle", "Physical", "Logos", "Crossway"})
		book.BoughtDate = null.TimeFrom(gofakeit.Date().Truncate(time.Second))
	}

	return book
}
