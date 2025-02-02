package cac

import (
	"strings"
)

// Parse parses the file content then turn to string slice.
func Parse(fileContent string) []string {
	list := strings.Split(fileContent, "\n")

	// Remove commands
	var listWithoutCommands []string
	for i := range list {
		trimmedSpace := strings.TrimSpace(list[i])
		markerIndex := strings.Index(trimmedSpace, "#")
		if markerIndex >= 0 {
			trimmedSpace = trimmedSpace[:markerIndex]
		}
		if trimmedSpace == "" {
			continue
		}
		listWithoutCommands = append(listWithoutCommands, trimmedSpace)
	}

	pureList := strings.Fields(strings.Join(listWithoutCommands, " "))

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
					break
				}
				needMergeCount++
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
