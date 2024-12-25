package coffee_request

import "mime/multipart"

type EditRequest struct {
	CoffeeID    uint                  `form:"coffeeId"`
	Name        string                `form:"name"`
	File        *multipart.FileHeader `form:"file"`
	ImageId     int                   `form:"imageId"`
}
