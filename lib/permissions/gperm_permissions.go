package permissions

func LoadGPermissions() {
	p := []*Permission{
		RegisterPermission("gperms.cmd.giverole", false),
		RegisterPermission("gperms.cmd.removerole", false),
		RegisterPermission("gperms.cmd.gplayerinfo", false),
	}

	RegisterPermission("gperms.cmd.all", false, p...)
}
