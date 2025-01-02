package sql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Sql struct {
}

func New() *Sql {
	return &Sql{}
}

func (s *Sql) Init(dbs *([]*gorm.DB), cfgs ...*Database) {
	// var err error
	for i := range cfgs {
		cfg := cfgs[i]
		log.Printf("initail database %s", cfg.DBName)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("connected to database %s", cfg.DBName)
		*dbs = append(*dbs, db)
	}
}
