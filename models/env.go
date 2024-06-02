package models

type Env struct {
	Host     string
	User     string
	DbName   string
	Password string
}

var Config Env
