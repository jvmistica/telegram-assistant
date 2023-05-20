package record

const (
	listItems = "/listitems - list items in your inventory\n"
)

var (
	ResponseStart = "Welcome. Here's a list of my commands:\n\n" +
		listItems +
		"/showitem - show an item's details\n" +
		"/additem - add an item\n" +
		"/updateitem - update an item\n" +
		"/deleteitem - delete an item\n"

	ResponseInvalid = "That's not a valid command. Here's a list of valid commands:\n\n" +
		listItems +
		"/showitem - show an item's details\n" +
		"/additem - add an item\n" +
		"/updateitem - update an item\n" +
		"/deleteitem - delete an item\n"

	ResponseInvalidList = "That's not a valid command. Here's a list of valid commands:\n\n" +
		"*List DB*\n" +
		listItems +
		"/listitems sort by _field_ - sort list (ascending)\n" +
		"/listitems sort by _field_ desc - sort list (descending)\n" +
		"/listitems filter by _field_ = _value_ - filter items in your inventory\n" +
		"_filter operations: =, <, >, <=, >=, <>, like_\n"

	ResponseAdd        = "What is this item called?"
	ResponseSuccessAdd = "Added \"<item>\" to the inventory.\n\n" +
		"Update this item's details using the commands below:\n" +
		"/updateitem _itemName_ description _itemDescription_\n" +
		"/updateitem _itemName_ calories _calories_\n" +
		"/updateitem _itemName_ amount _itemAmount_\n" +
		"/updateitem _itemName_ category _itemCategory_\n" +
		"/updateitem _itemName_ price _itemPrice_\n" +
		"/updateitem _itemName_ expiration _itemExpiration_\n"

	ResponseShow = "Which item do you want to see?"

	ResponseDelete        = "Which item do you want to delete?"
	ResponseSuccessDelete = "Removed \"<item>\" from the inventory."

	ResponseUpdate        = "Which item do you want to update?"
	ResponseSuccessUpdate = "Updated <item>'s <field>."
	ResponseInvalidUpdate = "Invalid parameters. To update an item, use the commands below:\n\n" +
		"*Update DB*\n" +
		"/updateitem _itemName_ description _itemDescription_\n" +
		"/updateitem _itemName_ calories _calories_\n" +
		"/updateitem _itemName_ amount _itemAmount_\n" +
		"/updateitem _itemName_ category _itemCategory_\n" +
		"/updateitem _itemName_ price _itemPrice_\n" +
		"/updateitem _itemName_ expiration _itemExpiration_\n"

	ResponseItemNotExist  = "Item \"<item>\" does not exist in the inventory."
	ResponseNoMatchFilter = "There are no items matching that filter."
	ResponseNoItems       = "There are no items in your inventory."
)
