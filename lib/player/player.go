package player

import (
	"github.com/arcoz0308/Gperms/lib/permissions"
	"github.com/arcoz0308/Gperms/lib/role"
	"github.com/df-mc/dragonfly/server/player"
)

var Players []*Player

type Player struct {
	p                                   *player.Player
	roles                               []*role.Role
	allowedPermission, dinnedPermission []string
}

func (p *Player) HasPermission(permName string) bool {
	for _, perm := range p.allowedPermission {
		if perm == permName {
			return true
		}
	}
	for _, perm := range p.dinnedPermission {
		if perm == permName {
			return false
		}
	}
	for _, r := range p.roles {
		ok1, ok2 := r.HasPermission(permName)
		if ok2 {
			return ok1
		}
	}
	perm := permissions.PermissionByName(permName)
	if perm != nil {
		return perm.IsByDefault()
	}
	return false
}
func GPlayer(p *player.Player) *Player {
	for _, p2 := range Players {
		if p.UUID() == p2.p.UUID() {
			return p2
		}
	}
	return nil
}
