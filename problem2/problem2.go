package problem2

func Run(numbers []int) (result []int) {
    length := len(numbers)
    result = make([]int, length, length)
    result[0] = 1
    sidekicks := make([]int, length, length)
    sidekicks[length-1] = 1

    for index, _ := range numbers {
        if index > 0 {
            result[index] = result[index-1] * numbers[index-1]
            reversedIndex := length - 1 - index
            sidekicks[reversedIndex] = sidekicks[reversedIndex+1] * numbers[reversedIndex+1]
        }
    }
    for index, _ := range result {
        result[index] *= sidekicks[index]
    }
    return
}
