package order

type Bucket struct {
	Identifier  string
	Items       []*Item
	OrderStatus int
}

type Item struct {
	Id int
}

func (bucket *Bucket) Add(item *Item) {
	bucket.Items = append(bucket.Items, item)
}

func (bucket *Bucket) Remove(itemId int) {
	for i := 0; i < len(bucket.Items); i++ {
		if bucket.Items[i].Id == itemId {
			copy(bucket.Items[i:], bucket.Items[i+1:])
			bucket.Items[len(bucket.Items)-1] = nil
			bucket.Items = bucket.Items[:len(bucket.Items)-1]
		}
	}
}
