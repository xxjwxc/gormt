## 目录
 - [sql code](#sql-code)
	- [Mult complex export without foreign key](#Mult-complex-export-without-foreign-key)
	- [Complex single table mode export](#Complex-single-table-mode-export)
	- [Simple export with JSON](#Simple-export-with-JSON)
	- [Simple export without JSON](#Simple-export-without-JSON)
	- [Simple with foreign key mode export](#Simple-with-foreign-key-mode-export)
 - [sql2](#sql2)
 	- [Support export gorm.model](#Support export gorm.model)


### sql code

- sql:
```
CREATE TABLE `user_account_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `account_type` int(11) NOT NULL DEFAULT '0' COMMENT '[@gormt default:'123456']帐号类型:0手机号，1邮件',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT 'authbucket_oauth2_client表的id',
  `user_info_tbl_id` int(11) NOT NULL,
  `reg_time` datetime DEFAULT NULL,
  `reg_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `bundle_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `describ` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account` (`account`) USING BTREE,
  UNIQUE KEY `UNIQ_5696AD037D3656A4` (`app_key`,`user_info_tbl_id`) USING BTREE,
  KEY `user_info_id` (`user_info_tbl_id`) USING BTREE,
  CONSTRAINT `user_account_tbl_ibfk_1` FOREIGN KEY (`user_info_tbl_id`) REFERENCES `user_info_tbl` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='[@gormt default:'admin']用户账号'
```

-------------

### Mult complex export without foreign key

- param: simple = false  is_foreign_key = false

###### --->export result

```
// UserAccountTbl 用户账号
type UserAccountTbl struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Account       string    `gorm:"unique;column:account;type:varchar(64);not null" json:"account"`
	Password      string    `gorm:"column:password;type:varchar(64);not null" json:"password"`
	AccountType   int       `gorm:"column:account_type;type:int(11);not null" json:"account_type"`                               // 帐号类型:0手机号，1邮件
	AppKey        string    `gorm:"unique_index:UNIQ_5696AD037D3656A4;column:app_key;type:varchar(255);not null" json:"app_key"` // authbucket_oauth2_client表的id
	UserInfoTblID int       `gorm:"unique_index:UNIQ_5696AD037D3656A4;index;column:user_info_tbl_id;type:int(11);not null" json:"user_info_tbl_id"`
	RegTime       time.Time `gorm:"column:reg_time;type:datetime" json:"reg_time"`
	RegIP         string    `gorm:"column:reg_ip;type:varchar(15)" json:"reg_ip"`
	BundleID      string    `gorm:"column:bundle_id;type:varchar(255)" json:"bundle_id"`
	Describ       string    `gorm:"column:describ;type:varchar(255)" json:"describ"`
}
```

-------------

### Simple-export-with-JSON

- param: simple = true is_web_tag = true  is_foreign_key = false

###### --->export result

```
// UserAccountTbl 用户账号
type UserAccountTbl struct {
	ID            int       `gorm:"primary_key" json:"-"`
	Account       string    `gorm:"unique" json:"account"`
	Password      string    `json:"password"`
	AccountType   int       `json:"account_type"`                                      // 帐号类型:0手机号，1邮件
	AppKey        string    `gorm:"unique_index:UNIQ_5696AD037D3656A4" json:"app_key"` // authbucket_oauth2_client表的id
	UserInfoTblID int       `gorm:"unique_index:UNIQ_5696AD037D3656A4;index" json:"user_info_tbl_id"`
	RegTime       time.Time `json:"reg_time"`
	RegIP         string    `json:"reg_ip"`
	BundleID      string    `json:"bundle_id"`
	Describ       string    `json:"describ"`
}

```
--------------

### Simple export without JSON

- param: simple = true is_web_tag = false  is_foreign_key = false

###### --->export result

```
// UserAccountTbl 用户账号
type UserAccountTbl struct {
	ID            int    `gorm:"primary_key"`
	Account       string `gorm:"unique"`
	Password      string
	AccountType   int    // 帐号类型:0手机号，1邮件
	AppKey        string `gorm:"unique_index:UNIQ_5696AD037D3656A4"` // authbucket_oauth2_client表的id
	UserInfoTblID int    `gorm:"unique_index:UNIQ_5696AD037D3656A4;index"`
	RegTime       time.Time
	RegIP         string
	BundleID      string
	Describ       string
}
```

--------------

### Simple with foreign key mode export

- param: simple = true is_web_tag = false  is_foreign_key = true

###### --->export result

```
// UserAccountTbl 用户账号
type UserAccountTbl struct {
	ID            int    `gorm:"primary_key"`
	Account       string `gorm:"unique"`
	Password      string
	AccountType   int         // 帐号类型:0手机号，1邮件
	AppKey        string      `gorm:"unique_index:UNIQ_5696AD037D3656A4"` // authbucket_oauth2_client表的id
	UserInfoTblID int         `gorm:"unique_index:UNIQ_5696AD037D3656A4;index"`
	UserInfoTbl   UserInfoTbl `gorm:"association_foreignkey:user_info_tbl_id;foreignkey:id"` // 用户信息
	RegTime       time.Time
	RegIP         string
	BundleID      string
	Describ       string
}
```

--------------

## sql2
```
CREATE TABLE `user_info_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `headurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id` (`id`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户信息'
```

### Support export gorm.model

- param: simple = true is_web_tag = false 

###### --->export result


```
// UserInfoTbl 用户信息
type UserInfoTbl struct {
	gorm.Model
	Nickname string
	Headurl  string
}
```