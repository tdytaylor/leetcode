package issue_0443

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {

	return 0
}

func compressToStr(chars []byte) int {
	nextCharIndex, writeIndex, length := 0, 0, len(chars)
	for index := 0; index < length; index++ {
		if index+1 == length || chars[index] != chars[index+1] {
			lastCharIndex := nextCharIndex
			nextCharIndex = index + 1
			subLength := nextCharIndex - lastCharIndex
			itoa := strconv.Itoa(subLength)

			chars[writeIndex] = chars[index]
			writeIndex++
			if subLength > 1 {
				for _, char := range itoa {
					chars[writeIndex] = byte(char)
					writeIndex++
				}
			}
		}
	}
	return writeIndex
}

func compressToStr2(chars []byte) []byte {
	nextCharIndex, writeIndex, length := 0, 0, len(chars)
	for index := 0; index < length; index++ {
		if index+1 == length || chars[index] != chars[index+1] {
			lastCharIndex := nextCharIndex
			nextCharIndex = index + 1
			subLength := nextCharIndex - lastCharIndex
			itoa := strconv.Itoa(subLength)

			chars[writeIndex] = chars[index]
			writeIndex++
			if subLength > 1 {
				for _, char := range itoa {
					chars[writeIndex] = byte(char)
					writeIndex++
				}
			}
		}
	}
	fmt.Println(string(chars[0:writeIndex]))
	return chars[0:writeIndex]
}
