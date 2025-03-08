package domain

type Message struct {
	ID    int     `json:"Id"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}
