package platform

import (
	"errors"
)

type Item struct {
	ID string `json:"id"`
	Title string `json:"title"`	
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

func (l *ItemList) GetOneItem(target string) (*Item, error) {
	for index, value := range l.Items{
		if value.ID == target{
			return &l.Items[index], nil
		}
	}
	return nil, errors.New("couldn't find the item'")
}

func (l *ItemList) RemoveItem(target string) {
	
}