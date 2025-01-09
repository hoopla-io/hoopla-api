package user_resource

type QrCodeResource struct {
	QrCode     string `json:"qrCode"`
	ExpireAt   int64  `json:"expireAt"`
	ExpireAtMs int64  `json:"expireAtMs"`
}
