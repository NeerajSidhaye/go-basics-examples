package main

import (
	"log"
	"net/http"
)

func main() {

	serveMux := http.NewServeMux()

	helloHandler := http.HandlerFunc(hello)

	// registers "helloHandler" function for the given pattern (/hello)
	// and also attach middleware to helloHandler.
	serveMux.Handle("/hello", logReq(helloHandler))

	log.Println("Server is listening at port 7070")
	http.ListenAndServe(":7070", serveMux)
	
}

// application handler
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Original handler called")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello there - after middleware execution"))
}

// middlewar function
func logReq(next http.Handler) http.Handler {  

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		log.Println("URL called -", r.URL ," ,before calling original handler")
		
		next.ServeHTTP(w, r) // calling original handler ( helloHandler)
		
		log.Println("After calling original handler")
	})
}
