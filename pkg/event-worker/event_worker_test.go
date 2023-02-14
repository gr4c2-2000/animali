package eventworker

import (
	"testing"
	"time"
)

var TestResultQueue = make(chan string)

func TestEventWoker_CommandWorker(t *testing.T) {
	type fields struct {
		Listiner []Listiner
		Queue    chan Event
	}

	TestQueue := []struct {
		name   string
		fields fields
	}{
		{"nil queue", fields{[]Listiner{}, nil}},
	}
	for _, tt := range TestQueue {
		t.Run(tt.name, func(t *testing.T) {
			cw := &EventWoker{
				Listiner: tt.fields.Listiner,
				Queue:    tt.fields.Queue,
			}
			err := cw.CommandWorker()
			if err == nil {
				t.Errorf("CommandWorker() nil Queue not detected")
			}
		})
	}
	TestResult := []struct {
		name   string
		fields fields
		result string
		event  Event
	}{
		{"Event procesing", fields{[]Listiner{testFuncType}, make(chan Event)}, "type", Event{"type", "value"}},
		{"Event procesing", fields{[]Listiner{testFuncValue}, make(chan Event)}, "value", Event{"type", "value"}},
	}

	for _, tt := range TestResult {
		t.Run(tt.name, func(t *testing.T) {
			cw := &EventWoker{
				Listiner: tt.fields.Listiner,
				Queue:    tt.fields.Queue,
			}
			err := cw.CommandWorker()
			if err != nil {
				t.Errorf("CommandWorker() Error : %v", err)
			}
			cw.Queue <- tt.event
			ticker := time.NewTicker(10 * time.Second)
			select {
			case <-ticker.C:
				t.Errorf("Did not revice Event in 10 sec")
			case val := <-TestResultQueue:
				if val != tt.result {
					t.Errorf("inccorect Value Reviced :%v want: %v ", val, tt.result)
				}
				return
			}
		})
	}

}

func testFuncType(a string, b string) {
	TestResultQueue <- a
}
func testFuncValue(a string, b string) {
	TestResultQueue <- b
}