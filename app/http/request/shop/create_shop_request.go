package shop_request

import "mime/multipart"

type CreateShopRequest struct {
	CompanyID int                   `form:"companyId"`
	File      *multipart.FileHeader `form:"file"`
	Name      string                `form:"name"`
	Location  string     			`form:"location"`
}