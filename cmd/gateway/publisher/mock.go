package publisher

import "github.com/sirupsen/logrus"

type mocked struct {
	events map[string]interface{}
}

func NewMocked() *mocked {
	return &mocked{
		events: make(map[string]interface{}),
	}
}

func (pub *mocked) Publish(event string, msg interface{}) error {
	pub.events[event] = msg
	logrus.Infof("event triggered: %s => message: %v", event, msg)
	return nil
}
