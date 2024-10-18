package cmd

import (
	"fmt"
	"gin-demo/global"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {

	if cfgFile != "" {
		// 使用用户提供的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找默认配置文件
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// 在用户主目录中查找配置文件
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
		viper.SetConfigType("yaml")
	}

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}

	// 解析配置文件
	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		fmt.Printf("无法解析配置文件: %v\n", err)
	}

	fmt.Println("配置文件解析成功")

	fmt.Println(global.CONFIG)
}

var RootCmd = &cobra.Command{
	Use:   "go-demo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
            examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 在这里处理命令行逻辑
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	},
}
