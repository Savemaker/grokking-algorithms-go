package main

func GetLongestPalindrome(input string) string {
	matrix := initMatrix(input)
	maxString := initLenghtOfOneAndTwo(&matrix, input)
	for i := 2; i < len(input); i++ {
		currentString := initLenghtOfThreeAndMore(&matrix, input, i)
		if len(currentString) > len(maxString) {
			maxString = currentString
		}
	}
	return maxString
}

func initLenghtOfOneAndTwo(matrix *[][]bool, input string) string {
	maxString := string(input[0])
	for i, v := range *matrix {
		v[i] = true
		if i+1 < len(input) {
			if input[i+1] == input[i] {
				v[i+1] = true
				maxString = input[i : i+2]
			}
		}
	}
	return maxString
}

func initLenghtOfThreeAndMore(matrix *[][]bool, input string, index int) string {
	ref := *matrix
	maxString := ""
	maxLen := 0
	for i := 0; i+index < len(input); i++ {
		if input[i] == input[i+index] {
			if ref[i+1][i+index-1] {
				if (index + 1) > maxLen {
					maxLen = index + 1
					maxString = input[i : i+index+1]
				}
				ref[i][i+index] = true
			}
		} else {
			ref[i][i+index] = false
		}
	}
	return maxString
}

func initMatrix(input string) [][]bool {
	matrix := make([][]bool, len(input))
	for i := range matrix {
		matrix[i] = make([]bool, len(input))
	}
	return matrix
}
