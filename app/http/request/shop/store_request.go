package shop_request

import "mime/multipart"

type StoreRequest struct {
	Name        string                `form:"name"`
	Location    string                `form:"location"`
	File        *multipart.FileHeader `form:"file"`
	ImageId     *int                  `form:"imageId"`
	CompanyId   int                   `form:"companyId"`
}
