package partners_request

type PartnerRequest struct {
	ID uint `form:"id" binding:"required"`
}
