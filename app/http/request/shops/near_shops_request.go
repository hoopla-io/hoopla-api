package shops_request

type NearShopsRequest struct {
	Lat  float64 `form:"lat" binding:"required"`
	Long float64 `form:"long" binding:"required"`
}
