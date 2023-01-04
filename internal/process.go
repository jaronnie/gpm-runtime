package internal

import (
	"syscall"
	"time"

	"github.com/rdsutbbp/logx"
	"github.com/shirou/gopsutil/process"

	"github.com/jaronnie/gpm-runtime/v2/utilx/server"
)

const ProcessTaskValue = "process-task"

type Options struct {
	RecycleZombiePeriod int
}

func ProcessTask(options *Options) error {
	logx.Debugf("Start to do process task server")

	svr, err := server.NewSvr(ProcessTaskValue, func(msg interface{}, num int) (resp interface{}, err error) {
		return nil, nil
	}, []server.TimedTask{
		{
			Task: recycleZombieProcess,
			Time: time.Duration(options.RecycleZombiePeriod) * time.Second,
		},
	})

	if err != nil {
		return err
	}

	svr.Go()

	return nil
}

func recycleZombieProcess(worker int) (err error) {
	logx.Debugf("Start to recycle zombie process")

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
			logx.Infof("Zombie process with PID [%d] has been recycled.", v.Pid)
		}
	}
	return nil
}
