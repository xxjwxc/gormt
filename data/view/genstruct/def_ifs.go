package genstruct

import "github.com/xxjwxc/gormt/data/view/generate"

// IPackage package of interface
type IPackage interface {
	generate.IGenerate
	// Defining package names.定义包名
	SetPackage(string)
	// Add import by type.通过类型添加import
	AddImport(string)
	// Add a structure.添加一个结构体
	AddStruct(IStruct)
}

// IStruct struct of interface
type IStruct interface {
	generate.IGenerate

	// Set up SQL create statement, backup use.设置创建语句，备份使用
	SetCreatTableStr(string)

	// Setting Structural Name设置结构体名字
	SetStructName(string)

	// setting the notes.设置注释
	SetNotes(string)

	// add one element.添加一个元素
	AddElement(...IElement)
}

// IElement element in stuct of interface.元素类
type IElement interface {
	generate.IGenerate

	// setting name of element.设置元素名字
	SetName(string)
	// Setting element type.设置元素类型
	SetType(string)

	// setting notes of element .设置注释
	SetNotes(string)

	// add one tag.添加一个tag标记
	AddTag(k string, v string)
}
