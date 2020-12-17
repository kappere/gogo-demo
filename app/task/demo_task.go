package task

import (
	"wataru.com/gogo/frame/component"
	"wataru.com/gogo/frame/context"
	"wataru.com/gogo/frame/task"
	"wataru.com/gogo/logger"
)

type DemoTask struct {
	*component.Component `autowired:""`
}

var DemoTaskBean = &DemoTask{}

func init() {
	context.RegistBean(&DemoTask{})
}

func (t *DemoTask) Initialize() {
	task.Tasklet("0 0/1 * * * *", t.Process)
}

func (t *DemoTask) Process() {
	logger.Info("DemoTask running!")
	logger.Info("DemoTask end!")
}
