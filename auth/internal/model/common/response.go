package common

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Details string      `json:"details"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
