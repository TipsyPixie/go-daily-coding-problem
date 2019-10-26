package problem2

func Run(numbers []int) (result []int) {
    length := len(numbers)
    result = make([]int, length, length)
    sidekicks := make([]int, length, length)

    for index, _ := range numbers {
        reversedIndex := length - 1 - index
        if index == 0 {
            result[index] = 1
            sidekicks[reversedIndex] = 1
        } else {
            result[index] = result[index-1] * numbers[index-1]
            sidekicks[reversedIndex] = sidekicks[reversedIndex+1] * numbers[reversedIndex+1]
        }
    }
    for index, _ := range result {
        result[index] *= sidekicks[index]
    }
    return
}
