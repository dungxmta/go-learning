package runner

import (
	"context"
	"log"
	"sync"
	msgQueue "testProject/learning/example/api_jwt_mongo/driver/redis"
	"testProject/learning/example/api_jwt_mongo/worker/queue2queue/models"
	"testProject/learning/example/api_jwt_mongo/worker/queue2queue/utils"
	"time"
)

type Worker struct {
	ctx      context.Context
	TenantId string
	RuleId   string
	queue    string
	NotifyCh chan<- models.NotifyMsg

	m sync.Mutex
}

func NewWorker(ctx context.Context, tenantId string, notifyCh chan<- models.NotifyMsg) *Worker {
	return &Worker{
		ctx:      ctx,
		TenantId: tenantId,
		queue:    models.QueueForWk(tenantId),
		NotifyCh: notifyCh,
	}
}

func (obj *Worker) Start() (err error) {
	log.Printf("[%v] Start worker", obj.TenantId)
	defer log.Printf("[%v] Done worker", obj.TenantId)

	timeout := time.After(time.Minute * models.WkTimeout)

	for {
		select {
		case <-obj.ctx.Done():
			return
		case <-timeout:
			return
		default:
		}

		ruleId, jobs, err := popJobs(obj.TenantId, obj.queue)
		if err != nil {
			log.Println("No data or error when get target from queue...")
			time.Sleep(time.Second * 30)
			continue
		}

		obj.m.Lock()
		obj.RuleId = ruleId
		obj.m.Lock()

		utils.UpdateStatus(obj.RuleId, models.RUNNING)
		obj.NotifyCh <- models.NotifyMsg{
			RuleId:   ruleId,
			TenantId: obj.TenantId,
			Status:   models.RUNNING,
		}

		for _, job := range jobs {
			utils.Update(job)
		}

		utils.UpdateStatus(obj.RuleId, models.IDLE)
		obj.NotifyCh <- models.NotifyMsg{
			RuleId:   ruleId,
			TenantId: obj.TenantId,
			Status:   models.IDLE,
		}

		obj.m.Lock()
		obj.RuleId = ""
		obj.m.Lock()

		// refresh timeout if has data
		timeout = time.After(time.Minute * models.WkTimeout)
	}
}

func popJobs(tenantId, queue string) (ruleId string, jobs []models.JobDetail, err error) {
	ruleId, err_ := msgQueue.GetInstance().SPop(queue)
	if err_ != nil {
		return "", nil, err
	}

	r, err := utils.Query(tenantId, ruleId)
	if err != nil {
		return ruleId, nil, err
	}

	if r.Running == nil || len(*r.Running) == 0 {
		return
	}

	return ruleId, *r.Running, nil
}
