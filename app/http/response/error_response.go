package response

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// var SessionExpired = ErrorResponse{
// 	Code: 422,
// 	Message: ErrorMessages{
// 		RU: "Сессия истекла!",
// 		UZ: "Seans muddati tugagan!",
// 		EN: "Session expired!",
// 	},
// }

// var InvalidCode = ErrorResponse{
// 	Code: 422,
// 	Message: ErrorMessages{
// 		RU: "Неверный код!",
// 		UZ: "Notog'ri kod!",
// 		EN: "Invalid code!",
// 	},
// }