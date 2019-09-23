package conf

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)
type App struct{
	MaxUploadSize int64
	PasswordPrefix string
}

type Server struct {
	RunMode		string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	DBName        string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Oss struct{
	OSSEndpoint string
	OSSBucket string
	OSSAccessKeyID string
	OSSAccessKeySecret string
}

var (
	AppSetting = &App{}
	ServerSetting = &Server{}
	DBSetting = &Database{}
	RedisSetting = &Redis{}
	OssSetting = &Oss{}
)


func init(){
	cfg,err := ini.Load("conf/app.ini")
	if err != nil{
		log.Fatalf(" fail to parse 'app.ini': %v", err)
	}
	cfg.Section("app").MapTo(AppSetting)
	cfg.Section("server").MapTo(ServerSetting)
	cfg.Section("database").MapTo(DBSetting)
	cfg.Section("redis").MapTo(RedisSetting)
	cfg.Section("oss").MapTo(OssSetting)

	AppSetting.MaxUploadSize = 1024*1024*AppSetting.MaxUploadSize
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}




