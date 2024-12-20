package company_request

import "mime/multipart"

type CreateCompanyRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	File        *multipart.FileHeader `form:"file"`
	ImageId     *int                  `form:"imageId"`
}
