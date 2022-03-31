package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func Open(config Config, models ...interface{}) *Database {
	if !config.Enable {
		return nil
	}

	var open = func(tag string, name string, list []Server) []gorm.Dialector {
		var dialectors = []gorm.Dialector{}
		for i,          v := range list {
			var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", []interface{}{
				v.User,
				v.Pwd,
				v.Host,
				v.Port,
				name,
				"charset=utf8mb4&parseTime=True&loc=Local",
			}...)
			fmt.Printf("dsn|%s[%d]: %s", tag, i, dsn)
			dialectors = append(dialectors, mysql.Open(dsn))
		}
		return dialectors
	}
	var data = dbresolver.Config{
		Sources : open("source", config.Name, config.Source),
		Replicas: open("replica", config.Name, config.Replica),
	}

	var db *gorm.DB
	var err error
	if db, err = gorm.Open(data.Sources[0], &gorm.Config{}); err != nil {
		fmt.Printf("%v", err)
	}
	db.Use(dbresolver.Register(data))

	if err = db.AutoMigrate(models...); err != nil {
		fmt.Printf("%v", err)
	}

	return &Database{impl: db}
}
