package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xxjwxc/gormt/model"
	"github.com/xxjwxc/public/mysqldb"
)

const (
	sqlConn = "root:qwer@tcp(localhost:3306)/oauth_db?charset=utf8&parseTime=true&loc=Local&timeout=10s"
)

func TestForeignKey(t *testing.T) {
	db, err := gorm.Open("mysql", sqlConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) // 全局禁用表名复数
	db.LogMode(true)

	var p0 model.UserAccountTbl
	p0.Account = "test"
	p0.RegTime = time.Now()
	p0.UserInfoTbl.Nickname = "testBBB"
	p0.UserInfoTblID = 1
	db.Save(&p0.UserInfoTbl)
	p0.UserInfoTblID = (int)(p0.UserInfoTbl.ID)
	db.Save(&p0)

	var p1 model.UserAccountTbl
	db.Where("account = ?", "test").Find(&p1)
	db.Model(&p1).Related(&p1.UserInfoTbl) // 外键关联
	fmt.Println(p1)
}

func TestForeignKeyOrm(t *testing.T) {
	orm := mysqldb.OnInitDBOrm(sqlConn)
	defer orm.OnDestoryDB()

	var p model.UserAccountTbl
	orm.Find(&p)
	orm.Model(&p).Related(&p.UserInfoTbl) // 外键关联
	fmt.Println(p)
}
