package cmd

import (
	"data/config"
	"fmt"
	"os"
	"public/tools"

	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

var mysqlInfo config.MysqlDbInfo

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "gorm mysql reflect tools",
	Long:  `base on gorm tools for mysql database to golang struct`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tools.GetJsonStr(config.GetMysqlDbInfo()))
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

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Host, "host", "H", "", "数据库地址.")
	rootCmd.MarkFlagRequired("host")
	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Username, "user", "u", "", "用户名.")
	rootCmd.MarkFlagRequired("user")

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Password, "password", "p", "", "密码.")
	rootCmd.MarkFlagRequired("password")

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Database, "database", "d", "", "数据库名")
	rootCmd.MarkFlagRequired("database")

	rootCmd.Flags().IntVar(&mysqlInfo.Port, "port", 3306, "端口号")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	validate := validator.New()
	err := validate.Struct(mysqlInfo)
	if err != nil {
		err1 := validate.Struct(config.GetMysqlDbInfo())
		if err1 != nil {
			fmt.Println("Can't read cmd: using （-h, --help) to get more imfo")
			fmt.Println("error info: ", err, err1)
			os.Exit(1)
		} else {
			fmt.Println("using default config info.")
		}
	} else {
		config.SetMysqlDbInfo(&mysqlInfo)
	}
}
