package gopm

import "strings"

type nodeValue struct {
	Values map[string]*struct{}
	Node   node
}

type node map[string]*nodeValue

func (m *Manager) load() {
	m.nodes = make(node) // refresh `nodes` when everytime loads

	for permission := range m.raw {
		attributes := strings.Split(permission, ".")

		var cur *nodeValue // cursor nodeValue
		n := m.nodes       // cursor node
		for _, attribute := range attributes {

			if cur != nil {
				cur.Values[attribute] = nil // update attributes to (previous) cursor nodeValue
			}

			_, ok := n[attribute] // check have next cursor nodeValue
			if !ok {
				n[attribute] = &nodeValue{
					Values: make(map[string]*struct{}),
					Node:   node{},
				}
			}

			cur = n[attribute]    // travel to next attribute nodeValue
			n = n[attribute].Node // travel to next attribute node
		}
	}
}

func mapToSlice(value map[string]*struct{}) []string {
	array := make([]string, 0)
	for k := range value {
		array = append(array, k)
	}
	return array
}
