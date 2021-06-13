package model

// Account [...]
type Account struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	AccountID int    `gorm:"uniqueIndex:account;column:account_id;type:int(11)" json:"accountId"`
	UserID    int    `gorm:"uniqueIndex:account;index:tp;column:user_id;type:int(11)" json:"userId"`
	User      User   `gorm:"joinForeignKey:user_id;foreignKey:user_id" json:"userList"`
	Type      int    `gorm:"index:tp;column:type;type:int(11)" json:"type"`
	Name      string `gorm:"column:name;type:varchar(255)" json:"name"`
}

// TableName get sql table name.获取数据库表名
func (m *Account) TableName() string {
	return "account"
}

// AccountColumns get sql column name.获取数据库列名
var AccountColumns = struct {
	ID        string
	AccountID string
	UserID    string
	Type      string
	Name      string
}{
	ID:        "id",
	AccountID: "account_id",
	UserID:    "user_id",
	Type:      "type",
	Name:      "name",
}

// User [...]
type User struct {
	UserID int    `gorm:"primaryKey;column:user_id;type:int(11);not null" json:"-"`
	Name   string `gorm:"column:name;type:varchar(30);not null" json:"name"`
	Sex    int    `gorm:"column:sex;type:int(11);not null" json:"sex"`
	Job    int    `gorm:"column:job;type:int(11);not null" json:"job"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	UserID string
	Name   string
	Sex    string
	Job    string
}{
	UserID: "user_id",
	Name:   "name",
	Sex:    "sex",
	Job:    "job",
}
