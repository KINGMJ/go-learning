package events

import "fmt"

// 事件订阅者
func HandleUserSignedUp(event Event) error {
	payload, ok := event.Payload.(*OrderSubmittedPayload)
	if !ok {
		fmt.Println("事件订阅失败")
	}
	a(payload)
	b(payload)
	return fmt.Errorf("出错了")
}

func a(payload *OrderSubmittedPayload) {
	fmt.Println(payload.Order)
}

func b(payload *OrderSubmittedPayload) {
	fmt.Println(payload.OrderProductStocks)
}
