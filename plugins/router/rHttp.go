package router

import "github.com/gin-gonic/gin"

type Http interface {
	AddRouter(router *gin.Engine)
}