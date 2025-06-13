package vendors_poster_request

type WebhookRequest struct {
	Account       string `json:"account"`
	AccountNumber string `json:"account_number"`
	Object        string `json:"object"`
	ObjectID      int64  `json:"object_id"`
	Action        string `json:"action"`
	Time          string `json:"time"`
	Verify        string `json:"verify"`
}
