package events

type EventType string

const (
	UserSignedUp EventType = "user.signedup"
)

type OrderSubmittedPayload struct {
	Order              string
	OrderProductStocks []int
}

type Event struct {
	Type    EventType
	Payload any
}
