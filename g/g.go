package g

import (
	"log"
	"os"
	"path/filepath"
)

var Modules map[string]bool
var BinOf map[string]string
var cfgOf map[string]string
var ModuleApps map[string]string
var logpathOf map[string]string
var PidOf map[string]string
var AllModulesInOrder []string
var Rootpath string

func Getpath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	return dir
}

func init() {

	Rootpath = Getpath()
	Modules = map[string]bool{
		"modgin":  true,
		"modgorm": true,
	}

	BinOf = map[string]string{
		"modgin":  Rootpath + "/modgin/bin/modginbin",
		"modgorm": Rootpath + "/modgorm/bin/modgormbin",
	}

	cfgOf = map[string]string{
		"modgin":  Rootpath + "/modgin/config/cfg.json",
		"modgorm": Rootpath + "/modgorm/config/cfg.json",
	}

	ModuleApps = map[string]string{
		"modgin":  "modgin",
		"modgorm": "modgorm",
	}

	logpathOf = map[string]string{
		"modgin":  Rootpath + "/modgin/logs/modgin.log",
		"modgorm": Rootpath + "/modgorm/logs/modgorm.log",
	}

	PidOf = map[string]string{
		"modgin":  "<NOT SET>",
		"modgorm": "<NOT SET>",
	}

	// Modules are deployed in this order
	AllModulesInOrder = []string{
		"modgin",
		"modgorm",
	}
}

func Bin(name string) string {
	p, _ := filepath.Abs(BinOf[name])
	return p
}

func Cfg(name string) string {
	p, _ := filepath.Abs(cfgOf[name])
	return p
}

func LogPath(name string) string {
	p, _ := filepath.Abs(logpathOf[name])
	return p
}

func LogDir(name string) string {
	d, _ := filepath.Abs(filepath.Dir(logpathOf[name]))
	return d
}
