package genmysql

import (
	"database/sql"
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
	m.getPackageInfo(orm, &dbInfo)
	dbInfo.PackageName = m.GetPkgName()
	dbInfo.DbName = m.GetDbName()
	return dbInfo
}

// GetDbName get database name.获取数据库名字
func (m *mysqlModel) GetDbName() string {
	return config.GetDbInfo().Database
}

// GetTableNames get table name.获取格式化后指定的表名
func (m *mysqlModel) GetTableNames() string {
	return config.GetTableNames()
}

// GetOriginTableNames get table name.获取原始指定的表名
func (m *mysqlModel) GetOriginTableNames() string {
	return config.GetOriginTableNames()
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

func (m *mysqlModel) getPackageInfo(orm *mysqldb.MySqlDB, info *model.DBInfo) {
	tabls := m.getTables(orm) // get table and notes
	// if m := config.GetTableList(); len(m) > 0 {
	// 	// 制定了表之后
	// 	newTabls := make(map[string]string)
	// 	for t := range m {
	// 		if notes, ok := tabls[t]; ok {
	// 			newTabls[t] = notes
	// 		} else {
	// 			fmt.Printf("table: %s not found in db\n", t)
	// 		}
	// 	}
	// 	tabls = newTabls
	// }
	for tabName, notes := range tabls {
		var tab model.TabInfo
		tab.Name = tabName
		tab.Notes = notes

		if config.GetIsOutSQL() {
			// Get create SQL statements.获取创建sql语句
			rows, err := orm.Raw("show create table " + assemblyTable(tabName)).Rows()
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
		tab.Em = m.getTableElement(orm, tabName)
		// --------end

		info.TabList = append(info.TabList, tab)
	}
	// sort tables
	sort.Slice(info.TabList, func(i, j int) bool {
		return info.TabList[i].Name < info.TabList[j].Name
	})
}

// getTableElement Get table columns and comments.获取表列及注释
func (m *mysqlModel) getTableElement(orm *mysqldb.MySqlDB, tab string) (el []model.ColumnsInfo) {
	keyNameCount := make(map[string]int)
	KeyColumnMp := make(map[string][]keys)
	// get keys
	var Keys []keys
	orm.Raw("show keys from " + assemblyTable(tab)).Scan(&Keys)
	for _, v := range Keys {
		keyNameCount[v.KeyName]++
		KeyColumnMp[v.ColumnName] = append(KeyColumnMp[v.ColumnName], v)
	}
	// ----------end

	var list []genColumns
	// Get table annotations.获取表注释
	orm.Raw("show FULL COLUMNS from " + assemblyTable(tab)).Scan(&list)
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
		sql := fmt.Sprintf(`select table_schema as table_schema,table_name as table_name,column_name as column_name,referenced_table_schema as referenced_table_schema,referenced_table_name as referenced_table_name,referenced_column_name as referenced_column_name
		from INFORMATION_SCHEMA.KEY_COLUMN_USAGE where table_schema = '%v' AND REFERENCED_TABLE_NAME IS NOT NULL AND TABLE_NAME = '%v'`, m.GetDbName(), tab)
		orm.Raw(sql).Scan(&foreignKeyList)
	}
	// ------------------end

	for _, v := range list {
		var tmp model.ColumnsInfo
		tmp.Name = v.Field
		tmp.Type = v.Type
		FixNotes(&tmp, v.Desc) // 分析表注释

		if v.Default != nil {
			if *v.Default == "" {
				tmp.Gormt = "default:''"
			} else {
				tmp.Gormt = fmt.Sprintf("default:%s", *v.Default)
			}
		}

		// keys
		if keylist, ok := KeyColumnMp[v.Field]; ok { // maybe have index or key
			for _, v := range keylist {
				if v.NonUnique == 0 { // primary or unique
					if strings.EqualFold(v.KeyName, "PRIMARY") { // PRI Set primary key.设置主键
						tmp.Index = append(tmp.Index, model.KList{
							Key:     model.ColumnsKeyPrimary,
							Multi:   (keyNameCount[v.KeyName] > 1),
							KeyType: v.IndexType,
						})
					} else { // unique
						if keyNameCount[v.KeyName] > 1 {
							tmp.Index = append(tmp.Index, model.KList{
								Key:     model.ColumnsKeyUniqueIndex,
								Multi:   (keyNameCount[v.KeyName] > 1),
								KeyName: v.KeyName,
								KeyType: v.IndexType,
							})
						} else { // unique index key.唯一复合索引
							tmp.Index = append(tmp.Index, model.KList{
								Key:     model.ColumnsKeyUnique,
								Multi:   (keyNameCount[v.KeyName] > 1),
								KeyName: v.KeyName,
								KeyType: v.IndexType,
							})
						}
					}
				} else { // mut
					tmp.Index = append(tmp.Index, model.KList{
						Key:     model.ColumnsKeyIndex,
						Multi:   true,
						KeyName: v.KeyName,
						KeyType: v.IndexType,
					})
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
func (m *mysqlModel) getTables(orm *mysqldb.MySqlDB) map[string]string {
	tbDesc := make(map[string]string)

	// Get column names.获取列名
	var tables []string

	if m.GetOriginTableNames() != "" {
		sarr := strings.Split(m.GetOriginTableNames(), ",")
		if len(sarr) != 0 {
			for _, val := range sarr {
				tbDesc[val] = ""
			}
		}
	} else {
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
	}

	// Get table annotations.获取表注释
	var err error
	var rows1 *sql.Rows
	if m.GetTableNames() != "" {
		rows1, err = orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbName() + "'and TABLE_NAME IN(" + m.GetTableNames() + ")").Rows()
		fmt.Println("getTables:" + m.GetTableNames())
		fmt.Println("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbName() + "'and TABLE_NAME IN(" + m.GetTableNames() + ")")
	} else {
		rows1, err = orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbName() + "'").Rows()
	}

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

func assemblyTable(name string) string {
	return "`" + name + "`"
}
