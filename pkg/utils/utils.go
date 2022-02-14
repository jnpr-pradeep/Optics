package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ContainsKey(key string, l []string) bool {
	for _, lItem := range l {
		if lItem == key {
			return true
		}
	}
	return false
}

func IsDistanceSupported(distance string) bool {
	// Only supported speeds - <100m, <500m
	d := strings.Split(distance, " ")
	if d[1] == "m" {
		n, err := strconv.ParseFloat(d[0], 32)
		if err != nil {
			fmt.Println(err)
		}
		// Rack height is close to 10 feet, so for fabric links we need optics between 10 m to 500 m.
		return 10.0 <= n && n <= 500.0
	}
	return false
}
