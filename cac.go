package cac

import (
	"strings"
)

// Parse parses the file content then turn to string slice.
func Parse(fileContent string) []string {
	list := strings.Split(fileContent, "\n")

	// Remove commands
	var listWihoutCommands []string
	for i := range list {
		trimedSpace := strings.TrimSpace(list[i])
		if strings.HasPrefix(trimedSpace, "#") {
			continue
		}
		listWihoutCommands = append(listWihoutCommands, trimedSpace)
	}
	clear(list)
	list = nil

	pureList := strings.Fields(strings.Join(listWihoutCommands, " "))

	return mergeLine(pureList)
}

// mergeLine merges near line if the first line ends with backslash.
func mergeLine(list []string) (mergedList []string) {
	length := len(list)
	for i := 0; i < length; i++ {
		if endWithValidBackslash(list[i]) {
			var needMergeCount int
			for j := i + 1; j < length; j++ {
				if !endWithValidBackslash(list[j]) {
					needMergeCount = j - i
					break
				}
			}
			mergedList = append(mergedList, list[i:i+needMergeCount]...)
			i += needMergeCount
			continue
		}
		mergedList = append(mergedList, list[i])
	}

	return
}

// endWithValidBackslash used to check wheaher the suffix is valid backslash.
func endWithValidBackslash(str string) bool {
	list := strings.Split(str, "")
	length := len(list)
	if length == 0 {
		return false
	}

	var backslashCount int
	for i := length - 1; i >= 0; i-- {
		if list[i] != "\\" {
			break
		}
		backslashCount++
	}

	return backslashCount%2 != 0
}
