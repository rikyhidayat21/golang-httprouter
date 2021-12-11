package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func (w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic: ", error)
	}

	router.GET("/", func (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// fmt.Fprint(w, "Hello World")
		panic("Ups")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Panic: Ups", string(body))
}