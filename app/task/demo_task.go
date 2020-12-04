package task

import (
	"wataru.com/gogo/frame/task"
	"wataru.com/gogo/logger"
)

type DemoTask struct {
	task.Task
}

var DemoTaskBean = &DemoTask{}

func (t *DemoTask) Process() {
	logger.Info("DemoTask running!")
	logger.Info("DemoTask end!")
}
