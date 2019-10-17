package genmysql

import "strings"

// filterModel filter.过滤 gorm.Model
func filterModel(list *[]genColumns) bool {
	var _temp []genColumns

	num := 0
	for _, v := range *list {
		if strings.EqualFold(v.Field, "id") ||
			strings.EqualFold(v.Field, "created_at") ||
			strings.EqualFold(v.Field, "updated_at") ||
			strings.EqualFold(v.Field, "deleted_at") {
			num++
		} else {
			_temp = append(_temp, v)
		}
	}

	if num >= 4 {
		*list = _temp
		return true
	}

	return false
}
