package modules

import (
	"fmt"
	"strings"
)

type Base struct {
	submodules map[string]Module
}

func (*Base) URI() string {
	return ""
}

func (*Base) Template() string {
	return ""
}

func (m *Base) Submodule(modules ...Module) {
	for _, module := range modules {
		if m.submodules == nil {
			m.submodules = make(map[string]Module)
		}

		parts := strings.Split(fmt.Sprintf("%T", module), ".")
		name := strings.ToLower(parts[len(parts)-1 : len(parts)][0])

		m.submodules[name] = module
	}
}

func (m *Base) Submodules() map[string]Module {
	return m.submodules
}

func (m *Base) Find(path []string) (Module, error) {
	var chosen Module = m

	for _, part := range path {
		module, ok := chosen.Submodules()[part]

		if !ok {
			return nil, fmt.Errorf("no module was found using the path %v", path)
		}

		chosen = module
	}

	return chosen, nil
}
