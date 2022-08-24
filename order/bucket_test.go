package order_test

import (
	"github.com/medium-stories/go-rabbitmq/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucket_Add(t *testing.T) {
	bucket := &order.Bucket{}
	assert.Equal(t, 0, len(bucket.Items))

	bucket.Add(&order.Item{
		Id: 1,
	})

	assert.Equal(t, 1, len(bucket.Items))
}

func TestBucket_Remove(t *testing.T) {
	table := map[string]struct {
		bucket         *order.Bucket
		expectedLength int
	}{
		"remove existing item": {
			bucket: &order.Bucket{
				Items: []*order.Item{
					{
						Id: 1,
					},
					{
						Id: 2,
					},
					{
						Id: 3,
					},
				},
			},
			expectedLength: 2,
		},
	}

	for name, suite := range table {
		t.Run(name, func(t *testing.T) {
			suite.bucket.Remove(2)
			assert.Equal(t, suite.expectedLength, len(suite.bucket.Items))

			item1 := suite.bucket.Items[0]
			assert.Equal(t, 1, item1.Id)

			item2 := suite.bucket.Items[1]
			assert.Equal(t, 3, item2.Id)
		})
	}
}
