## 5. 导出

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
  UNIQUE KEY `UNIQ_5696AD037D3656A4` (`app_key`,`user_info_tbl_id`) USING BTREE,
  KEY `user_info_id` (`user_info_tbl_id`) USING BTREE,
  CONSTRAINT `user_account_tbl_ibfk_1` FOREIGN KEY (`user_info_tbl_id`) REFERENCES `user_info_tbl` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户账号'
```

-------------

- 参数:singular_table = false simple = false  is_foreign_key = false

###### --->导出结果

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

- 参数:singular_table = true simple = false  is_foreign_key = false

###### --->导出结果

```
// UserAccountTbl 用户账号
type User_account_tbl struct {
	Id               int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Account          string    `gorm:"unique;column:account;type:varchar(64);not null" json:"account"`
	Password         string    `gorm:"column:password;type:varchar(64);not null" json:"password"`
	Account_type     int       `gorm:"column:account_type;type:int(11);not null" json:"account_type"`                               // 帐号类型:0手机号，1邮件
	App_key          string    `gorm:"unique_index:UNIQ_5696AD037D3656A4;column:app_key;type:varchar(255);not null" json:"app_key"` // authbucket_oauth2_client表的id
	User_info_tbl_id int       `gorm:"unique_index:UNIQ_5696AD037D3656A4;index;column:user_info_tbl_id;type:int(11);not null" json:"user_info_tbl_id"`
	Reg_time         time.Time `gorm:"column:reg_time;type:datetime" json:"reg_time"`
	Reg_ip           string    `gorm:"column:reg_ip;type:varchar(15)" json:"reg_ip"`
	Bundle_id        string    `gorm:"column:bundle_id;type:varchar(255)" json:"bundle_id"`
	Describ          string    `gorm:"column:describ;type:varchar(255)" json:"describ"`
}
```

-------------

- 参数:singular_table = false simple = true is_json_tag = true  is_foreign_key = false

###### --->导出结果

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

- 参数:singular_table = false simple = true is_json_tag = false  is_foreign_key = false

###### --->导出结果

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

- 参数:singular_table = false simple = true is_json_tag = false  is_foreign_key = true

###### --->导出结果

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

- sql:
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

- 参数:singular_table = false simple = true is_json_tag = false 

###### --->导出结果


```
// UserInfoTbl 用户信息
type UserInfoTbl struct {
	gorm.Model
	Nickname string
	Headurl  string
}
```