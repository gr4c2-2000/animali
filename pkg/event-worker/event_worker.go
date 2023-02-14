package eventworker

import "errors"

type Listiner func(typ string, value string)

type Event struct {
	Type  string
	Value string
}

type EventWoker struct {
	Listiner []Listiner
	Queue    chan Event
}

func (cw *EventWoker) AddListiner(ce ...Listiner) {
	for _, function := range ce {
		cw.Listiner = append(cw.Listiner, function)
	}
}

func NewEvent(typ string, value string) Event {
	return Event{typ, value}
}

func (cw *EventWoker) CommandWorker() error {
	if cw.Queue == nil {
		return errors.New("NO_QUEUE")
	}
	go func() {
		for event := range cw.Queue {
			for _, reciver := range cw.Listiner {
				reciver(event.Type, event.Value)
			}
		}
	}()

	return nil
}
