package main

// ListItems returns all the items in the inventory
func (i *Items) ListItems(params string) string {
	var (
		items     []Item
		itemsList string
	)

	switch params {
	case "sort by name":
		i.db.Order("name").Find(&items)
	case "sort by amount":
		i.db.Order("amount").Find(&items)
	case "sort by category":
		i.db.Order("category").Find(&items)
	case "sort by price":
		i.db.Order("price").Find(&items)
	case "sort by expiration":
		i.db.Order("expiration").Find(&items)
	default:
		i.db.Find(&items)
	}

	for _, item := range items {
		itemsList += item.Name + "\n"
	}

	return itemsList
}
