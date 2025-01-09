package partner_resource

type DrinksCollection struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	PictureUrl *string `json:"pictureUrl"`
}
