package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	DBType      string
	User        string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
}

var (
	vp              *viper.Viper
	ServerSetting   *Server
	DatabaseSetting *Database
)

func InitSetting(env string) {
	vp = viper.New()
	vp.SetConfigFile("conf/config.yaml")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	InitServerSetting()
	InitDatabaseSetting()
}

func InitServerSetting() {
	err := vp.UnmarshalKey("Server", &ServerSetting)
	if err != nil {
		fmt.Println(err)
	}
}

func InitDatabaseSetting() {
	err := vp.UnmarshalKey("Database", &DatabaseSetting)
	if err != nil {
		fmt.Println(err)
	}
}
