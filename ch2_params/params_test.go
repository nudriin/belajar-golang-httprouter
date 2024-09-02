package ch2_params

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func HandlerParam(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	category := param.ByName("category") //Get param category
	id := param.ByName("id")             //Get param id

	fmt.Fprint(writer, "Products categories : ", category, " Products id : ", id)
}

func TestParams(t *testing.T) {
	router := httprouter.New()

	// ! Membuat params di url
	router.GET("/products/:category/:id", HandlerParam)

	req := httptest.NewRequest("GET", "http://localhost/products/books/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Products categories : books Products id : 1", string(body))

}
