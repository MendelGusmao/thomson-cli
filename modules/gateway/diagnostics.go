package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type diagnostics struct {
	modules.Base
}

func (*diagnostics) URI() string {
	return "/RgDiagnostics.asp"
}
