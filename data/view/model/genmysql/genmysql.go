package genmysql

import (
	"fmt"
	"sort"
	"strings"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/model"
	"github.com/xxjwxc/public/mysqldb"
	"github.com/xxjwxc/public/tools"
)

// MySQLModel mysql model from IModel
var MySQLModel mysqlModel

type mysqlModel struct {
}

// GenModel get model.DBInfo info.获取数据库相关属性
func (m *mysqlModel) GenModel() model.DBInfo {
	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	var dbInfo model.DBInfo
	getPackageInfo(orm, &dbInfo)
	dbInfo.PackageName = m.GetPkgName()
	dbInfo.DbName = m.GetDbName()
	return dbInfo
}

// GetDbName get database name.获取数据库名字
func (m *mysqlModel) GetDbName() string {
	return config.GetMysqlDbInfo().Database
}

// GetPkgName package names through config outdir configuration.通过config outdir 配置获取包名
func (m *mysqlModel) GetPkgName() string {
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
		list = strings.Split(tools.GetModelPath(), "/")
		if len(list) > 0 {
			pkgName = list[len(list)-1]
		}
	}

	return pkgName
}

func getPackageInfo(orm *mysqldb.MySqlDB, info *model.DBInfo) {
	tables := getTables(orm) // get table and notes

	for tabName, notes := range tables {
		var tab model.TabInfo
		tab.Name = tabName
		tab.Notes = notes

		if config.GetIsOutSQL() {
			// Get create SQL statements.获取创建sql语句
			rows, err := orm.Raw("show create table " + tabName).Rows()
			//defer rows.Close()
			if err == nil {
				if rows.Next() {
					var table, CreateTable string
					rows.Scan(&table, &CreateTable)
					tab.SQLBuildStr = CreateTable
				}
			}
			rows.Close()
			// ----------end
		}

		// build element.构造元素
		tab.Em = getTableElement(orm, tabName)
		// --------end

		info.TabList = append(info.TabList, tab)
	}
	// sort tables
	sort.Slice(info.TabList, func(i, j int) bool {
		return info.TabList[i].Name < info.TabList[j].Name
	})
}

// getTableElement Get table columns and comments.获取表列及注释
func getTableElement(orm *mysqldb.MySqlDB, tab string) (el []model.ColumnsInfo) {
	keyNums := make(map[string]int)
	// get keys
	var Keys []keys
	orm.Raw("show keys from " + tab).Scan(&Keys)
	for _, v := range Keys {
		keyNums[v.KeyName]++
	}
	// ----------end

	var list []genColumns
	// Get table annotations.获取表注释
	orm.Raw("show FULL COLUMNS from " + tab).Scan(&list)
	// filter gorm.Model.过滤 gorm.Model
	if filterModel(&list) {
		el = append(el, model.ColumnsInfo{
			Type: "gorm.Model",
		})
	}
	// -----------------end

	// ForeignKey
	var foreignKeyList []genForeignKey
	if config.GetIsForeignKey() {
		orm.Raw(fmt.Sprintf(`select table_schema,table_name,column_name,referenced_table_schema,referenced_table_name,referenced_column_name from INFORMATION_SCHEMA.KEY_COLUMN_USAGE
		where table_schema = '%v' AND REFERENCED_TABLE_NAME IS NOT NULL AND TABLE_NAME = '%v'`, config.GetMysqlDbInfo().Database, tab)).Scan(&foreignKeyList)
	}
	// ------------------end

	for _, v := range list {
		var tmp model.ColumnsInfo
		tmp.Name = v.Field
		tmp.Notes = v.Desc
		tmp.Type = v.Type

		// keys
		if strings.EqualFold(v.Key, "PRI") { // Set primary key.设置主键
			tmp.Index = append(tmp.Index, model.KList{
				Key: model.ColumnsKeyPrimary,
			})
		} else if strings.EqualFold(v.Key, "UNI") { // unique
			tmp.Index = append(tmp.Index, model.KList{
				Key: model.ColumnsKeyUnique,
			})
		} else {
			for _, v1 := range Keys {
				if strings.EqualFold(v1.ColumnName, v.Field) {
					var k model.KList
					if v1.NonUnique == 1 { // index
						k.Key = model.ColumnsKeyIndex
					} else {
						k.Key = model.ColumnsKeyUniqueIndex
					}
					if keyNums[v1.KeyName] > 1 { // Composite index.复合索引
						k.KeyName = v1.KeyName
					}
					tmp.Index = append(tmp.Index, k)
				}
			}
		}

		tmp.IsNull = strings.EqualFold(v.Null, "YES")

		// ForeignKey
		fixForeignKey(foreignKeyList, tmp.Name, &tmp.ForeignKeyList)
		// -----------------end
		el = append(el, tmp)
	}
	return
}

// getTables Get columns and comments.获取表列及注释
func getTables(orm *mysqldb.MySqlDB) map[string]string {
	tbDesc := make(map[string]string)

	// Get column names.获取列名
	var tables []string

	rows, err := orm.Raw("show tables").Rows()
	if err != nil {
		if !config.GetIsGUI() {
			fmt.Println(err)
		}
		return tbDesc
	}

	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
		tbDesc[table] = ""
	}
	rows.Close()

	// Get table annotations.获取表注释
	rows1, err := orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + config.GetMysqlDbInfo().Database + "'").Rows()
	if err != nil {
		if !config.GetIsGUI() {
			fmt.Println(err)
		}
		return tbDesc
	}

	for rows1.Next() {
		var table, desc string
		rows1.Scan(&table, &desc)
		tbDesc[table] = desc
	}
	rows1.Close()

	return tbDesc
}
