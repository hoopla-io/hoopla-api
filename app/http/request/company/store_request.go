package company_request

import "mime/multipart"

type StoreRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	File        *multipart.FileHeader `form:"file"`
	ImageId     *int                  `form:"imageId"`
}
