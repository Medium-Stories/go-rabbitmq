package order

import (
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

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
}

func (svc *service) Create(bucket *Bucket) error {
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
