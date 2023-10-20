package config

import "go-project-example/internal/pkg/log"

type Config struct {
	Server struct {
		HostName string             `json:"-"`
		Ip       string             `json:"-"`
		Port     uint16             `json:"port"`
		Log      map[log.Log]string `json:"log"`
	} `json:"server"`
	DB struct {
		Mysql struct {
			Host     string `json:"host"`
			Database string `json:"database"`
			User     string `json:"user"`
			Password string `json:"password"`
		} `json:"mysql"`
	} `json:"db"`
	Cache struct {
		Redis struct {
			Host     string `json:"host"`
			Prefix   string `json:"prefix"`
			Database int    `json:"database"`
			Password string `json:"password"`
		} `json:"redis"`
	} `json:"cache"`
}
