package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/*
func EnableCorsMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w *http.ResponseWriter, r *http.Request) {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Headers", "*")
		log.Println("Executing middlewareOne")
		handler.ServeHTTP(*w, r)
	})
}
*/

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next.ServeHTTP(w, r)
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the index.html page from public folder
	http.ServeFile(w, r, "./public/index.html")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(CorsMiddleware)
	Routes.HandleRoutes(router, "/api")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func main() {
	handleRequests()
}
