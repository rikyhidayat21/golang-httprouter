package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr: "localhost:4000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}