package ch1_router

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

// ! Membuat hanlder untuk http router
func HandlerRouter(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprint(writer, "Hello, this is home")
}

func TestRouter(t *testing.T) {
	// ! Membuat handler router
	router := httprouter.New()

	router.GET("/home", HandlerRouter)

	server := http.Server{
		Handler: router,
		Addr:    "localhost:5000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRouterUnitTest(t *testing.T) {
	router := httprouter.New()

	router.GET("/", HandlerRouter)

	request := httptest.NewRequest("GET", "http://localhost:5000/", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP menggunakan Router
	router.ServeHTTP(recorder, request)

	res := recorder.Result()

	body, _ := io.ReadAll(res.Body)

	require.Equal(t, "Hello, this is home", string(body))
	fmt.Println("Hallo")
}
