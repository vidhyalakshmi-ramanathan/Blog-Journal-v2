package main

import (
	"backend/db"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/routes"
	"backend/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	// PORT SETUP
	port := os.Getenv("PORT")
	log.Printf("Starting up on http://localhost:%s", port)

	db.ConnectDB()
	userModel := models.NewUserModel(db.DB)
	postModel := models.NewPostModel(db.DB)

	userService := services.NewUserService(userModel)
	postService := services.NewPostService(postModel)

	postHandler := handlers.NewPostHandler(postService)
	userHandler := handlers.NewUserHandler(userService)

	// ROUTER SETUP
	r := routes.SetupRouter(userHandler, postHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
