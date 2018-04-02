package worker

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

const (
	maxQueue = 10
	id       = "sample"
	testData = "test"
)

func TestWorker(t *testing.T) {
	// Create JobQueue
	jq := NewJobQueue(maxQueue)

	// Start JobQueue
	jq.Start()
	defer jq.Stop()

	// Register "sample" routing
	sample := SampleJob{}
	jq.Route(id, sample)

	// Push Job
	input := JobInput{
		ID:   id,
		Data: []byte(testData),
	}
	if err := jq.Push(input); err != nil {
		t.Errorf("Error JobQueue Push %v", err)
	}

	time.Sleep(1 * time.Second)
}

type SampleJob struct {
}

func (j SampleJob) Serve(data []byte) {
	log.Print("SampleJob Serve")

	expected := string(data)

	actual := testData

	if e, a := expected, actual; !reflect.DeepEqual(e, a) {
		log.Print(pretty.Compare(e, a))
	}
}
