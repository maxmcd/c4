package config

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/kelseyhightower/envconfig"
)

// API defines environment for the core api server
var API struct {
	Port      int `default:"8083"`
	ProxyPort int `default:"8085"`
}

// DB defines environment for the database
var DB struct {
	Username   string `default:"c4"`
	Password   string `default:"password"`
	Host       string `default:"postgres"`
	Port       int    `default:"26257"`
	Database   string `default:"c4"`
	DisableSSL bool   `default:"true"`
}

// Secret defines external service api keys
var ApiKeys struct {
	NexmoKey    string `default:"notakey"`
	NexmoSecret string `default:"notasecret"`
}

// Auth defines environment for authentication
var Auth struct {
	Secret string `default:"notasecret"`
}

// Load env variables on import
func init() {
	if err := envconfig.Process("API", &API); err != nil {
		glog.Fatal(err)
	}
	if err := envconfig.Process("APIKEY", &ApiKeys); err != nil {
		glog.Fatal(err)
	}
	if err := envconfig.Process("DB", &DB); err != nil {
		glog.Fatal(err)
	}
	if err := envconfig.Process("AUTH", &Auth); err != nil {
		glog.Fatal(err)
	}
}

func DatabaseUrl(protocol string) string {
	url := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s",
		protocol,
		DB.Username,
		DB.Password,
		DB.Host,
		DB.Port,
		DB.Database,
	)
	if DB.DisableSSL {
		url += "?sslmode=disable"
	}
	return url
}
