package problem017

import (
	"bufio"
	"regexp"
	"strings"
)

var fileRegex = regexp.MustCompile("^\t*[^\n]*\\..+\n?$")

func isFilename(s string) bool {
	return fileRegex.MatchString(s)
}

func Run(path string) int {
	return longestAbsPathLength(0, path)
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestAbsPathLength(level int, path string) int {
	scanner := bufio.NewScanner(strings.NewReader(path))
	_, topMostName := scanner.Scan(), scanner.Text()
	if isFilename(topMostName) {
		return len(topMostName[level:])
	}

	pathLength := len(topMostName[level:]) + 1
	longestSubPathLength := 0
	subPathBuilder := strings.Builder{}
	for scanned, subPath := scanner.Scan(), scanner.Text(); scanned; scanned, subPath = scanner.Scan(), scanner.Text() {
		tabCount := 0
		for []rune(subPath)[tabCount] == '\t' {
			tabCount++
		}
		if tabCount == level+1 && subPathBuilder.Len() > 0 {
			longestSubPathLength = maxInt(longestSubPathLength, longestAbsPathLength(level+1, subPathBuilder.String()))
			subPathBuilder.Reset()
		}
		subPathBuilder.WriteString(subPath + "\n")
	}

	if subPathBuilder.Len() > 0 {
		longestSubPathLength = maxInt(longestSubPathLength, longestAbsPathLength(level+1, subPathBuilder.String()))
	}

	if longestSubPathLength > 0 {
		return pathLength + longestSubPathLength
	} else {
		return 0
	}
}
