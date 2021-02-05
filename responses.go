package main

var (
	startMsg = "Welcome. Here's a list of my commands:\n" +
		"/list items	- List items in your inventory\n" +
		"/add item	- Add an item to your inventory\n" +
		"/update item	- Update an item in your inventory\n" +
		"/delete item	- Delete an item in your inventory\n\n" +
		"To view the full list of commands and sub-commands, enter /list commands"

	invalidMsg = "That's not a valid command. Here's a list of valid commands:\n" +
		"/list items	- List items in your inventory\n" +
		"/add item	- Add an item to your inventory\n" +
		"/update item	- Update an item in your inventory\n" +
		"/delete item	- Delete an item in your inventory\n"

	addChoose  = "What is this item called?"
	addSuccess = "Added \"<item>\" to the inventory.\n\n" +
		"Update this item's details using the commands below:\n" +
		"/update item description <item> <description>\n" +
		"/update item amount <item> <amount>\n" +
		"/update item category <item> <category>\n" +
		"/update item price <item> <price>\n" +
		"/update item expiration <item> <expiration>\n"

	deleteChoose  = "Which item do you want to delete?"
	deleteSuccess = "Removed \"<item>\" from the inventory."

	updateChoose  = "Which item do you want to update?"
	updateSuccess = "Updated <item>'s <field>."

	itemNotExist = "Item \"<item>\" does not exist in the inventory."
	noItems      = "There are no items in your inventory."
)
