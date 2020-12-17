package boot

import (
	assets "gogo-demo/assets"

	_ "gogo-demo/app/controller"
	_ "gogo-demo/app/task"

	config "wataru.com/gogo/config"
)

func init() {
	config.SetAssets(assets.Assets)
	config.InitConfig()
}
