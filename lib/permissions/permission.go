package permissions

import (
	"fmt"
	"regexp"
	"strings"
)

const PermMaxLength int = 128

var Permissions []*Permission

type Permission struct {
	name        string
	isByDefault bool
	children    []*Permission
	parent      []*Permission
}

func RegisterPermission(name string, isByDefault bool, children ...*Permission) *Permission {
	if len(name) > PermMaxLength {
		panic(fmt.Sprintf("%s are a to long permission name, the limit are %d and %s have a size of %d", name, PermMaxLength, name, len(name)))
	}
	name = strings.ToLower(name)
	if PermissionByName(name) != nil {
		panic(fmt.Sprintf("permission with name %s already loaded", name))
	}
	if !ValidName(name) {
		panic(fmt.Sprintf("permission name of %s are invalid, please only use lowercase characters, \"-\", \".\" and \"_\"", name))
	}
	p := &Permission{name: name, isByDefault: isByDefault, children: children}
	for _, c := range children {
		c.parent = append(c.parent, p)
	}
	Permissions = append(Permissions, p)
	return p
}
func (p *Permission) Name() string {
	return p.name
}
func (p *Permission) IsByDefault() bool {
	return p.isByDefault
}
func (p *Permission) Children() []*Permission {
	return p.children
}
func PermissionByName(name string) *Permission {
	for _, p := range Permissions {
		if p.Name() == name {
			return p
		}
	}
	return nil
}
func ValidName(name string) bool {
	ok, err := regexp.MatchString(`([a-z]|\-|\.|\_*)+`, name)
	if err != nil {
		panic(err)
	}
	return ok
}
