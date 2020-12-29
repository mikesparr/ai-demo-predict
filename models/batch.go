package models

import (
	"fmt"
	"net/http"
)

type Batch struct {
	BatchID     string   `json:"batch_id"`
	Subjects    []string `json:"subjects"`
	Predictions []string `json:"predictions"`
	Created     string   `json:"created,omitempty"`
}
type BatchList struct {
	Batches []Batch `json:"batches"`
}
type BatchFeedback struct {
	BatchID  string   `json:"batch_id,omitempty"` // inject after request to send to pubsub
	Subjects []string `json:"subjects"`
	Ratings  []int    `json:"ratings"`
}

func (bf *BatchFeedback) Bind(r *http.Request) error {
	if bf.BatchID == "" {
		return fmt.Errorf("batch_id is a required field")
	}
	if len(bf.Subjects) <= 0 {
		return fmt.Errorf("subjects must have one or more record")
	}
	if len(bf.Ratings) != len(bf.Subjects) {
		return fmt.Errorf("ratings count must equal subjects")
	}
	for _, rating := range bf.Ratings {
		if !(rating == 1 || rating == 0) {
			return fmt.Errorf("ratings must be a 1 or 0 value")
		}
	}
	return nil
}
func (*BatchFeedback) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BatchList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Batch) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
