package pttw

func FindParam(key string, params []string) string {
	for i := 0; i < len(params) - 1; i++ {
		if (params[i] == key) {
			return params[i + 1]
		}
	}
	return ""
}
