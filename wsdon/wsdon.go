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
	i := 0
	end := findEnd(lines, i)
	if cleanString(lines[i]) == arrayString {
		wsdonItem.array = parseArray(lines[i:end])
		i = end
	} else if titleRegex.MatchString(lines[i]) {
		wsdonItem.object = parseObject(lines[i:end])
		i = end
	} else {
		wsdonItem.simple = cleanString(lines[i])
		i++
	}
	return &wsdonItem
}

func parseObject(lines []string) map[string]*Item {
	itemsObject := map[string]*Item{}
	linesLen := len(lines)
	for i := 0; i < linesLen; {
		item := lines[i]
		end := findEnd(lines, i+1)
		if end != -1 {
			if titleRegex.MatchString(item) && i+1 < end {
				itemsObject[getTitle(lines[i])] = parseWsdonObject(lines[i+1 : end])
				i = end - 1
			} else {
				i++
			}
		}
	}
	return itemsObject
}

func parseArray(lines []string) []*Item {
	wsdonList := make([]*Item, len(lines))
	counter := 0
	i := 1
	linesLength := len(lines)
	for i < linesLength {
		end := findEnd(lines, i+1)
		wsdonList[counter] = parseWsdonObject(lines[i+1 : end])
		counter++
		i = end
	}
	array := make([]*Item, counter)
	for i := 0; i < counter; i++ {
		array[i] = wsdonList[i]
	}
	return array
}

func findEnd(lines []string, startIndex int) int {
	if startIndex >= len(lines) {
		return startIndex
	}
	startIndentation := getIndentCount(lines[startIndex])
	if startIndex+1 == len(lines) {
		return startIndex + 1
	}
	for i := startIndex + 1; i < len(lines); i++ {
		if getIndentCount(lines[i]) < startIndentation {
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
	if strings.Index(in, "\t") != -1 {
		for i := 0; i < len(in); i++ {
			if in[i] == '\t' || in[i] == '\n' {
				start = i + 1
			} else {
				break
			}
		}
	}
	return in[start:]
}
func getTitle(line string) string {
	line = cleanString(line)
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	return line
}
