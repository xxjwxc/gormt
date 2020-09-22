package cmd

import (
	"os"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/gormt/data/view/gtools"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

var mysqlInfo config.DBInfo
var outDir string
var singularTable bool
var foreignKey bool
var funcKey bool
var ui bool
var urlTag string

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "gorm mysql reflect tools",
	Long:  `base on gorm tools for mysql database to golang struct`,
	Run: func(cmd *cobra.Command, args []string) {
		gtools.Execute()
		// Start doing things.开始做事情
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Host, "host", "H", "", "数据库地址.(注意-H为大写)")
	rootCmd.MarkFlagRequired("host")
	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Username, "user", "u", "", "用户名.")
	rootCmd.MarkFlagRequired("user")

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Password, "password", "p", "", "密码.")
	rootCmd.MarkFlagRequired("password")

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Database, "database", "d", "", "数据库名")
	rootCmd.MarkFlagRequired("database")

	rootCmd.PersistentFlags().StringVarP(&outDir, "outdir", "o", "", "输出目录")
	rootCmd.MarkFlagRequired("outdir")

	rootCmd.PersistentFlags().BoolVarP(&singularTable, "singular", "s", true, "是否禁用表名复数")
	rootCmd.MarkFlagRequired("singular")

	rootCmd.PersistentFlags().BoolVarP(&foreignKey, "foreign", "f", false, "是否导出外键关联")
	rootCmd.MarkFlagRequired("foreign key")

	rootCmd.PersistentFlags().BoolVarP(&funcKey, "fun", "F", false, "是否导出函数")
	rootCmd.MarkFlagRequired("func export")

	rootCmd.PersistentFlags().BoolVarP(&ui, "gui", "g", false, "是否ui显示模式")
	rootCmd.MarkFlagRequired("show on gui")

	rootCmd.PersistentFlags().StringVarP(&urlTag, "url", "l", "", "url标签(json,url)")
	rootCmd.MarkFlagRequired("url tag")

	rootCmd.Flags().IntVar(&mysqlInfo.Port, "port", 3306, "端口号")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	MergeMysqlDbInfo()
	validate := validator.New()
	err := validate.Struct(config.GetDbInfo())
	if err != nil {
		mylog.Info("Can't read cmd: using （-h, --help) to get more info")
		mylog.Error(err)
		os.Exit(1)
	} else {
		mylog.Info("using database info:")
		mylog.JSON(config.GetDbInfo())
	}
}

// MergeMysqlDbInfo merge parm
func MergeMysqlDbInfo() {
	var tmp = config.GetDbInfo()
	if len(mysqlInfo.Database) > 0 {
		tmp.Database = mysqlInfo.Database
	}
	if len(mysqlInfo.Host) > 0 {
		tmp.Host = mysqlInfo.Host
	}
	if len(mysqlInfo.Password) > 0 {
		tmp.Password = mysqlInfo.Password
	}
	if mysqlInfo.Port != 3306 {
		tmp.Port = mysqlInfo.Port
	}
	if len(mysqlInfo.Username) > 0 {
		tmp.Username = mysqlInfo.Username
	}
	if len(urlTag) > 0 {
		config.SetURLTag(urlTag)
	}

	config.SetMysqlDbInfo(&tmp)

	if len(outDir) > 0 {
		config.SetOutDir(outDir)
	}

	if singularTable {
		config.SetSingularTable(singularTable)
	}

	if foreignKey {
		config.SetForeignKey(foreignKey)
	}

	if funcKey {
		config.SetIsOutFunc(funcKey)
	}

	if ui {
		config.SetIsGUI(ui)
	}

}
