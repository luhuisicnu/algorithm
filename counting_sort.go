package main

import (
	"flag"
	"fmt"
	"strconv"
)

const intMin = ^int(^uint(0) >> 1)

func countingSort(source []int, result []int, max int, show, desc bool) {
	temp := make([]int, max+1)
	for _, value := range source {
		temp[value]++
	}
	if desc {
		for i := max - 1; i >= 0; i-- {
			temp[i] += temp[i+1]
		}
	} else {
		for i := 1; i <= max; i++ {
			temp[i] += temp[i-1]
		}
	}

	if show {
		fmt.Println("sort: temp counting array init success", temp)
	}
	for j := len(source) - 1; j >= 0; j-- {
		if show {
			fmt.Println("sort: result index", temp[source[j]], "get value", source[j])
		}
		result[temp[source[j]]-1] = source[j]
		temp[source[j]]--
		if show {
			fmt.Println("sort: temp counting array state", temp)
			fmt.Println("sort: result state", result)
		}
	}
}

func max(source []int) int {
	maxValue := intMin
	for _, value := range source {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	show := flag.Bool("s", false, "show detail flag")
	desc := flag.Bool("d", false, "desc flag")
	flag.Parse()

	source := []int{}
	for _, arg := range flag.Args() {
		i, _ := strconv.Atoi(arg)
		source = append(source, i)
	}
	result := make([]int, len(source))

	countingSort(source, result, max(source), *show, *desc)
	fmt.Println(result)
}
