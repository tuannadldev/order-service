package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v <= first {
			first = v
		} else if v <= second {
			second = v
		} else {
			fmt.Println("====> ", first, v, second)
		}

		fmt.Println(first, v, second)
	}
	return false
}
func increasingTriplets(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			//
			fmt.Println("====> ", first, num, second)
			return true
		} // 5 5 12
		fmt.Println("==== ", first, num, second)
	}
	return false
}

func compress(chars []byte) int {
	write, read := 0, 0
	for read < len(chars) {
		char := chars[read]
		count := 1
		for read+1 < len(chars) && chars[read+1] == char {
			count++
			read++
		}
		chars[write] = char
		write++
		if count > 1 {
			countStr := []byte(strconv.Itoa(count))
			for _, c := range countStr {
				chars[write] = c
				write++
			}
		}
		read++
	}
	return write
}

func main() {
	chars := []string{"a", "a", "b", "b", "c", "c", "c"}
	fmt.Println(strings.Join(chars, ""))
	bytes := []byte(strings.Join(chars, ""))
	fmt.Println(bytes)
	for _, v := range bytes {
		fmt.Println(string(v))
	}
	fmt.Println(compress(bytes))
}
