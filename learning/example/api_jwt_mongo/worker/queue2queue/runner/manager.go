package runner

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	msgQueue "testProject/learning/example/api_jwt_mongo/driver/redis"
	"testProject/learning/example/api_jwt_mongo/worker/queue2queue/models"
	"testProject/learning/example/api_jwt_mongo/worker/queue2queue/utils"
	"time"
)

type Boss interface {
	Monitor()
	Notify()
	Listening()
	HandleJob()
}

type Manager struct {
	ctx      context.Context
	cancelFn context.CancelFunc

	WkMap    map[string]*Worker // {tenant: wk}
	NotifyCh chan models.NotifyMsg
	// StopCh   chan bool
	m sync.Mutex
}

func NewManager() *Manager {
	ctx, cancelFn := context.WithCancel(context.TODO())

	return &Manager{
		ctx:      ctx,
		cancelFn: cancelFn,

		WkMap:    make(map[string]*Worker),
		NotifyCh: make(chan models.NotifyMsg, 10),
		// StopCh:   make(chan bool),
	}
}

// run when init
// query all running in db
// check is <id> being processing?
// N -> send to queue (set) for worker
// Y -> skip
func (obj *Manager) Monitor() {
	lst, err := utils.QueryRunning()
	if err != nil {
		log.Println(err)
		return
	}

	for _, r := range lst {

		if obj.isProcessing(r.TenantId, r.RuleId) {
			continue
		}

		if r.Running == nil || len(*r.Running) == 0 {
			continue
		}

		job := models.JobInQueue{
			RuleId:   r.RuleId,
			TenantId: r.TenantId,
		}

		obj.HandleJob(job)
	}
}

func (obj *Manager) getWorker(tenantId string) (wk *Worker, ok bool) {
	// obj.m.Lock()
	// defer obj.m.Unlock()
	wk, ok = obj.WkMap[tenantId]
	return
}

// run new worker if not available
func (obj *Manager) runWorker(tenantId string) {
	obj.m.Lock()
	defer obj.m.Unlock()

	if _, ok := obj.getWorker(tenantId); ok {
		log.Println("worker already running")
		return
	}

	wk := NewWorker(obj.ctx, tenantId, obj.NotifyCh)

	go func(obj *Manager, wk *Worker, tenantId string) {
		wk.Start()

		// free worker after done
		if obj != nil {
			obj.m.Lock()
			delete(obj.WkMap, tenantId)
			obj.m.Unlock()
		}
		wk = nil
	}(obj, wk, tenantId)

	// add worker to map
	obj.WkMap[tenantId] = wk

	return
}

func (obj *Manager) isProcessing(tenantId, ruleId string) bool {
	obj.m.Lock()
	defer obj.m.Unlock()

	if wk, ok := obj.getWorker(tenantId); ok {
		return ruleId == wk.RuleId
	}

	return false
}

// add to set -> wait worker to pop
func (obj *Manager) sendJobToWkQueue(tenantId string, ruleId string) error {
	queue := models.QueueForWk(tenantId)

	_, err := msgQueue.GetInstance().SAdd(queue, ruleId)

	return err
}

func (obj *Manager) Notify() {
	for {
		select {
		case msg := <-obj.NotifyCh:
			// TODO: send notify to socket here
			log.Println(msg)
		default:
		}
	}
}

func (obj *Manager) HandleJob(job models.JobInQueue) (err error) {
	err = obj.sendJobToWkQueue(job.TenantId, job.RuleId)
	if err != nil {
		log.Printf("Error when send job to worker queue! %v", err)
	}

	obj.runWorker(job.TenantId)
	return
}

func (obj *Manager) Listening() {

	for {
		job, err := popJobFromQueue()
		if err != nil {
			log.Println("No data or error when get target from queue...")
			time.Sleep(time.Second * 30)
			continue
		}

		obj.HandleJob(job)
	}
}

func (obj *Manager) Stop() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[panic] stop manager ! %s\n", r)
		}
	}()
	log.Println("Try stop manager ...")
	// close(obj.StopCh)
	obj.cancelFn()
}

func popJobFromQueue() (job models.JobInQueue, err error) {
	rs, err_ := msgQueue.GetInstance().RPop(models.QueueJob)
	if err_ != nil {
		err = err_
		return
	}

	err = json.Unmarshal([]byte(rs), &job)
	return
}
