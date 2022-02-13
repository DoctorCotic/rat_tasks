package main

import (
	"bufio"
	"io"
	"strconv"
)

//func main() {
//	get2(os.Stdin, os.Stdout)
//}

func get2(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	maximum := scanner.Text()
	max, _ := strconv.Atoi(maximum)

	nums := make([]int32, 0, max)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, int32(num))
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
