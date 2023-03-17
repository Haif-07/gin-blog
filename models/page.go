package models

type Page struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Link      string `json:"link"`
	Status    int    `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
