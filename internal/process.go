package internal

import (
	"github.com/jaronnie/gpm-runtime/utilx/server"
	"github.com/rdsutbbp/logx"
	"github.com/shirou/gopsutil/process"
	"syscall"
	"time"
)

const ProcessTaskValue = "process-task"

func ProcessTask() error {
	svr, err := server.NewSvr(ProcessTaskValue, func(msg interface{}, num int) (resp interface{}, err error) {
		return nil, nil
	}, []server.TimedTask{
		{
			Task: recycleZombieProcess,
			Time: time.Minute * 1,
		},
	})

	if err != nil {
		return err
	}

	svr.Go()

	return nil
}

func recycleZombieProcess(worker int) (err error) {
	processes, err := process.Processes()
	if err != nil {
		logx.Errorf("get processes meet error. Err: [%v]", err)
		return
	}
	for _, v := range processes {
		status, err := v.Status()
		if err != nil {
			logx.Errorf("get process status with PID [%d] meet error. Err: [%v]", v.Pid, err)
			continue
		}
		if status == "zombie" || status == "defunct" || status == "Z" {
			logx.Warnf("Process with PID [%d] is a zombie process", v.Pid)
			// wait process zombie to recycle
			var waitStatus syscall.WaitStatus
			_, err = syscall.Wait4(int(v.Pid), &waitStatus, 0, nil)
			if err != nil {
				logx.Errorf("syscall wait pid [%d]. Err: [%v]", v.Pid, err)
				continue
			}
		}
	}
	return nil
}
