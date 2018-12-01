package main;

import (
	"strconv"
	"bufio"
	"log"
	"os"
)


func main(){


	val := int64(0)
	values := make([]int64,0)
	values = append(values, val)
	dup := int64(0)
	for dup == 0 {
		input, err := os.Open("input")
		if err != nil {
			log.Fatal(err)
		}
		val, values = scan(input, val, values)
		dup = findFirstDup(values)
		input.Close()
	}
	log.Println(dup)
}
func scan(file *os.File, val  int64, values  []int64) (int64, []int64) {
	scanner := bufio.NewScanner(file)
	for(scanner.Scan()) {
		line, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Println(err)
		}
		val+=line
		values = append(values, val)
	}
	return val, values
}

func findFirstDup(values []int64) int64 {
	count := make(map[int64]bool)
	for _, val := range values  {
		if count[val]{
			return val
		}
		count[val] = true
	} 
	return 0
}
