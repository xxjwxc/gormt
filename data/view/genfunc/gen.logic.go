package genfunc

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Example demo of root
type Example struct {
	ID       int    `json:"-"`
	UserID   int    `json:"userId"`
	UserList []User `gorm:"association_foreignkey:userId;foreignkey:job" json:"user_list"`
}

// User demo of user
type User struct {
	UserID int `json:"userId"`
	Job    int `json:"job"`
}

////////////////////////////////////////////-----logic------

type _ExampleMgr struct {
	*_BaseMgr
}

// ExampleMgr open func
func ExampleMgr(db *gorm.DB) *_ExampleMgr {
	if db == nil {
		panic(fmt.Errorf("ExampleMgr init need db"))
	}
	return &_ExampleMgr{_BaseMgr: &_BaseMgr{DB: db}}
}

// GetFromID 通过id获取内容
func (obj *_ExampleMgr) GetFromID(id int) (results []*Example, err error) {
	err = obj.DB.Table("example").Where("id = ?", id).Find(&results).Error
	if err == nil {
		for i := 0; i < len(results); i++ {
			var userList []User
			err = obj.DB.Where("job = ?", results[i].UserID).Find(&userList).Error
			if err != nil {
				return
			}
			results[i].UserList = userList
		}
	}
	return
}

// GetByPrimaryKey 唯一主键查找
func (obj *_ExampleMgr) GetByPrimaryKey(id int64) (result Example, err error) {
	err = obj.DB.Table("example").Where("id = ?", id).Find(&result).Error
	if err == nil {
		var info []User
		err = obj.DB.Where("job = ?", result.UserID).Find(&info).Error
		if err != nil {
			return
		}
		result.UserList = info
	}
	return
}

// GetByPrimaryKey 批量唯一主键查找
func (obj *_ExampleMgr) GetByPrimaryKeys(ids []int64) (results []*Example, err error) {
	err = obj.DB.Table("example").Where("id IN (?)", ids).Find(&results).Error
	if err == nil {
		for i := 0; i < len(results); i++ {
			var userList []User
			err = obj.DB.Where("job = ?", results[i].UserID).Find(&userList).Error
			if err != nil {
				return
			}
			results[i].UserList = userList
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// GetByPrimaryKey 功能选项模式获取
func (obj *_ExampleMgr) GetByOption(opts ...Option) (result Example, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table("example").Where(options.query).Find(&result).Error
	if err == nil {
		var info []User
		err = obj.DB.Where("job = ?", result.UserID).Find(&info).Error
		if err != nil {
			return
		}
		result.UserList = info
	}
	return
}

// GetByPrimaryKey 批量功能选项模式获取
func (obj *_ExampleMgr) GetByOptions(opts ...Option) (results []*Example, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table("example").Where(options.query).Find(&results).Error
	if err == nil {
		for i := 0; i < len(results); i++ {
			var userList []User
			err = obj.DB.Where("job = ?", results[i].UserID).Find(&userList).Error
			if err != nil {
				return
			}
			results[i].UserList = userList
		}
	}
	return
}

// WithID id获取
func (obj *_ExampleMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) {
		o.query["id"] = id
	})
}

func (obj *_ExampleMgr) WithUserID(id int64) Option {
	return optionFunc(func(o *options) {
		o.query["user_id"] = id
	})
}

///////////////////////////////////////////////////
