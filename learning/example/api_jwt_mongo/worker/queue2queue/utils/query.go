package utils

import (
	"log"
	"testProject/learning/example/api_jwt_mongo/worker/queue2queue/models"
)

func Query(tenantId, ruleId string) (rs models.Rule, err error) {
	rs = models.Rule{
		RuleId:   "1",
		TenantId: "1",
		Running: &[]models.JobDetail{
			models.JobDetail{
				Conf: "1",
			},
		},
	}
	return
}

func QueryRunning() (rs []models.Rule, err error) {
	rs = []models.Rule{
		models.Rule{
			RuleId:   "1",
			TenantId: "1",
			Running: &[]models.JobDetail{
				models.JobDetail{
					Conf: "1",
				},
			},
		},
		models.Rule{
			RuleId:   "2",
			TenantId: "2",
			Running: &[]models.JobDetail{
				models.JobDetail{
					Conf: "2",
				},
			},
		},
	}
	return
}

func UpdateStatus(id string, status int32) {
	log.Printf("update status=%v id=%v", status, id)
}
