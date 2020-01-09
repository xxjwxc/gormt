package genfunc

const (
	genBase = `
package {{.PackageName}}
import (
	"context"

	"github.com/jinzhu/gorm"
)

// prepare for outher
type _BaseMgr struct {
	*gorm.DB
	ctx       *context.Context
	isRelated bool
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c *context.Context) {
	obj.ctx = c
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// IsRelated Query foreign key Association.是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) IsRelated(b bool) {
	obj.isRelated = b
}

type options struct {
	query map[string]interface{}
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}
	`

	genlogic = `{{$obj := .}}{{$list := $obj.Em}}
type _{{$obj.StructName}}Mgr struct {
	*_BaseMgr
}

// {{$obj.StructName}}Mgr open func
func {{$obj.StructName}}Mgr(db *gorm.DB) *_{{$obj.StructName}}Mgr {
	if db == nil {
		panic(fmt.Errorf("{{$obj.StructName}}Mgr need init by db"))
	}
	return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_{{$obj.StructName}}Mgr) GetTableName() string {
	return "{{$obj.TableName}}"
}

// Get 获取
func (obj *_{{$obj.StructName}}Mgr) Get() (result {{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}

// Gets 获取批量结果
func (obj *_{{$obj.StructName}}Mgr) Gets() (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}

//////////////////////////option case ////////////////////////////////////////////
{{range $oem := $obj.Em}}
// With{{$oem.ColStructName}} {{$oem.ColName}}获取 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) With{{$oem.ColStructName}}({{$oem.ColStructName}} {{$oem.Type}}) Option {
	return optionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{$oem.ColStructName}} })
}
{{end}}

// GetByOption 功能选项模式获取
func (obj *_{{$obj.StructName}}Mgr) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_{{$obj.StructName}}Mgr) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	{{GenPreloadList $obj.PreloadList true}}
	return
}
//////////////////////////enume case ////////////////////////////////////////////

{{range $oem := $obj.Em}}
// GetFrom{{$oem.ColStructName}} 通过{{$oem.ColName}}获取内容 {{$oem.Notes}} {{if $oem.IsMulti}}
func (obj *_{{$obj.StructName}}Mgr) GetFrom{{$oem.ColStructName}}({{$oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{$oem.ColName}} = ?", {{$oem.ColStructName}}).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
{{else}}
func (obj *_{{$obj.StructName}}Mgr)  GetFrom{{$oem.ColStructName}}({{$oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{$oem.ColName}} = ?", {{$oem.ColStructName}}).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}
{{end}}
// GetsBatchFrom{{$oem.ColStructName}} 批量唯一主键查找 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) GetsBatchFrom{{$oem.ColStructName}}({{$oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{$oem.ColName}} IN (?)", {{$oem.ColStructName}}s).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
 {{end}}
`
	genPreload = `if err == nil && obj.isRelated { {{range $obj := .}}{{if $obj.IsMulti}}
		{
			var info []{{$obj.ForeignkeyStructName}}  // {{$obj.Notes}} 
			err = obj.DB.Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", result.{{$obj.ColStructName}}).Find(&info).Error
			if err != nil {
				return
			}
			result.{{$obj.ForeignkeyStructName}}List = info
		}  {{else}} 
		{
			var info {{$obj.ForeignkeyStructName}}  // {{$obj.Notes}} 
			err = obj.DB.Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", result.{{$obj.ColStructName}}).Find(&info).Error
			if err != nil {
				return
			}
			result.{{$obj.ForeignkeyStructName}} = info
		} {{end}} {{end}}
	}
`
	genPreloadMulti = `if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ { {{range $obj := .}}{{if $obj.IsMulti}}
		{
			var info []{{$obj.ForeignkeyStructName}}  // {{$obj.Notes}} 
			err = obj.DB.Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", results[i].{{$obj.ColStructName}}).Find(&info).Error
			if err != nil {
				return
			}
			results[i].{{$obj.ForeignkeyStructName}}List = info
		}  {{else}} 
		{
			var info {{$obj.ForeignkeyStructName}}  // {{$obj.Notes}} 
			err = obj.DB.Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", results[i].{{$obj.ColStructName}}).Find(&info).Error
			if err != nil {
				return
			}
			results[i].{{$obj.ForeignkeyStructName}} = info
		} {{end}} {{end}}
	}
}`
)
