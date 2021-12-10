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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Products " + id + " Item " + itemId
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Products 1 Item 1", string(body))
}
func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}