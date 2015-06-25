package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type eventLog struct {
	modules.Base
}

func (*eventLog) URI() string {
	return "/RgEventLog.asp"
}
