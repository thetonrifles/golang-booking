package main

import (
	"fmt"
  "strings"
  "strconv"
  "bufio"
  "os"
  "github.com/thetonrifles/golang-booking/booking"
)

var sums map[int64]bool

func main() {
	var cities int
	var scores string
	fmt.Scanf("%d", &cities)

	consoleReader := bufio.NewReader(os.Stdin)
	scores, _ = consoleReader.ReadString('\n')
	scores = strings.Replace(scores, "\n", "", 1)

	array := strings.Split(scores, " ")
	numbers := []int64{}
	for _, value := range array {
		num, _ := strconv.Atoi(value)
		numbers = append(numbers, int64(num))
	}

	count := booking.HappingessScores(numbers)
	fmt.Printf("%v\n", count)
}
