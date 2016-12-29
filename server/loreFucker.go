package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateSubnet(start string, group int, end string) ([]string, error) {
	startIpSplit := strings.Split(start, ".")
	endIpSplit := strings.Split(end, ".")
	result := make([]string, 0)
	k := 0
	firstStratValue, err := strconv.Atoi(startIpSplit[0])
	if err != nil {
		return nil, err
	}
	firstEndValue, err := strconv.Atoi(endIpSplit[0])
	if err != nil {
		return nil, err
	}
	SecondStratValue, err := strconv.Atoi(startIpSplit[1])
	if err != nil {
		return nil, err
	}
	SecondEndValue, err := strconv.Atoi(endIpSplit[1])
	if err != nil {
		return nil, err
	}
	ThirdStratValue, err := strconv.Atoi(startIpSplit[2])
	if err != nil {
		return nil, err
	}
	ThirdEndValue, err := strconv.Atoi(endIpSplit[2])
	if err != nil {
		return nil, err
	}
	last := ThirdStratValue + group
	if SecondEndValue != SecondStratValue {
		ThirdEndValue = 256
	}
	for i := firstStratValue; i <= firstEndValue; i++ {
		for j := SecondStratValue; j <= SecondEndValue; j++ {
			for k = ThirdStratValue + group; k < ThirdEndValue; k += group {
				result = append(result, fmt.Sprintf("%d.%d.%d-%d.1-254", i, j, (k-last), k))
			}
			last = ThirdStratValue + group
			if k >= 256 {
				k = 0
			}
			result = append(result, fmt.Sprintf("%d.%d.253-256.1-254", i, j))
		}

	}

	return result, nil
}
