package gplayerinfo

import (
	"fmt"
	"github.com/arcoz0308/Gperms/lib"
	player2 "github.com/arcoz0308/Gperms/lib/player"
	"github.com/arcoz0308/Gperms/lib/role"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"strings"
)

var GPLAYERINFO = cmd.New("gplayerinfo", "get gperms info about a player", []string{"gpi"}, GPlayerInfo{})

type GPlayerInfo struct {
	Player string
}

func (g GPlayerInfo) Run(_ cmd.Source, output *cmd.Output) {
	if p, ok := lib.PlayerDB().PlayerByName(g.Player); ok {
		var r []*role.Role
		for _, id := range p.Roles {
			r = append(r, role.Roles[id])
		}
		if len(r) == 0 {
			if len(p.AllowedPermissions) == 0 && len(p.DinnedPermissions) == 0 {
				output.Printf("%s's gperms infos : \nroles : this player don't have any roles\npermissions : this player don't have personal permissions", g.Player)
			} else {
				var permissions []string
				for _, perm := range p.AllowedPermissions {
					permissions = append(permissions, fmt.Sprintf("%s : true", perm))
				}
				for _, perm := range p.DinnedPermissions {
					permissions = append(permissions, fmt.Sprintf("%s : false", perm))
				}
				output.Printf("%s's gperms infos : \nroles : this player don't have any roles\npermissions : %s", g.Player, strings.Join(permissions, ", "))
			}
		} else {
			r = role.SortByPosition(r)
			var rstring []string
			for _, role2 := range r {
				rstring = append(rstring, fmt.Sprintf("%s(%d)", role2.Name(), role2.Id()))
			}
			if len(p.AllowedPermissions) == 0 && len(p.DinnedPermissions) == 0 {
				output.Printf("%s's gperms infos : \nroles : %s\npermissions : this player don't have personal permissions", g.Player, strings.Join(rstring, ", "))
			} else {
				var permissions []string
				for _, perm := range p.AllowedPermissions {
					permissions = append(permissions, fmt.Sprintf("%s : true", perm))
				}
				for _, perm := range p.DinnedPermissions {
					permissions = append(permissions, fmt.Sprintf("%s : false", perm))
				}
				output.Printf("%s's gperms infos : \nroles : %s\npermissions : %s", g.Player, strings.Join(rstring, ", "), strings.Join(permissions, ", "))
			}
		}
	} else {
		output.Errorf("Player with username \"%s\" not found", g.Player)
	}
}
func (g GPlayerInfo) Allow(source cmd.Source) bool {
	if p, ok := source.(*player.Player); ok {
		gp := player2.GPlayer(p)
		if gp != nil {
			return gp.HasPermission("gperms.cmd.gplayerinfo")
		}
		return false
	} else {
		return true
	}
}
