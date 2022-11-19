package client

import 	"github.com/gin-gonic/gin"


func PrintClientError(err error, message string, context *gin.Context) {
	if err != nil {
		context.JSON(400, gin.H{
			"error": message + err.Error(),
		})
		return;
	}
}