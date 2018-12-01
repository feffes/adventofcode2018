package main

import (
	"testing"
)

func TestFindFirstDup(t *testing.T){
	values := make([]int64, 5)
	values[0] = 1
	values[1] = 3
	values[2] = 1
	values[3] = 3
	values[4] = 4
	
	dup := findFirstDup(values)

	if dup != int64(1) {
		t.Errorf("Value incorrect, got %d", dup)
	}
}