package genfunc

import (
	"fmt"
	"testing"

	"github.com/xxjwxc/public/mysqldb"
)

func TestFunc(t *testing.T) {
	orm := mysqldb.OnInitDBOrm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local")
	defer orm.OnDestoryDB()

	mgr := ExampleMgr(orm.DB)
	obj, err := mgr.GetFromID(1)
	fmt.Println(obj, err)

	obj1, err := mgr.GetByPrimaryKey(1)
	fmt.Println(obj1, err)

	obj2, err := mgr.GetByPrimaryKeys([]int64{1, 2})
	fmt.Println(obj2, err)

	obj3, err := mgr.GetByOptions(mgr.WithID(1), mgr.WithUserID(1))
	fmt.Println(obj3, err)

	obj4, err := mgr.GetByOption(mgr.WithID(1), mgr.WithUserID(1))
	fmt.Println(obj4, err)
}
