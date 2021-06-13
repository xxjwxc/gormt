package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AccountMgr struct {
	*_BaseMgr
}

// AccountMgr open func
func AccountMgr(db *gorm.DB) *_AccountMgr {
	if db == nil {
		panic(fmt.Errorf("AccountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AccountMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Account{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AccountMgr) GetTableName() string {
	return "account"
}

// Get 获取
func (obj *_AccountMgr) Get() (result Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user").Where("user_id = ?", result.UserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AccountMgr) Gets() (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AccountMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAccountID account_id获取
func (obj *_AccountMgr) WithAccountID(accountID int) Option {
	return optionFunc(func(o *options) { o.query["account_id"] = accountID })
}

// WithUserID user_id获取
func (obj *_AccountMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithType type获取
func (obj *_AccountMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithName name获取
func (obj *_AccountMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_AccountMgr) GetByOption(opts ...Option) (result Account, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user").Where("user_id = ?", result.UserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AccountMgr) GetByOptions(opts ...Option) (results []*Account, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AccountMgr) GetFromID(id int) (result Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user").Where("user_id = ?", result.UserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_AccountMgr) GetBatchFromID(ids []int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAccountID 通过account_id获取内容
func (obj *_AccountMgr) GetFromAccountID(accountID int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`account_id` = ?", accountID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAccountID 批量查找
func (obj *_AccountMgr) GetBatchFromAccountID(accountIDs []int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`account_id` IN (?)", accountIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AccountMgr) GetFromUserID(userID int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量查找
func (obj *_AccountMgr) GetBatchFromUserID(userIDs []int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromType 通过type获取内容
func (obj *_AccountMgr) GetFromType(_type int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`type` = ?", _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromType 批量查找
func (obj *_AccountMgr) GetBatchFromType(_types []int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`type` IN (?)", _types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_AccountMgr) GetFromName(name string) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找
func (obj *_AccountMgr) GetBatchFromName(names []string) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AccountMgr) FetchByPrimaryKey(id int) (result Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user").Where("user_id = ?", result.UserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByAccount primary or index 获取唯一内容
func (obj *_AccountMgr) FetchUniqueIndexByAccount(accountID int, userID int) (result Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`account_id` = ? AND `user_id` = ?", accountID, userID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user").Where("user_id = ?", result.UserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByTp  获取多个内容
func (obj *_AccountMgr) FetchIndexByTp(userID int, _type int) (results []*Account, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ? AND `type` = ?", userID, _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
