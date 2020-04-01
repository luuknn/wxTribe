package initial

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"wxTribe/utils"
)

func InitConfig() (err error) {
	viper.AddConfigPath(fmt.Sprintf("%s/../conf", utils.GetCurrentPath()))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err = InitConfig()
	})
	return
}
