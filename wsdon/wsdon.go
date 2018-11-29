package wsdon

import (
	"regexp"
	"strings"
)

const (
	titleRegexString = "\\[.+\\]"
	arrayString      = "(array)"
)

var (
	titleRegex *regexp.Regexp
)

func init() {
	titleRegex, _ = regexp.Compile(titleRegexString)
}

func ParseWsdon(document string) *Item {
	lines := strings.Split(document, "\n")
	result := parseWsdonObject(lines)
	return result
}

func parseWsdonObject(lines []string) *Item {
	var wsdonItem = Item{}
	if cleanString(lines[0]) == arrayString {
		wsdonItem.array = parseArray(lines[0:findEnd(lines, 0)])
	} else if titleRegex.MatchString(lines[0]) {
		wsdonItem.object = parseObject(lines[0:findEnd(lines, 0)])
	} else {
		wsdonItem.simple = cleanString(lines[1])
	}
	return &wsdonItem
}

func parseObject(lines []string) map[string]*Item {
	itemsObject := map[string]*Item{}
	for i := 0; i < len(lines); i++ {
		item := lines[i]
		if titleRegex.MatchString(item) {
			itemsObject[getTitle(lines[i])] = parseWsdonObject(lines[i:findEnd(lines, i)])
		}
	}
	return itemsObject
}

func parseArray(lines []string) []*Item {
	wsdonList := make([]*Item, len(lines))
	counter := 0
	i := 0
	linesLength := len(lines)
	for i >= linesLength {
		end := findEnd(lines, i)
		wsdonList[counter] = parseWsdonObject(lines[i:end])
		counter++
		i += end
	}
	array := make([]*Item, counter)
	for i := 0; i < len(wsdonList); i++ {
		array[i] = wsdonList[i]
	}
	return array
}

func findEnd(lines []string, startIndex int) int {
	startIndentation := getIndentCount(lines[startIndex])
	for i := startIndex + 1; i < len(lines); i++ {
		if getIndentCount(lines[i]) <= startIndentation {
			return i
		}
	}
	return len(lines)
}

func getIndentCount(line string) int {
	return strings.Count(line, "\t")
}
func cleanString(in string) string {
	start := 0
	for i := 0; i < len(in); i++ {
		if in[i] != '\t' && in[i] != '\n' {
			start = i
		}
	}
	return in[start : len(in)-1]
}
func getTitle(line string) string {
	line = cleanString(line)
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	return line
}
