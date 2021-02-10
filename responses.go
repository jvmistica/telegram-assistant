package main

var (
	startMsg = "Welcome. Here's a list of my commands:\n" +
		"/listitems - list items in your inventory\n" +
		"/showitem - show an item's details\n" +
		"/additem - add an item\n" +
		"/updateitem - update an item\n" +
		"/deleteitem - delete an item\n"

	invalidMsg = "That's not a valid command. Here's a list of valid commands:\n" +
		"/listitems - list items in your inventory\n" +
		"/showitem - show an item's details\n" +
		"/additem - add an item\n" +
		"/updateitem - update an item\n" +
		"/deleteitem - delete an item\n"

	addChoose  = "What is this item called?"
	addSuccess = "Added \"<item>\" to the inventory.\n\n" +
		"Update this item's details using the commands below:\n" +
		"/updateitem _itemName_ description _itemDescription_\n" +
		"/updateitem _itemName_ amount _itemAmount_\n" +
		"/updateitem _itemName_ category _itemCategory_\n" +
		"/updateitem _itemName_ price _itemPrice_\n" +
		"/updateitem _itemName_ expiration _itemExpiration_\n"

	deleteChoose  = "Which item do you want to delete?"
	deleteSuccess = "Removed \"<item>\" from the inventory."

	updateChoose  = "Which item do you want to update?"
	updateSuccess = "Updated <item>'s <field>."
	updateInvalid = "Invalid parameters. To update an item, use the commands below:\n" +
		"/updateitem _itemName_ description _itemDescription_\n" +
		"/updateitem _itemName_ amount _itemAmount_\n" +
		"/updateitem _itemName_ category _itemCategory_\n" +
		"/updateitem _itemName_ price _itemPrice_\n" +
		"/updateitem _itemName_ expiration _itemExpiration_\n"

	itemNotExist = "Item \"<item>\" does not exist in the inventory."
	noItems      = "There are no items in your inventory."
)
