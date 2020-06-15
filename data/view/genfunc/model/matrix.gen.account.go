package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _AccountMgr struct {
	*BaseMgr
}

// AccountMgr open func
func AccountMgr(db *gorm.DB) *_AccountMgr {
	if db == nil {
		panic(fmt.Errorf("AccountMgr need init by db"))
	}
	return &_AccountMgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AccountMgr) GetTableName() string {
	return "account"
}

// Get 获取
func (obj *_AccountMgr) Get() (result Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AccountMgr) Gets() (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AccountMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithAccountID account_id获取
func (obj *_AccountMgr) WithAccountID(AccountID int) Option {
	return optionFunc(func(o *options) { o.query["account_id"] = AccountID })
}

// WithUserID user_id获取
func (obj *_AccountMgr) WithUserID(UserID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithType type获取
func (obj *_AccountMgr) WithType(Type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = Type })
}

// WithName name获取
func (obj *_AccountMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// GetByOption 功能选项模式获取
func (obj *_AccountMgr) GetByOption(opts ...Option) (result Account, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
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

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AccountMgr) GetFromID(ID int) (result Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_AccountMgr) GetBatchFromID(IDs []int) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

// GetFromAccountID 通过account_id获取内容
func (obj *_AccountMgr) GetFromAccountID(AccountID int) (result Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account_id = ?", AccountID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
		}
	}

	return
}

// GetBatchFromAccountID 批量唯一主键查找
func (obj *_AccountMgr) GetBatchFromAccountID(AccountIDs []int) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account_id IN (?)", AccountIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AccountMgr) GetFromUserID(UserID int) (result Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
		}
	}

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_AccountMgr) GetBatchFromUserID(UserIDs []int) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

// GetFromType 通过type获取内容
func (obj *_AccountMgr) GetFromType(Type int) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", Type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_AccountMgr) GetBatchFromType(Types []int) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", Types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_AccountMgr) GetFromName(Name string) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_AccountMgr) GetBatchFromName(Names []string) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AccountMgr) FetchByPrimaryKey(ID int) (result Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
		}
	}

	return
}

// FetchUniqueIndexByAccount primay or index 获取唯一内容
func (obj *_AccountMgr) FetchUniqueIndexByAccount(AccountID int, UserID int) (result Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account_id = ? AND user_id = ?", AccountID, UserID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info User //
			err = obj.DB.New().Table("user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.User = info
		}
	}

	return
}

// FetchIndexByTp  获取多个内容
func (obj *_AccountMgr) FetchIndexByTp(UserID int, Type int) (results []*Account, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ? AND type = ?", UserID, Type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info User //
				err = obj.DB.New().Table("user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].User = info
			}
		}
	}
	return
}
