package env

type Vars interface {
	Get(key string) string
} 