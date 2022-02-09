package genfunc

const (
	genTnf = `
// TableName get sql table name.获取数据库表名
func (m *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
	genColumn = `
// {{.StructName}}Columns get sql column name.获取数据库列名
var {{.StructName}}Columns = struct { {{range $em := .Em}}
	{{$em.StructName}} string{{end}}    
	}{ {{range $em := .Em}}
		{{$em.StructName}}:"{{$em.ColumnName}}",  {{end}}           
	}
`
	genBase = `
package {{.PackageName}}
import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var globalIsRelated bool = true  // 全局预加载

// prepare for other
type _BaseMgr struct {
	*gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

// SetTimeOut set timeout
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c context.Context) {
	if c != nil {
		obj.ctx = c
	}
}

// GetCtx get context
func (obj *_BaseMgr) GetCtx() context.Context {
	return obj.ctx
}

// Cancel cancel context
func (obj *_BaseMgr) Cancel(c context.Context) {
	obj.cancel()
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// GetIsRelated Query foreign key Association.获取是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

// SetIsRelated Query foreign key Association.设置是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) SetIsRelated(b bool) {
	obj.isRelated = b
}

// New new gorm.新gorm,重置条件
func (obj *_BaseMgr) New() {
	obj.DB = obj.NewDB()
}

// NewDB new gorm.新gorm
func (obj *_BaseMgr) NewDB() *gorm.DB {
	return obj.DB.Session(&gorm.Session{NewDB: true, Context: obj.ctx})
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


// 自定义sql查询
type Condition struct {
	list []*conditionInfo
}

func (c *Condition) AndWithCondition(condition bool,column string, cases string, value interface{}) (*Condition) {
	if condition {
		c.list = append(c.list, &conditionInfo{
			andor:  "and",
			column: column, // 列名
			case_:  cases,  // 条件(and,or,in,>=,<=)
			value:  value,
		})
	}
	return c
}


// And a Condition by and .and 一个条件
func (c *Condition) And(column string, cases string, value interface{}) (*Condition) {
	return c.AndWithCondition(true,column,cases,value)
}


func (c *Condition) OrWithCondition(condition bool,column string, cases string, value interface{}) (*Condition)  {
	if condition {
		c.list = append(c.list, &conditionInfo{
			andor:  "or",
			column: column, // 列名
			case_:  cases,  // 条件(and,or,in,>=,<=)
			value:  value,
		})
	}
	return c
}

// Or a Condition by or .or 一个条件
func (c *Condition) Or(column string, cases string, value interface{}) (*Condition) {
	return c.OrWithCondition(true,column,cases,value)
}

func (c *Condition) Get() (where string, out []interface{}) {
	firstAnd := -1
	for i := 0; i < len(c.list); i++ { // 查找第一个and
		if c.list[i].andor == "and" {
			where = fmt.Sprintf("{{GetVV }} %v ?", c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
			firstAnd = i
			break
		}
	}

	if firstAnd < 0 && len(c.list) > 0 { // 补刀
		where = fmt.Sprintf("{{GetVV }} %v ?", c.list[0].column, c.list[0].case_)
		out = append(out, c.list[0].value)
		firstAnd = 0
	}

	for i := 0; i < len(c.list); i++ { // 添加剩余的
		if firstAnd != i {
			where += fmt.Sprintf(" %v {{GetVV }} %v ?", c.list[i].andor, c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
		}
	}

	return
}

type conditionInfo struct {
	andor  string
	column string // 列名
	case_  string // 条件(in,>=,<=)
	value  interface{}
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
	ctx, cancel := context.WithCancel(context.Background())
	return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("{{GetTablePrefixName $obj.TableName}}"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_{{$obj.StructName}}Mgr) GetTableName() string {
	return "{{GetTablePrefixName $obj.TableName}}"
}

// Reset 重置gorm会话
func (obj *_{{$obj.StructName}}Mgr) Reset() *_{{$obj.StructName}}Mgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_{{$obj.StructName}}Mgr) Get() (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}

// Gets 获取批量结果
func (obj *_{{$obj.StructName}}Mgr) Gets() (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_{{$obj.StructName}}Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////
{{range $oem := $obj.Em}}
// With{{$oem.ColStructName}} {{$oem.ColName}}获取 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
	return optionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
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

	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where(options.query).Find(&result).Error
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

	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where(options.query).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}

{{if $obj.IsOutPage}}
// SelectPage 分页查询
func (obj *_{{$obj.StructName}}Mgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]{{$obj.StructName}},0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	resultPage.SetRecords(results)
	return
}
{{end}}

//////////////////////////enume case ////////////////////////////////////////////

{{range $oem := $obj.Em}}
// GetFrom{{$oem.ColStructName}} 通过{{$oem.ColName}}获取内容 {{$oem.Notes}} {{if $oem.IsMulti}}
func (obj *_{{$obj.StructName}}Mgr) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
{{else}}
func (obj *_{{$obj.StructName}}Mgr)  GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).First(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}
{{end}}
// GetBatchFrom{{$oem.ColStructName}} 批量查找 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} IN (?)", {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
 {{end}}
 //////////////////////////primary index case ////////////////////////////////////////////
 {{range $ofm := $obj.Primary}}
 // {{GenFListIndex $ofm 1}} primary or index 获取唯一内容
 func (obj *_{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&result).Error
	{{GenPreloadList $obj.PreloadList false}}
	return
}
 {{end}}

 {{range $ofm := $obj.Index}}
 // {{GenFListIndex $ofm 1}}  获取多个内容
 func (obj *_{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&results).Error
	{{GenPreloadList $obj.PreloadList true}}
	return
}
 {{end}}

`
	genPreload = `if err == nil && obj.isRelated { {{range $obj := .}}{{if $obj.IsMulti}}
		if err = obj.NewDB().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", result.{{$obj.ColStructName}}).Find(&result.{{$obj.ForeignkeyStructName}}List).Error;err != nil { // {{$obj.Notes}}
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}	
			} {{else}} 
		if err = obj.NewDB().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", result.{{$obj.ColStructName}}).Find(&result.{{$obj.ForeignkeyStructName}}).Error; err != nil { // {{$obj.Notes}} 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}{{end}} {{end}}}
`
	genPreloadMulti = `if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ { {{range $obj := .}}{{if $obj.IsMulti}}
		if err = obj.NewDB().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", results[i].{{$obj.ColStructName}}).Find(&results[i].{{$obj.ForeignkeyStructName}}List).Error;err != nil { // {{$obj.Notes}}
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} {{else}} 
		if err = obj.NewDB().Table("{{$obj.ForeignkeyTableName}}").Where("{{$obj.ForeignkeyCol}} = ?", results[i].{{$obj.ColStructName}}).Find(&results[i].{{$obj.ForeignkeyStructName}}).Error; err != nil { // {{$obj.Notes}} 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} {{end}} {{end}}
	}
}`
	genPage = `package {{.PackageName}}

