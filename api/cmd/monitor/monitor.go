// Copyright © 2023 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitor

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"mm-ndj/config"
	"mm-ndj/server/dao"
	"mm-ndj/server/task"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "starland",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		serverCtx := dao.GetServiceCtx()
		if serverCtx == nil {
			panic("err config")
		}
		//r := router.NewRouter(serverCtx)
		//appRouter, err := app.NewPlatform(serverCtx.C, r, serverCtx)
		//if err != nil {
		//	config.Logger.Error("init error", zap.Error(err))
		//	return
		//}
		//if err := utils.SonyFlakeInit(config.Conf.Common.StartTime, 1); err != nil {
		//	config.Logger.Error("SonyFlakeInit init failed", zap.Error(err))
		//	return
		//}
		config.InitPrizeConfig()
		Pprof()
		//appRouter.AppStart()
		chSig := make(chan os.Signal)
		signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
		<-chSig

		//appRouter.AppClose()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.starland.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Crontab
	c := newCronWithSeconds()
	fmt.Println("crontab")
	//_, _ = c.AddFunc("*/6, *, *, *, *, ?", task.CrontabForTest)
	_, _ = c.AddFunc("*/6, *, *, *, *, ?", task.MonitorGacha)

	c.Start()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".starland" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".starland")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// 监控分析
func Pprof() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
}

func newCronWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
