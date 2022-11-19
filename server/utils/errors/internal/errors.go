package internal

import 	"github.com/gin-gonic/gin"


func PrintServerError(err error, message string, context *gin.Context) {
	if err != nil {
		context.JSON(500, gin.H{
			"internal error": message + err.Error(),
		})
		return;
	}
}