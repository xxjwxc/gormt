package genfunc

import (
	"fmt"
	"testing"
	"time"

	"github.com/xxjwxc/gormt/data/view/genfunc/model"
	"github.com/xxjwxc/public/mysqldb"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

/**
测试数据库地址:https://github.com/xxjwxc/gormt/blob/master/data/view/genfunc/model/matrix.sql
*/

func GetGorm(dataSourceName string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{PrepareStmt: false})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db.Debug()
}

// func NewDB(){
// 	db, _ := gorm.Open(...)
// 	db.Model(&AAA).Where("aaa = ?", 2)
// 	CallFunc(db)
// }

// func CallFunc(db *gorm.DB){
// 	// select a...
// 	var bbb BBB
// 	db.Table("bbb").Where("bbb = ?", 2).Find(&bbb)// in this case aaa = ?  valid
// 	// in this func how to us db to query BBB
// }

// TestFuncGet 测试条件获(Get/Gets)
func TestFuncGet(t *testing.T) {
	model.OpenRelated() // 打开全局预加载 (外键)

	db := GetGorm("root:123456@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

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
	// defer func() {
	// 	sqldb, _ := db.DB()
	// 	sqldb.Close()
	// }()
	orm := mysqldb.OnInitDBOrm("root:123456@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True") // 推荐方式
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
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

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
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

	accountMgr := model.AccountMgr(db)
	accountMgr.SetIsRelated(true) // 打开预加载 (外键)

	account, err := accountMgr.FetchByPrimaryKey(2) // primary key
	fmt.Println(err)
	fmt.Println(account)

	account1, err := accountMgr.FetchUniqueIndexByAccount(2, 2) // unique index
	fmt.Println(err)
	fmt.Println(account1)

	accounts, err := accountMgr.FetchIndexByTp(2, 2)
	fmt.Println(err)
	fmt.Println(accounts)
}

// TestCondetion 测试sql构建
func TestCondetion(t *testing.T) {
	condetion := model.Condetion{}
	condetion.And(model.AccountColumns.AccountID, ">=", "1")
	condetion.And(model.AccountColumns.UserID, "in", "1", "2", "3")
	condetion.Or(model.AccountColumns.Type, "in", "1", "2", "3")

	where, obj := condetion.Get()
	fmt.Println(where)
	fmt.Println(obj...)

	db := GetGorm("root:qwer@tcp(127.0.0.1:3306)/matrix?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

	accountMgr := model.AccountMgr(db.Where(condetion.Get()))
	accountMgr.Gets()
}
