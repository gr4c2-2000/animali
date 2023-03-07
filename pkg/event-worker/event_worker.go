package eventworker

import (
	"context"
	"errors"
)

type Listiner func(typ string, value string)

type Event struct {
	Type  string
	Value string
}

// TODO : Work with typos
type EventWoker struct {
	ctx      context.Context
	Listiner []Listiner
	Queue    chan Event
}

func New(ctx context.Context, Queue chan Event, listiner ...Listiner) *EventWoker {
	ew := EventWoker{ctx: ctx, Queue: Queue, Listiner: listiner}
	return &ew
}

func (cw *EventWoker) AddListiner(ce ...Listiner) {
	for _, function := range ce {
		cw.Listiner = append(cw.Listiner, function)
	}
}

func NewEvent(typ string, value string) Event {
	return Event{typ, value}
}

func (cw *EventWoker) Worker() error {
	if cw.Queue == nil {
		return errors.New("NO_QUEUE")
	}
	go func() {
		for {
			select {
			case event := <-cw.Queue:
				for _, reciver := range cw.Listiner {
					reciver(event.Type, event.Value)
				}
			case <-cw.ctx.Done():
				return
			}
		}
	}()
	return nil
}
