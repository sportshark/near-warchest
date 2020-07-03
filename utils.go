package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Note. most copy from https://github.com/masknetgoal634/near-go-warchest/blob/master/common/utils.go

func intFromString(s string) int {
	value := strings.Replace(s, ",", "", -1)
	value = strings.TrimSpace(value)
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(v)
}

func stakeFromString(s string) int {
	if len(s) == 1 {
		return 0
	}
	l := len(s) - 19 - 5
	v, err := strconv.ParseFloat(s[0:l], 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(v)
}

func stakeFromNearView(s string) int {
	s2 := strings.Split(s, "})")
	if len(s2) > 1 {
		s3 := s2[1]
		s4 := strings.Split(s3, "m")
		if len(s4) > 1 {
			s5 := strings.Replace(s4[1], "'", "", -1)
			s6 := s5[0 : len(s5)-4]
			return stakeFromString(s6)
		}
	}
	return 0
}

func stringFromStake(stake int) string {
	return fmt.Sprintf("%d%s", stake, "000000000000000000000000")
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}
