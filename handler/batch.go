package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mikesparr/ai-demo-predict/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type key int

const batchIDPathKey = "batchId"
const batchCtxKey key = iota

func batches(router chi.Router) {
	router.Get("/", getAllBatches)
	router.Route("/{batchId}", func(router chi.Router) {
		router.Use(BatchContext)
		router.Patch("/", updateBatch) // add is_correct to one or more predictions in batch
	})
}

// BatchContext handle input parameters
func BatchContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		batchID := chi.URLParam(r, batchIDPathKey)
		if batchID == "" {
			err := render.Render(w, r, ErrorRenderer(fmt.Errorf("batch ID is required")))
			if err != nil {
				fmt.Println("Error rendering")
			}
			return
		}
		ctx := context.WithValue(r.Context(), batchCtxKey, batchID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBatches(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all batches")
	batches, err := client.GetAllBatches()
	if err != nil {
		err := render.Render(w, r, ServerErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := render.Render(w, r, batches); err != nil {
		err := render.Render(w, r, ErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
}

func updateBatch(w http.ResponseWriter, r *http.Request) {
	batchID := r.Context().Value(batchCtxKey)
	fmt.Printf("Updating batch (%s) with prediction ratings\n", batchID)

	feedback := &models.BatchFeedback{}

	// assert batchId is string and inject into feedback
	if id, ok := batchID.(string); ok {
		feedback.BatchID = id
	} else {
		err := render.Render(w, r, ErrorRenderer(fmt.Errorf("batch ID must be string")))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := render.Bind(r, feedback); err != nil {
		err := render.Render(w, r, ErrorBadRequest(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := producer.UpdateBatch(feedback); err != nil {
		err := render.Render(w, r, ErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := render.Render(w, r, feedback); err != nil {
		err := render.Render(w, r, ServerErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
}
