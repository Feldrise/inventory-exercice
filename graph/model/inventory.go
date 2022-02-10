package model

type Inventory struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	UserID  *string   `json:"user"`
	ItemsID []*string `json:"items"`
}
