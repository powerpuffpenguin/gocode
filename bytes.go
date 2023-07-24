package gocode

import "unsafe"

// Convert byte slice to string
//
// It doesn't reallocate memory so it's faster than standard conversion
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Convert string to readonly byte slice
//
// It doesn't reallocate memory so it's faster than standard conversion.
// But be aware that slices are read-only if written behavior will be unknown.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
