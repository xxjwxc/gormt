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

// GetTableName get sql table name.获取数据库名字
func (obj *_ExampleMgr) GetTableName() string {
	return "example"
}

// Get 获取
func (obj *_ExampleMgr) Get() (result Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		var info []User
		err = obj.DB.Where("job = ?", result.UserID).Find(&info).Error
		if err != nil {
			return
		}
		result.UserList = info
	}
	return
}

// Gets 获取批量结果
func (obj *_ExampleMgr) Gets() (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
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

// GetFromID 通过id获取内容
func (obj *_ExampleMgr) GetFromID(id int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error
	if err == nil && obj.isRelated {
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
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
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
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
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

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
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

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	if err == nil && obj.isRelated {
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
	return optionFunc(func(o *options) { o.query["id"] = id })
}

func (obj *_ExampleMgr) WithUserID(id int64) Option {
	return optionFunc(func(o *options) {
		o.query["user_id"] = id
	})
}

///////////////////////////////////////////////////
