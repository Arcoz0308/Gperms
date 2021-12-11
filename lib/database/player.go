package database

type PlayerInfos struct {
	UUID               string   `json:"uuid" orm:"uuid"`
	Name               string   `json:"name" orm:"name"`
	Roles              []int    `json:"roles" orm:"roles"`
	AllowedPermissions []string `json:"allowed_permissions" orm:"allowed_permissions"`
	DinnedPermissions  []string `json:"dinned_permissions" orm:"dinned_permissions"`
}
type PlayerDB interface {
	AddPlayer(infos PlayerInfos)
	PlayerByUUID(uuid string) (PlayerInfos, bool)
	PlayerByName(name string) (PlayerInfos, bool)
	AddRole(role int)
	RemoveRole(role int)
	ClearRoles()
	AddPermission(permName string, allow bool)
	RemovePermission(permName string)
	ClearPermissions()
	ClearRolesAndPermissions()
}
