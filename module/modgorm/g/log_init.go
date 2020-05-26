package g

import (
	"encoding/json"
	"github.com/go-frame/go-base-framework/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"logger"
	"time"
)

func InitLog() (err error) {
	var logConfig model.LogConfig
	b, _ := json.Marshal(viper.Get("log_config"))
	json.Unmarshal(b, &logConfig)
	logger.ConfigLocalFilesystemLogger("./module/modgorm/logs", "gorm-log", time.Second*60*60*24*time.Duration(logConfig.MaxAge), time.Second*60*60*24*time.Duration(logConfig.RotationTime))
	loglevel := logConfig.LogLevel
	switch loglevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	return
}
