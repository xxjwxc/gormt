package genmssql

import (
	"fmt"
	"sort"
	"strings"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/model"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// MssqlModel mysql model from IModel
var MssqlModel mssqlModel

type mssqlModel struct {
}

// GenModel get model.DBInfo info.获取数据库相关属性
func (m *mssqlModel) GenModel() model.DBInfo {
	dsn := fmt.Sprintf("server=%v;database=%v;user id=%v;password=%v;port=%v;encrypt=disable",
		config.GetDbInfo().Host, config.GetDbInfo().Database, config.GetDbInfo().Username, config.GetDbInfo().Password, config.GetDbInfo().Port)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		mylog.Error(err)
		return model.DBInfo{}
	}
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

	var dbInfo model.DBInfo
	m.getPackageInfo(db, &dbInfo)
	dbInfo.PackageName = m.GetPkgName()
	dbInfo.DbName = m.GetDbName()
	return dbInfo
}

// GetDbName get database name.获取数据库名字
func (m *mssqlModel) GetDbName() string {
	return config.GetDbInfo().Database
}

// GetTableNames get table name.获取格式化后指定的表名
func (m *mssqlModel) GetTableNames() string {
	return config.GetTableNames()
}

// GetOriginTableNames get table name.获取原始指定的表名
func (m *mssqlModel) GetOriginTableNames() string {
	return config.GetOriginTableNames()
}

// GetPkgName package names through config outdir configuration.通过config outdir 配置获取包名
func (m *mssqlModel) GetPkgName() string {
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

func (m *mssqlModel) getPackageInfo(orm *gorm.DB, info *model.DBInfo) {
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
			// TODO:获取创建sql语句
			// Get create SQL statements.获取创建sql语句
			// rows, err := orm.Raw("show create table " + assemblyTable(tabName)).Rows()
			// //defer rows.Close()
			// if err == nil {
			// 	if rows.Next() {
			// 		var table, CreateTable string
			// 		rows.Scan(&table, &CreateTable)
			// 		tab.SQLBuildStr = CreateTable
			// 	}
			// }
			// rows.Close()
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
func (m *mssqlModel) getTableElement(orm *gorm.DB, tab string) (el []model.ColumnsInfo) {
	sql := fmt.Sprintf(`SELECT 
    id   = a.colorder,
    name     = a.name,
    pk       = case when exists(SELECT 1 FROM sysobjects where xtype='PK' and parent_obj=a.id and name in (
                     SELECT name FROM sysindexes WHERE indid in( SELECT indid FROM sysindexkeys WHERE id = a.id AND colid=a.colid))) then 1 else 0 end,
    tp       = b.name,
    len       = COLUMNPROPERTY(a.id,a.name,'PRECISION'),
    isnull     = a.isnullable,
    des   = isnull(g.[value],'')
FROM 
    syscolumns a
left join 
    systypes b 
on 
    a.xusertype=b.xusertype
inner join 
    sysobjects d 
on 
    a.id=d.id  and d.xtype='U' and  d.name<>'dtproperties'
left join 
sys.extended_properties   g 
on 
    a.id=G.major_id and a.colid=g.minor_id  
left join
sys.extended_properties f
on 
    d.id=f.major_id and f.minor_id=0
where 
    d.name='%v' 
order by 
    a.id,a.colorder`, tab)

	lenPk := 0
	// get keys
	var Keys []ColumnKeys
	orm.Raw(sql).Scan(&Keys)
	for i := 0; i < len(Keys); i++ {
		v := &Keys[i]
		if v.Pk == 1 {
			lenPk++
		}
		if strings.EqualFold(v.Type, "varchar") { // 字符串
			v.Type = fmt.Sprintf("varchar(%v)", v.Length)
		} else if strings.EqualFold(v.Type, "int") { // int
			v.Type = fmt.Sprintf("int(%v)", v.Length)
		}
	}
	// ----------end

	// TODO:ForeignKey

	for _, v := range Keys {
		var tmp model.ColumnsInfo
		tmp.Name = v.Name
		tmp.Type = v.Type
		FixNotes(&tmp, v.Desc) // 分析表注释

		if v.Pk > 0 { // 主键，或者联合组件
			if lenPk <= 1 { // 主键
				tmp.Index = append(tmp.Index, model.KList{
					Key:     model.ColumnsKeyPrimary,
					Multi:   false,
					KeyType: "primaryKey",
				})
			} else {
				tmp.Index = append(tmp.Index, model.KList{
					Key:     model.ColumnsKeyPrimary,
					Multi:   true,
					KeyType: "FULLTEXT",
				})
			}
		}

		tmp.IsNull = (v.Isnull == 1)

		el = append(el, tmp)
	}
	return
}

// getTables Get columns and comments.获取表列及注释
func (m *mssqlModel) getTables(orm *gorm.DB) map[string]string {
	tbDesc := make(map[string]string)

	// Get column names.获取列名
	if m.GetOriginTableNames() != "" {
		sarr := strings.Split(m.GetOriginTableNames(), ",")
		if len(sarr) != 0 {
			for _, val := range sarr {
				tbDesc[val] = ""
			}
		}
	} else {
		var list []TableDescription
		err := orm.Raw(`SELECT DISTINCT
		d.name,
		f.value
		FROM
		syscolumns a
		LEFT JOIN systypes b ON a.xusertype= b.xusertype
		INNER JOIN sysobjects d ON a.id= d.id
		AND d.xtype= 'U'
		AND d.name<> 'dtproperties'
		LEFT JOIN syscomments e ON a.cdefault= e.id
		LEFT JOIN sys.extended_properties g ON a.id= G.major_id
		AND a.colid= g.minor_id
		LEFT JOIN sys.extended_properties f ON d.id= f.major_id
		AND f.minor_id= 0 ;`).Scan(&list).Error
		if err != nil {
			if !config.GetIsGUI() {
				fmt.Println(err)
			}
			return tbDesc
		}

		for _, v := range list {
			tbDesc[v.Name] = v.Value
		}
	}

	return tbDesc
}

func assemblyTable(name string) string {
	return "`" + name + "`"
}