import (
	"fmt"
	"strings"
)

type IPage interface {
	GetRecords() interface{}      // 获取查询的记录
	SetRecords(interface{})       // 设置查询的记录
	GetTotal() int64              // 获取总记录数
	SetTotal(int64)               // 设置总记录数
	GetCurrent() int64            // 获取当前页
	SetCurrent(int64)             // 设置当前页
	GetSize() int64               // 获取每页显示大小
	SetSize(int64)                // 设置每页显示大小
	AddOrderItem(OrderItem)       // 设置排序条件
	AddOrderItems([]OrderItem)    // 批量设置排序条件
	GetOrederItemsString() string // 将排序条件拼接成字符串
	Offset() int64                // 获取偏移量
	GetPages() int64              // 获取总的分页数
}

type Page struct {
	total   int64       // 总的记录数
	size    int64       // 每页显示的大小
	current int64       // 当前页
	orders  []OrderItem // 排序条件
	Records interface{} // 查询数据列表

}

func (page *Page) GetRecords() interface{} {
	return page.Records
}

func (page *Page) SetRecords(records interface{}) {
	page.Records = records
}

func (page *Page) GetTotal() int64 {
	return page.total
}

func (page *Page) SetTotal(total int64) {
	page.total = total

}

func (page *Page) GetCurrent() int64 {
	return page.current
}

func (page *Page) SetCurrent(current int64) {
	page.current = current
}

func (page *Page) GetSize() int64 {
	return page.size
}
func (page *Page) SetSize(size int64) {
	page.size = size

}

func (page *Page) AddOrderItem(orderItem OrderItem) {
	page.orders = append(page.orders, orderItem)
}

func (page *Page) AddOrderItems(orderItems []OrderItem) {
	page.orders = append(page.orders, orderItems...)
}

func (page *Page) GetOrederItemsString() string {
	arr := make([]string, 0)
	var order string

	for _, val := range page.orders {
		if val.asc {
			order = ""
		} else {
			order = "desc"
		}
		arr = append(arr, fmt.Sprintf("%s %s", val.column, order))
	}
	return strings.Join(arr, ",")
}

func (page *Page) Offset() int64 {
	if page.GetCurrent() > 0 {
		return (page.GetCurrent() - 1) * page.GetSize()
	} else {
		return 0
	}
}

func (page *Page) GetPages() int64 {
	if page.GetSize() == 0 {
		return 0
	}
	pages := page.GetTotal() / page.GetSize()
	if page.GetTotal()%page.size != 0 {
		pages++
	}

	return pages
}

type OrderItem struct {
	column string // 需要排序的字段
	asc    bool   // 是否正序排列，默认true
}

func BuildAsc(column string) OrderItem {
	return OrderItem{column: column, asc: true}
}

func BuildDesc(column string) OrderItem {
	return OrderItem{column: column, asc: false}
}

func BuildAscs(columns ...string) []OrderItem {
	items := make([]OrderItem, 0)
	for _, val := range columns {
		items = append(items, BuildAsc(val))
	}
	return items
}

func BuildDescs(columns ...string) []OrderItem {
	items := make([]OrderItem, 0)
	for _, val := range columns {
		items = append(items, BuildDesc(val))
	}
	return items
}

func NewPage(size, current int64, orderItems ...OrderItem) *Page {
	return &Page{size: size, current: current, orders: orderItems}
}
`
)
