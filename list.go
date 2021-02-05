package main

// ListItems returns all the items in the inventory
func (i *Items) ListItems(params string) string {
	var (
		items     []Item
		itemsList string
	)

	if params == "sort by name" {
		i.db.Order("name").Find(&items)
	} else {
		i.db.Find(&items)
	}

	for _, item := range items {
		itemsList += item.Name + "\n"
	}

	return itemsList
}
