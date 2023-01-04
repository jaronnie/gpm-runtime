/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Long:  `get version`,
	RunE:  version,
}

func version(cmd *cobra.Command, args []string) error {
	if Version != "" && Commit != "" {
		fmt.Printf("%s-%s\n", Version, Commit[0:6])
	}
	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
