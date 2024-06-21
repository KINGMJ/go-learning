package events

import "fmt"

type UserDeletedSubsrciber struct {
}

func NewUserDeletedSubsrciber() *UserDeletedSubsrciber {
	return &UserDeletedSubsrciber{}
}

// 事件订阅者
func (s *UserDeletedSubsrciber) HandleUserDeleted(event Event) {
	payload, ok := event.Payload.(*OrderSubmittedPayload)
	if !ok {
		fmt.Println("事件订阅失败")
	}
	s.a(payload)
	s.b(payload)
}

func (s *UserDeletedSubsrciber) a(payload *OrderSubmittedPayload) error {
	fmt.Println(payload.Order)
	return nil
}

func (s *UserDeletedSubsrciber) b(payload *OrderSubmittedPayload) error {
	fmt.Println(payload.OrderProductStocks)
	return nil
}
