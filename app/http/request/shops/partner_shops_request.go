package shops_request

type PartnerShopsRequest struct {
	PartnerID uint `form:"partnerId" binding:"required"`
}
