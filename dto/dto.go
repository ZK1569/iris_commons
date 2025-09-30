package dto

type Action string

const (
	Init   Action = "INIT"
	Update Action = "UPDATE"
	Delete Action = "DELETE"
)

type MangaAction struct {
	Action Action  `json:"action"`
	Title  *string `json:"title"`
	Slug   *string `json:"slug"`
}
