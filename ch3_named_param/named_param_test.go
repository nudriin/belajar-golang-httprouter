package ch3_named_param

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func HandleNamedParam(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	itemId := param.ByName("itemId")

	fmt.Fprint(writer, "Products id is ", id, " and item id is ", itemId)
}

func HandleFullpathParam(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	path := param.ByName("path")

	fmt.Fprint(writer, "File path is : ", path)
}

func TestNamedParam(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/items/:itemId", HandleNamedParam)
	router.GET("/file/*path", HandleFullpathParam)

	// ! TEST NAMED PARAM
	request := httptest.NewRequest("GET", "http://localhost:5000/products/2/items/24", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Products id is 2 and item id is 24", string(body))

	// ! TEST FULLPATH PARAM
	request = httptest.NewRequest("GET", "http://localhost:5000/file/products/2/items/24", nil)
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response = recorder.Result()
	body, _ = io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(t, "File path is : /products/2/items/24", string(body))

}
