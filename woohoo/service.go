package woohoo

type service struct {
	items []*Item
}

func NewService() *service {
	return &service{
		items: []*Item{
			{
				Sku:      "item1",
				Price:    10,
				Category: "cat1",
			},
			{
				Sku:      "item2",
				Price:    20,
				Category: "cat1",
			},
			{
				Sku:      "item3",
				Price:    30,
				Category: "cat2",
			},
		},
	}
}

func (svc *service) GetItemBySku(sku string) *Item {
	for _, item := range svc.items {
		if item.Sku == sku {
			return item
		}
	}
	return nil
}

func (svc *service) GetAvailableItems() []*Item {
	return svc.items
}
