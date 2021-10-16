package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST string
	DB_NAME string
	DB_PORT string

	// KHUSUS UNTUK ORACLE
	ORACLE_USERNAME string
	ORACLE_PASSWORD string
	ORACLE_HOST string
	ORACLE_DB string
	ORACLE_PORT string
}

func GetConfig()  Configuration{
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}