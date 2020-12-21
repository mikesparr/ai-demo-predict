package models

import (
	"fmt"
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
}
type JobList struct {
	Jobs []Job `json:"jobs"`
}

func (j *Job) Bind(r *http.Request) error {
	if j.ModelFileName == "" {
		return fmt.Errorf("model_file_name is a required field")
	}
	if j.Records >= 0 {
		return fmt.Errorf("records (count) is a required field")
	}
	if j.Accuracy <= 1.0 {
		return fmt.Errorf("accuracy is a required field")
	}
	return nil
}
func (*JobList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Job) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
