
# [gormt](https://github.com/xxjwxc/gormt)

## [中文文档](README_zh_cn.md)

###  mysql database to goang struct conversion tools base on [gorm](https://github.com/jinzhu/gorm)，You can automatically generate golang sturct from MySQL database.

### big Camel-Case Name Rule
### JSON tag 



--------

## 1. Configure default configuration items through the current directory config.yml file
```
out_dir : "."  # out dir
singular_table : false  # Table name plural (big Camel-Case):gorm.SingularTable
simple : false #simple output
is_json_tag : true # Whether to mark JSON or not
is_foreign_key : true # Whether to mark foreign key or not
[mysql_info]
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
  -h, --help              help for main
  -H, --host string       数据库地址.(注意-H为大写)
  -o, --outdir string     输出目录
  -p, --password string   密码.
      --port int          端口号 (default 3306)
  -s, --singular          是否禁用表名复数
  -u, --user string       用户名.
  
```
## 3. Can be updated configuration items using command line tools
```
./gormt -H=127.0.0.1 -d=oauth_db -p=qwer -u=root --port=3306
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

## 6. build
```
make windows
make linux
make mac
```
or

```
go generate
```

## 7. Next step 

- Add common function (ormfunc)

## 8. one windows gui tools

![1](/image/gormt/1.png)

![2](/image/gormt/2.jpg)

![3](/image/gormt/3.jpg)

![4](/image/gormt/4.jpg)

[Download](https://github.com/xxjwxc/gormt/releases/download/v1.1.0/v1.0.zip)



- ###### [link](https://xxjwxc.github.io/post/gormtools/)
