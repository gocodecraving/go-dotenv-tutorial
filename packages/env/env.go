package env

import (
	"bufio"
	"os"
	"strings"
)

func Get(field string) string {
	value := ""

	file, err := os.Open(".env")
	if err != nil {
		panic(".env file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		segments := strings.Split(line, "=")
		if len(segments) == 2 && segments[0] == field {
			value = strings.ReplaceAll(segments[1], `"`, ``)
		}

	}

	return value
}
