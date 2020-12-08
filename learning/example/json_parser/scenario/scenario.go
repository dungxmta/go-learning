package scenario

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"testProject/learning/example/json_parser/scenario/action"
	"testProject/learning/example/json_parser/scenario/model"
)

const limitLoop = 1

type track struct {
	Name    string `json:"name"`
	ActType string `json:"act_type"`

	RunCount  int `json:"run_count"`
	BackCount int `json:"back_count"`

	Prev string `json:"prev"`
	Next string `json:"next"`

	Pass     map[string]struct{} `json:"pass"`
	OtherWay map[string]struct{} `json:"other_way"`
}

func newTrack() *track {
	return &track{
		Pass:     make(map[string]struct{}),
		OtherWay: make(map[string]struct{}),
	}
}

func (obj *track) isExceedLimit() bool {
	switch obj.ActType {
	case model.Condition:
		return obj.RunCount > limitLoop
	default:
		return obj.RunCount >= limitLoop
	}
}
func (obj *track) incrRun() {
	obj.RunCount++
}
func (obj *track) incrBack() {
	obj.BackCount++
}
func (obj *track) resetBack() {
	obj.BackCount = 0
}
func (obj *track) setNext(s string) {
	obj.Next = s
}
func (obj *track) setPrev(s string) {
	obj.Prev = s
}
func (obj *track) addOthers(lst ...string) {
	for _, v := range lst {
		obj.OtherWay[v] = struct{}{}
	}
}
func (obj *track) addPass(lst ...string) {
	for _, v := range lst {
		if v == "" {
			continue
		}
		if _, ok := obj.OtherWay[v]; ok {
			delete(obj.OtherWay, v)
		}
		obj.Pass[v] = struct{}{}
	}
}

func (obj *track) getBreak() (s string, ok bool) {
	if len(obj.OtherWay) == 0 {
		return
	}

	for k, _ := range obj.OtherWay {
		return k, true
	}

	return
}

type Flow struct {
	Config  model.Config
	StepMap map[string]model.Step

	FirstStep model.Step
	ParamMap  map[string]model.ActionResult
	stepTrack map[string]*track
}

func NewFlow(conf *model.Config) (*Flow, error) {
	var firstStep model.Step
	stepTrack := make(map[string]*track, len(conf.Output.FlowConfig.UpdateStep))

	stepMap := make(map[string]model.Step)
	for _, step := range conf.Output.FlowConfig.UpdateStep {
		stepMap[step.Name] = step
		if step.FirstStep {
			firstStep = step
		}

		tr := newTrack()
		tr.ActType = step.Action
		tr.Name = step.Name
		switch step.Action {
		case model.Condition:
			if step.TrueStep != nil {
				tr.addOthers(*step.TrueStep)
			} else {
				tr.addOthers("")
			}
			if step.FalseStep != nil {
				tr.addOthers(*step.FalseStep)
			} else {
				tr.addOthers("")
			}
			break
		default:
			tr.addOthers(step.NextStep)
		}
		stepTrack[step.Name] = tr
	}

	if firstStep.Action == "" {
		return nil, errors.New("cant find first step")
	}

	return &Flow{
		Config:    *conf,
		StepMap:   stepMap,
		FirstStep: firstStep,
		ParamMap:  nil,
		stepTrack: stepTrack,
	}, nil
}

func (obj *Flow) Start(newData *map[string]interface{}) error {
	// refresh data
	obj.ParamMap = make(map[string]model.ActionResult)
	obj.ParamMap[obj.Config.Input.NewDataName] = model.ActionResult{
		ActionType: model.NewData,
		Value:      *newData,
	}

	// let the show begin
	err := obj.Process(&obj.FirstStep)
	if err != nil {
		log.Print(err)
	}

	b, _ := json.MarshalIndent(obj.stepTrack, "", "   ")
	fmt.Println(string(b))

	return err
}

func (obj *Flow) Process(step *model.Step) error {
	if step == nil {
		log.Println("empty step, done!")
		return nil
	}

	// get action
	actionFn, ok := action.FnMap[step.Action]
	if !ok {
		return fmt.Errorf("invalid action %v", step.Action)
	}

	// check limit run
	tr := obj.stepTrack[step.Name]

	nextStepName, err := actionFn(step, &obj.Config, &obj.ParamMap)
	if err != nil {
		return fmt.Errorf("error when exec action %v, %v", step.Action, err)
	}

	tr.incrRun()
	// tr.Prev =

	if nextStepName == "" {
		log.Println("finished!")
		return nil
	}

	// set current track.Next
	tr.setNext(nextStepName)
	tr.addPass(nextStepName)

	nextStep, ok := obj.StepMap[nextStepName]
	if !ok {
		return fmt.Errorf("invalid step %v", nextStepName)
	}

	// get Next step track
	nextTr := obj.stepTrack[nextStepName]
	nextTr.setPrev(step.Name)

	if nextTr.isExceedLimit() {
		// todo: calc next_step
		ns, cs, err := getOtherWay(tr, &obj.stepTrack)
		if err != nil {
			return fmt.Errorf("limit loop exceed")
		}
		log.Printf("Set new next_step=%v instead of %v to break loop\n", ns, nextStepName)
		nextStep, ok = obj.StepMap[ns]
		if !ok {
			return fmt.Errorf("invalid step %v", ns)
		}

		// get Next step track
		nextTr := obj.stepTrack[ns]
		nextTr.setPrev(cs)
	}

	return obj.Process(&nextStep)
}

func getOtherWay(curTrack *track, stepTrack *map[string]*track) (nextStepName, curStepName string, err error) {
	var ok bool
	curTrack.incrBack()
	fmt.Println("increase back++", curTrack.Name, "->", curTrack.BackCount)

	nextTr := (*stepTrack)[curTrack.Prev]
	nextStepName, ok = nextTr.getBreak()
	if !ok {
		fmt.Printf("no break from %v (prev of %v) ...\n", curTrack.Prev, curTrack.Name)
		nextStepName, curStepName, err = getOtherWay(nextTr, stepTrack)
	} else {
		curStepName = nextTr.Name
		fmt.Printf("found break from %v (prev of %v)\n", curTrack.Prev, curTrack.Name)
	}

	nextTr.incrBack()
	return
}
