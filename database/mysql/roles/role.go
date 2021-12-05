package roles

import (
	"database/sql"
	"github.com/arcoz0308/Gperms/database/mysql/resources"
	_ "github.com/go-sql-driver/mysql"
)

type RoleDB struct {
	con *sql.DB
}

func Load(dsn string, autoCreateTable bool) *RoleDB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	if autoCreateTable {
		_, err := db.Exec(resources.Roles)
		if err != nil {
			panic(err)
		}
	}
	return &RoleDB{con: db}
}
