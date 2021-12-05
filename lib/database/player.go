package database

type PlayerInfos struct {
	UUID               string   `json:"uuid"`
	Name               string   `json:"name"`
	Roles              []int    `json:"roles"`
	AllowedPermissions []string `json:"allowed_permissions"`
	DinnedPermissions  []string `json:"dinned_permissions"`
}
type PlayerDB interface {
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
