package gopm

type Manager struct {
	raw   map[string]*struct{}
	nodes node
}

func New() *Manager {
	m := new(Manager)
	m.raw = make(map[string]*struct{})
	m.nodes = node{}
	return m
}
