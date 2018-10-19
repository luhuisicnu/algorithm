package main

import (
	"flag"
	"fmt"
	"strconv"
)

const intMax = int(^uint(0) >> 1)
const intMin = ^intMax

func merge(source []int, p, q, r int, show, desc bool) {
	L := []int{}
	R := []int{}
	L = append(L, source[p:q]...)
	R = append(R, source[q:r]...)
	i := 0
	j := 0
	if desc {
		L = append(L, intMin)
		R = append(R, intMin)
	} else {
		L = append(L, intMax)
		R = append(R, intMax)
	}
	for k := p; k < r; k++ {
		if desc {
			if L[i] > R[j] {
				source[k] = L[i]
				if show {
					fmt.Println("select left", L[i])
				}
				i++
			} else {
				source[k] = R[j]
				if show {
					fmt.Println("select right", R[j])
				}
				j++
			}
		} else {
			if L[i] <= R[j] {
				source[k] = L[i]
				if show {
					fmt.Println("select left", L[i])
				}
				i++
			} else {
				source[k] = R[j]
				if show {
					fmt.Println("select right", R[j])
				}
				j++
			}
		}
	}
}

func mergeSort(source []int, p, r int, show, desc bool) {
	if p < r-1 {
		q := (p + r) / 2
		if show {
			fmt.Println("split", source[p:r], "after index", (r-p)/2-1, "value", source[q-1],
				"left", source[p:q], "right", source[q:r])
		}
		mergeSort(source, p, q, show, desc)
		mergeSort(source, q, r, show, desc)
		if show {
			fmt.Println("merge left", source[p:q], "and right", source[q:r])
		}
		merge(source, p, q, r, show, desc)
		if show {
			fmt.Println("status", source[p:r])
		}
	}
}

func main() {
	show := flag.Bool("s", false, "show detaiL flag")
	desc := flag.Bool("d", false, "desc flag")
	flag.Parse()

	source := []int{}
	for _, arg := range flag.Args() {
		i, _ := strconv.Atoi(arg)
		source = append(source, i)
	}
	fmt.Println("start sort:", source)
	fmt.Println()
	mergeSort(source, 0, len(source), *show, *desc)
	fmt.Println()
	fmt.Println(source)
}
