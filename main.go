package main

import (
	"fmt"
	"sort"
)

func main() {

	var purchaseIDs []int

	purchaseIDs = append(purchaseIDs, 12, 11, 2, 11, 112)
	sort.Ints(purchaseIDs)

	fmt.Println(purchaseIDs)

	sort.Slice(purchaseIDs, func(i, j int) bool {
		return purchaseIDs[i] < purchaseIDs[j]
	})

	fmt.Println(purchaseIDs)

}
