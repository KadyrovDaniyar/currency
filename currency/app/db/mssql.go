package db

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"projects/currency/app/config"
)

type SqlDB struct {
	*gorm.DB
}

func Dial(cfg *config.Config) (*SqlDB,error) {

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", cfg.MsSql.User, cfg.MsSql.Password, cfg.MsSql.Server, cfg.MsSql.Port,cfg.MsSql.Db)
	conn, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	// Test if the connection is OK or not
	if err != nil {
		panic("Cannot connect to database")
	} else {
		fmt.Println("Connected!")
	}

	return &SqlDB{conn}, nil
}