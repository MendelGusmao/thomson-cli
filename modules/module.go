package modules

var Root Module = new(Base)

type Module interface {
	URI() string
	Template() string
	Submodule(modules ...Module)
	Submodules() map[string]Module
	Find([]string) (Module, error)
}
