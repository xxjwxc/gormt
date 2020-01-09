package genfunc

import (
	"fmt"
	"testing"

	"github.com/xxjwxc/gormt/data/view/genfunc/model"

	"github.com/xxjwxc/public/mysqldb"
)

func TestFunc(t *testing.T) {
	orm := mysqldb.OnInitDBOrm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local")
	defer orm.OnDestoryDB()
	mgr := model.OrganMgr(orm.DB)
	mgr.IsRelated(true)              // 设置允许加载外键
	res, err := mgr.GetFromUserID(2) // 通过列获取
	fmt.Println(res, err)

	obj1, err := mgr.GetByOptions(mgr.WithID(1), mgr.WithUserID(1)) // 批量获取
	fmt.Println(obj1, err)

	obj2, err := mgr.GetByOption(mgr.WithID(1), mgr.WithUserID(1)) // 多条件获取一条
	fmt.Println(obj2, err)

	// 复合键获取
}
