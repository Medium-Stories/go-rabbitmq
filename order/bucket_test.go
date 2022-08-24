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
		Sku: "item1",
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
						Sku: "item1",
					},
					{
						Sku: "item2",
					},
					{
						Sku: "item3",
					},
				},
			},
			expectedLength: 2,
		},
	}

	for name, suite := range table {
		t.Run(name, func(t *testing.T) {
			suite.bucket.Remove("item2")
			assert.Equal(t, suite.expectedLength, len(suite.bucket.Items))

			item1 := suite.bucket.Items[0]
			assert.Equal(t, "item1", item1.Sku)

			item2 := suite.bucket.Items[1]
			assert.Equal(t, "item3", item2.Sku)
		})
	}
}
