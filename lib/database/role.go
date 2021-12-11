package database

type RoleInfos struct {
	// the id of the role
	Id int `json:"id" orm:"id"`
	// the name of the role
	Name int `json:"name" orm:"name"`
	// position start with 0
	Position int `json:"position" orm:"position"`
	// all permission that the role have
	AllowedPermission []string `json:"allowed_permission" orm:"allowed_permission"`
	// all the permission don't have (like if a parent role have it or the permission are allowed by default)
	DinnedPermission []string `json:"dinned_permission" orm:"dinned_permission"`
	// the id of the parent role, -1 if he doesn't have parent role
	ParentId int `json:"parent_id" orm:"parent_id"`
}
type RoleDB interface {
	RoleById(id int) RoleInfos
	CreateRole(infos RoleInfos) RoleInfos
	Roles() []RoleInfos
	DeleteRole(id int)
	EditPosition(id, newPos int)
	EditParent(id, parentId int)
	AddPermission(id int, permName string, allow bool)
	removePermission(id int, permName string)
	ClearPermissions(id int)
}
