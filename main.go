package main

import (
	"fmt"
	"github.com/dayemsiddiqui/eventloop/eventbus"
	"time"
)

type Data struct {
	title string
}

func main() {
	EventBus := eventbus.New()

	EventBus.Subscribe("topic1", func(data interface{}) {
		time.Sleep(4 * time.Second)
		if d, ok := data.(Data); ok {
			fmt.Println("topic1:", d.title)
			return
		}
		fmt.Println("topic1:", data)
	})

	EventBus.Subscribe("topic2", func(data interface{}) {
		fmt.Println("topic2:", data)
	})

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Hello from goroutine from main")
	}()
	EventBus.Publish("topic1", Data{title: "This is a title"})

	fmt.Println("Hello from main")

	EventBus.Publish("topic2", "world")

	fmt.Println("Goodbye from main")

	EventBus.Wait()
}
