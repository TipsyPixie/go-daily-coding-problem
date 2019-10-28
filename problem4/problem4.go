package problem4

import (
    "math"
)

func Run(input []int) int {
    var l, r = 0, len(input)
    for l < r {
        if input[l] > 0 {
            l++
        } else {
            if r < len(input) && input[r] > 0 {
                input[l], input[r] = input[r], input[l]
            } else {
                r--
            }
        }
    }
    input = input[:l]

    for _, value := range input {
        if absoluteValue := int(math.Abs(float64(value))); absoluteValue <= len(input) && input[absoluteValue-1] > 0 {
            input[absoluteValue-1] *= -1
        }
    }

    smallestMissing := 1
    for smallestMissing <= len(input) && input[smallestMissing-1] < 0 {
        smallestMissing++
    }
    return smallestMissing
}
