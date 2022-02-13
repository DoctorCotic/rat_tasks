package main

import (
	"io"
	"strconv"
	"strings"
)

//func main() {
//	fmt.Println(strings.Join(get(os.Stdin), " "))
//}

func get(r io.Reader) []string {
	rawData, _ := io.ReadAll(r)
	lines := strings.Split(string(rawData), "\n")
	strs := strings.Split(lines[1], " ")

	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}

	counter := 100000
	result := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			counter = 0
			continue
		}
		counter++
		result[i] = counter
	}

	counter = 10000
	for i := len(result) - 1; i >= 0; i-- {
		num := result[i]
		if num == 0 {
			counter = 0
		}

		if num >= counter {
			result[i] = counter
		}
		counter++
	}

	resultStrs := make([]string, len(result))
	for i, num := range result {
		resultStrs[i] = strconv.Itoa(num)
	}
	return resultStrs
}
