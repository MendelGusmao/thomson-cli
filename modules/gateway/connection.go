package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type connection struct {
	modules.Base
}

func (*connection) URI() string {
	return "/RgConnect.asp"
}
