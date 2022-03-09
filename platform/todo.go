package platform

type Item struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Done bool `json:"done"`
}

type ItemList struct {
	Items []Item 
}

func New() *ItemList {
	return &ItemList{
		Items: []Item{},
	}
}

func (l *ItemList) Add(item Item){
	l.Items = append(l.Items, item)
}

func (l *ItemList) GetAll() []Item {
	return l.Items
}

func (l *ItemList) GetOneItem(target string) Item {
	for _, value := range l.Items{
		if value.Title == target{
			return value
		}
	}
	return Item{}
}

func (l *ItemList) RemoveItem(target string) {
	
}