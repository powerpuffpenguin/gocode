package gocode

func IsExport(name string) bool {
	r := []rune(name)
	if len(r) < 1 {
		return false
	}
	return r[0] >= 'A' && r[0] <= 'Z'
}
