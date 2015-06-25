package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type gateway struct {
	modules.Base
}

func (*gateway) URI() string {
	return "/RgSwInfo.asp"
}

func init() {
	gateway := new(gateway)

	gateway.Submodule(
		new(backupRestore),
		new(connection),
		new(diagnostics),
		new(eventLog),
		new(password),
		new(software),
	)

	modules.Root.Submodule(gateway)
}
