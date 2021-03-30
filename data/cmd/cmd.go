package cmd

import (
	"os"
	"strings"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/gormt/data/view/gtools"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/spf13/cobra"
	"github.com/xxjwxc/public/mycobra"
	"gopkg.in/go-playground/validator.v9"
)

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

	rootCmd.PersistentFlags().StringP("host", "H", "", "数据库地址.(注意-H为大写)")
	rootCmd.MarkFlagRequired("host")
	rootCmd.PersistentFlags().StringP("user", "u", "", "用户名.")
	rootCmd.MarkFlagRequired("user")

	rootCmd.PersistentFlags().StringP("password", "p", "", "密码.")
	rootCmd.MarkFlagRequired("password")

	rootCmd.PersistentFlags().StringP("database", "d", "", "数据库名")
	rootCmd.MarkFlagRequired("database")

	rootCmd.PersistentFlags().StringP("outdir", "o", "", "输出目录")
	rootCmd.MarkFlagRequired("outdir")

	rootCmd.PersistentFlags().BoolP("singular", "s", true, "是否禁用表名复数")
	rootCmd.MarkFlagRequired("singular")

	rootCmd.PersistentFlags().BoolP("foreign", "f", false, "是否导出外键关联")
	rootCmd.MarkFlagRequired("foreign key")

	rootCmd.PersistentFlags().BoolP("fun", "F", false, "是否导出函数")
	rootCmd.MarkFlagRequired("func export")

	rootCmd.PersistentFlags().BoolP("gui", "g", false, "是否ui显示模式")
	rootCmd.MarkFlagRequired("show on gui")

	rootCmd.PersistentFlags().StringP("url", "l", "", "url标签(json,url)")
	rootCmd.MarkFlagRequired("url tag")

	rootCmd.Flags().Int("port", 3306, "端口号")

	rootCmd.Flags().StringP("table_prefix", "t", "", "表前缀")
	//add table name. 增加表名称
	rootCmd.Flags().StringP("table_names", "b", "", "表名称")

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
	mycobra.IfReplace(rootCmd, "database", &tmp.Database) // 如果设置了，更新
	mycobra.IfReplace(rootCmd, "host", &tmp.Host)         // 如果设置了，更新
	mycobra.IfReplace(rootCmd, "password", &tmp.Password) // 如果设置了，更新
	mycobra.IfReplace(rootCmd, "port", &tmp.Port)         // 如果设置了，更新
	mycobra.IfReplace(rootCmd, "user", &tmp.Username)     // 如果设置了，更新
	config.SetMysqlDbInfo(&tmp)

	url := config.GetURLTag()
	mycobra.IfReplace(rootCmd, "url", &url) // 如果设置了，更新
	config.SetURLTag(url)

	dir := config.GetOutDir()
	mycobra.IfReplace(rootCmd, "outdir", &dir) // 如果设置了，更新
	config.SetOutDir(dir)

	fk := config.GetIsForeignKey()
	mycobra.IfReplace(rootCmd, "foreign", &fk) // 如果设置了，更新
	config.SetForeignKey(fk)

	funcKey := config.GetIsOutFunc()
	mycobra.IfReplace(rootCmd, "fun", &funcKey) // 如果设置了，更新
	config.SetIsOutFunc(funcKey)

	ig := config.GetIsGUI()
	mycobra.IfReplace(rootCmd, "gui", &ig) // 如果设置了，更新
	config.SetIsGUI(ig)

	tablePrefix := config.GetTablePrefix()
	mycobra.IfReplace(rootCmd, "table_prefix", &tablePrefix) // 如果设置了，更新
	config.SetTablePrefix(tablePrefix)

	//update tableNames. 更新tableNames
	tableNames := config.GetTableNames()
	if tableNames != "" {
		tableNames = strings.Replace(tableNames, "'", "", -1)
	}
	mycobra.IfReplace(rootCmd, "table_names", &tableNames) // 如果设置了，更新
	config.SetTableNames(tableNames)

}
