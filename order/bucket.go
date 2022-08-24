package order

type Bucket struct {
	Identifier  string
	Items       []*Item
	OrderStatus int
}

type Item struct {
	Sku string
}

func (bucket *Bucket) Add(item *Item) {
	bucket.Items = append(bucket.Items, item)
}

func (bucket *Bucket) Remove(sku string) {
	for i := 0; i < len(bucket.Items); i++ {
		if bucket.Items[i].Sku == sku {
			copy(bucket.Items[i:], bucket.Items[i+1:])
			bucket.Items[len(bucket.Items)-1] = nil
			bucket.Items = bucket.Items[:len(bucket.Items)-1]
		}
	}
}
