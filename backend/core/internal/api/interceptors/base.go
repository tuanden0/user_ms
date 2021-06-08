package interceptors

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}
