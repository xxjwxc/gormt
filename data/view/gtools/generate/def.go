package generate

//包类
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

//结构体类
type IStruct interface {
	//设置结构体名字
	SetStructName(string)

	//设置注释
	SetNotes(string)

	//添加一个元素
	AddElement(IElement)

	//获取结果数据
	Generate() []string
}

//元素类
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

//元素类
type GenElement struct {
	Name  string              //元素名
	Type  string              //类型标记
	Notes string              //注释
	Tags  map[string][]string //标记
}

//结构体
type GenStruct struct {
	Name  string       //名字
	Notes string       //注释
	Em    []GenElement //元素组合
}

//包体
type GenPackage struct {
	Name    string            //名字
	Imports map[string]string //元素组合
	Structs []GenStruct       //结构体组合
}

//间隔
var _interval = "\t"

var EImportsHead = map[string]string{
	"stirng": "string",
}

var isGoKeyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

var isGoPredeclaredIdentifier = map[string]bool{
	"append":     true,
	"bool":       true,
	"byte":       true,
	"cap":        true,
	"close":      true,
	"complex":    true,
	"complex128": true,
	"complex64":  true,
	"copy":       true,
	"delete":     true,
	"error":      true,
	"false":      true,
	"float32":    true,
	"float64":    true,
	"imag":       true,
	"int":        true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"int8":       true,
	"iota":       true,
	"len":        true,
	"make":       true,
	"new":        true,
	"nil":        true,
	"panic":      true,
	"print":      true,
	"println":    true,
	"real":       true,
	"recover":    true,
	"rune":       true,
	"string":     true,
	"true":       true,
	"uint":       true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uint8":      true,
	"uintptr":    true,
}

type PrintAtom struct {
	lines []string
}
