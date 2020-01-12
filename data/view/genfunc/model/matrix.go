package model

// Account [...]
type Account struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	AccountID int    `gorm:"unique_index:account;column:account_id;type:int(11)" json:"account_id"`
	UserID    int    `gorm:"unique_index:account;index:tp;column:user_id;type:int(11)" json:"user_id"`
	User      User   `gorm:"association_foreignkey:user_id;foreignkey:user_id" json:"user_list"`
	Type      int    `gorm:"index:tp;column:type;type:int(11)" json:"type"`
	Name      string `gorm:"column:name;type:varchar(255)" json:"name"`
}

// User [...]
type User struct {
	UserID int    `gorm:"primary_key;column:user_id;type:int(11);not null" json:"user_id"`
	Name   string `gorm:"column:name;type:varchar(30);not null" json:"name"`
	Sex    int    `gorm:"column:sex;type:int(11);not null" json:"sex"`
	Job    int    `gorm:"column:job;type:int(11);not null" json:"job"`
}
