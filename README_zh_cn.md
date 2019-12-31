[![Build Status](https://travis-ci.org/xxjwxc/gormt.svg?branch=master)](https://travis-ci.org/xxjwxc/gormt)
[![Go Report Card](https://goreportcard.com/badge/github.com/xxjwxc/gormt)](https://goreportcard.com/report/github.com/xxjwxc/gormt)
[![GoDoc](https://godoc.org/github.com/xxjwxc/gormt?status.svg)](https://godoc.org/github.com/xxjwxc/gormt)

# [gorm-tools](https://github.com/xxjwxc/gormt)

--------

#### mysql数据库转 struct 工具,可以将mysql数据库自动生成golang sturct结构，带大驼峰命名规则。带json标签


[English](README.md)

--------

## 1. 通过当前目录 config.yml 文件配置默认配置项
```
out_dir : "."  # 输出目录
singular_table : false  # 表名复数,是否大驼峰构建 参考:gorm.SingularTable
simple : false #简单输出
is_out_sql : false # 是否输出 sql 原信息
is_json_tag : false #是否打json标记
is_foreign_key : true #是否导出外键关联
mysql_info :
    host : "127.0.0.1"
    port : 3306
    username : "root"
    password : "qwer"
    database : "oauth_db"
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
  -f, --foreign           是否导出外键关联
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
- gorm.Model 基本模型   [支持gorm.Model模式导出>>>](https://github.com/xxjwxc/gormt/tree/master/doc/export_cn.md)
- PRIMARY_KEY	将列指定为主键
- UNIQUE	将列指定为唯一
- NOT NULL	将列指定为非 NULL
- INDEX	创建具有或不带名称的索引, 如果多个索引同名则创建复合索引
- UNIQUE_INDEX	和 INDEX 类似，只不过创建的是唯一索引
- 支持外键相关属性 [简单带外键模式导出>>>](https://github.com/xxjwxc/gormt/tree/master/doc/export_cn.md)

### 您可以在这里丰富数据映射类型 [def](data/view/cnf/def.go) 。

## 5. 示例展示
sql:
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

###### --->导出结果

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

### [更多>>>](https://github.com/xxjwxc/gormt/tree/master/doc/export_cn.md)

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

## 7. 下一步计划

- 加入相关快捷函数(OrmFunc)

## 8. 提供一个windows 可视化工具

![图片描述](/image/gormt/1.png)

![图片描述](/image/gormt/2.jpg)

![图片描述](/image/gormt/3.jpg)

![图片描述](/image/gormt/4.jpg)

[下载地址](https://github.com/xxjwxc/gormt/releases/download/v1.1.0/v1.0.zip)


- ###### [传送门](https://xxjwxc.github.io/post/gormtools/)
