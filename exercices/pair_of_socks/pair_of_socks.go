package pair_of_socks

import (
	"sort"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func getPairOfSocksWithSort(n int, ar []int) int {
	sort.Ints(ar)
	totalOfSocks := 0
	for i:= 0; i < n; i++ {
		if i >= n-1 {
			return totalOfSocks
		}

		if ar[i] == ar[i+1] {
			totalOfSocks++
			i = i+1
		}
	}

	return totalOfSocks
}

func getPairOfSocksWithoutSort(n int, ar []int) int {
	totalOfSocks := 0
	var indixesUsed []int
	for i:= 0; i < n-1; i++ {
		if contains(indixesUsed, i) {
			continue
		}

		for j := i+1; j < n; j++ {
			if contains(indixesUsed, j) {
				continue
			}

			if ar[i] == ar[j] {
				indixesUsed = append(indixesUsed, i, j)
				totalOfSocks++
				break
			}
		}
	}


	return totalOfSocks
}


func contains(data []int, value int) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}
	return false
}