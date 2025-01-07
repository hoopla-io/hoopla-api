package shops_request

type ShopRequest struct {
	ShopId uint `form:"shopId" binding:"required"`
}
