[![Build Status](https://travis-ci.org/xxjwxc/gormt.svg?branch=master)](https://travis-ci.org/xxjwxc/gormt)
[![Go Report Card](https://goreportcard.com/badge/github.com/xxjwxc/gormt)](https://goreportcard.com/report/github.com/xxjwxc/gormt)
[![GoDoc](https://godoc.org/github.com/xxjwxc/gormt?status.svg)](https://godoc.org/github.com/xxjwxc/gormt)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go) 
 
# [gormt](https://github.com/xxjwxc/gormt)

## [中文文档](README_zh_cn.md)

###  mysql database to goang struct conversion tools base on [gorm](https://github.com/jinzhu/gorm)，You can automatically generate golang sturct from MySQL database. big Camel-Case Name Rule, JSON tag. 

## gui support

![show](/image/gormt/ui_en.gif)

```
./gormt -g=true
```

## cmd support

![show](/image/gormt/out.gif)
```
./gormt -g=false
```
--------

## 1. Configure default configuration items through the current directory config.yml file
```
out_dir : "."  # out dir
url_tag : json # web url tag(json,db(https://github.com/google/go-querystring))
language :  # language(English,中 文)
db_tag : gorm # DB tag(gorm,db)
singular_table : true  # Table name plural (big Camel-Case):gorm.SingularTable
simple : false #simple output
is_out_sql : false # Whether to output sql
is_out_func : true # Whether to output function
is_url_tag : true # Whether to mark web or not
is_foreign_key : true # Whether to mark foreign key or not
is_gui : false # Whether to operate on gui

mysql_info :
    host : "127.0.0.1"
    port : 3306
    username : "root"
    password : "qwer"
    database : "oauth_db"

```
## 2. get help
```
./gormt --help
or
./gormt -h

-------------------------------------------------------
base on gorm tools for mysql database to golang struct

Usage:
  main [flags]

Flags:
  -d, --database string   数据库名
  -f, --foreign           是否导出外键关联
  -F, --fun               是否导出函数
  -g, --gui               是否ui显示模式
  -h, --help              help for main
  -H, --host string       数据库地址.(注意-H为大写)
  -o, --outdir string     输出目录
  -p, --password string   密码.
      --port int          端口号 (default 3306)
  -s, --singular          是否禁用表名复数
  -l, --url string        url标签(json,url)
  -u, --user string       用户名.
  
```
## 3. Can be updated configuration items using command line tools
```
./gormt -H=127.0.0.1 -d=oauth_db -p=qwer -u=root --port=3306 -F=true
```

## 4. Support for gorm attributes
   
- Database tables, column field annotation support
- singular_table, Table name plural (big Camel-Case)
- json tag json tag output
- gorm.Model [Support export gorm.model>>>](doc/export.md)
- PRIMARY_KEY	Specifies column as primary key
- UNIQUE	Specifies column as unique
- NOT NULL	Specifies column as NOT NULL
- INDEX	Create index with or without name, same name creates composite indexes
- UNIQUE_INDEX	Like INDEX, create unique index
- Support foreign key related properties [Support export gorm.model>>>](doc/export.md)
- Support function export (foreign key, association, index , unique and more)[Support export function >>>](https://github.com/xxjwxc/gormt/blob/master/data/view/genfunc/genfunc_test.go)

### You can enrich data types in [def](data/view/cnf/def.go) 

## 5. Demonstration

- sql:
```
CREATE TABLE `user_account_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `account_type` int(11) NOT NULL DEFAULT '0' COMMENT '帐号类型:0手机号，1邮件',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT 'authbucket_oauth2_client表的id',
  `user_info_tbl_id` int(11) NOT NULL,
  `reg_time` datetime DEFAULT NULL,
  `reg_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `bundle_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `describ` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account` (`account`) USING BTREE,
  KEY `user_info_id` (`user_info_tbl_id`) USING BTREE,
  CONSTRAINT `user_account_tbl_ibfk_1` FOREIGN KEY (`user_info_tbl_id`) REFERENCES `user_info_tbl` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户账号'
```

###### --->Derived results

```
// UserAccountTbl 用户账号
type UserAccountTbl struct {
	ID            int    `gorm:"primary_key"`
	Account       string `gorm:"unique"`
	Password      string
	AccountType   int         // 帐号类型:0手机号，1邮件
	AppKey        string      // authbucket_oauth2_client表的id
	UserInfoTblID int         `gorm:"index"`
	UserInfoTbl   UserInfoTbl `gorm:"association_foreignkey:user_info_tbl_id;foreignkey:id"` // 用户信息
	RegTime       time.Time
	RegIP         string
	BundleID      string
	Describ       string
}
```

### [more>>>](doc/export.md)

## 6. support func export
### The exported function is only the auxiliary class function of Gorm, and calls Gorm completely
```
// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserAccountTblMgr) FetchByPrimaryKey(ID int) (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

```

### [more>>>](https://github.com/xxjwxc/gormt/tree/master/doc/func.md)
### [how to use call style>>>](https://github.com/xxjwxc/gormt/blob/master/data/view/genfunc/genfunc_test.go)

## 7. build
```
make windows
make linux
make mac
```
or

```
go generate
```

## 8. Next step 
- update，delete support
- revew

## 9. one windows gui tools

![1](/image/gormt/1.png)

![2](/image/gormt/2.jpg)

![3](/image/gormt/3.jpg)

![4](/image/gormt/4.jpg)

[Download](https://github.com/xxjwxc/gormt/releases/download/v1.1.0/v1.0.zip)



- ###### [link](https://xxjwxc.github.io/post/gormtools/)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/xxjwxc/gormt.svg)](https://starchart.cc/xxjwxc/gormt)

## 20200228 llllancelot foked
