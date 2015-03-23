package termeter

import (
	"sort"
	"strconv"
)

type numerics []string

func (ns numerics) Len() int {
	return len(ns)
}

func (ns numerics) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns numerics) Less(i, j int) bool {
	nI, errI := strconv.ParseFloat(ns[i], 64)
	nJ, errJ := strconv.ParseFloat(ns[j], 64)

	if errI == nil && errJ == nil {
		return nI < nJ
	}
	if errI != nil && errJ != nil {
		return ns[i] < ns[j]
	}
	if errI == nil && errJ != nil {
		return false
	}
	// if errI != nil && errJ == nil
	return true
}

func numericalSort(ss []string) {
	sort.Sort(numerics(ss))
}
