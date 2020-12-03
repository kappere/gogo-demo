package task

import (
	"wataru.com/gogo/frame/task"
	"wataru.com/gogo/logger"
)

var DemoTask_ = &DemoTask{}

type DemoTask struct {
	task.Task
}

func (t *DemoTask) Process() {
	logger.Info("DemoTask running!")
	logger.Info("DemoTask end!")
}
