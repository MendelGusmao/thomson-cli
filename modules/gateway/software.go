package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type software struct {
	modules.Base
}

func (*software) URI() string {
	return "/RgSwInfo.asp"
}
