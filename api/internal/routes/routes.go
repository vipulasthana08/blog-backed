package routes

import (
	blogcontroller "blog-backend/api/internal/controller"
	"log"
	"net/http"
)

// RegisterRoutes sets up the HTTP routes for the application and starts the server.
// It registers a handler for the "/home" endpoint, which responds with a simple "hello" message.
// The server listens on port 8080. If the server fails to start, it logs a fatal error.
func RegisterRoutes() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	http.Handle("/new", corsMiddleware(http.HandlerFunc(blogcontroller.CreateNewBlog)))
	http.Handle("/check", corsMiddleware(http.HandlerFunc(blogcontroller.CheckBlogExist)))
	http.Handle("/get", corsMiddleware(http.HandlerFunc(blogcontroller.GetBlog)))
	http.Handle("/delete", corsMiddleware(http.HandlerFunc(blogcontroller.DeleteBlog)))

	log.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
