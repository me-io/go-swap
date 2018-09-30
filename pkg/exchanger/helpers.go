package exchanger

// ReverseMap ... Reverse a map in <value, key>
func ReverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

// MapKeyArrInterface ... Reverse a map in <value, key>
func MapKeyArrInterface(m map[string]string) []interface{} {
	var n []interface{}
	for k := range m {
		n = append(n, k)
	}
	return n
}
