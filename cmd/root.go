package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// RootCmd -
var RootCmd = &cobra.Command{
	Use:     "dowmuzic-app",
	Short:   "Download muzics app",
	Long:    `Download muzics app`,
	Version: "0.0.0",
}

// Execute -
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "run app withconfig file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("download-muzic")
	}

	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
