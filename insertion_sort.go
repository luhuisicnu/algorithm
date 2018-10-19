package main

import (
	"flag"
	"fmt"
	"strconv"
)

func insertionSort(source []int, show, desc bool) {
	if show {
		fmt.Println("start sort:", source)
		fmt.Println()
	}

	for i := 1; i < len(source); i++ {
		key := source[i]
		if show {
			fmt.Println("now insert index", i, "value", key)
		}
		j := i - 1
		if desc {
			for ; j >= 0 && source[j] < key; j-- {
				if show {
					fmt.Println("index:", j, "value:", source[j], "back 1 position")
				}
				source[j+1] = source[j]
			}
		} else {
			for ; j >= 0 && source[j] > key; j-- {
				if show {
					fmt.Println("index:", j, "value:", source[j], "back 1 position")
				}
				source[j+1] = source[j]
			}
		}
		if show && j == i-1 {
			fmt.Println("keep status quo")
		}

		source[j+1] = key
		if show {
			fmt.Println("status:", source)
			fmt.Println()
		}
	}
	return
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

	insertionSort(source, *show, *desc)
	fmt.Println(source)
}
