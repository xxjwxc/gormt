package genfunc

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

type _BaseMgr struct {
	*gorm.DB
	ctx *context.Context
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c *context.Context) {
	obj.ctx = c
}

////////////////////////////////////////////logic

type _ExampleMgr struct {
	*_BaseMgr
}

// ExampleMgr open func
func ExampleMgr(db *gorm.DB) *_ExampleMgr {
	if db == nil {
		panic(fmt.Errorf("ExampleMgr init need db"))
	}
	return &_ExampleMgr{_BaseMgr: &_BaseMgr{DB: db}}
}

///////////////////////////////////////////////////
