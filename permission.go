package gopm

import (
	"sort"
	"strings"
)

func (m *Manager) AddPermissions(scopes []string) {
	for _, scope := range scopes {
		m.AddPermission(scope)
	}
}

func (m *Manager) RemovePermissions(scopes []string) {
	for _, scope := range scopes {
		m.RemovePermission(scope)
	}
}

func (m *Manager) AddPermission(scope string) {
	m.raw[scope] = nil
}

func (m *Manager) RemovePermission(scope string) {
	delete(m.raw, scope)
}

func (m *Manager) SavePermissions() []string {
	nodes := make([]string, 0)
	for node := range m.raw {
		nodes = append(nodes, node)
	}
	return nodes
}

func (m *Manager) LoadPermissions() {
	m.load()
}

func (m *Manager) GetPermissionChilds(permission string) []string {
	nodes := strings.Split(permission, ".")

	first, other := nodes[0], nodes[1:]
	cur := m.nodes[first]
	for _, node := range other {
		n, ok := cur.Node[node]
		if !ok {
			return nil
		}
		cur = n
	}

	return mapToSlice(cur.Values)
}

func (m *Manager) HasPermissions(scopes []string) bool {
	for _, scope := range scopes {
		if !m.HasPermission(scope) {
			return false
		}
	}
	return true
}

func (m *Manager) HasPermission(scope string) bool {
	// asterisk check always take priority, when any of the asterisk exists
	// we will not needed go through permission node check
	return m.hasAsterisk(scope) || m.hasPermission(scope)
}

func (m *Manager) hasAsterisk(scope string) bool {
	nodes := strings.Split(scope, ".")

	// Example Node  - merchant.create.5
	// Check Node    - merchant.*, merchant.create.*
	// last node always removed, asterisk will only replace all other than last node
	nodes = nodes[:len(nodes)-1]

	cursor := make([]string, 0)
	for _, node := range nodes {
		cursor = append(cursor, node)
		permission := strings.Join(append(cursor, "*"), ".")
		if m.hasPermission(permission) {
			return true
		}
	}
	return false
}

func (m *Manager) hasPermission(scope string) bool {
	nodes := strings.Split(scope, ".")

	// Check last node is an integer value, so we can do comparison based on number node
	// when it's integer we will reserved last node for final check, after traverse all other nodes
	// and only final node can be number node, other will be ignored as treat as a permission node
	last := nodes[len(nodes)-1]
	i := m.tryParseInt(last)
	if i != nil {
		nodes = nodes[:len(nodes)-1]
	}

	first, other := nodes[0], nodes[1:]
	cur := m.nodes[first]

	for _, node := range other {
		n, ok := cur.Node[node]
		if !ok {
			return false
		}
		cur = n
	}

	// after traverse all permission node, and come to final check the nodeValues
	// most of the time the nodeValues will be list of integers or only one integers
	// that will be how much users can be done
	if i != nil {
		lists := m.trySliceInt(mapToSlice(cur.Values))
		sort.Ints(lists)
		return *i <= lists[len(lists)-1]
	}

	return cur != nil
}
