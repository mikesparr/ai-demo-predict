package models

import (
	"fmt"
	"net/http"
)

type Batch struct {
	BatchID     string   `json:"batch_id"`
	Subjects    []string `json:"subjects"`
	Predictions []string `json:"predictions"`
}
type BatchList struct {
	Batches []Batch `json:"batches"`
}
type BatchResponse struct {
	BatchID  string   `json:"batch_id,omitempty"` // inject after request to send to pubsub
	Subjects []string `json:"subjects"`
	Ratings  []int    `json:"ratings"`
}

func (i *BatchResponse) Bind(r *http.Request) error {
	if i.BatchID == "" {
		return fmt.Errorf("batch_id is a required field")
	}
	if len(i.Subjects) <= 0 {
		return fmt.Errorf("subjects must have one or more record")
	}
	if len(i.Ratings) != len(i.Subjects) {
		return fmt.Errorf("ratings count must equal subjects")
	}
	return nil
}
func (*BatchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BatchList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Batch) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
