//
// map.go
//
// Distributed under terms of the MIT license.
//

package util

func CopyMapString(m map[string]string) map[string]string {
	if m == nil {
		return map[string]string{}
	}
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

func CopyMapInterface(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return map[string]interface{}{}
	}
	n := make(map[string]interface{}, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

func CopyMap2Interface(m map[string]string) map[string]interface{} {
	if m == nil {
		return map[string]interface{}{}
	}
	n := make(map[string]interface{}, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

// ExtendMapString merge n to m
func ExtendMapString(m, n map[string]string) map[string]string {
	if m == nil {
		m = make(map[string]string, len(n))
	}
	for k, v := range n {
		m[k] = v
	}
	return m
}

// MergeMapString return new map
func MergeMapString(m, n map[string]string) map[string]string {
	a := make(map[string]string, len(m)+len(n))
	for k, v := range m {
		a[k] = v
	}
	for k, v := range n {
		a[k] = v
	}
	return a
}
