package main

import (
	"fmt"
	"sort"
)

func main() {

	var purchaseIDs []int

	purchaseIDs = append(purchaseIDs, 11, 122, 2, 99, 53)
	fmt.Println("Original list: ", purchaseIDs)

	sort.Ints(purchaseIDs) // sort values in an ascending manner

	fmt.Println("Ascending Order: ", purchaseIDs)

	sort.Slice(purchaseIDs, func(i, j int) bool { // sort values in a descending order
		return purchaseIDs[i] > purchaseIDs[j] // checks for each value based on criteria
	})

	fmt.Println("Descending Order", purchaseIDs)

	fmt.Println("Value at Index 4: ", purchaseIDs[2])
	value_removed := purchaseIDs[2]
	purchaseIDs = append(purchaseIDs[:2], purchaseIDs[2+1:]...) // remove value at the nth index
	fmt.Printf("removing: %d "+" new values: %d\n", value_removed, purchaseIDs)

}
