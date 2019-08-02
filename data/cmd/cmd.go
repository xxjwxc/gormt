package cmd

import (
	"fmt"
	"os"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/gormt/data/view/gtools"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

var mysqlInfo config.MysqlDbInfo
var outDir string
var singularTable bool

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "gorm mysql reflect tools",
	Long:  `base on gorm tools for mysql database to golang struct`,
	Run: func(cmd *cobra.Command, args []string) {
		gtools.Execute()
		//开始做事情
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

	rootCmd.PersistentFlags().BoolVarP(&singularTable, "singular", "s", false, "是否禁用表名复数")
	rootCmd.MarkFlagRequired("singular")

	rootCmd.Flags().IntVar(&mysqlInfo.Port, "port", 3306, "端口号")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	MergeMysqlDbInfo()
	validate := validator.New()
	err := validate.Struct(config.GetMysqlDbInfo())
	if err != nil {
		fmt.Println("Can't read cmd: using （-h, --help) to get more imfo")
		fmt.Println("error info: ", err, err)
		os.Exit(1)
	} else {
		fmt.Println("using config info:")
		fmt.Println(tools.GetJsonStr(config.GetMysqlDbInfo()))
	}
}

//MergeMysqlDbInfo 合并
func MergeMysqlDbInfo() {
	var tmp = config.GetMysqlDbInfo()
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

	config.SetMysqlDbInfo(&tmp)

	if len(outDir) > 0 {
		config.SetOutDir(outDir)
	}

	if singularTable {
		config.SetSingularTable(singularTable)
	}
}
