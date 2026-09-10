package router

import (
	"net/http"

	"github.com/Dew-F/v-chat/internal/auth"

	"github.com/go-chi/chi/v5"
)

func New(
	authHandler *auth.Handler,
) http.Handler {

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ok":true}`))
	})

	r.Route("/api", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
		})

	})

	return r
}
