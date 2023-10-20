package main

import (
	"go-project-example/internal/pkg/config"
	"go-project-example/internal/pkg/log"
)

func load(configPath string) (err error) {
	if err = config.Initialize(configPath); err != nil {
		return
	}
	var conf = config.Get()

	var logConf = make(map[log.Log]log.Config)
	for l, path := range conf.Server.Log {
		logConf[l] = log.Config{
			Path: path,
		}
	}
	if err = log.Initialize(logConf, "info"); err != nil {
		return
	}

	//if err = db.Initialize(
	//	conf.DB.Mysql.Host,
	//	conf.DB.Mysql.Database,
	//	conf.DB.Mysql.User,
	//	conf.DB.Mysql.Password); err != nil {
	//	return
	//}
	//
	//if err = cache.Initialize(
	//	conf.Cache.Redis.Host,
	//	conf.Cache.Redis.Password,
	//	conf.Cache.Redis.Prefix,
	//	conf.Cache.Redis.Database); err != nil {
	//	return
	//}

	return
}
