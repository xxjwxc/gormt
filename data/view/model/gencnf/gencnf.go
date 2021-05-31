package gencnf

import (
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/model"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

// GetCnfModel get model interface. 获取model接口
func GetCnfModel() model.IModel {
	//now just support mysql
	return &CnfModel
}

// GenOutPut 输出
func GenOutPut(info *model.DBInfo) {
	path := path.Join(config.GetOutDir(), info.DbName+".yml")
	out, _ := yaml.Marshal(info)

	flag := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	f, err := os.OpenFile(path, flag, 0666)
	if err != nil {
		mylog.Error(err)
		return
	}
	defer f.Close()
	f.Write(out)
}

// CnfModel yaml model from IModel
var CnfModel cnfModel

type cnfModel struct {
}

// GenModel get model.DBInfo info.获取数据库相关属性
func (m *cnfModel) GenModel() model.DBInfo {
	var dbInfo model.DBInfo
	// getPackageInfo(orm, &dbInfo)
	// 添加逻辑
	dbInfo.PackageName = m.GetPkgName()
	dbInfo.DbName = m.GetDbName()
	return dbInfo
}

// GetTableNames get table name.获取指定的表名
func (m *cnfModel) GetTableNames() string {
	return config.GetTableNames()
}

// GetDbName get database name.获取数据库名字
func (m *cnfModel) GetDbName() string {
	dir := config.GetDbInfo().Host
	dir = strings.Replace(dir, "\\", "/", -1)
	if len(dir) > 0 {
		if dir[len(dir)-1] == '/' {
			dir = dir[:(len(dir) - 1)]
		}
	}
	var dbName string
	list := strings.Split(dir, "/")
	if len(list) > 0 {
		dbName = list[len(list)-1]
	}
	list = strings.Split(dbName, ".")
	if len(list) > 0 {
		dbName = list[0]
	}

	if len(dbName) == 0 || dbName == "." {
		panic(fmt.Sprintf("%v : db host config err.must file dir", dbName))
	}

	return dbName
}

// GetPkgName package names through config outdir configuration.通过config outdir 配置获取包名
func (m *cnfModel) GetPkgName() string {
	dir := config.GetOutDir()
	dir = strings.Replace(dir, "\\", "/", -1)
	if len(dir) > 0 {
		if dir[len(dir)-1] == '/' {
			dir = dir[:(len(dir) - 1)]
		}
	}
	var pkgName string
	list := strings.Split(dir, "/")
	if len(list) > 0 {
		pkgName = list[len(list)-1]
	}

	if len(pkgName) == 0 || pkgName == "." {
		curDir := tools.GetModelPath()
		curDir = strings.Replace(curDir, "\\", "/", -1)
		list = strings.Split(curDir, "/")
		if len(list) > 0 {
			pkgName = list[len(list)-1]
		}
	}

	return pkgName
}
