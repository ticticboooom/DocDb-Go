package wsdon

import "strconv"

func Stringify(item *Item) string {
	return stringifyWsdon(item, 0)
}

func stringifyWsdon(item *Item, depth int) string {
	result := ""
	if item.itemType == object {
		result += stringifyObject(item.object, depth)
	} else if item.itemType == array {
		result += stringifyArray(item.array, depth)
	} else if item.itemType == simple {
		result += getTabs(depth) + item.simple
	}
	return result
}

func stringifyArray(arr []*Item, depth int) string {
	arrayLength := len(arr)
	result := getTabs(depth) + "(array)\n"
	depth++
	for i := 0; i < arrayLength; i++ {
		result += getTabs(depth)
		result += "[" + strconv.Itoa(i) + "]\n"
		result += stringifyWsdon(arr[i], depth+1)
		result += "\n"
	}
	return result
}

func stringifyObject(obj map[string]*Item, depth int) string {
	result := ""
	for k := range obj {
		result += getTabs(depth) + "[" + k + "]\n"
		result += stringifyWsdon(obj[k], depth+1)
		result += "\n"
	}
	return result
}

func getTabs(depth int) string {
	result := ""
	for i := 0; i < depth; i++ {
		result += "\t"
	}
	return result
}
