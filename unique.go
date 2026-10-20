package main

import "slices"

//unique
func unique(a []int64) []int64 {
	slices.Sort(a)
	return slices.Compact(a)
}
