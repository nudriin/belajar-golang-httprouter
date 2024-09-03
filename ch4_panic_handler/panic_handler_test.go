package ch4_panic_handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// ! Membuat panic handler
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Terjadi Panic : ", i)
	}

	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Server sedang error")
	})

	request := httptest.NewRequest("GET", "http://localhost:5000/panic", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
