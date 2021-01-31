package main

func (i *Items) CheckCommand(cmd string) string {
	var msg string
	if cmd == "/start" {
		msg = "Welcome. Here's a list of my commands:\n" +
			"/list items	- List items in your inventory\n" +
			"/add item	- Add an item to your inventory\n" +
			"/edit item	- Edit an item in your inventory\n" +
			"/delete item	- Delete an item in your inventory\n\n" +
			"To view the full list of commands and sub-commands, enter /list commands"
	} else if cmd == "/list items" {
		msg = i.ListItems()
	} else if cmd == "/add item" {
		msg = "What is this item called?"
	} else if cmd == "/edit item" {
		msg = "Which item do you want to edit?"
	} else if cmd == "/delete item" {
		msg = "Which item do you want to delete?"
	} else {
		msg = "That's not a valid command. Here's a list of valid commands:\n" +
			"/list items	- List items in your inventory\n" +
			"/add item	- Add an item to your inventory\n" +
			"/edit item	- Edit an item in your inventory\n" +
			"/delete item	- Delete an item in your inventory\n"
	}

	return msg
}

func (i *Items) CheckItem(item string) int64 {
	var rec Item
	res := i.db.Where("name = ?", item).Find(&rec)
	return res.RowsAffected
}

func (i *Items) ListItems() string {
	var items []Item
	var st string
	i.db.Find(&items)

	for _, item := range items {
		st += item.Name + "\n"
	}
	return st
}

func (i *Items) AddItem(item string) error {
	rec := Item{Name: item}
	res := i.db.Create(&rec)
	return res.Error
}

func (i *Items) EditItem(item, newItem string) string {
	i.db.Model(&Item{}).Where("name = ?", item).Update("name", newItem)
	return ""
}

func (i *Items) DeleteItem(item string) int64 {
	res := i.db.Where("name = ?", item).Delete(Item{})
	return res.RowsAffected
}
