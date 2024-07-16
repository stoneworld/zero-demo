package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/WTFAcademy/platform/internal/config"
	"github.com/WTFAcademy/platform/internal/handler"
	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	confPath string
	c        config.Config

	configFile = flag.String("f", "/etc/platform.yaml", "the config file")

	// Name is the name of the project
	rootCmd = &cobra.Command{
		Use:   "WTF",
		Short: "WTF is a platform for learning",
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the application",
		Run: func(cmd *cobra.Command, args []string) {
			runApp()
		},
	}

	// 执行一个临时的任务或者定时任务 demo 示例
	taskCmd = &cobra.Command{
		Use:   "task",
		Short: "Run the application",
		Run: func(cmd *cobra.Command, args []string) {
			confPath := cmd.Flag("config")
			fmt.Println("confPath:", confPath.Value.String())
			runTestCmd() // 执行命令
		},
	}
)

func runTestCmd() {
	fmt.Println("run test cmd")
}

func init() {
	// init config
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "etc/platform.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(taskCmd)
}

func initConfig() {
	fmt.Println("confPath:", confPath)
	conf.MustLoad(confPath, &c)
}

func runApp() {
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
