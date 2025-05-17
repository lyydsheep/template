package middleware

import "github.com/gin-gonic/gin"

func GetHandlerFunc() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		StartTrace(),
		LogAccess(),
		PanicRecorder(),
	}
}
