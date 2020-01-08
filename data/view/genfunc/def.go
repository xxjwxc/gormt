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
		panic(fmt.Errorf("{{$obj.StructName}}Mgr init need db"))
	}
	return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_{{$obj.StructName}}Mgr) GetTableName() string {
	return "{{$obj.TableName}}"
}
`
)
