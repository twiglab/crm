package cmdb

import (
	"log"

	"github.com/spf13/viper"
)

func InitViperConfig() {

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	log.Println("Using config file:", viper.ConfigFileUsed())
}

func Cmdb(config any) {
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(err)
	}
}

func CmdbKey(key string, config any) {
	if err := viper.UnmarshalKey(key, config); err != nil {
		log.Fatal(err)
	}
}
