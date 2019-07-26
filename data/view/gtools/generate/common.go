package generate

import (
	"fmt"
	"strings"

	"gormt/data/config"

	"github.com/xxjwxc/public/tools"
)

// 打印
func (p *PrintAtom) Add(str ...interface{}) {
	var tmp string
	for _, v := range str {
		tmp += tools.AsString(v) + _interval
	}
	p.lines = append(p.lines, tmp)
}

// 打印
func (p *PrintAtom) Generate() []string {
	return p.lines
}

//设置元素名字
func (e *GenElement) SetName(name string) {
	e.Name = name
}

//设置元素类型
func (e *GenElement) SetType(tp string) {
	e.Type = tp
}

//设置注释
func (e *GenElement) SetNotes(notes string) {
	e.Notes = notes
}

//添加一个tag标记
func (e *GenElement) AddTag(k string, v string) {
	if e.Tags == nil {
		e.Tags = make(map[string][]string)
	}
	e.Tags[k] = append(e.Tags[k], v)
}

//获取结果数据
func (e *GenElement) Generate() string {
	tag := ""
	var tags []string
	if e.Tags != nil {
		for k, v := range e.Tags {
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, k, strings.Join(v, ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p PrintAtom
	p.Add(e.Name, e.Type, tag, "//", e.Notes)
	return p.Generate()[0]
}

//////////////////////////////////////////////////////////////////////////////
// struct
//////////////////////////////////////////////////////////////////////////////

//设置创建语句，备份使用
func (s *GenStruct) SetCreatTableStr(sql string) {
	s.SqlBuildStr = sql
}

//获取结果数据
func (s *GenStruct) SetStructName(name string) {
	s.Name = name
}

//设置注释
func (e *GenStruct) SetNotes(notes string) {
	e.Notes = notes
}

//添加一个/或多个元素
func (s *GenStruct) AddElement(e ...GenElement) {
	s.Em = append(s.Em, e...)
}

//获取结果数据
func (s *GenStruct) Generate() []string {
	var p PrintAtom
	if !config.GetSimple() {
		p.Add("/******sql******")
		p.Add(s.SqlBuildStr)
		p.Add("******sql******/")
	}
	p.Add("//", s.Notes)
	p.Add("type", s.Name, "struct {")
	for _, v := range s.Em {
		p.Add(v.Generate())
	}
	p.Add("}")

	return p.Generate()
}

//////////////////////////////////////////////////////////////////////////////
// package
//////////////////////////////////////////////////////////////////////////////

//定义包名
func (p *GenPackage) SetPackage(pname string) {
	p.Name = pname
}

//通过类型添加import
func (p *GenPackage) AddImport(imp string) {
	if p.Imports == nil {
		p.Imports = make(map[string]string)
	}
	p.Imports[imp] = imp
}

//添加一个结构体
func (p *GenPackage) AddStruct(st GenStruct) {
	p.Structs = append(p.Structs, st)
}

//获取结果数据
func (p *GenPackage) Generate() string {
	p.genimport() //补充 import

	var pa PrintAtom
	pa.Add("package", p.Name)
	//add import
	if p.Imports != nil {
		pa.Add("import (")
		for _, v := range p.Imports {
			pa.Add(v)
		}
		pa.Add(")")
	}
	//-----------end
	//add struct
	for _, v := range p.Structs {
		for _, v1 := range v.Generate() {
			pa.Add(v1)
		}
	}
	//-----------end

	//输出
	strOut := ""
	for _, v := range pa.Generate() {
		strOut += v + "\n"
	}

	return strOut
}

//获取结果数据
func (p *GenPackage) genimport() {
	for _, v := range p.Structs {
		for _, v1 := range v.Em {
			if v2, ok := EImportsHead[v1.Type]; ok {
				if len(v2) > 0 {
					p.AddImport(v2)
				}
			}
		}
	}
}
