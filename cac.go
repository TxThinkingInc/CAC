package cac

import (
	"slices"
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
	var currentMergedLineParts []string
	for _, item := range list {
		if endWithValidBackslash(item) {
			currentMergedLineParts = append(currentMergedLineParts, strings.TrimSuffix(item, "\\"))
		} else {
			currentMergedLineParts = append(currentMergedLineParts, item)
			mergedList = append(mergedList, strings.Join(currentMergedLineParts, ""))
			currentMergedLineParts = nil // Reset for next merged line
		}
	}
	return mergedList
}

// endWithValidBackslash checks whether the suffix is valid backslash.
func endWithValidBackslash(str string) bool {
	list := strings.Split(str, "") // Can handle different encoding.
	if len(list) == 0 {
		return false
	}

	slices.Reverse(list)
	var backslashCount int
	for _, word := range list {
		if word != "\\" {
			break
		}
		backslashCount++
	}

	return backslashCount%2 != 0
}
