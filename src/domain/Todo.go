package domain

type Todo struct {
	ID          int    `gorm:"primary_key;not null"       json:"id"`
	Name        string `gorm:"type:varchar(200);not null" json:"name"`
	Description string `gorm:"type:varchar(400)"          json:"description"`
	State       string `gorm:"type:varchar(200);not null" json:"state"`
}

type Todos []*Todo
