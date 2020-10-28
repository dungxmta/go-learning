package models

import "fmt"

type JobInQueue struct {
	RuleId   string `json:"rule_id"`
	TenantId string `json:"tenant_id"`
}

type Rule struct {
	RuleId   string       `json:"rule_id"`
	TenantId string       `json:"tenant_id"`
	Status   int32        `json:"status"`
	Running  *[]JobDetail `json:"running"`
}

// type JobForWorker struct {
// 	RuleId   string      `json:"rule_id"`
// 	TenantId string      `json:"tenant_id"`
// 	Jobs     []JobDetail `json:"jobs"`
// }

// input of update func
type JobDetail struct {
	Conf string
}

const (
	QueueJob  = "test"
	WkTimeout = 10 // minute
)

func QueueForWk(tenantId string) string {
	return fmt.Sprintf("%v_%v", QueueJob, tenantId)
}

const (
	IDLE    = 0
	RUNNING = 1
)

// notify
type NotifyMsg struct {
	RuleId   string `json:"rule_id"`
	TenantId string `json:"tenant_id"`
	Status   int32  `json:"status"`
}
