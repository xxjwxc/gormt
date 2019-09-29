
# [gorm-tools](https://github.com/xxjwxc/gorm-tools)
### [gorm](https://github.com/jinzhu/gorm) mysql数据库转 struct 工具,可以将mysql数据库自动生成golang sturct结构，带大驼峰命名规则。带json标签


[English](README.md)
--------

## 1. 通过当前目录config.toml文件配置默认配置项
```
out_dir = "."  # 输出目录
singular_table = false  # 表名复数,是否大驼峰构建 参考:gorm.SingularTable
simple = false #简单输出
isJsonTag = true #是否打json标记
[mysql_info]
    host = "127.0.0.1"
    port = 3306
    username = "root"
    password = "qwer"
    database = "oauth_db"

```
## 2. 可以使用命令行工具更新配置项
```
./gormt -H=127.0.0.1 -d=oauth_db -p=qwer -u=root --port=3306
```
## 3. 查看帮助
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

## 4. 支持gorm 相关属性 
   
- 数据库表,列字段注释支持
- singular_table 表名复数(大驼峰)
- json tag json标签输出
- gorm.Model 基本模型
- PRIMARY_KEY	将列指定为主键
- UNIQUE	将列指定为唯一
- NOT NULL	将列指定为非 NULL
- INDEX	创建具有或不带名称的索引, 如果多个索引同名则创建复合索引
- UNIQUE_INDEX	和 INDEX 类似，只不过创建的是唯一索引

### 您可以在这里丰富数据映射类型 [def](https://github.com/xxjwxc/gormt/blob/master/data/view/gtools/def.go) 。

## 5. 示例展示

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
- 参数:singular_table = false simple = false 

###### --->导出结果

```
//	用户信息
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

- 参数:singular_table = true simple = false 

###### --->导出结果

```
type User_account_tbl struct {
	Id           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`                                                   //
	Account      string    `gorm:"unique;column:account;type:varchar(64);not null" json:"account"`                                         //
	Password     string    `gorm:"column:password;type:varchar(64);not null" json:"password"`                                              //
	Account_type int       `gorm:"column:account_type;type:int(11);not null" json:"account_type"`                                          //	帐号类型:0手机号，1邮件
	App_key      string    `gorm:"unique_index:UNIQ_5696AD037D3656A4;column:app_key;type:varchar(255);not null" json:"app_key"`            //	authbucket_oauth2_client表的id
	User_info_id int       `gorm:"unique_index:UNIQ_5696AD037D3656A4;index;column:user_info_id;type:int(11);not null" json:"user_info_id"` //
	Reg_time     time.Time `gorm:"column:reg_time;type:datetime" json:"reg_time"`                                                          //
	Reg_ip       string    `gorm:"column:reg_ip;type:varchar(15)" json:"reg_ip"`                                                           //
	Bundle_id    string    `gorm:"column:bundle_id;type:varchar(255)" json:"bundle_id"`                                                    //
	Describ      string    `gorm:"column:describ;type:varchar(255)" json:"describ"`                                                        //
}
```
- 参数:singular_table = false simple = true isJsonTag = true

###### --->导出结果

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

- 参数:singular_table = false simple = true isJsonTag = false

###### --->导出结果

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

- 参数:singular_table = false simple = true isJsonTag = false

###### --->导出结果


```
//	用户信息
type UserInfoTbl struct {
	gorm.Model        //
	Nickname   string //
	Headurl    string //
}
```
## 6. 构建
```
make windows
make linux
make mac
```
or

```
go generate
```

## 7. 提供一个windows 可视化工具

![图片描述](/image/gormt/1.png)

![图片描述](/image/gormt/2.jpg)

![图片描述](/image/gormt/3.jpg)

![图片描述](/image/gormt/4.jpg)

[下载地址](https://github.com/xxjwxc/gormt/releases/download/v1.1.0/v1.0.zip)

## 7. 下一步计划

- 加入外键支持(ForeignKey)

- ###### [传送门](https://xxjwxc.github.io/post/gormtools/)
