package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GormDB struct {
	*gorm.DB
}

var gormInstance *GormDB = nil

func GetDBInstance() *GormDB{
	if gormInstance == nil {
		if db,err:= gorm.Open("postgres", "host=159.89.205.12 port=5432 user=merchant dbname=web-gateway password=merchant") ; err == nil{
			gormInstance = &GormDB{
				db,
			}
			return gormInstance
		}
	}
	return gormInstance
}

