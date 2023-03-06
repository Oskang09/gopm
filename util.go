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
		i, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		r = append(r, i)
	}
	return r
}
