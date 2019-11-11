package config

import (
	"log"
	"path"
	"runtime"
	"time"

	"gopkg.in/ini.v1"
)

type MySqlConf struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
}

type LogConfig struct {
	LogFilePath string `ini:"log_file_path"`
	LogFileName string `ini:"log_file_name"`
}

type HashIdsConfig struct {
	Salt      string `ini:"salt"`
	MinLength int    `ini:"min_length"`
}

type JwtConfig struct {
	Secret     string        `ini:"secret"`
	ExpireTime time.Duration `ini:"expire_time"`
}

type AllConfig struct {
	MySqlConf
	LogConfig
	HashIdsConfig
	JwtConfig
}

var (
	Config  *AllConfig
	EnvMode string = "dev"
)

func init() {
	// Parse GIN_MODE
	// flag.StringVar(&EnvMode, "env", "dev", "--env dev/test/prod")
	// flag.Parse()

	Config = &AllConfig{MySqlConf{
		Host: "localhost",
		Port: "3306",
	},
		LogConfig{},
		HashIdsConfig{},
		JwtConfig{},
	}

	confPath := path.Join(getCurrentPath(), EnvMode+"_conf.ini")

	cfg, err := ini.Load(confPath)
	if err != nil {
		log.Print("Load conf_*.ini error.")
	}

	err = cfg.Section("mysql").MapTo(&Config.MySqlConf)
	err = cfg.Section("log").MapTo(&Config.LogConfig)
	err = cfg.Section("hashIds").MapTo(&Config.HashIdsConfig)
	err = cfg.Section("jwt").MapTo(&Config.JwtConfig)
	if err != nil {
		log.Print("Mapping config error.")
	}
}

// get current *.go path
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
