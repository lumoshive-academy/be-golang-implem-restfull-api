package router

import (
	"net/http"
	"session-18/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api/v1", Apiv1(handler))
	r.Mount("/api/v2", Apiv2(handler))

	// //menu
	// r.Route("/user", func(r chi.Router) {
	// 	r.Use(middleware.AuthMiddleware)
	// 	r.Get("/assignments", handler.AssignmentHandler.List)
	// 	r.Get("/success-submit", handler.AssignmentHandler.SuccessSubmit)
	// 	r.Post("/submit-assignment", handler.AssignmentHandler.SubmitAssignment)
	// 	r.Get("/grade", handler.HandlerMenu.GradeView)
	// 	r.Get("/logout", handler.HandlerAuth.LogoutView)
	// })
	// r.Get("/page401", handler.HandlerMenu.PageUnauthorized)

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	return r
}

func Apiv1(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	//authentication
	r.Post("/login", handler.HandlerAuth.Login)
	r.Post("/logout", handler.HandlerAuth.Logout)

	r.Route("/assignment", func(r chi.Router) {
		r.Get("/", handler.AssignmentHandler.List)
		r.Post("/", handler.AssignmentHandler.Create)
		r.Route("/{assignment_id}", func(r chi.Router) {
			r.Get("/", handler.AssignmentHandler.GetByID)
			r.Put("/", handler.AssignmentHandler.Update)
			r.Delete("/", handler.AssignmentHandler.Delete)
		})
	})

	return r
}

func Apiv2(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/assignment", func(r chi.Router) {
		r.Post("/", handler.AssignmentHandler.Create)
	})

	return r
}
