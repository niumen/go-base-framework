package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-frame/go-base-framework/model"
	"github.com/spf13/viper"
	"logger"
	"net/http"
	"time"
)

func StartGin(port string) {
	var logConfig model.LogConfig
	b, _ := json.Marshal(viper.Get("log_config"))
	json.Unmarshal(b, &logConfig)
	fmt.Println("log_file",logConfig)
	r := gin.New()
	r.Use(logger.Logger("./module/modgin/logs", "gin", time.Second*60*60*24*time.Duration(logConfig.MaxAge), time.Second*60*60*24*time.Duration(logConfig.RotationTime)))

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin test")
	})
	apiv1.POST("/group/create", CreateGroup)
	apiv1.GET("/group/search", SearchGroup)
	r.Run(port)
}