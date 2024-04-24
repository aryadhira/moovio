package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	APIBaseURL := os.Getenv("SERVICE_BASE_URL")

	// Define Service URL
	GrabberServiceURL, _ := url.Parse(APIBaseURL + ":9002")
	MovieServiceURL, _ := url.Parse(APIBaseURL + ":9003")
	StreamerServiceURL, _ := url.Parse(APIBaseURL + ":9004")

	// Define Service Reverse Proxy
	GrabberServiceProxy := httputil.NewSingleHostReverseProxy(GrabberServiceURL)
	MovieServiceProxy := httputil.NewSingleHostReverseProxy(MovieServiceURL)
	StreamerServiceProxy := httputil.NewSingleHostReverseProxy(StreamerServiceURL)

	// Define route handlers
	http.Handle("/", loggingMiddleware(http.DefaultServeMux))
	http.HandleFunc("/grabber/", func(w http.ResponseWriter, r *http.Request) {
		// r.URL.Path = strings.TrimPrefix(r.URL.Path, "/extractor")
		GrabberServiceProxy.ServeHTTP(w, r)
	})
	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		// r.URL.Path = strings.TrimPrefix(r.URL.Path, "/extractor")
		MovieServiceProxy.ServeHTTP(w, r)
	})
	http.HandleFunc("/streamer/", func(w http.ResponseWriter, r *http.Request) {
		// r.URL.Path = strings.TrimPrefix(r.URL.Path, "/extractor")
		StreamerServiceProxy.ServeHTTP(w, r)
	})

	log.Println("Starting API Gateway on port 9001...")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Forward the request to the next handler
	})
}
