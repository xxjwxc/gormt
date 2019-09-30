package generate

// IGenerate Generate Printing Interface.生成打印接口
type IGenerate interface {
	// Get the generate data .获取结果数据
	Generate() string
}
