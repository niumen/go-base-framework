package main

import (
	"flag"
	"fmt"
	"github.com/go-frame/go-base-framework/db"
	"github.com/go-frame/go-base-framework/module/modgorm/controller"
	"github.com/go-frame/go-base-framework/module/modgorm/g"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func main() {
	cfgTmp := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	cfg := *cfgTmp
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath("/")
	viper.AddConfigPath("./config")
	cfg = strings.Replace(cfg, ".json", "", 1)
	viper.SetConfigName(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	g.InitLog()
	db.InitDB()
	controller.TimeCycle()
}
