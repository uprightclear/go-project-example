package db

import (
	"go-project-example/internal/pkg/db/mysql"
	"gorm.io/gorm"
)

func Initialize(host, database, user, password string) (err error) {
	return mysql.Initialize(host, database, user, password)
}

func GetMySQL() *gorm.DB {
	return mysql.Get()
}
