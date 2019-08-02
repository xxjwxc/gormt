package generate

//IPackage 包类
type IPackage interface {
	//定义包名
	SetPackage(string)
	//通过类型添加import
	AddImport(string)
	//添加一个结构体
	AddStruct(IStruct)
	//获取结果数据
	Generate() string
}

//IStruct 结构体类
type IStruct interface {
	//设置创建语句，备份使用
	SetCreatTableStr(string)

	//设置结构体名字
	SetStructName(string)

	//设置注释
	SetNotes(string)

	//添加一个元素
	AddElement(...IElement)

	//获取结果数据
	Generate() []string
}

//IElement 元素类
type IElement interface {
	//设置元素名字
	SetName(string)
	//设置元素类型
	SetType(string)

	//设置注释
	SetNotes(string)

	//添加一个tag标记
	AddTag(k string, v string)

	//获取结果数据
	Generate() string
}

//////////////////////////////////////////////////////////////////////////////
//
//////////////////////////////////////////////////////////////////////////////

//GenElement 元素类
type GenElement struct {
	Name  string              //元素名
	Type  string              //类型标记
	Notes string              //注释
	Tags  map[string][]string //标记
}

//GenStruct 结构体
type GenStruct struct {
	SQLBuildStr string       //创建sql语句
	Name        string       //名字
	Notes       string       //注释
	Em          []GenElement //元素组合
}

//GenPackage 包体
type GenPackage struct {
	Name    string            //名字
	Imports map[string]string //元素组合
	Structs []GenStruct       //结构体组合
}

//间隔
var _interval = "\t"

//EImportsHead .
var EImportsHead = map[string]string{
	"stirng":     `"string"`,
	"time.Time":  `"time"`,
	"gorm.Model": `"github.com/jinzhu/gorm"`,
}

//PrintAtom .
type PrintAtom struct {
	lines []string
}
