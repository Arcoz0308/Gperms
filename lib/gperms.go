package lib

import (
	"errors"
	"github.com/arcoz0308/Gperms/lib/commands/giverole"
	"github.com/arcoz0308/Gperms/lib/commands/gplayerinfo"
	"github.com/arcoz0308/Gperms/lib/database"
	"github.com/arcoz0308/Gperms/lib/permissions"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

var (
	roleDB   database.RoleDB
	playerDB database.PlayerDB
	loaded   = false
	Serv     *server.Server
)

func Load(roleDB database.RoleDB, db database.PlayerDB) error {
	if loaded {
		return errors.New("package are already loaded")
	}
	permissions.LoadGPermissions()
	return nil
}
func LoadCommands(serv *server.Server) error {
	if !loaded {
		return errors.New("you need to load package before load commands")
	}
	Serv = serv
	cmd.Register(giverole.GIVEROLE)
	cmd.Register(gplayerinfo.GPLAYERINFO)
	return nil
}
func RoleDB() database.RoleDB {
	return roleDB
}
func PlayerDB() database.PlayerDB {
	return playerDB
}
