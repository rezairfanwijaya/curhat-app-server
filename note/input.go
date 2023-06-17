package note

type InputNewNote struct {
	Author string `json:"author"`
	Note   string `json:"note" binding:"required"`
}
