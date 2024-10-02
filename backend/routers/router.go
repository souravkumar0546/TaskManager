package routers

import (
	"net/http"
	"task-manager-backend/controllers"
	"task-manager-backend/middlewares"

	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
func InitRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	api.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	api.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	api.HandleFunc("/user", controllers.GetUserProfile).Methods("GET")
	api.HandleFunc("/user/avatar", controllers.UpdateAvatar).Methods("POST")

	// Serve static files from ./user_data/avatar/ under /avatars/
	avatarsDir := "./user_data/avatar/"
	avatarsHandler := http.StripPrefix("/avatars/", http.FileServer(http.Dir(avatarsDir)))
	router.PathPrefix("/avatars/").Handler(avatarsHandler)

	// CORS setup
	cors := handlers.CORS(
		handlers.AllowedOrigins(getAllowedOrigins()),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	// return router
	return cors(router)
}

// func getAllowedOrigins() []string {
// 	return []string{os.Getenv("APP_URL")}
// }
func getAllowedOrigins() []string {
    appUrl := os.Getenv("APP_URL")  // Get the live backend URL from environment variables
    return []string{
        appUrl,                     // Live backend URL
        "http://localhost:3000",     // Localhost for React, Vue, etc. (adjust the port as needed)
        "http://127.0.0.1:3000",     // Localhost IP address (alternative to localhost)
    }
}
