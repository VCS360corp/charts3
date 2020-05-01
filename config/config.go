package config

import (
	"log"

	"github.com/vrischmann/envconfig"
)

//Current holds currently loaded configuration values
var Current Config

//Config holds application configuration values
type Config struct {
	AppName string
	Server  ServerConfig
	URL     URLConfig
	Path    PathConfig
	Mysql   MysqlConfig
}

//ServerConfig holds configuration for server
type ServerConfig struct {
	Address string
	Port    int
}

//URLConfig holds urls required by the application
type URLConfig struct {
	Asset string
	App   string
}

//PathConfig holds paths required by the application
type PathConfig struct {
	Template string
	Asset    string
}

//MysqlConfig holds configuration of mysql
type MysqlConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string `envconfig:"optional"`
}

//Load loads configuration values
func Load() {
	err := envconfig.InitWithPrefix(&Current, "CHARTS_DEMO")
	if err != nil {
		log.Fatalln("Error while loading configuration\n", err)
	}
}
