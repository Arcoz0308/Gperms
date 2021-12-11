package players

import (
	"database/sql"
	"github.com/arcoz0308/Gperms/database/mysql/resources"
	"github.com/arcoz0308/Gperms/database/mysql/utils"
	"github.com/arcoz0308/Gperms/lib/database"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type PlayerDB struct {
	con      *sql.DB
	prepares map[string]*sql.Stmt
}
type playerInfos struct {
	UUID              string `orm:"uuid"`
	Username          string `orm:"username"`
	Roles             string `orm:"roles"`
	AllowPermissions  string `orm:"allow_permissions"`
	DinnedPermissions string `orm:"dinned_permissions"`
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
func (db *PlayerDB) LoadPrepares() {
	p, err := db.con.Prepare("SELECT * where uuid=?")
	if err != nil {
		panic(err)
	}
	db.prepares["select_by_uuid"] = p
}
func (db *PlayerDB) AddPlayer(infos playerInfos) {

}
func (db *PlayerDB) PlayerByUUID(uuid string) (database.PlayerInfos, bool) {
	var p playerInfos
	err := db.prepares["select_by_uuid"].QueryRow(uuid).Scan(&p)
	if err != nil {
		if err == sql.ErrNoRows {
			return database.PlayerInfos{}, false
		}
		panic(err)
	}
	return formatToGperms(p), true
}
func (db *PlayerDB) PlayerByName(name string) (database.PlayerInfos, bool) {
	var p playerInfos
	err := db.con.QueryRow("SELECT * where username=?", name).Scan(&p)
	if err != nil {
		if err == sql.ErrNoRows {
			return database.PlayerInfos{}, false
		}
		panic(err)
	}
	return formatToGperms(p), true
}
func (db *PlayerDB) AddRole(id int) {

}

func formatToGperms(p playerInfos) database.PlayerInfos {
	roles := utils.StringSliceToIntSlice(strings.Split(p.Roles, ";"))
	aPerms := strings.Split(p.AllowPermissions, ";")
	dPerms := strings.Split(p.DinnedPermissions, ";")
	return database.PlayerInfos{
		UUID:               p.UUID,
		Name:               p.Username,
		Roles:              roles,
		AllowedPermissions: aPerms,
		DinnedPermissions:  dPerms,
	}
}
