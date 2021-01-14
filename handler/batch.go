package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mikesparr/ai-demo-predict/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var batchIDKey = "batchId"

func batches(router chi.Router) {
	router.Get("/", getAllBatches)
	router.Route("/{batchId}", func(router chi.Router) {
		router.Use(BatchContext)
		router.Patch("/", updateBatch) // add is_correct to one or more predictions in batch
	})
}

func BatchContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		batchId := chi.URLParam(r, "batchId")
		if batchId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("batch ID is required")))
			return
		}
		ctx := context.WithValue(r.Context(), batchIDKey, batchId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBatches(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all batches")
	batches, err := client.GetAllBatches()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, batches); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func updateBatch(w http.ResponseWriter, r *http.Request) {
	batchId := r.Context().Value(batchIDKey)
	fmt.Printf("Updating batch (%s) with prediction ratings\n", batchId)

	feedback := &models.BatchFeedback{}

	// assert batchId is string and inject into feedback
	if id, ok := batchId.(string); ok {
		feedback.BatchID = id
	} else {
		err := render.Render(w, r, ErrorRenderer(fmt.Errorf("batch ID must be string")))
		if err != nil {
			fmt.Printf("Error rendering", err)
		}
		return
	}
	if err := render.Bind(r, feedback); err != nil {
		err := render.Render(w, r, ErrorBadRequest(err))
		if err != nil {
			fmt.Printf("Error rendering", err)
		}
		return
	}
	if err := producer.UpdateBatch(feedback); err != nil {
		err := render.Render(w, r, ErrorRenderer(err))
		if err != nil {
			fmt.Printf("Error rendering", err)
		}
		return
	}
	if err := render.Render(w, r, feedback); err != nil {
		err := render.Render(w, r, ServerErrorRenderer(err))
		if err != nil {
			fmt.Printf("Error rendering", err)
		}
		return
	}
}
