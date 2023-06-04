package env

import(
	"os"
	"strings"
	"bufio"
)

func _read(name string) (map[string]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return map[string]string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	vars := make(map[string]string)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}
		keyVal := strings.Split(line, "=")
		if len(keyVal) != 2 {
			continue
		}
		key := strings.TrimSpace(keyVal[0])
		val := strings.TrimSpace(keyVal[1])
		vars[key]=val
	}
	if err := scanner.Err(); err != nil {
		return map[string]string{}, err
	}
	return vars, nil
}