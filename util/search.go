package util

// @params int target
// @params []int source
//
// @return int index. -1 is not find.
func FastSearchWithInt(target int, source []int) int {
	if len(source) < 1 {
		return -1
	}

	for idx, v := range source {
		if target == v {
			return idx
		}
	}
	return -1
}
