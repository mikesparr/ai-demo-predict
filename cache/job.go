package cache

import (
	"encoding/json"
	"fmt"
	"github.com/mikesparr/ai-demo-predict/models"

	"github.com/gomodule/redigo/redis"
)

func (client Client) GetAllJobs() (*models.JobList, error) {
	fmt.Println("Fetching training jobs !!!")

	conn := client.Pool.Get()
	list := &models.JobList{}
	const jobsKey = "jobs" // demo purpose only
	defer conn.Close()

	jobs, err := redis.Strings(conn.Do("LRANGE", jobsKey, 0, 24)) // most recent 25
	if err != nil {
		fmt.Println("Error retrieving jobs list from cache")
		return list, err
	}
	for i, b := range jobs {
		fmt.Println(i, b)

		job := models.Job{}
		err := json.Unmarshal([]byte(b), &job)
		if err != nil {
			fmt.Println("Error unmarshalling json object")
			return list, err
		}

		list.Jobs = append(list.Jobs, job)
	}

	return list, nil
}
