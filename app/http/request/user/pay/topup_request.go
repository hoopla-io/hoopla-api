package user_pay_request

type TopUpRequest struct {
	ID     uint64  `form:"id" binding:"required"`
	Amount float64 `form:"amount" binding:"required"`
}
