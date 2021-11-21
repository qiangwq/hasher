package main

import "fmt"

type Stats struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
}

var total int64
var totalMs int64

func getStats() Stats {
	return Stats{
		Total:   total,
		Average: (int64)(totalMs / total)}
}

func addOP(mili int64) {
	total += 1
	totalMs += mili
	fmt.Printf("op# %d\n", total)
}
