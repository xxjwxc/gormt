package genfunc

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/xxjwxc/gormt/data/view/genfunc/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xxjwxc/public/mysqldb"
)

/**
测试数据库地址:https://github.com/xxjwxc/gormt/blob/master/data/view/genfunc/model/matrix.sql
*/

func GetGorm(dataSourceName string) *gorm.DB {
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

// TestFuncGet 测试条件获(Get/Gets)
func TestFuncGet(t *testing.T) {
	model.OpenRelated() // 打开全局预加载 (外键)

	db := GetGorm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	defer db.Close()

	accountMgr := model.AccountMgr(db.Where("account_id = ?", 2))
	account, err := accountMgr.Get() // 单条获取
	fmt.Println(err)
	fmt.Println(account)

	dbs := db.Where("name = ?", "bbbb")
	accountMgr.UpdateDB(dbs)           // 更新数据库
	accounts, err := accountMgr.Gets() // 多个获取
	fmt.Println(err)
	fmt.Println(accounts)
}

// TestFuncOption 功能选项方式获取
func TestFuncOption(t *testing.T) {
	// db := GetGorm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	// defer db.Close()
	orm := mysqldb.OnInitDBOrm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True") // 推荐方式
	defer orm.OnDestoryDB()
	db := orm.DB

	accountMgr := model.AccountMgr(db)
	accountMgr.SetIsRelated(true)                                                          // 打开预加载 (外键)
	account, err := accountMgr.GetByOption(accountMgr.WithID(1), accountMgr.WithUserID(1)) // 多case条件获取单个
	fmt.Println(err)
	fmt.Println(account)

	accounts, err := accountMgr.GetByOptions(accountMgr.WithName("bbbb")) // 多功能选项获取
	fmt.Println(err)
	fmt.Println(accounts)
}

// TestFuncFrom 单元素方式获取
func TestFuncFrom(t *testing.T) {
	db := GetGorm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	defer db.Close()

	accountMgr := model.AccountMgr(db)
	accountMgr.SetIsRelated(true) // 打开预加载 (外键)

	account, err := accountMgr.GetFromAccountID(2)
	fmt.Println(err)
	fmt.Println(account)

	accounts, err := accountMgr.GetFromName("bbbb")
	fmt.Println(err)
	fmt.Println(accounts)
}

// TestFuncFetchBy 索引方式获取
func TestFuncFetchBy(t *testing.T) {
	db := GetGorm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	defer db.Close()

	accountMgr := model.AccountMgr(db)
	accountMgr.SetIsRelated(true) // 打开预加载 (外键)

	account, err := accountMgr.FetchByPrimaryKey(2) // primay key
	fmt.Println(err)
	fmt.Println(account)

	account1, err := accountMgr.FetchUniqueIndexByAccount(2, 2) // unique index
	fmt.Println(err)
	fmt.Println(account1)

	accounts, err := accountMgr.FetchIndexByTp(2, 2)
	fmt.Println(err)
	fmt.Println(accounts)
}
