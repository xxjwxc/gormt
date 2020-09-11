package genmysql

import (
	"strings"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/xxjwxc/gormt/data/view/model"
)

// filterModel filter.过滤 gorm.Model
func filterModel(list *[]genColumns) bool {
	if config.GetDBTag() != "gorm" {
		return false
	}

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

// fixForeignKey fix foreign key.过滤外键
func fixForeignKey(list []genForeignKey, columuName string, result *[]model.ForeignKey) {
	for _, v := range list {
		if strings.EqualFold(v.ColumnName, columuName) { // find it .找到了
			*result = append(*result, model.ForeignKey{
				TableName:  v.ReferencedTableName,
				ColumnName: v.ReferencedColumnName,
			})
		}
	}
}

// GetMysqlModel get model interface. 获取model接口
func GetMysqlModel() model.IModel {
	//now just support mysql
	return &MySQLModel
}

// FixElementNote 分析元素表注释
func FixElementNote(em *model.ColumnsInfo, note string) {
	matches := noteRegex.FindStringSubmatch(note)
	if len(matches) < 2 {
		em.Notes = note
		return
	}
	em.Notes = note[len(matches[0]):]

	list := strings.Split(matches[1], ";")
	for _, v := range list {
		tmp := strings.Split(v, ":")
		if len(tmp) == 2 {
			if strings.EqualFold(tmp[0], "default") { // 默认值
				em.Default = tmp[1]
			}
		}
	}
}
