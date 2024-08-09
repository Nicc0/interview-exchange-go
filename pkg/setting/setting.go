package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type OpenExchangeApi struct {
	Url          string
	Key          string
	BaseCurrency string
	Timeout      time.Duration
}

var OpenExchangeApiSetting = &OpenExchangeApi{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini.dist")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini.dist': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("openexchangeapi", OpenExchangeApiSetting)

	OpenExchangeApiSetting.Timeout = OpenExchangeApiSetting.Timeout * time.Second
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
