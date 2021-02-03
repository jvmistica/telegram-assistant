package main

var (
	startMsg = "Welcome. Here's a list of my commands:\n" +
		"/list items	- List items in your inventory\n" +
		"/add item	- Add an item to your inventory\n" +
		"/edit item	- Edit an item in your inventory\n" +
		"/delete item	- Delete an item in your inventory\n\n" +
		"To view the full list of commands and sub-commands, enter /list commands"

	invalidMsg = "That's not a valid command. Here's a list of valid commands:\n" +
		"/list items	- List items in your inventory\n" +
		"/add item	- Add an item to your inventory\n" +
		"/edit item	- Edit an item in your inventory\n" +
		"/delete item	- Delete an item in your inventory\n"

	addChoose  = "What is this item called?"
	addSuccess = "Added \"<item>\" to the inventory.\n\n" +
		"Add more details about this item using the commands below:\n" +
		"/add item description <item> <description>\n" +
		"/add item amount <item> <amount>\n" +
		"/add item category <item> <category>\n" +
		"/add item price <item> <price>\n" +
		"/add item expiration <item> <expiration>\n"

	deleteChoose  = "Which item do you want to delete?"
	deleteSuccess = "Removed \"<item>\" from the inventory."

	editChoose  = "Which item do you want to edit?"
	editPrompt  = "What do you want to change \"<item>\" to?"
	editSuccess = "Item \"<oldItem>\" has been changed to \"<newItem>\"."

	itemNotExist = "Item \"<item>\" does not exist in the inventory."
	noItems      = "There are no items in your inventory."
)
