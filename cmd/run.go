/*
Copyright © 2023 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"github.com/pkg/errors"
	"github.com/rdsutbbp/logx"
	"github.com/spf13/cobra"

	"github.com/jaronnie/gpm-runtime/v2/internal"
	"github.com/jaronnie/gpm-runtime/v2/utilx/server"
)

var (
	Retry      bool
	RetryDelay int
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run -- <your program with args>",
	Short: "run your program",
	Long:  `run your program.`,
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		panic(errors.New("empty command"))
	}

	// get command
	command := args[0]

	if len(args) >= 1 {
		args = args[1:]
	}

	if err := internal.ProcessTask(&internal.Options{
		RecycleZombiePeriod: RecycleZombiePeriod,
	}); err != nil {
		logx.Errorf("Error running for process task: %v\n", err)
		return err
	}

	if err := internal.RunCommand(command, args...); err != nil {
		logx.Errorf("Error running for command: %v\n", err)
		return nil
	}

	logx.Debugf("Stop process task server")

	server.StopServer(internal.ProcessTaskValue)

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&Retry, "retry", "r", false, "retry run")
	runCmd.Flags().IntVarP(&RetryDelay, "retry-delay", "", 5, "retry delay second")
}
