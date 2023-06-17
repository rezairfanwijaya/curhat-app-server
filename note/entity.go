package note

type Note struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Author string `json:"author"`
	Note   string `json:"note"`
}
