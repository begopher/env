package env

import (
	"os"
)

func File(name string) (Vars, error) {
	vars, err := _read(name)
	return _file{vars}, err
}

type _file struct {
	vars map[string]string
}

func (f _file) Get(key string) string {
	if val, ok := f.vars[key]; ok {
		return val
	}
	return os.Getenv(key)
}
