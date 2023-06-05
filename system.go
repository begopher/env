package env

import (
	"os"
)

func System(file string) (Vars, error) {
	vars, err := _read(file)
	return _system{vars}, err
}

type _system struct {
	vars map[string]string
}

func (s _system) Get(key string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return s.vars[key]
}
