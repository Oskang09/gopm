package gopm

import "strconv"

func (m *Manager) tryParseInt(val string) *int {
	i, e := strconv.Atoi(val)
	if e != nil {
		return nil
	}
	return &i
}

func (m *Manager) trySliceInt(values []string) []int {
	r := make([]int, 0)
	for _, val := range values {
		i, _ := strconv.Atoi(val)
		r = append(r, i)
	}
	return r
}
