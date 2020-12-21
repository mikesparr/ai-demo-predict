package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mikesparr/ai-demo-predict/cache"
	"github.com/mikesparr/ai-demo-predict/message"
	"net/http"
)

var client cache.Client
var producer message.Producer

func NewHandler(c cache.Client, p message.Producer) http.Handler {
	router := chi.NewRouter()
	client = c
	producer = p
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/batches", batches)
	router.Route("/jobs", jobs)
	return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
