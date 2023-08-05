package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/cpglsn/readinglist/internal/data"
)

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"Status":      "available",
		"Environment": app.config.environment,
		"Version":     version,
	}

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *Application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		books := []data.Book{
			{
				ID:        2,
				CreatedAt: time.Now(),
				Title:     "Vendetta",
				Published: 2022,
				Pages:     22,
				Genres:    []string{"Action", "Thriller"},
				Rating:    2.2,
				Version:   2,
			},
			{
				ID:        3,
				CreatedAt: time.Now(),
				Title:     "La Tela di Carlotta",
				Published: 2033,
				Pages:     33,
				Genres:    []string{"Fiabe", "Thriller"},
				Rating:    3.3,
				Version:   3,
			},
		}

		if err := WriteJSON(w, envelop{"books": books}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {

		var input struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float32  `json:"rating"`
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%v\n", input)

		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (app *Application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		app.getBook(w, r)

	case http.MethodPut:
		app.updateBook(w, r)

	case http.MethodDelete:
		app.deleteBook(w, r)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *Application) getBook(w http.ResponseWriter, r *http.Request) {
	tmpID := r.URL.Path[len("/vi/books/"):]
	id, err := strconv.ParseInt(tmpID, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Bello Libro",
		Published: 2019,
		Pages:     33,
		Genres:    []string{"Fiction", "Thriller"},
		Rating:    4.5,
		Version:   1,
	}

	if err := WriteJSON(w, envelop{"book": book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *Application) updateBook(w http.ResponseWriter, r *http.Request) {
	tmpID := r.URL.Path[len("/vi/books/"):]
	id, err := strconv.ParseInt(tmpID, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Update details of book with id: %d", id)
}

func (app *Application) deleteBook(w http.ResponseWriter, r *http.Request) {
	tmpID := r.URL.Path[len("/vi/books/"):]
	id, err := strconv.ParseInt(tmpID, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Delete details of book with id: %d", id)
}
