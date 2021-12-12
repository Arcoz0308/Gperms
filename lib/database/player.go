package database

type PlayerInfos struct {
	UUID               string   `json:"uuid" orm:"uuid"`
	Username           string   `json:"username" orm:"username"`
	Roles              []int    `json:"roles" orm:"roles"`
	AllowedPermissions []string `json:"allowed_permissions" orm:"allowed_permissions"`
	DinnedPermissions  []string `json:"dinned_permissions" orm:"dinned_permissions"`
}
type PlayerDB interface {
	// AddPlayer this function are called when a new player join the server
	AddPlayer(infos PlayerInfos)
	// PlayerByUUID this function
	PlayerByUUID(uuid string) (PlayerInfos, bool)
	PlayerByName(name string) (PlayerInfos, bool)
	AddRole(infos PlayerInfos, role int)
	RemoveRole(infos PlayerInfos, role int)
	ClearRoles(infos PlayerInfos)
	AddPermission(infos PlayerInfos, permName string, allow bool)
	RemovePermission(infos PlayerInfos, permName string)
	ClearPermissions(infos PlayerInfos)
	ClearRolesAndPermissions(infos PlayerInfos)
}
