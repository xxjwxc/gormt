package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserMgr struct {
	*BaseMgr
}

// NewUserMgr open func
func NewUserMgr(db *gorm.DB) *UserMgr {
	if db == nil {
		panic(fmt.Errorf("NewUserMgr need init by db"))
	}
	return &UserMgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *UserMgr) Get() (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *UserMgr) WithUserID(UserID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithName name获取
func (obj *UserMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// WithSex sex获取
func (obj *UserMgr) WithSex(Sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = Sex })
}

// WithJob job获取
func (obj *UserMgr) WithJob(Job int) Option {
	return optionFunc(func(o *options) { o.query["job"] = Job })
}

// GetByOption 功能选项模式获取
func (obj *UserMgr) GetByOption(opts ...Option) (result User, err error) {
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
func (obj *UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
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

// GetFromUserID 通过user_id获取内容
func (obj *UserMgr) GetFromUserID(UserID int) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *UserMgr) GetBatchFromUserID(UserIDs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *UserMgr) GetFromName(Name string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *UserMgr) GetBatchFromName(Names []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *UserMgr) GetFromSex(Sex int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", Sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找
func (obj *UserMgr) GetBatchFromSex(Sexs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", Sexs).Find(&results).Error

	return
}

// GetFromJob 通过job获取内容
func (obj *UserMgr) GetFromJob(Job int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("job = ?", Job).Find(&results).Error

	return
}

// GetBatchFromJob 批量唯一主键查找
func (obj *UserMgr) GetBatchFromJob(Jobs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("job IN (?)", Jobs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *UserMgr) FetchByPrimaryKey(UserID int) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}
