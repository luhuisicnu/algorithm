package main

import (
	"flag"
	"fmt"
	"strconv"
)

func partition(source []int, startIndex, endIndex int, show, desc bool) int {
	if show {
		fmt.Println("partition: source", source[startIndex:endIndex])
	}
	keyValue := source[endIndex-1]
	if show {
		fmt.Println("partition: find last value as key", keyValue)
	}
	i := startIndex - 1
	if show {
		fmt.Println("partition: start part")
	}
	for j := startIndex; j < endIndex-1; j++ {
		if desc {
			if source[j] > keyValue {
				i++
				if show {
					fmt.Println("partition: index", j, "value", source[j], "is more than key value")
					fmt.Println("partition: exchange ", source[i], source[j])
				}
				source[i], source[j] = source[j], source[i]
				if show {
					fmt.Println("partition: middle result", source[startIndex:endIndex])
				}
			} else {
				if show {
					fmt.Println("partition: nothing to do with index", j, "value", source[j])
				}
			}
		} else {
			if source[j] <= keyValue {
				i++
				if show {
					fmt.Println("partition: index", j, "value", source[j], "is less than key value")
					fmt.Println("partition: exchange ", source[i], source[j])
				}
				source[i], source[j] = source[j], source[i]
				if show {
					fmt.Println("partition: middle result", source[startIndex:endIndex])
				}
			} else {
				if show {
					fmt.Println("partition: nothing to do with index", j, "value", source[j])
				}
			}
		}
	}
	if show {
		fmt.Println("partition: exchange key value", keyValue, "and first bigger value", source[i+1])
	}
	source[i+1], source[endIndex-1] = source[endIndex-1], source[i+1]
	if show {
		fmt.Println("partition: result", source[startIndex:endIndex])
	}
	return i + 1
}

func quickSort(source []int, startIndex, endIndex int, show, desc bool) {
	if show {
		fmt.Println("sort: source", source[startIndex:endIndex], "start index", startIndex, "end index", endIndex)
	}
	if startIndex < endIndex && startIndex < endIndex-1 {
		middleIndex := partition(source, startIndex, endIndex, show, desc)
		if show {
			fmt.Println("sort: find a middle index", middleIndex)
		}
		quickSort(source, startIndex, middleIndex, show, desc)
		quickSort(source, middleIndex+1, endIndex, show, desc)
	} else {
		if show {
			fmt.Println("sort: nothing to do")
		}
	}
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

	quickSort(source, 0, len(source), *show, *desc)
	fmt.Println(source)
}
