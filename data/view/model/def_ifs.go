package model

// IModel Implement the interface to acquire database information and initialize it.实现接口获取数据库信息获取并初始化
type IModel interface {
	GenModel() DBInfo
	GetDbName() string
	GetPkgName() string    // Getting package names through config outdir configuration.通过config outdir 配置获取包名
	GetTableNames() string // Getting tableNames by config. 获取设置的表名
}
