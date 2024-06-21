package main

import (
	"fmt"

	"github.com/KINGMJ/go-learning/tutorial9/events"
)

func init() {
	events.InitEventBus()
}

func main() {
	event := events.Event{
		Type: events.UserSignedUp,
		Payload: &events.OrderSubmittedPayload{
			Order:              "haha",
			OrderProductStocks: []int{1, 2, 3},
		},
	}
	res := events.Publish(event)
	fmt.Println(res)
}
