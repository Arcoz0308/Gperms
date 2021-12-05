package giverole

import "github.com/df-mc/dragonfly/server/cmd"

var GIVEROLE = cmd.New("giverole", "give a role to a player", nil, GiveRole{})

type GiveRole struct {
	Player string
}

func (g GiveRole) Run(source cmd.Source, output *cmd.Output) {

}
func (g GiveRole) Allow(source cmd.Source) bool {
	return false
}
