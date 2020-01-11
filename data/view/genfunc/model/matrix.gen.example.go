package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _ExampleMgr struct {
	*_BaseMgr
}

// ExampleMgr open func
func ExampleMgr(db *gorm.DB) *_ExampleMgr {
	if db == nil {
		panic(fmt.Errorf("ExampleMgr need init by db"))
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

	return
}

// Gets 获取批量结果
func (obj *_ExampleMgr) Gets() (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_ExampleMgr) WithUserID(UserID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithName name获取
func (obj *_ExampleMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// WithSex sex获取
func (obj *_ExampleMgr) WithSex(Sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = Sex })
}

// WithJob job获取
func (obj *_ExampleMgr) WithJob(Job int) Option {
	return optionFunc(func(o *options) { o.query["job"] = Job })
}

// WithID id获取
func (obj *_ExampleMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// GetByOption 功能选项模式获取
func (obj *_ExampleMgr) GetByOption(opts ...Option) (result Example, err error) {
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
func (obj *_ExampleMgr) GetByOptions(opts ...Option) (results []*Example, err error) {
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
func (obj *_ExampleMgr) GetFromUserID(UserID int) (result Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_ExampleMgr) GetBatchFromUserID(UserIDs []int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ExampleMgr) GetFromName(Name string) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ExampleMgr) GetBatchFromName(Names []string) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_ExampleMgr) GetFromSex(Sex int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", Sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找
func (obj *_ExampleMgr) GetBatchFromSex(Sexs []int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", Sexs).Find(&results).Error

	return
}

// GetFromJob 通过job获取内容
func (obj *_ExampleMgr) GetFromJob(Job int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("job = ?", Job).Find(&results).Error

	return
}

// GetBatchFromJob 批量唯一主键查找
func (obj *_ExampleMgr) GetBatchFromJob(Jobs []int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("job IN (?)", Jobs).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_ExampleMgr) GetFromID(ID int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ExampleMgr) GetBatchFromID(IDs []int) (results []*Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ExampleMgr) FetchByPrimaryKey(UserID int) (result Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// FetchByIndex primay or index 获取唯一内容
func (obj *_ExampleMgr) FetchByIndex(Sex int) (result Example, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", Sex).Find(&result).Error

	return
}
