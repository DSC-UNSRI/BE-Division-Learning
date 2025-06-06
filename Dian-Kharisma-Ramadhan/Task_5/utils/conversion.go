package utils

import (
	"errors"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func SplitPath(path string) ([]string, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return []string{}, errors.New("request is not valid")
	}
	return parts, nil
}