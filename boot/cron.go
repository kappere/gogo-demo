package boot

import (
	apptask "gogo-demo/app/task"

	"wataru.com/gogo/frame/context"
	"wataru.com/gogo/frame/task"
)

func InitCron() {
	context.RegistInitializer(func() {
		task.NewTask("0/10 * * * * *", apptask.DemoTaskBean)
	})
}
