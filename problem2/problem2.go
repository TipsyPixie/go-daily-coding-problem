package problem2

func Run(numbers []int) (result []int) {
    length := len(numbers)
    lefties := make([]int, length, length)
    lefties[0] = 1
    righties := make([]int, length, length)
    righties[length - 1] = 1

    for index, _ := range numbers {
        if index > 0 {
            lefties[index] = lefties[index - 1] * numbers[index - 1]
            reversedIndex := length - 1 - index
            righties[reversedIndex] = righties[reversedIndex + 1] * numbers[reversedIndex + 1]
        }
    }

    result = make([]int, length, length)
    for index, _ := range lefties {
        result[index] = lefties[index] * righties[index]
    }
    return
}
