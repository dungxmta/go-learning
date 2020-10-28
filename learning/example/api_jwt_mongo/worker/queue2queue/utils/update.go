package utils

import (
	"log"
	"testProject/learning/example/api_jwt_mongo/worker/queue2queue/models"
	"time"
)

func Update(job models.JobDetail) {
	time.Sleep(time.Second * 5)
	log.Println("update with job:", job)
	return
}
