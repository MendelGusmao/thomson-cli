package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type password struct {
	modules.Base
}

func (*password) URI() string {
	return "/RgSecurity.asp"
}
