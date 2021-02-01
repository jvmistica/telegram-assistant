package main

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Items) CheckCommand(cmd string) string {
	var msg string

	switch cmd {
	case "/start":
		msg = "Welcome. Here's a list of my commands:\n" +
			"/list items	- List items in your inventory\n" +
			"/add item	- Add an item to your inventory\n" +
			"/edit item	- Edit an item in your inventory\n" +
			"/delete item	- Delete an item in your inventory\n\n" +
			"To view the full list of commands and sub-commands, enter /list commands"
	case "/list items":
		msg = i.ListItems()
		if msg == "" {
			msg = "There are no items in your inventory."
		}
	case "/add item":
		msg = "What is this item called?"
	case "/edit item":
		msg = "Which item do you want to edit?"
	case "/delete item":
		msg = "Which item do you want to delete?"
	default:
		msg = "That's not a valid command. Here's a list of valid commands:\n" +
			"/list items	- List items in your inventory\n" +
			"/add item	- Add an item to your inventory\n" +
			"/edit item	- Edit an item in your inventory\n" +
			"/delete item	- Delete an item in your inventory\n"
	}
	return msg
}

// CheckItem returns true if item exists
func (i *Items) CheckItem(item string) bool {
	var rec Item
	res := i.db.Where("name = ?", item).Find(&rec)
	return res.RowsAffected == 1
}

// ListItems returns all the items in the inventory
func (i *Items) ListItems() string {
	var (
		items     []Item
		itemsList string
	)

	i.db.Find(&items)
	for _, item := range items {
		itemsList += item.Name + "\n"
	}
	return itemsList
}

// AddItem adds an item to the inventory
func (i *Items) AddItem(item string) error {
	rec := Item{Name: item}
	res := i.db.Create(&rec)
	return res.Error
}

// EditItem updates an item's properties
func (i *Items) EditItem(item, newItem string) {
	i.db.Model(&Item{}).Where("name = ?", item).Update("name", newItem)
}

// DeleteItem deletes an item
func (i *Items) DeleteItem(item string) int64 {
	res := i.db.Where("name = ?", item).Delete(Item{})
	return res.RowsAffected
}
