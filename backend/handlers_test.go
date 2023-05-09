package main

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/leeyenter/books/backend/books"
	"gopkg.in/guregu/null.v4"
	"time"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Books Handlers", func() {
	var boughtBook books.Book
	var unboughtBook books.Book

	BeforeEach(func() {
		boughtBook = books.CreateRandom(true)
		unboughtBook = books.CreateRandom(false)

		_ = books.Add(context.Background(), &boughtBook)
		_ = books.Add(context.Background(), &unboughtBook)
	})

	It("gets books", func() {
		r, err := http.NewRequest("GET", "/book/", nil)
		Expect(err).To(BeNil())

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		Expect(w.Code).To(Equal(http.StatusOK))

		var resp []books.Book
		err = json.NewDecoder(w.Body).Decode(&resp)
		Expect(err).To(BeNil())
		Expect(resp).To(ContainElement(boughtBook))
		Expect(resp).To(ContainElement(unboughtBook))
	})

	It("adds book", func() {
		newBook := books.CreateRandom(true)
		body, _ := json.Marshal(newBook)
		r, err := http.NewRequest("POST", "/book/", bytes.NewReader(body))
		Expect(err).To(BeNil())

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusOK))

		var resp map[string]string
		err = json.NewDecoder(w.Body).Decode(&resp)
		Expect(err).To(BeNil())
		newBook.ID = resp["id"]

		booksFound, err := books.GetAll(context.Background())
		Expect(err).To(BeNil())
		Expect(booksFound).To(ContainElement(boughtBook))
		Expect(booksFound).To(ContainElement(newBook))
	})

	It("updates book", func() {
		editedBook := books.CreateRandom(true)
		body, _ := json.Marshal(editedBook)
		r, err := http.NewRequest("POST", "/book/"+unboughtBook.ID, bytes.NewReader(body))
		Expect(err).To(BeNil())

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusOK))

		editedBook.ID = unboughtBook.ID
		booksFound, err := books.GetAll(context.Background())
		Expect(err).To(BeNil())
		Expect(booksFound).To(ContainElement(editedBook))
		Expect(booksFound).ToNot(ContainElement(unboughtBook))
	})

	It("sets price for book", func() {
		source := gofakeit.Vegetable()
		price := gofakeit.Number(100, 10000)
		unboughtBook.Prices[source] = price

		body, _ := json.Marshal(gin.H{"source": source, "price": price})
		r, _ := http.NewRequest("POST", "/book/"+unboughtBook.ID+"/price", bytes.NewReader(body))

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusOK))

		comparePrices(unboughtBook, source)
	})

	It("removes price from book", func() {
		var removedSource string
		for k, _ := range unboughtBook.Prices {
			delete(unboughtBook.Prices, k)
			removedSource = k
			break
		}

		body, _ := json.Marshal(gin.H{"source": removedSource})
		r, _ := http.NewRequest("DELETE", "/book/"+unboughtBook.ID+"/price", bytes.NewReader(body))

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusOK))

		comparePrices(unboughtBook, "")
	})

	It("marks book as bought", func() {
		unboughtBook.BoughtDate = null.TimeFrom(time.Now())
		unboughtBook.BoughtType = gofakeit.Breakfast()

		body, _ := json.Marshal(gin.H{
			"boughtType": unboughtBook.BoughtType,
		})
		r, _ := http.NewRequest("POST", "/book/"+unboughtBook.ID+"/bought", bytes.NewReader(body))

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusOK))

		booksFound, err := books.GetAll(context.Background())
		Expect(err).To(BeNil())
		for _, book := range booksFound {
			if book.ID != unboughtBook.ID {
				Expect(unboughtBook.BoughtType).ToNot(Equal(book.BoughtType))
				continue
			}

			Expect(unboughtBook.Title).To(Equal(book.Title))
			Expect(unboughtBook.Authors).To(Equal(book.Authors))
			Expect(unboughtBook.BoughtType).To(Equal(book.BoughtType))
			Expect(book.BoughtDate.Time.Sub(unboughtBook.BoughtDate.Time).Seconds()).To(BeNumerically("~", 0, 1))
		}
	})

	It("sets book status", func() {
		boughtBook.ReadStatus = gofakeit.PhrasePreposition()
		body, _ := json.Marshal(gin.H{
			"status": boughtBook.ReadStatus,
		})
		r, _ := http.NewRequest("PUT", "/book/"+unboughtBook.ID+"/status", bytes.NewReader(body))

		router := setupRouter()
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusOK))

		booksFound, err := books.GetAll(context.Background())
		Expect(err).To(BeNil())
		for _, book := range booksFound {
			if book.ID != unboughtBook.ID {
				Expect(book.ReadStatus).ToNot(Equal(boughtBook.ReadStatus))
				continue
			}

			Expect(book.ReadStatus).To(Equal(boughtBook.ReadStatus))
		}
	})
})

func comparePrices(unboughtBook books.Book, source string) {
	booksFound, err := books.GetAll(context.Background())
	Expect(err).To(BeNil())
	for _, book := range booksFound {
		if book.ID != unboughtBook.ID {
			if source != "" {
				_, found := book.Prices[source]
				Expect(found).To(BeFalse())
			}
			continue
		}

		Expect(book.Prices).To(HaveLen(len(unboughtBook.Prices)))

		for expSource, expPrice := range unboughtBook.Prices {
			foundPrice, found := book.Prices[expSource]
			Expect(found).To(BeTrue())
			Expect(foundPrice).To(Equal(expPrice))
		}
	}
}
