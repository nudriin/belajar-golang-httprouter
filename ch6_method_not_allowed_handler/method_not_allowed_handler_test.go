package ch6_method_not_allowed_handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowedHandler(t *testing.T) {
	router := httprouter.New()

	// ! Membuat method not allowed handler
	// ! Apabila method yang di gunakan berbeda dan tidak match maka function ini akan di eksekusi
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Method ", r.Method, " Tidak di izinkan untuk handler ini")
	})

	router.GET("/halo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hallo")
	})

	request := httptest.NewRequest("POST", "http://localhost:5000/halo", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Method POST Tidak di izinkan untuk handler ini", string(body))
	fmt.Println(string(body))
}
