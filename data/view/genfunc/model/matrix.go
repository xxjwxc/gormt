package model

// Example [...]
type Example struct {
	UserID int    `gorm:"primary_key" json:"user_id"`
	Name   string `json:"name"`
	Sex    int    `gorm:"index" json:"sex"`
	Job    int    `json:"job"`
	ID     int    `json:"-"`
}

// Organ [...]
type Organ struct {
	ID       int    `gorm:"primary_key" json:"-"`
	UserID   int    `gorm:"index" json:"user_id"`
	UserList []User `gorm:"association_foreignkey:user_id;foreignkey:sex" json:"user_list"`
	Type     int    `json:"type"`
	Score    int    `json:"score"`
}

// User [...]
type User struct {
	UserID int    `gorm:"primary_key" json:"user_id"`
	Name   string `json:"name"`
	Sex    int    `gorm:"index" json:"sex"`
	Job    int    `json:"job"`
}
