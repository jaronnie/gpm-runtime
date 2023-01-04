/*
Copyright © 2023 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"os"

	"github.com/rdsutbbp/logx"
	"github.com/spf13/cobra"
)

var (
	Loglevel string

	RecycleZombiePeriod int // second
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpm-runtime",
	Short: "used to manager process in container in order to recycle zombie process.",
	Long:  `used to manager process in container in order to recycle zombie process.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	logx.SetLogLevel(logx.Level(Loglevel))
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&Loglevel, "loglevel", "", "error", "set log level")
	rootCmd.PersistentFlags().IntVarP(&RecycleZombiePeriod, "recycle-zombie-period", "p", 60, "set recycle zombie period")
}
