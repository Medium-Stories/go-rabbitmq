package repository

import "github.com/medium-stories/go-rabbitmq/order"

type inmemory struct {
	orders []*order.Bucket
}

func NewInMemory() *inmemory {
	return &inmemory{}
}

func (repo *inmemory) Save(bucket *order.Bucket) error {
	repo.orders = append(repo.orders, bucket)
	return nil
}

func (repo *inmemory) GetByIdentifier(identifier string) *order.Bucket {
	for _, o := range repo.orders {
		if o.Identifier == identifier {
			return o
		}
	}

	return nil
}

func (repo *inmemory) UpdateStatus(identifier string, status int) error {
	for _, o := range repo.orders {
		if o.Identifier == identifier {
			o.OrderStatus = status
		}
	}

	return nil
}
