package main;

import (
	"strconv"
	"bufio"
	"log"
	"os"
)


func main(){
	input, err := os.Open("input")
	defer input.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	
	value := int64(0);

	for(scanner.Scan()) {
		line, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Println(err)
		}
		value+=line

		log.Println(value)
	}

	
}


