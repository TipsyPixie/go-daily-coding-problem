package problem046

func GetPalindromicSubstring(values string) string {
	table := make([][]int, len(values), len(values))
	palindromeEndsAt, palindromeLength := 0, 0
	for rowIndex := range values {
		rowCharacter := rune(values[len(values)-1-rowIndex])
		table[rowIndex] = make([]int, len(values), len(values))
		for columnIndex, columnCharacter := range values {
			if columnCharacter == rowCharacter {
				if columnIndex == 0 || rowIndex == 0 {
					table[rowIndex][columnIndex] = 1
				} else {
					table[rowIndex][columnIndex] = table[rowIndex-1][columnIndex-1] + 1
				}
				if palindromeLength < table[rowIndex][columnIndex] {
					palindromeEndsAt, palindromeLength = columnIndex, table[rowIndex][columnIndex]
				}
			} else {
				table[rowIndex][columnIndex] = 0
			}
		}
	}
	return values[palindromeEndsAt+1-palindromeLength : palindromeEndsAt+1]
}
