package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"os"
	"strconv"
	"strings"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func colEmpty(numbers []int) int {
	t := [9]int{}
	for _, n := range(numbers){
		t[(n-1)/10] += 1
	}

	for i, x := range(t) {
		if x == 0 {
			return i
		}
	}
	return -1
}

func getNumbers() ([]int, []int, []int, bool) {
	ret := []int{}
	alreadyOneRow := false
	fillOrder := []int{}
	Used := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for len(ret) < 15 {
		val := rand.Intn(90) + 1
		usedIdx := (val - 1) / 10
		if Used[usedIdx] == 3 {
			continue
		}

		if Used[usedIdx] == 2 && alreadyOneRow {
			continue
		}

		found := false
		for _, x := range ret {
			if x == val {
				found = true
				break
			}
		}
		if !found {
			Used[usedIdx] = Used[usedIdx] + 1
			if Used[usedIdx] == 3 {
				alreadyOneRow = true
			}
			ret = append(ret, val)
		}
	}
	sort.Ints(ret)
	for _, chk := range []int{3, 2, 1} {
		for ix, r := range Used {
			if r == chk {
				fillOrder = append(fillOrder, ix+1)
			}
		}
	}

	valid := true
	if colEmpty(ret) >= 0 {
		valid = false
	}
	return ret, fillOrder, Used, valid
}

func rowFull(input [9]int) bool {
	num := 0
	for _, i := range input {
		if i > 0 {
			num++
			if num == 5 {
				return true
			}
		}
	}
	return false
}

func rowSize(input [9]int) int {
	num := 0
	for _, i := range input {
		if i > 0 {
			num++
		}
	}
	return num
}


func mostFitRow(table [3][9]int, col int, colSize int) int {
	if colSize == 3{
		for row := 0; row <3; row++ {
			if table[row][col] == 0 {
				return row
			}
		}
	}

	if colSize == 2 {
		rowUsage := [3]int{}
		for row := 0; row <3; row++ {
			rowUsage[row] = rowSize(table[row])
		}

		if table[0][col] == 0 && table[1][col] == 0 &&
			table[2][col] == 0 {
				if rowUsage[0] < rowUsage[1] {
					if rowSize(table[0]) < 5 {
					 return 0
					} else {
					 return 1
					}
				} else {
					if rowSize(table[1]) < 5 {
					 return 1
					} else {
					 return 0
					}
				}
			return 0
		}

		if (table[0][col] > 0 ||  table[1][col] > 0) {
			if table[1][col] > 0 {
				if rowSize(table[2]) < 5 {
					return 2
				}
				return -1
			}
			if rowUsage[1] < rowUsage[2] && rowSize(table[1]) < 5{
				return 1
			} else {
				if rowSize(table[2]) < 5 {
					return 2
				}
				return -1
			}

		}

	}

	if colSize == 1 {
		rowUsage := [3]int{}
		for row := 0; row <3; row++ {
			rowUsage[row] = rowSize(table[row])
		}
		if rowUsage[0] < rowUsage[1] && rowUsage[0] < rowUsage[2] && rowSize(table[0]) < 5 {
			return 0
		}
		if rowUsage[1] < rowUsage[0] && rowUsage[0] < rowUsage[2] && rowSize(table[1]) < 5 {
			return 1
		}
                if rowUsage[2] < rowUsage[0] && rowUsage[0] < rowUsage[1] && rowSize(table[2]) < 5 {
		  return 2
		}
		for ans := 0; ans <3 ; ans ++ {
			if rowSize(table[ans]) < 5 {
				return ans
			}
		}
	}

	panic("Invalid colSize")
}


func createTicket() int {
	ticket := [3][9]int{}
	numbers, fillOrder, used, valid := getNumbers()

	for !valid {
		numbers, fillOrder, used, valid = getNumbers()
	}

	rowUsed := []bool{}
	rowUsed = append(rowUsed, false, false, false)
	for _, i := range fillOrder {
		for _, num := range numbers {
			if num <= i*10 && num > (i-1)*10 {
				idx := mostFitRow(ticket, i-1, used[i-1])
				if idx < 0 {
					return -1
				}
				if ticket[idx][i-1] > 0 {
					panic("Chosen full slot")
				}
				ticket[idx][i-1] = num
				if rowFull(ticket[idx]) {
					rowUsed[idx] = true
				}
			}
		}
	}

	for row := 0; row < 3; row++ {
		for _,v := range(ticket[row]) {
			if v > 0 {
				fmt.Printf("|%2d",v)
			} else {
				fmt.Printf("|  ")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("\n\n")
	return 0
}

func main() {
	rand.Seed(makeTimestamp())
	for _, name := range os.Args[1:] {
		details := strings.Split(name,":")
		fmt.Println(details[0]+":")
		numTickets, _ := strconv.Atoi(details[1])
		for i:=0; i < numTickets; i++ {
		for createTicket() != 0 {
		}
		}
	}
}
