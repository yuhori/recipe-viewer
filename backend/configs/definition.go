package configs

import (
	"github.com/spf13/viper"
)

var (
	exists bool
)

func setDefaultConfig() {
	// default 設定を書く場合には、ここに書く
}

func loadConfig() error {
	setDefaultConfig()

	// 環境変数を優先
	viper.AutomaticEnv()

	exists = true
	return nil
}

func GetString(k string) string {
	if !exists {
		if err := loadConfig(); err != nil {
			panic(err)
		}
	}
	return viper.GetString(k)
}
