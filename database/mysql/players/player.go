package players

import (
	"database/sql"
	"github.com/arcoz0308/Gperms/database/mysql/resources"
	_ "github.com/go-sql-driver/mysql"
)

type PlayerDB struct {
	con      *sql.DB
	prepares map[string]*sql.Stmt
}

func Load(dsn string, autoCreateTable bool) *PlayerDB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	if autoCreateTable {
		_, err := db.Exec(resources.Player)
		if err != nil {
			panic(err)
		}
	}
	p := &PlayerDB{con: db}
	p.LoadPrepares()
	return p
}
func (db PlayerDB) LoadPrepares() {
	p, err := db.con.Prepare("SELECT * where uuid=?")
	if err != nil {
		panic(err)
	}
	db.prepares["select_by_uuid"] = p
}
