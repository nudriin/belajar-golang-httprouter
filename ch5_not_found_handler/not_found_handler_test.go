package ch5_not_found_handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()

	// ! Membuat not found handler
	// ! Apablia status code error 404 atau file tidak ketemu maka function ini akan di eksekusi
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "404 Data tidak ada")
	})

	request := httptest.NewRequest("GET", "http://localhost:5000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "404 Data tidak ada", string(body))
	fmt.Println(string(body))
}
