package main

import (
	"errors"
	"github.com/jaronnie/gpm-runtime/internal"
	"github.com/rdsutbbp/logx"
	"os"
)

func main() {
	logx.Logger()

	if len(os.Args) < 1 {
		panic(errors.New("empty command"))
	}

	// get command
	command := os.Args[1]

	// get args
	var args []string
	if len(os.Args) >= 2 {
		args = os.Args[2:]
	}

	exit := make(chan bool, 0)

	if err := internal.RunCommand(command, args...); err != nil {
		logx.Errorf("Error running for command: %v\n", err)
		return
	}

	if err := internal.ProcessTask(); err != nil {
		logx.Errorf("Error running for process task: %v\n", err)
		return
	}

	<-exit

}
