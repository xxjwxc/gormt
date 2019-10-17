package generate

import "github.com/xxjwxc/public/tools"

// Add add one to print.打印
func (p *PrintAtom) Add(str ...interface{}) {
	var tmp string
	for _, v := range str {
		tmp += tools.AsString(v) + _interval
	}
	p.lines = append(p.lines, tmp)
}

// Generates Get the generated list.获取生成列表
func (p *PrintAtom) Generates() []string {
	return p.lines
}
