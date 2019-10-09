package generate

// interval.间隔
var _interval = "\t"

// IGenerate Generate Printing Interface.生成打印接口
type IGenerate interface {
	// Get the generate data .获取结果数据
	Generate() string
}

// PrintAtom . atom print .原始打印
type PrintAtom struct {
	lines []string
}
