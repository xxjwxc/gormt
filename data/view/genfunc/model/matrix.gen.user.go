package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID userId获取
func (obj *_UserMgr) WithUserID(UserID int) Option {
	return optionFunc(func(o *options) { o.query["userId"] = UserID })
}

// WithName name获取
func (obj *_UserMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// WithSex sex获取
func (obj *_UserMgr) WithSex(Sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = Sex })
}

// WithJob job获取
func (obj *_UserMgr) WithJob(Job int) Option {
	return optionFunc(func(o *options) { o.query["job"] = Job })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserID 通过userId获取内容
func (obj *_UserMgr) GetFromUserID(UserID int) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("userId = ?", UserID).Find(&result).Error

	return
}

// GetsBatchFromUserID 批量唯一主键查找
func (obj *_UserMgr) GetsBatchFromUserID(UserIDs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("userId IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_UserMgr) GetFromName(Name string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error

	return
}

// GetsBatchFromName 批量唯一主键查找
func (obj *_UserMgr) GetsBatchFromName(Names []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_UserMgr) GetFromSex(Sex int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", Sex).Find(&results).Error

	return
}

// GetsBatchFromSex 批量唯一主键查找
func (obj *_UserMgr) GetsBatchFromSex(Sexs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", Sexs).Find(&results).Error

	return
}

// GetFromJob 通过job获取内容
func (obj *_UserMgr) GetFromJob(Job int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("job = ?", Job).Find(&results).Error

	return
}

// GetsBatchFromJob 批量唯一主键查找
func (obj *_UserMgr) GetsBatchFromJob(Jobs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("job IN (?)", Jobs).Find(&results).Error

	return
}
