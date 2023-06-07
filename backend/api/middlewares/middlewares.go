package middlewares

import (	
	"net/http"	
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// allowedDomains := "http://127.0.0.1:5173"
		// allowedHeaders := "Accept, content-type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
		allowedMethods := "POST, GET, OPTIONS, PUT, DELETE"
		allowedDomains := "*"		
		allowedHeaders := "*"		
		//allowedMethods := "*"		
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)		
		w.Header().Set("Access-Control-Allow-Origin", allowedDomains)

		if r.Method == "OPTIONS" {
			// Set the necessary headers to allow the DELETE method and other required headers
			w.Header().Set("Allow", "DELETE")
			w.Header().Set("Access-Control-Allow-Methods", "DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
			// Respond with a success status code
			w.WriteHeader(http.StatusOK)
			return
		}				

		next(w, r)
	}
}

