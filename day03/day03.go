package day03

import (
	"fmt"
	"strconv"
	"log"
	"bufio"
	"os"
	"strings"
)

const width = 1000

type line struct{
	id int
	x int
	y int
	w int
	h int
}

func day03_1(){
	file, err := os.Open("input")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	input := scan(file)
	parea := make([][]string, width)
	for i := range parea {
		parea[i] = make([]string, width)
	}
	lines := split(input)
	for _, val := range lines{
		parea = occupy(val, parea)
	}

	cnt := countX(parea)
	log.Println(cnt)
	for _, val := range lines {
		i1, b := findNoOverlap(parea, val)
		if b {
			fmt.Println(strconv.Itoa(i1))

		}

	}
}

func findNoOverlap(parea [][]string, ln line) (int, bool){
	for i:=0; i<ln.w; i++  {
		for j:=0; j<ln.h; j++ {
			if parea[ln.x+i][ln.y+j] != strconv.Itoa(ln.id) {
				return 0,false
			}
		}
	}
	return ln.id, true
}

func countX(parea [][]string) int64 {
	count := int64(0)
	for _, val1 := range parea{
		for _, val2 := range val1{
			if val2 == "X"{
				count++
			}
		}
	}
	return count
}
func occupy(ln line, parea [][]string) [][]string{
	//offs := parea[ln.x][ln.y]
	for i:=0; i<ln.w; i++  {
		for j:=0; j<ln.h; j++ {
			if len(parea[ln.x+i][ln.y+j]) > 0 {
				parea[ln.x+i][ln.y+j] = "X"
			} else {
				parea[ln.x+i][ln.y+j] = strconv.Itoa(ln.id)
			}
		}
	}
	return parea
}
func scan(file *os.File) []string {
	rtrn := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		rtrn = append(rtrn,scanner.Text())
	}
	return rtrn;
}
func split(input []string) ([]line) {
	lines := make([]line,0)
	
	for _, val1 := range input{
		ln := line{}
		val1 = strings.Replace(val1, " ", "", -1)
		tmp := strings.FieldsFunc(val1, delimiter)
		for index, val2 := range tmp {
			if index == 0{
				ln.id = myStrtoInt(val2)
			}
			if index == 1{
				ln.x = myStrtoInt(val2)
			}
			if index == 2{
				ln.y = myStrtoInt(val2)
			}
			if index == 3{
				ln.w = myStrtoInt(val2)
			}
			if index == 4{
				ln.h = myStrtoInt(val2)
			}
		}
		lines = append(lines, ln)
	}
	return lines
}
func delimiter(str rune) bool {
	return str=='#' || str=='@' || str==','|| str==':' || str=='x'
}

func myStrtoInt(str string) int {
	i, err := strconv.ParseInt(str, 10, 0)
	if err != nil{
		log.Println(err)
	}
	return int(i)

}