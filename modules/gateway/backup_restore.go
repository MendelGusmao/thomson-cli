package gateway

import "github.com/MendelGusmao/thomson-cli/modules"

type backupRestore struct {
	modules.Base
}

func (*backupRestore) URI() string {
	return "/RgBackupRestore.asp"
}
