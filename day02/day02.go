package day02

import (
	"log"
	"os"
	"bufio"
)

func scan(file *os.File, count2 int, count3 int) (int, int){
	scanner := bufio.NewScanner(file)
	cnt2:=0
	cnt3:=0
	for scanner.Scan() {
		cntln2, cntln3 := countLine(scanner.Text())
		cnt2+=cntln2
		cnt3+=cntln3
	}
	return cnt2, cnt3
}

func day02(){
	count2 := 0;
	count3 := 0;
	input, err := os.Open("input")
	defer input.Close()
	if err!=nil {
		log.Fatal(err)
	}
	count2, count3 = scan(input, count2, count3)
	input.Close()
	sumck := count2*count3
	log.Println(sumck)
}

func countLine(line string) (int, int){
	cnt := make(map[rune]int, len(line))
	cnt2, cnt3 := 0, 0
	for _, val := range line {
		cnt[val]++
	}
	for _, val := range cnt{
		if val == 2{
			cnt2 = 1
		}
		if val ==3{
			cnt3 = 1
		}
	}
	return cnt2, cnt3
}