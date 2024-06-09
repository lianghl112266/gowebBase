package conf

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
