package boot

import (
	apptask "gogo-demo/app/task"

	"wataru.com/gogo/frame/task"
)

func InitCron() {
	task.NewTask("0/10 * * * * *", apptask.DemoTask{})
}
