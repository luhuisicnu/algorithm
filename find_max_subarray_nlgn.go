package main

import (
	"flag"
	"fmt"
	"strconv"
)

const intMin = ^int(^uint(0) >> 1)

func findMaxCrossingSubArray(source []int, low, mid, high int, show bool) (int, int, int) {
	leftSum := intMin
	maxLeft := mid
	sum := 0
	if show {
		fmt.Println("cross start: left index", low, "right index", high, "source", source[low:high])
	}
	for i := mid - 1; i >= low; i-- {
		sum += source[i]
		if show {
			fmt.Println("cross get: left extend, index", i, "temp sum", sum, "left sum", leftSum)
		}
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
			if show {
				fmt.Println("cross get: find bigger sum, left index", i, "mid index", mid, "sum", sum)
			}
		} else {
			if show {
				fmt.Println("cross pass: cannot find bigger sum, left index", i, "mid index", mid,
					"temp sum", sum, "left sum", leftSum)
			}
		}
	}

	rightSum := intMin
	maxRight := mid
	sum = 0
	for j := mid; j < high; j++ {
		sum += source[j]
		if show {
			fmt.Println("cross get: right extend, index", j, "temp sum", sum, "right sum", rightSum)
		}
		if sum > rightSum {
			rightSum = sum
			maxRight = j
			if show {
				fmt.Println("cross get: find bigger sum, mid index", mid, "right index", j, "sum", sum)
			}
		} else {
			if show {
				fmt.Println("cross pass: cannot find bigger sum, mid index", mid, "right index", j,
					"temp sum", sum, "right sum", rightSum)
			}
		}
	}

	if show {
		fmt.Println("cross end: left index", maxLeft, "right index", maxRight, "sum", leftSum+rightSum)
	}
	return maxLeft, maxRight, leftSum + rightSum
}

func findMaxSubArray(source []int, low, high int, show bool) (targetLow, targetHigh, targetSum int) {
	if low == high || len(source[low:high]) == 1 {
		if show {
			fmt.Println("side get: index", low, "sum", source[low])
		}
		return low, low, source[low]
	}
	mid := (low + high) / 2
	if show {
		fmt.Println("split array to left", source[low:mid], "and right", source[mid:high])
	}
	leftLow, leftHigh, leftSum := findMaxSubArray(source, low, mid, show)
	rightLow, rightHigh, rightSum := findMaxSubArray(source, mid, high, show)
	crossLow, crossHigh, crossSum := findMaxCrossingSubArray(source, low, mid, high, show)
	if show {
		fmt.Println("left array", source[low:mid], "left max sum", leftSum,
			"right array", source[mid:high], "right max sum", rightSum,
			"cross array", source[low:high], "cross max sum", crossSum)
	}
	if leftSum > rightSum && leftSum > crossSum {
		targetLow, targetHigh, targetSum = leftLow, leftHigh, leftSum
	} else if rightSum > leftSum && rightSum > crossSum {
		targetLow, targetHigh, targetSum = rightLow, rightHigh, rightSum
	} else {
		targetLow, targetHigh, targetSum = crossLow, crossHigh, crossSum
	}
	if show {
		fmt.Println("get: low index", targetLow, "high index", targetHigh, "sum", targetSum)
	}
	return
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
	startIndex, endIndex, maxSum := findMaxSubArray(source, 0, len(source), *show)
	fmt.Println()
	fmt.Println("start index:", startIndex, "end index:", endIndex, "sum:", maxSum)
}
