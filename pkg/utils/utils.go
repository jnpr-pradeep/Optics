package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GetOpticsQueryParams(cableSpeed string) (cableType, connectorType, speed string, postFix []string) {
	if cableSpeed == "SMD_100G" {
		cableType = "SMF"
		connectorType = "Duplex LC"
		speed = "100 Gigabit Ethernet"
		postFix = []string{"DR", "CWDM-C", "CWDM"}
	} else if cableSpeed == "SMP_100G" {
		cableType = "SMF"
		connectorType = "MPO-12"
		speed = "100 Gigabit Ethernet"
		postFix = []string{"PSM4"}
	} else if cableSpeed == "MMD_100G" {
		cableType = "MMF"
		connectorType = "Duplex LC"
		speed = "100 Gigabit Ethernet"
		postFix = []string{"BSXR"}
	} else if cableSpeed == "MMP_100G" {
		cableType = "MMF"
		connectorType = "MPO-12"
		speed = "100 Gigabit Ethernet"
		postFix = []string{"BSXR"}
	} else if cableSpeed == "SMD_400G" {
		cableType = "SMF"
		connectorType = "Duplex LC"
		speed = "400 Gigabit Ethernet"
		postFix = []string{"FR4"}
	} else if cableSpeed == "SMP_400G" {
		cableType = "SMF"
		connectorType = "MPO-12"
		speed = "400 Gigabit Ethernet"
		postFix = []string{"DR4"}
	} else if cableSpeed == "MMD_400G" {
		cableType = "MMF"
		connectorType = "Duplex LC"
		speed = "400 Gigabit Ethernet"
	} else if cableSpeed == "MMP_400G" {
		cableType = "MMF"
		connectorType = "MPO-12"
		speed = "400 Gigabit Ethernet"
	}
	return
}

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
	// TODO: disabling the distnace filtering to show all.
	return false
}

func EndsWithKey(key string, l []string) string {
	for _, lItem := range l {
		if strings.HasSuffix(lItem, key) {
			return lItem
		}
	}
	return ""
}
