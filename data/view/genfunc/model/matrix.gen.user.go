package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Model(User{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_UserMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithName name获取
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSex sex获取
func (obj *_UserMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithJob job获取
func (obj *_UserMgr) WithJob(job int) Option {
	return optionFunc(func(o *options) { o.query["job"] = job })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

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

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserID 通过user_id获取内容
func (obj *_UserMgr) GetFromUserID(userID int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_UserMgr) GetBatchFromUserID(userIDs []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_UserMgr) GetFromSex(sex int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找
func (obj *_UserMgr) GetBatchFromSex(sexs []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromJob 通过job获取内容
func (obj *_UserMgr) GetFromJob(job int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`job` = ?", job).Find(&results).Error

	return
}

// GetBatchFromJob 批量查找
func (obj *_UserMgr) GetBatchFromJob(jobs []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`job` IN (?)", jobs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(userID int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&result).Error

	return
}
