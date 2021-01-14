package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func jobs(router chi.Router) {
	router.Get("/", getAllJobs)
}

func getAllJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all jobs")
	jobs, err := client.GetAllJobs()
	if err != nil {
		err := render.Render(w, r, ServerErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := render.Render(w, r, jobs); err != nil {
		err := render.Render(w, r, ErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
}
