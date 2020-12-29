package models

import (
	"net/http"
)

type Job struct {
	ID            string  `json:"job_id"`
	ModelFileName string  `json:"model_file_name"`
	Records       int     `json:"records"`
	Accuracy      float64 `json:"accuracy"`
	DataPrepTime  float64 `json:"data_prep_time,omitempty"`
	TrainingTime  float64 `json:"training_time,omitempty"`
	TestingTime   float64 `json:"testing_time,omitempty"`
	Created       string  `json:"created,omitempty"`
}
type JobList struct {
	Jobs []Job `json:"jobs"`
}

func (*JobList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Job) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
