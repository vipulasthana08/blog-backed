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
	http.HandleFunc("/new", blogcontroller.CreateNewBlog)
	http.HandleFunc("/check", blogcontroller.CheckBlogExist)
	http.HandleFunc("/get", blogcontroller.GetBlog)
	http.HandleFunc("/delete", blogcontroller.DeleteBlog)

	log.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
