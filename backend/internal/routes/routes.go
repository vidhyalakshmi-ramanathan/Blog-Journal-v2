package routes

import (
	"backend/internal/handlers"
	"backend/middleware"

	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router, userHandler *handlers.UserHandler) {
	r.Post("/signup", userHandler.CreateUser)
	r.Post("/login", userHandler.LoginUser)
}

func PostRoutes(r chi.Router, postHandler *handlers.PostHandler) {
	r.Get("/posts", postHandler.GetAllPosts)
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/userDetails", postHandler.GetUserDetails)
		r.Route("/post", func(r chi.Router) {
			r.Post("/", postHandler.CreatePost)
			r.Get("/", postHandler.GetPosts)
			r.Get("/{id}", postHandler.GetPost)
			r.Put("/{id}", postHandler.UpdatePost)
			r.Delete("/{id}", postHandler.DeletePost)
		})
	})
}
