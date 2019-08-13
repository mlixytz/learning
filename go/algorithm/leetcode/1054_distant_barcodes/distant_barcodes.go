package leetcode1054

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	counter := make(map[int]int)
	for _, v := range barcodes {
		counter[v]++
	}
	orderKey := make([][]int, 0, len(counter))
	for k, v := range counter {
		orderKey = append(orderKey, []int{k, v})
	}

	sort.Slice(orderKey, func(i, j int) bool { return orderKey[i][1] > orderKey[j][1] })

	j := 0
	for _, val := range orderKey {
		for val[1] > 0 {
			barcodes[j] = val[0]
			j += 2
			val[1]--
			if j >= len(barcodes) {
				j = 1
			}
		}
	}
	return barcodes
}
