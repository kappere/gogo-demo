package boot

import (
	assets "gogo-demo/assets"

	config "wataru.com/gogo/config"
)

func init() {
	config.SetAssets(assets.Assets)
	config.InitConfig()
	InitRouter()
	InitCron()
}
