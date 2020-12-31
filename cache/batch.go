package cache

import (
	"encoding/json"
	"fmt"
	"github.com/mikesparr/ai-demo-predict/models"

	"github.com/gomodule/redigo/redis"
)

func (client Client) GetAllBatches() (*models.BatchList, error) {
	fmt.Println("Fetching batches of predictions !!!")

	conn := client.Pool.Get()
	list := &models.BatchList{}
	const batchKey = "batches" // demo purpose only
	defer conn.Close()

	batches, err := redis.Strings(conn.Do("LRANGE", batchKey, 0, -1))
	if err != nil {
		fmt.Println("Error retrieving batches list from cache")
		return list, err
	}
	for i, b := range batches {
		fmt.Println(i, b)

		batch := models.Batch{}
		err := json.Unmarshal([]byte(b), &batch)
		if err != nil {
			fmt.Println("Error unmarshalling json object")
			return list, err
		}

		list.Batches = append(list.Batches, batch)
	}

	return list, nil
}
