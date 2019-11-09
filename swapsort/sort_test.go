package main

import (
	"testing"
)

func TestIsSorted(tst *testing.T) {
	cases := [][]int{
		{1, 2, 3, 5, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 8, 7, 6, 4, 3, 9},
		{1, 2, 3, 4, 6, 7, 8, 9},
		{1, 2, 8, 4, 5, 6, 3, 9},
		{1, 2, 3, 4, 5, 6, 8, 9},
	}

	for idx, t := range cases {
		tst.Logf("Test isSorted %d: source %#v \n", idx, t)
		tst.Logf("\tresult %#v \n", isSorted(t, 1))
	}
}

func TestAlmostSorted(tst *testing.T) {
	cases := []map[string][]int{
		{"in": {1, 2, 3, 5, 4},
			"out": {1, 2, 3, 4, 5},
		},
		{"in": {1, 2, 8, 7, 6, 4, 3, 9},
			"out": {1, 2, 3, 4, 6, 7, 8, 9},
		},
		{"in": {1, 2, 8, 4, 5, 6, 3, 9},
			"out": {1, 2, 3, 4, 5, 6, 8, 9},
		},
	}

	for idx, t := range cases {
		tst.Logf("Test almostSorted %d: source %#v \n", idx, t["in"])
		almostSorted(t["in"])
		tst.Logf("\tresult %#v \n", t["out"])
	}
}
