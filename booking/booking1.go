package booking

import "math"

var sums map[int64]bool

func HappingessScores(scores []int64) int {
  sums = make(map[int64]bool)
  sumScores(scores, 0, 0)
  count := 0
  for key, _ := range sums {
    if isPrime(key) {
      count++
    }
  }
  return count
}

func sumScores(numbers []int64, starting int, sum int64) {
    if len(numbers) == starting {
        if sum != 0 {
          sums[sum] = true
        }
        return;
    }
    value := sum + numbers[starting]
    sumScores(numbers, starting + 1, value)
    sumScores(numbers, starting + 1, sum)
}

func isPrime(a int64) bool {
	if a <= 1 {
		return false
	}
	for b := sqrt(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a % b == 0 {
			return false
		}
	}

	return true
}

func sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}
