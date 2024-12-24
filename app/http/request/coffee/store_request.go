package coffee_request

import "mime/multipart"

type StoreRequest struct {
	Name        string                `form:"name"`
	File        *multipart.FileHeader `form:"file"`
	ImageId     *int                  `form:"imageId"`
}
