package event

const (
	OrderCreated = "order_created"
	OrderPaid    = "order_paid"
	OrderShipped = "order_shipped"
)

type Publisher interface {
	Publish(event string, payload interface{}) error
}
