package main

import (
	"flag"
	"fmt"
	"strconv"
)

const intMin = ^int(^uint(0) >> 1)

func findMaxSubArray(source []int, show bool) (int, int, int) {
	if len(source) == 1 {
		if show {
			fmt.Println("only one element")
		}
		return 0, 0, source[0]
	}
	startIndex := 0
	endIndex := 0
	maxSum := intMin
	for i := 0; i < len(source); i++ {
		tempSum := source[i]
		if tempSum > maxSum {
			startIndex = i
			endIndex = i
			maxSum = tempSum
			if show {
				fmt.Println("get: start index", i, "end index", i, "current max sum", maxSum)
			}
		} else {
			if show {
				fmt.Println("pass: this index", i, "element", source[i], " is less than max sum", maxSum)
			}
		}
		for j := i + 1; j < len(source); j++ {
			tempSum += source[j]
			if tempSum > maxSum {
				startIndex = i
				endIndex = j
				maxSum = tempSum
				if show {
					fmt.Println("get: start index", i, "end index", j, "current max sum", maxSum)
				}
			} else {
				if show {
					fmt.Println("pass: start index", i, "end index", j, "sum", tempSum,
						"is less than max sum", maxSum)
				}
			}
		}
	}
	return startIndex, endIndex, maxSum
}

func main() {
	show := flag.Bool("s", false, "show detaiL flag")
	flag.Parse()

	source := []int{}
	for _, arg := range flag.Args() {
		i, _ := strconv.Atoi(arg)
		source = append(source, i)
	}
	fmt.Println("start find:", source)
	fmt.Println()
	startIndex, endIndex, maxSum := findMaxSubArray(source, *show)
	fmt.Println()
	fmt.Println("start index:", startIndex, "end index:", endIndex, "sum:", maxSum)
}
