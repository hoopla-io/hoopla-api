package shop_request

import "mime/multipart"

type EditRequest struct {
	ShopID      uint                  `form:"shopId"`
	CompanyID   uint                  `form:"companyId"`
	Name        string                `form:"name"`
	Location    string                `form:"location"`
	File        *multipart.FileHeader `form:"file"`
	ImageId     int                   `form:"imageId"`
	CoffeeIds   []uint                `form:"coffeeIds"`
}
