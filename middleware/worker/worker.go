package worker

import "log"

type JobInput struct {
	ID   string
	Data []byte
}

type Job interface {
	Serve([]byte)
}

func NewJobQueue(maxQueue int) *JobQueue {
	return &JobQueue{
		queue:     make(chan JobInput, maxQueue),
		quitChan:  make(chan bool),
		jobRouter: make(map[string]Job),
	}
}

type JobQueue struct {
	queue     chan JobInput
	quitChan  chan bool
	jobRouter jobRouter
}

func (jq JobQueue) Route(id string, job Job) {
	jq.jobRouter[id] = job
}

func (jq JobQueue) Push(input JobInput) error {
	id := input.ID
	if _, ok := jq.jobRouter[id]; ok {
		jq.queue <- input
	}

	return nil
}

func (jq JobQueue) Start() {
	go func() {
		for {
			select {
			case input := <-jq.queue:
				i, d := input.ID, input.Data
				job := jq.jobRouter[i]

				log.Print("Exec job")
				job.Serve(d)

			case <-jq.quitChan:
				log.Print("Stop Worker")
				return
			}
		}
	}()
}

func (jq JobQueue) Stop() {
	go func() {
		jq.quitChan <- true
	}()
}

type jobRouter map[string]Job
