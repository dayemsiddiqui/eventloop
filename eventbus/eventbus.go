package eventbus

import "sync"

type Handler struct {
	topic    string
	callback func(interface{})
}

type EventBus struct {
	handlers  []Handler
	waitGroup sync.WaitGroup
}

func New() *EventBus {
	return &EventBus{}
}

func (e *EventBus) Subscribe(topic string, callback func(interface{})) {
	handler := Handler{topic: topic, callback: callback}
	e.handlers = append(e.handlers, handler)
}

func (e *EventBus) Publish(topic string, data interface{}) {
	for _, handler := range e.handlers {
		if handler.topic == topic {
			e.waitGroup.Add(1)
			go func(handler Handler) {
				defer e.waitGroup.Done()
				handler.callback(data)
			}(handler)
		}
	}
}

func (e *EventBus) Wait() {
	e.waitGroup.Wait()
}
