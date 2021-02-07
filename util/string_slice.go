package util

func StrSliceContains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func StrSliceIntersects(a, b []string) []string {
	intersects := []string{}
	for _, v := range b {
		if StrSliceContains(a, v) {
			intersects = append(intersects, v)
		}
	}
	return intersects
}
