package db

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"projects/currency/app/config"
)

type SqlDB struct {
	*sql.DB
}

func Dial(cfg *config.Config) (*SqlDB,error) {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		cfg.MsSql.Server, cfg.MsSql.User, cfg.MsSql.Password, cfg.MsSql.Port, cfg.MsSql.Db)
	conn, err := sql.Open("mssql", connString)

	// Test if the connection is OK or not
	if err != nil {
		panic("Cannot connect to database")
	} else {
		fmt.Println("Connected!")
	}

	return &SqlDB{conn}, nil
}
