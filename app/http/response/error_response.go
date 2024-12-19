package response

type ErrorMessages struct {
	RU string `json:"ru"`
	UZ string `json:"uz"`
	EN string `json:"en"`
}

type ErrorResponse struct {
	Code    int           `json:"code"`
	Message ErrorMessages `json:"message"`
	Data    interface{}   `json:"data"`
}