package resources

var Roles = "CREATE TABLE IF NOT EXISTS `roles` (" +
	"`id` int(11) NOT NULL," +
	"`name` varchar(50) NOT NULL DEFAULT ''," +
	"`position` int(11) NOT NULL," +
	"`allowed_permissions` longtext DEFAULT NULL," +
	"`dinned_permissions` longtext DEFAULT NULL," +
	"`parent_id` int(11) DEFAULT NULL" +
	")"
var Player = "CREATE TABLE IF NOT EXISTS `players` (" +
	"`uuid` varchar(50) NOT NULL COMMENT 'player uuid'," +
	"`username` varchar(50) NOT NULL COMMENT 'player username'," +
	"`roles` varchar(50) DEFAULT NULL COMMENT 'player roles, syntax : \"id;id;id\"'," +
	"`allow_permission` longtext DEFAULT '' COMMENT 'player allowed permissions, syntax : \"perm_1;perm_2;other_perm\"'," +
	"`dinned_permissions` longtext DEFAULT '' COMMENT 'player dinned permissions, syntax : \"perm_1;perm_2;other_perm\"'" +
	")"
