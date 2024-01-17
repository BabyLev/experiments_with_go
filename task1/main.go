package main

import (
	"fmt"
	"sort"
)

type Order struct {
	ID       int
	name     string
	category string
}

func main() {
	orders := []Order{
		{
			ID:       5,
			name:     "Pavel",
			category: "нижнее бельё",
		},
		{
			ID:       4,
			name:     "Pavel1",
			category: "нижнее бельё",
		},
		{
			ID:       7,
			name:     "Pavel2",
			category: "electronics",
		},
	}
	res := ordersGroupAndSort(orders)
	fmt.Println(res)
	// map {"нижнее бельё":[{
	// ID:       5,
	// name:     "Pavel",
	// category: "нижнее бельё"
	// }, {
	// ID:       4,
	// name:     "Pavel1",
	// category: "нижнее бельё"
	// }]
	// 	}
}

func ordersGroupAndSort(orders []Order) map[string][]Order {
	ordersMap := make(map[string][]Order)
	for _, order := range orders {
		ordersMap[order.category] = append(ordersMap[order.category], order)
	}
	for key := range ordersMap {
		sort.Slice(ordersMap[key], func(i, j int) bool {
			return ordersMap[key][i].ID < ordersMap[key][j].ID
		})
	}

	return ordersMap
}
