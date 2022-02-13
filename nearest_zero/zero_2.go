package main

import (
	"io"
	"strconv"
	"strings"
)

//func main() {
//	get1(os.Stdin, os.Stdout)
//}

func get1(r io.Reader, w io.Writer) {
	rawData, _ := io.ReadAll(r)
	lines := strings.Split(string(rawData), "\n")
	strs := strings.Split(lines[1], " ")

	nums := make([]int32, len(strs))
	for i, str := range strs {
		num, _ := strconv.Atoi(str)
		nums[i] = int32(num)
	}

	counter := int32(100000)
	result := make([]int32, len(nums))
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

	for i, num := range result {
		if i != 0 {
			w.Write([]byte(" "))
		}
		w.Write([]byte(strconv.Itoa(int(num))))
	}
}
