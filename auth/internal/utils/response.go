package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/model/common"
)

// SuccessResponse generates a standard success response
func DataSuccessResponse(c *gin.Context, message string, data interface{}) {

	response := common.Response{
		Success: true,
		Code:    200,
		Details: "Request processed successfully",
		Message: message,
		Data:    data,
	}

	c.JSON(response.Code, response)

	// writeJSONResponse(w, response)

}

func SuccessResponse(c *gin.Context, message string) {

	response := common.Response{
		Success: true,
		Code:    200,
		Details: "Request processed successfully",
		Message: message,
		Data:    nil,
	}

	c.JSON(response.Code, response)
}

// ErrorResponse generates a standard error response
func ErrorResponse(c *gin.Context, message string, details string, code ...int) {
	statusCode := 400
	if len(code) > 0 && code[0] != 0 {
		statusCode = code[0]
	}
	response := common.Response{
		Success: false,
		Code:    statusCode,
		Details: details,
		Message: message,
		Data:    nil,
	}
	c.JSON(200, response)
	// c.JSON(response.Code, response)
}

// writeJSONResponse writes the response as JSON to the client
// func writeJSONResponse(w http.ResponseWriter, response common.Response) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(response.Code)
// 	json.NewEncoder(w).Encode(response)
// }
