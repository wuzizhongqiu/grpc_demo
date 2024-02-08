package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

var (
	config GlobalConfig
	once   sync.Once
)

type SvrConfig struct {
	SvrName     string `mapstructure:"svr_name"`
	Port        int    `mapstructure:"port"`
	UserSvrName string `mapstructure:"user_svr_name"` // 用户服务name
}

type DbConfig struct {
	Url         string `mapstructure:"url"`
	User        string `mapstructure:"user"`          // 用户名
	Password    string `mapstructure:"password"`      // 密码
	Dbname      string `mapstructure:"dbname"`        // db名
	MaxIdleConn int    `mapstructure:"max_idle_conn"` // 最大空闲连接数
	MaxOpenConn int    `mapstructure:"max_open_conn"` // 最大打开的连接数
	MaxIdleTime int    `mapstructure:"max_idle_time"` // 连接最大空闲时间
}

type ConsulConfig struct {
	Host string   `mapstructure:"host" json:"host" yaml:"host"`
	Port int      `mapstructure:"port" json:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" json:"tags" yaml:"tags"`
}

type GlobalConfig struct {
	SvrConfig    SvrConfig    `mapstructure:"svr_config"`
	DbConfig     DbConfig     `mapstructure:"db"`
	ConsulConfig ConsulConfig `mapstructure:"consul"`
}

func GetGlobalConfig() *GlobalConfig {
	once.Do(readConfig)
	return &config
}

func readConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./gatewaysvr/config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file err:" + err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic("config file unmarshal err:" + err.Error())
	}

	log.Infof("config === %+v", config)
}
