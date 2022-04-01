package main

import (
	"log"
	"net/http"

	"github.com/cfabrica46/myrouter/cfrouter"
)

func main() {
	myRouter()
}

func myRouter() {
	router := cfrouter.NewCfRouter()
	router.Methods(http.MethodGet).HandlerFunc(myHandlerFunc1).Path("/1")
	router.Methods(http.MethodPost).HandlerFunc(myHandlerFunc2).Path("/2")
	router.Use(myMiddleware1, myMiddleware2)

	log.Println("Listen And Serve on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func myHandlerFunc1(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("1"))
}

func myHandlerFunc2(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("2"))
}

func myMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 1 :D")
		next.ServeHTTP(w, r)
	})
}

func myMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 2 :D")
		next.ServeHTTP(w, r)
	})
}
