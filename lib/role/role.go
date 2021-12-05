package role

import (
	"sort"
)

var Roles []*Role

type Role struct {
	id                 int
	name               string
	position           int
	allowedPermissions []string
	dinnedPermission   []string
	parentId           int
}

func (r *Role) Id() int {
	return r.id
}
func (r *Role) Name() string {
	return r.name
}
func (r *Role) Position() int {
	return r.position
}
func (r *Role) ParentRole() *Role {
	if r.parentId == -1 {
		return nil
	}
	return Roles[r.id]
}

// HasPermission the bool n°2 say if he has found info about perm in role
func (r *Role) HasPermission(permName string) (bool, bool) {
	for _, p := range r.allowedPermissions {
		if p == permName {
			return true, true
		}
	}
	for _, p := range r.dinnedPermission {
		if p == permName {
			return false, true
		}
	}
	parentRole := r.ParentRole()
	if parentRole != nil {
		return parentRole.HasPermission(permName)
	}
	return false, false
}
func SortByPosition(roles []*Role) []*Role {
	if len(roles) < 2 {
		return roles
	}
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].position < roles[j].position
	})
	return Roles
}
