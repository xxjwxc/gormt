package genfunc

const (
	genTnf = `
// TableName get sql table name.获取数据库表名
func (m *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
	genBase = `
package {{.PackageName}}
import (
	"context"

	"github.com/jinzhu/gorm"
)

var globalIsRelated bool = true  // 全局预加载

// prepare for other
type BaseMgr struct {
	*gorm.DB
	ctx       *context.Context
	isRelated bool
}

// SetCtx set context
func (obj *BaseMgr) SetCtx(c *context.Context) {
	obj.ctx = c
}

// GetDB get gorm.DB info
func (obj *BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// GetIsRelated Query foreign key Association.获取是否查询外键关联(gorm.Related)
func (obj *BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

// SetIsRelated Query foreign key Association.设置是否查询外键关联(gorm.Related)
func (obj *BaseMgr) SetIsRelated(b bool) {
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


// OpenRelated 打开全局预加载
func OpenRelated() {
	globalIsRelated = true
}

// CloseRelated 关闭全局预加载
func CloseRelated() {
	globalIsRelated = true
}

	`

	genlogic = `{{$obj := .}}{{$list := $obj.Em}}
type {{$obj.StructName}}Mgr struct {
	*BaseMgr
}

// New{{$obj.StructName}}Mgr open func
func New{{$obj.StructName}}Mgr(db *gorm.DB) *{{$obj.StructName}}Mgr {
	if db == nil {
		panic(fmt.Errorf("{{$obj.StructName}}Mgr need init by db"))
	}
	return &{{$obj.StructName}}Mgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *{{$obj.StructName}}Mgr) GetTableName() string {
	return "{{$obj.TableName}}"
}

// Get 获取
func (obj *{{$obj.StructName}}Mgr) Get() (result {{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}

// Gets 获取批量结果
func (obj *{{$obj.StructName}}Mgr) Gets() (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}

//////////////////////////option case ////////////////////////////////////////////
{{range $oem := $obj.Em}}
// With{{$oem.ColStructName}} {{$oem.ColName}}获取 {{$oem.Notes}}
func (obj *{{$obj.StructName}}Mgr) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
	return optionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
}
{{end}}

// GetByOption 功能选项模式获取
func (obj *{{$obj.StructName}}Mgr) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
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
func (obj *{{$obj.StructName}}Mgr) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
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
func (obj *{{$obj.StructName}}Mgr) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{$oem.ColName}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
{{else}}
func (obj *{{$obj.StructName}}Mgr)  GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{$oem.ColName}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}
{{end}}
// GetBatchFrom{{$oem.ColStructName}} 批量唯一主键查找 {{$oem.Notes}}
func (obj *{{$obj.StructName}}Mgr) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{$oem.ColName}} IN (?)", {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
 {{end}}
 //////////////////////////primary index case ////////////////////////////////////////////
 {{range $ofm := $obj.Primay}}
 // {{GenFListIndex $ofm 1}} primay or index 获取唯一内容
 func (obj *{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}
 {{end}}

 {{range $ofm := $obj.Index}}
 // {{GenFListIndex $ofm 1}}  获取多个内容
 func (obj *{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
 {{end}}

`
	genPreload = `if err == nil && obj.isRelated { {{range $obj := .}}{{if $obj.IsMulti}}
		err = obj.DB.New().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", result.{{$obj.ColStructName}}).Find(&result.{{$obj.ForeignkeyStructName}}List).Error // {{$obj.Notes}}
		{{else}} 
		err = obj.DB.New().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", result.{{$obj.ColStructName}}).Find(&result.{{$obj.ForeignkeyStructName}}).Error // {{$obj.Notes}} 
		{{end}} {{end}}}
`
	genPreloadMulti = `if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ { {{range $obj := .}}{{if $obj.IsMulti}}
		if err = obj.DB.New().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", results[i].{{$obj.ColStructName}}).Find(&results[i].{{$obj.ForeignkeyStructName}}List).Error;err != nil { // {{$obj.Notes}}
				return	
			} {{else}} 
		if err = obj.DB.New().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", results[i].{{$obj.ColStructName}}).Find(&results[i].{{$obj.ForeignkeyStructName}}).Error; err != nil { // {{$obj.Notes}} 
				return
			} {{end}} {{end}}
	}
}`
)
