package problem7

import (
    "strconv"
)

func Run(input string) int {
    length := len(input)
    switch {
    case length == 1:
        if asInteger, _ := strconv.Atoi(input); asInteger < 1 {
            return 0
        } else {
            return 1
        }
    case length == 2:
        if asInteger, _ := strconv.Atoi(input); asInteger < 10 || asInteger > 26 {
            return 0
        } else if input[1] == '0' {
            return 1
        } else {
            return 2
        }
    case length >= 3:
        answer := Run(input[:length-1]) * Run(input[length-1:])
        if asInteger, _ := strconv.Atoi(input[length-2:]); asInteger >= 10 || asInteger >= 26 {
            answer += Run(input[:length-2])
        }
        return answer
    default:
        return 0
    }
}
