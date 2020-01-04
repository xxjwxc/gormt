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
)
