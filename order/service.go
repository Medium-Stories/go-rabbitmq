package order

import (
	"github.com/google/uuid"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

// listen for events: order paid and order shipped, and update order status accordingly

type service struct {
	repo Repository
	pub  event.Publisher
}

func NewService(repo Repository, pub event.Publisher) *service {
	return &service{
		repo: repo,
		pub:  pub,
	}
}

type Repository interface {
	Save(bucket *Bucket) error
	GetByIdentifier(identifier string) *Bucket
	UpdateStatus(identifier string, status int) error
}

func (svc *service) Create(bucket *Bucket) error {
	bucket.Identifier = uuid.NewString()

	if err := svc.repo.Save(bucket); err != nil {
		return err
	}

	go func(identifier string) {
		if err := svc.pub.Publish(event.OrderCreated, identifier); err != nil {
			logrus.Error(err)
		}
	}(bucket.Identifier)

	return nil
}

func (svc *service) CheckStatus(identifier string) *Bucket {
	return svc.repo.GetByIdentifier(identifier)
}
