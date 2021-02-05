package main

// CheckItem returns true if item exists
func (i *Items) CheckItem(item string) bool {
	var rec Item
	res := i.db.Where("name = ?", item).Find(&rec)
	return res.RowsAffected == 1
}
