package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		txt := ""
		if i%3 == 0 {
			txt += "Fizz"
		}
		if i%5 == 0 {
			txt += "Buzz"
		}
		if txt == "" {
			txt = strconv.Itoa(i)
		}
		answer = append(answer, txt)
	}
	return answer
}

func main() {
	fmt.Println(fizzBuzz(3))

}
