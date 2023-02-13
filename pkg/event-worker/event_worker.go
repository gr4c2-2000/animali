package eventworker

type Listiner func(typ string, value string)

type Event struct {
	Type  string
	Value string
}

type EventWorder struct {
	Listiner []Listiner
	Queue    chan Event
}

func (cw *EventWorder) AddListiner(ce Listiner) {
	cw.Listiner = append(cw.Listiner, ce)
}

func (cw *EventWorder) CommandWorker() {

	go func() {
		for Event := range cw.Queue {
			for _, revicer := range cw.Listiner {
				revicer(Event.Type, Event.Value)
			}
		}
	}()
}
