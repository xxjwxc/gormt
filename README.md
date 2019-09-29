
# [gormt](https://github.com/xxjwxc/gormt)
###  mysql database to goang struct conversion tools base on [gorm](https://github.com/jinzhu/gorm)，You can automatically generate golang sturct from MySQL database.

### big Camel-Case Name Rule
### JSON tag 

[中文文档](README_zh_cn.md)

--------

## 1. Configure default configuration items through the current directory config. toml file
```
out_dir = "."  # out dir
singular_table = false  # Table name plural (big Camel-Case):gorm.SingularTable
simple = false #simple output
isJsonTag = true # Whether to mark JSON or not
[mysql_info]
    host = "127.0.0.1"
    port = 3306
    username = "root"
    password = "qwer"
    database = "oauth_db"

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
- gorm.Model 
- PRIMARY_KEY	Specifies column as primary key
- UNIQUE	Specifies column as unique
- NOT NULL	Specifies column as NOT NULL
- INDEX	Create index with or without name, same name creates composite indexes
- UNIQUE_INDEX	Like INDEX, create unique index

### You can enrich data types in [def](https://github.com/xxjwxc/gormt/blob/master/data/view/gtools/def.go)


## 5. build
```
make windows
make linux
make mac
```
or

```
go generate
```

## 6. Demonstration

- sql:
```
CREATE TABLE `user_account_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `account_type` int(11) NOT NULL DEFAULT '0' COMMENT '帐号类型:0手机号，1邮件',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT 'authbucket_oauth2_client表的id',
  `user_info_id` int(11) NOT NULL,
  `reg_time` datetime DEFAULT NULL,
  `reg_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `bundle_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `describ` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account` (`account`),
  UNIQUE KEY `UNIQ_5696AD037D3656A4` (`app_key`,`user_info_id`) USING BTREE,
  KEY `user_info_id` (`user_info_id`),
  CONSTRAINT `1` FOREIGN KEY (`user_info_id`) REFERENCES `user_info_tbl` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户信息'
```
- param :singular_table = false simple = false 

###### --->Derived results

```
//	User information
type UserAccountTbl struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`                                                   //
	Account     string    `gorm:"unique;column:account;type:varchar(64);not null" json:"account"`                                         //
	Password    string    `gorm:"column:password;type:varchar(64);not null" json:"password"`                                              //
	AccountType int       `gorm:"column:account_type;type:int(11);not null" json:"account_type"`                                          //	帐号类型:0手机号，1邮件
	AppKey      string    `json:"app_key" gorm:"unique_index:UNIQ_5696AD037D3656A4;column:app_key;type:varchar(255);not null"`            //	authbucket_oauth2_client表的id
	UserInfoID  int       `gorm:"unique_index:UNIQ_5696AD037D3656A4;index;column:user_info_id;type:int(11);not null" json:"user_info_id"` //
	RegTime     time.Time `gorm:"column:reg_time;type:datetime" json:"reg_time"`                                                          //
	RegIP       string    `gorm:"column:reg_ip;type:varchar(15)" json:"reg_ip"`                                                           //
	BundleID    string    `json:"bundle_id" gorm:"column:bundle_id;type:varchar(255)"`                                                    //
	Describ     string    `gorm:"column:describ;type:varchar(255)" json:"describ"`                                                        //
}
```

###### --->erived results

```
//	用户信息
type UserAccountTbl struct {
	ID          int       `json:"-" gorm:"primary_key"`                                         //
	Account     string    `gorm:"unique" json:"account"`                                        //
	Password    string    `json:"password"`                                                     //
	AccountType int       `json:"account_type"`                                                 //	帐号类型:0手机号，1邮件
	AppKey      string    `gorm:"unique_index:UNIQ_5696AD037D3656A4" json:"app_key"`            //	authbucket_oauth2_client表的id
	UserInfoID  int       `json:"user_info_id" gorm:"unique_index:UNIQ_5696AD037D3656A4;index"` //
	RegTime     time.Time `json:"reg_time"`                                                     //
	RegIP       string    `json:"reg_ip"`                                                       //
	BundleID    string    `json:"bundle_id"`                                                    //
	Describ     string    `json:"describ"`                                                      //
}
```

- param :singular_table = false simple = true isJsonTag = false

###### --->erived results

```
//	用户信息
type UserAccountTbl struct {
	ID          int       `gorm:"primary_key"` //
	Account     string    `gorm:"unique"`      //
	Password    string    //
	AccountType int       //	帐号类型:0手机号，1邮件
	AppKey      string    `gorm:"unique_index:UNIQ_5696AD037D3656A4"`       //	authbucket_oauth2_client表的id
	UserInfoID  int       `gorm:"unique_index:UNIQ_5696AD037D3656A4;index"` //
	RegTime     time.Time //
	RegIP       string    //
	BundleID    string    //
	Describ     string    //
}
```

- sql:
```
CREATE TABLE `user_info_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `headurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户信息'
```

- param :singular_table = false simple = true isJsonTag = false

###### --->erived results


```
//	用户信息
type UserInfoTbl struct {
	gorm.Model        //
	Nickname   string //
	Headurl    string //
}
```


## 7. one windows gui tools

![1](/image/gormt/1.png)

![2](/image/gormt/2.jpg)

![3](/image/gormt/3.jpg)

![4](/image/gormt/4.jpg)

[Download](https://github.com/xxjwxc/gormt/releases/download/v1.1.0/v1.0.zip)

## 8. Next step 

- support (ForeignKey)

- ###### [link](https://xxjwxc.github.io/post/gormtools/)
