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
	p1, err := db.con.Prepare("SELECT * FROM players WHERE uuid=?")
	if err != nil {
		panic(err)
	}
	p2, err := db.con.Prepare("SELECT * FROM players WHERE username=?")
	if err != nil {
		panic(err)
	}
	p3, err := db.con.Prepare("INSERT INTO players VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	p4, err := db.con.Prepare("UPDATE players SET roles=? WHERE uuid=?")
	if err != nil {
		panic(err)
	}

	db.prepares["select_by_uuid"] = p1
	db.prepares["select_by_username"] = p2
	db.prepares["insert"] = p3
	db.prepares["edit_roles"] = p4
}
func (db *PlayerDB) AddPlayer(infos database.PlayerInfos) {
	_, err := db.prepares["insert"].Exec(
		infos.UUID,
		infos.Name,
		strings.Join(utils.IntSliceToStringSlice(infos.Roles), ";"),
		strings.Join(infos.AllowedPermissions, ";"),
		strings.Join(infos.DinnedPermissions, ";"),
	)
	if err != nil {
		panic(err)
	}
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
	err := db.prepares["select_by_username"].QueryRow(name).Scan(&p)
	if err != nil {
		if err == sql.ErrNoRows {
			return database.PlayerInfos{}, false
		}
		panic(err)
	}
	return formatToGperms(p), true
}
func (db *PlayerDB) AddRole(infos database.PlayerInfos, id int) {
	_, err := db.prepares["edit_roles"].Exec(strings.Join(utils.IntSliceToStringSlice(append(infos.Roles, id)), ";"), infos.UUID)
	if err != nil {
		panic(err)
	}
}
func (db *PlayerDB) RemoveRole(infos database.PlayerInfos, id int) {
	var r []int
	for _, i := range infos.Roles {
		if i != id {
			r = append(r, i)
		}
	}
	_, err := db.prepares["edit_roles"].Exec(strings.Join(utils.IntSliceToStringSlice(r), ";"), infos.UUID)
	if err != nil {
		panic(err)
	}
}
func (db *PlayerDB) ClearRoles(infos database.PlayerInfos) {
	_, err := db.prepares["edit_roles"].Exec("", infos.UUID)
	if err != nil {
		panic(err)
	}
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
