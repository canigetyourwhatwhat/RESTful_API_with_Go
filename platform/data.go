package platform

type Item struct {
	ID string `json:"id"`
	Title string `json:"title"`	
	Done bool `json:"done"`
}

type ItemList struct {
	Items []Item 
}