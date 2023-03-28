package redis

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runcli/internal/utils"
)

type RedisConf struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

var configPath = "data/redisConf.json"

func New() RedisConf {
	return RedisConf{}
}

var redisConfigMap = make(map[string]RedisConf)

func InitConfig() {
	path := configPath
	if utils.Exists(path) {
		configBytes, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("reading config file error: %+v", err)
		}
		err = utils.Json.Unmarshal(configBytes, &redisConfigMap)
		if err != nil {
			log.Fatalf("load config error: %+v", err)
		}

	}
}

func GetConfigIfExist() map[string]RedisConf {
	return redisConfigMap
}

func SaveConfig(conf RedisConf) {
	if !utils.Exists(configPath) {
		log.Infof("config file not exists, creating default config file")
		_, err := utils.CreateNestedFile(configPath)
		if err != nil {
			log.Fatalf("failed to create config file: %+v", err)
		}
	}
	name := conf.Name
	redisConfigMap[name] = conf
	utils.WriteJsonToFile(configPath, redisConfigMap)
}
