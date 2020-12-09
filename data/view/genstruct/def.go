package genstruct

// GenElement element of sturct.元素类
type GenElement struct {
	Name       string              // Name.元素名
	ColumnName string              // table name.表名
	Type       string              // Type.类型标记
	Notes      string              // Notes.注释
	Tags       map[string][]string // tages.标记
}

// GenStruct struct of IStruct .结构体
type GenStruct struct {
	SQLBuildStr string       // Create SQL statements.创建sql语句
	TableName   string       // table_name.表名
	Name        string       // name.名字
	Notes       string       // notes.注释
	Em          []GenElement // em.元素组合
}

// GenPackage package of IPackage.包体
type GenPackage struct {
	Name        string            // name.名字
	Imports     map[string]string // Inclusion term.元素组合
	Structs     []GenStruct       // struct list .结构体组合
	FuncStrList []string          // func of template on string. 函数的最终定义
}
