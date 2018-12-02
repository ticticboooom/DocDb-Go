package wsdon

type wsdonItemType int
type Item struct {
	array    []*Item
	object   map[string]*Item
	simple   string
	itemType wsdonItemType
}

const (
	array  wsdonItemType = 0
	simple wsdonItemType = 1
	object wsdonItemType = 2
)

func GetObject(item *Item, name string) *Item {
	if val, ok := item.object[name]; ok {
		return val
	}
	return &Item{}
}
func Set(item *Item, name string, value *Item) {
	item.object[name] = value
}

func GetArray(item *Item, index int) *Item {
	count := len(item.array)
	if count > index && index > 0 {
		return item.array[index]
	}
	return &Item{}
}

func GetSimple(item *Item) string {
	return item.simple
}
